
Instalado:

    sudo apt install qemu qemu-kvm rsync

    echo <Nombre del equipo> > src/etc/hostname
    sudo systemd-nspawn -bD src

        apt install grub-pc linux-image-amd64 lvm2
        echo "/dev/sda1    /    ext4    defaults    0    1" > /etc/fstab

    fallocate --length 500M disk.img
    sudo fdisk disk.img

        g - Tabla de particiones GPT
        n - Nueva partición
            [Enter] - Número
            [Enter] - Sector inicial
            [Enter] - Tamaño
        x - Entrar en modo experto
            A - Activar flag de boot para soporte a MBR
            r - Regresar al menú normal
        w - Guardar y salir

    sudo losetup --partscan --show --find disk.img
    sudo mkfs.ext4 <primera partición>
    sudo mount /mnt <primera partición>
    sudo rsync -ah --info=progress2 src/ /mnt/
    sudo umount /mnt
    sudo losetup -d <unidad>
    sudo qemu-system-x86_64 -m 512 -kernel src/boot/vmlinuz-4.9.0-3-amd64 -append "root=/dev/sda1" -initrd src/boot/initrd.img-4.9.0-3-amd64 -hda <unidad>

        grub-install /dev/sda
        update-grub

    sudo qemu-system-x86_64 -enable-kvm -cpu host -smp 2 -m 512 -hda <unidad>
