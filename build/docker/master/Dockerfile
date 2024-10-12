# Copyright 2024 Benjamin Lee <cyan0908@163.com>. All rights reserved.
# Use of this source code is governed by a MIT style
# license that can be found in the LICENSE file.


FROM BASE_IMAGE
LABEL maintainer="cyan0908@163.com"

WORKDIR /opt/hurricane

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
      echo "Asia/Shanghai" > /etc/timezone && \
      mkdir -p /var/log/hurricane

COPY master /opt/hurricane/bin/

ENTRYPOINT ["/opt/hurricane/bin/master"]
CMD ["-c", "/etc/hurricane/master.yaml"]
