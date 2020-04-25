# $FreeBSD: head/sysutils/packer/Makefile 520924 2019-12-26 15:41:09Z lifanov $

PORTNAME=	packer
PORTVERSION=	1.5.1
DISTVERSIONPREFIX=	v
CATEGORIES=	sysutils

MAINTAINER=	lifanov@FreeBSD.org
COMMENT=	Tool for creating machine images

LICENSE=	MPL20
LICENSE_FILE=	${WRKSRC}/LICENSE

PLIST_FILES=	bin/${PORTNAME}

USES=	go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	hashicorp

.include <bsd.port.mk>
