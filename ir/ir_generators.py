# ir_generators.py
from __future__ import annotations
from dataclasses import dataclass
from enum import StrEnum
from typing import Literal, Optional, Dict
from .ir_base import Vec3, Size3, Pose, MaterialRef, Anchor, FaceID
from .ir_schema import Part, Panel, Door

# -------- Panel generator --------
@dataclass(frozen=True)
class PanelParams:
    id: str
    size: Size3                # final X,Y,Z of the slab
    pose: Pose                 # where to place in parent space
    material: MaterialRef
    # optional precomputed anchors you want on every panel:
    want_edge_anchors: bool = True

def gen_panel(p: PanelParams) -> Panel:
    part = Panel(
        id=p.id, kind="panel", size=p.size, pose=p.pose, material=p.material
    )
    if p.want_edge_anchors:
        # edge midpoints for easy mating (left/right/front/back/top/bottom)
        sx, sy, sz = p.size.x, p.size.y, p.size.z
        ox, oy, oz = p.pose.origin.x, p.pose.origin.y, p.pose.origin.z

        def anchor(name: str, pos_local: tuple[float,float,float]) -> None:
            ax, ay, az = pos_local
            part.add_anchor(name, Anchor(
                id=name,
                pose=Pose(origin=Vec3(ox+ax, oy+ay, oz+az))
            ))

        anchor("LeftMid",   (0,       sy/2, sz/2))
        anchor("RightMid",  (sx,      sy/2, sz/2))
        anchor("FrontMid",  (sx/2,    0,    sz/2))
        anchor("BackMid",   (sx/2,    sy,   sz/2))
        anchor("TopMid",    (sx/2,    sy/2, sz))
        anchor("BottomMid", (sx/2,    sy/2, 0))
    return part

# -------- Door generator --------
class DoorStyle(StrEnum):
    OVERLAY = "overlay"
    INSET   = "inset"

@dataclass(frozen=True)
class DoorEdgeRoles:
    left:  Literal["external","shared","hinge"] = "external"
    right: Literal["external","shared","hinge"] = "external"
    top:   Literal["external","shared","hinge"] = "external"
    bottom:Literal["external","shared","hinge"] = "external"

@dataclass(frozen=True)
class Opening:
    x0: float; z0: float; w: float; h: float

@dataclass(frozen=True)
class DoorParams:
    id: str
    thickness: float                         # Y (depth)
    style: DoorStyle = DoorStyle.OVERLAY
    reveal: float = 2.0
    gap_mid: float = 2.0                     # used when an edge is shared (pair)
    edge_role: DoorEdgeRoles = DoorEdgeRoles()
    hinge_side: Literal["left","right"] = "left"
    opening: Opening = Opening(0,0,0,0)      # rectangle on the carcass front
    # optional:
    material: Optional[MaterialRef] = None

    # advanced per-edge overrides; if None → fall back to reveal
    reveal_edges: Optional[Dict[str, float]] = None
    # inset conceal for bottom (optional)
    conceal_bottom_lip: float = 0.0

def _edge_margin(edge: str, params: DoorParams) -> float:
    # base margin
    base = params.reveal_edges.get(edge, params.reveal) if params.reveal_edges else params.reveal
    role = getattr(params.edge_role, edge)
    if role == "shared":
        # split the middle gap/reveal between neighbors (half each)
        return (params.gap_mid if params.style == DoorStyle.OVERLAY else base) * 0.5
    return base

def gen_door(p: DoorParams) -> Door:
    # 1) compute effective margins
    L = _edge_margin("left", p)
    R = _edge_margin("right", p)
    T = _edge_margin("top", p)
    B = _edge_margin("bottom", p)

    # 2) compute slab size (X,Y,Z) and position (Pose.origin)
    door_w = max(0.0, p.opening.w - L - R)
    door_h = max(0.0, p.opening.h - T - B)
    t      = p.thickness

    if p.style == DoorStyle.OVERLAY:
        # sits in front: negative Y offset by thickness
        origin = Vec3(p.opening.x0 - L, -t, p.opening.z0 - T)
    else:
        # inset: sits inside opening
        origin = Vec3(p.opening.x0 + L, 0.0, p.opening.z0 + T)
        if p.conceal_bottom_lip > 0:
            origin = Vec3(origin.x, origin.y, origin.z - p.conceal_bottom_lip)
            door_h += p.conceal_bottom_lip

    size = Size3(door_w, t, door_h)

    mat = p.material or MaterialRef(id="GenericDoor", thickness_mm=t, color_rgb=None)

    part = Door(
        id=p.id, kind="door", size=size, pose=Pose(origin=origin), material=mat,
        meta={"style": p.style, "hinge_side": p.hinge_side}
    )

    # 3) add anchors (hinge & face center)
    hx = origin.x if p.hinge_side == "left" else origin.x + size.x
    hy = origin.y + (size.y if p.style == DoorStyle.OVERLAY else 0.0)
    hz = origin.z + size.z * 0.5

    part.add_anchor("Hinge", Anchor("Hinge", Pose(Vec3(hx, hy, hz))))
    part.add_anchor("FaceCenter", Anchor("FaceCenter", Pose(Vec3(origin.x + size.x/2,
                                                                 hy,
                                                                 origin.z + size.z/2))))
    return part