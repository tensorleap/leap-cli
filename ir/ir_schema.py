# ir_schema.py
from __future__ import annotations
from dataclasses import dataclass, field
from typing import Dict, List, Optional
from .ir_base import Vec3, Size3, Pose, MaterialRef, Anchor, MachiningOp

@dataclass
class Part:
    id: str
    kind: str                 # "panel", "shelf", "door", ...
    size: Size3               # local bounding box
    pose: Pose                # in parent coords
    material: MaterialRef
    anchors: Dict[str, Anchor] = field(default_factory=dict)
    ops: List[MachiningOp] = field(default_factory=list)
    meta: Dict[str, object] = field(default_factory=dict)

    def add_anchor(self, name: str, anchor: Anchor) -> None:
        self.anchors[name] = anchor

    def add_op(self, op: MachiningOp) -> None:
        self.ops.append(op)

# If you want light type aliases for readability:
@dataclass
class Panel(Part):
    pass

@dataclass
class Shelf(Part):
    pass

@dataclass
class Door(Part):
    # door-specific meta is optional; most geometry is in Part
    pass