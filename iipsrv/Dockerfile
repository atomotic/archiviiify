FROM debian
RUN apt update && apt install -y build-essential git libtiff-dev autoconf libtool pkg-config spawn-fcgi cmake libmemcached-dev
RUN git clone https://github.com/uclouvain/openjpeg /tmp/openjpeg; cd /tmp/openjpeg; cmake .; make; make install
RUN printf "include /etc/ld.so.conf.d/*.conf\ninclude /usr/local/lib\n" > /etc/ld.so.conf && ldconfig
RUN git clone https://github.com/ruven/iipsrv.git /tmp/iipsrv
WORKDIR /tmp/iipsrv
RUN ./autogen.sh && ./configure && make
WORKDIR /iipsrv
RUN cp /tmp/iipsrv/src/iipsrv.fcgi .
