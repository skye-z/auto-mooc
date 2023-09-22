FROM debian:bullseye
MAINTAINER skye-z <skai-zhang@hotmail.com>

WORKDIR /data
COPY auto-mooc .

RUN chmod +x auto-mooc && \
    apt-get update && \
    apt-get install wget -y && \
    ./auto-mooc init
    
RUN apt-get install libsoup2.4-1 libgstreamer1.0-0 libgtk-3-0 libpangocairo-1.0-0 \
    libpango-1.0-0 libharfbuzz0b libcairo-gobject2 libcairo2 libgdk-pixbuf2.0-0 libatomic1 \
    libsqlite3-0 libxslt1.1 libepoxy0 liblcms2-2 libevent-2.1-7 libopus0 \
    libfontconfig1 libfreetype6 libharfbuzz-icu0 libgstreamer-plugins-base1.0-0 \
    libgstreamer-gl1.0-0 libgstreamer-plugins-bad1.0-0 libjpeg62-turbo \
    libpng16-16 libopenjp2-7 libwebpdemux2 libenchant-2-2 libsecret-1-0 libhyphen0 \
    libx11-6 libxcomposite1 libxdamage1 libxrender1 libxt6 libwayland-server0 \
    libwayland-egl1 libwayland-client0 libmanette-0.2-0 libflite1 \
    libgbm1 libdrm2 libxkbcommon0 libdbus-1-3 libgles2 libwoff1 libx264-160 -y && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

EXPOSE 80

ENTRYPOINT [ "/data/auto-mooc" ]