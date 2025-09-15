# ir_base.py
from __future__ import annotations
from dataclasses import dataclass, field
from enum import Enum, StrEnum
from typing import Dict, List, Tuple, Optional

# ---------- math primitives ----------
@dataclass(frozen=True)
class Vec3:
    x: float
    y: float
    z: float

    def __add__(self, o: "Vec3") -> "Vec3": return Vec3(self.x+o.x, self.y+o.y, self.z+o.z)
    def __sub__(self, o: "Vec3") -> "Vec3": return Vec3(self.x-o.x, self.y-o.y, self.z-o.z)

@dataclass(frozen=True)
class Size3:
    x: float  # width  (local X)
    y: float  # depth  (local Y)
    z: float  # height (local Z)

@dataclass(frozen=True)
class Pose:
    origin: Vec3                   # local origin in parent space
    # simple orientation: local axes expressed in parent space
    x_axis: Vec3 = Vec3(1,0,0)
    y_axis: Vec3 = Vec3(0,1,0)
    z_axis: Vec3 = Vec3(0,0,1)

# ---------- materials & appearance ----------
class GrainAxis(StrEnum):
    x = "x"; y = "y"; z = "z"

@dataclass(frozen=True)
class MaterialRef:
    id: str                              # e.g. "Mel18", "Oak20", "HDF4"
    thickness_mm: float                  # critical for joinery & sizing
    color_rgb: Tuple[float,float,float] | None = None
    grain_axis: GrainAxis | None = None  # preferred grain direction

# ---------- anchoring & faces ----------
class FaceID(StrEnum):
    LEFT="left"; RIGHT="right"; FRONT="front"; BACK="back"; TOP="top"; BOTTOM="bottom"

@dataclass(frozen=True)
class Anchor:
    id: str
    pose: Pose

# ---------- machining / finishing operations ----------
class OpKind(StrEnum):
    DRILL = "drill"
    POCKET = "pocket"
    EDGEBAND = "edgeband"   # on an edge (face+segment)
    NOTE = "note"           # metadata only

@dataclass(frozen=True)
class DrillOp:
    face: FaceID
    at: Vec3               # local coords on that face (x,y,z in part space)
    diameter: float
    depth: float

@dataclass(frozen=True)
class PocketOp:
    face: FaceID
    rect_origin: Vec3      # local
    rect_size: Size3       # use x,z on a face; y is depth of pocket cut

@dataclass(frozen=True)
class EdgeBandOp:
    edge: FaceID           # e.g. apply band on LEFT/RIGHT/FRONT/BACK edges
    thickness: float
    material: str | None = None

MachiningOp = DrillOp | PocketOp | EdgeBandOp

# ---------- utility id ----------
def make_id(prefix: str, counter: int) -> str:
    return f"{prefix}_{counter:04d}"


    