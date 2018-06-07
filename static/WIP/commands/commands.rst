========
Commands
========

BTRFS
=====

Verify
------

.. code:: sh

    su

.. code:: sh

    btrfs check <device>

If errors were found

.. code:: sh

    su

.. code:: sh

    btrfs check --repair <device>

.. code:: sh

    btrfs rescue super-recover -y <device>

.. code:: sh

    btrfs scrub start -B <device>

If you get errors like:

.. code::

    ? replay_one_dir_item+0xb5/0xb5 [btrfs]
    ? walk_log_tree+0x9c/0x19d [btrfs]
    ? btrfs_read_fs_root_no_radix+0x169/0x1a1 [btrfs]
    ? btrfs_recover_log_trees+0x195/0x29c [btrfs]
    ? replay_one_dir_item+0xb5/0xb5 [btrfs]
    ? btree_read_extent_buffer_pages+0x76/0xbc [btrfs]
    ? open_ctree+0xff6/0x132c [btrfs]

.. code:: sh

    su

.. code:: sh

    btrfs rescue zero-log <device>