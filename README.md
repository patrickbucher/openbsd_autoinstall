# OpenBSD Autoinstall

Run a web server to provide an OpenBSD `autoinstall(8)` file:

    $ go run openbsd_autoinstall.go install.conf
    serving /install.conf on port 8000

Modify the `install.conf` file as needed.

Make sure to use port 8000 when defining the http source for the autoinstall
file, e.g. `http://192.168.1.1:8000/install.conf`.
