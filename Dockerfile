# SPDX-License-Identifier: UNLICENSED
#
# Copyright (c) 2026, Unikraft GmbH.  All rights reserved.
#
# This software and related documentation ("Unikraft Software") are protected
# under relevant copyright laws.  The information contained herein is
# confidential and proprietary to Unikraft GmbH and/or its licensors.  Without
# the prior written permission of Unikraft GmbH and/or its licensors, any
# reproduction, modification, use or disclosure of Unikraft Software, and
# information contained herein, in whole or in part, shall be strictly
# prohibited.
#
# BY OPENING THIS FILE, RECEIVER HEREBY UNEQUIVOCALLY ACKNOWLEDGES AND AGREES
# THAT THE SOFTWARE/FIRMWARE AND ITS DOCUMENTATIONS ("UNIKRAFT SOFTWARE")
# RECEIVED FROM UNIKRAFT AND/OR ITS REPRESENTATIVES ARE PROVIDED TO RECEIVER ON
# AN "AS-IS" BASIS ONLY.  UNIKRAFT EXPRESSLY DISCLAIMS ANY AND ALL WARRANTIES,
# EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE IMPLIED WARRANTIES OF
# MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE OR NONINFRINGEMENT.
# NEITHER DOES UNIKRAFT PROVIDE ANY WARRANTY WHATSOEVER WITH RESPECT TO THE
# SOFTWARE OF ANY THIRD PARTY WHICH MAY BE USED BY, INCORPORATED IN, OR
# SUPPLIED WITH THE UNIKRAFT SOFTWARE, AND RECEIVER AGREES TO LOOK ONLY TO SUCH
# THIRD PARTY FOR ANY WARRANTY CLAIM RELATING THERETO.  RECEIVER EXPRESSLY
# ACKNOWLEDGES THAT IT IS RECEIVER'S SOLE RESPONSIBILITY TO OBTAIN FROM ANY
# THIRD PARTY ALL PROPER LICENSES CONTAINED IN UNIKRAFT SOFTWARE.  UNIKRAFT
# SHALL ALSO NOT BE RESPONSIBLE FOR ANY UNIKRAFT SOFTWARE RELEASES MADE TO
# RECEIVER'S SPECIFICATION OR TO CONFORM TO A PARTICULAR STANDARD OR OPEN
# FORUM.  RECEIVER'S SOLE AND EXCLUSIVE REMEDY AND UNIKRAFT'S ENTIRE AND
# CUMULATIVE LIABILITY WITH RESPECT TO THE UNIKRAFT SOFTWARE RELEASED HEREUNDER
# WILL BE, AT UNIKRAFT'S OPTION, TO REVISE OR REPLACE THE UNIKRAFT SOFTWARE AT
# ISSUE, OR REFUND ANY SOFTWARE LICENSE FEES OR SERVICE CHARGE PAID BY RECEIVER
# TO UNIKRAFT FOR SUCH UNIKRAFT SOFTWARE AT ISSUE.

FROM debian:bookworm AS user

# Create an unprivileged user for production runtime
ENV USER=k8s-operator
ENV UID=10001
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --no-create-home \
    --shell "/sbin/nologin" \
    --uid "${UID}" \
    "${USER}"

FROM alpine:3.17.2 as ca-certificates

RUN apk add -U --no-cache ca-certificates

FROM debian:bullseye AS prod

# Import the user and group files from the build
COPY --from=user /etc/passwd /etc/passwd
COPY --from=user /etc/group /etc/group

# Copy over binary and license
COPY k8s-operator /usr/local/bin/k8s-operator
COPY LICENSE.txt /

# Copy CA certificates
COPY --from=ca-certificates /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

LABEL org.opencontainers.image.licenses=UNLICENSED

USER k8s-operator:k8s-operator

ENTRYPOINT [ "k8s-operator" ]
