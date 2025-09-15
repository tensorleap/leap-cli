from ir_base import Vec3, Size3, Pose, MaterialRef
from ir_generators import PanelParams, gen_panel, DoorParams, Opening, gen_door, DoorStyle, DoorEdgeRoles

MEL18 = MaterialRef(id="Mel18", thickness_mm=18.0, color_rgb=(0.92,0.92,0.92))
HDF4  = MaterialRef(id="HDF4",  thickness_mm=4.0,  color_rgb=(0.78,0.78,0.78))

# Side panel
side_left = gen_panel(PanelParams(
    id="Side_L",
    size=Size3(18, 600, 2000),
    pose=Pose(origin=Vec3(0,0,0)),
    material=MEL18
))

# A single inset door that covers an opening
door = gen_door(DoorParams(
    id="Door_L",
    thickness=20,
    style=DoorStyle.INSET,
    reveal=2.0,
    hinge_side="left",
    edge_role=DoorEdgeRoles(left="hinge", right="shared", top="external", bottom="external"),
    opening=Opening(x0=18, z0=18, w=764, h=1964),  # e.g. inner opening of a carcass
    material=MEL18,
    conceal_bottom_lip=3.0
))

print(side_left)
print(door)