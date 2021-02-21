Safe Channel
------------

The library tries to remove some `dangerous` behaviors on normal channel:
+ Write on a closed channel
+ Close a closed channel

These above operations will be silently discarded.
