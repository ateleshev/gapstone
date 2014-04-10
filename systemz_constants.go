package gapstone

// #cgo pkg-config: capstone
// #include <stdlib.h>
// #include <capstone.h>
import "C"

const (
	SYSZ_CC_INVALID = C.SYSZ_CC_INVALID
	SYSZ_CC_O       = C.SYSZ_CC_O
	SYSZ_CC_H       = C.SYSZ_CC_H
	SYSZ_CC_NLE     = C.SYSZ_CC_NLE
	SYSZ_CC_L       = C.SYSZ_CC_L
	SYSZ_CC_NHE     = C.SYSZ_CC_NHE
	SYSZ_CC_LH      = C.SYSZ_CC_LH
	SYSZ_CC_NE      = C.SYSZ_CC_NE
	SYSZ_CC_E       = C.SYSZ_CC_E
	SYSZ_CC_NLH     = C.SYSZ_CC_NLH
	SYSZ_CC_HE      = C.SYSZ_CC_HE
	SYSZ_CC_NL      = C.SYSZ_CC_NL
	SYSZ_CC_LE      = C.SYSZ_CC_LE
	SYSZ_CC_NH      = C.SYSZ_CC_NH
	SYSZ_CC_NO      = C.SYSZ_CC_NO
)
const (
	// Operands
	SYSZ_OP_INVALID = C.SYSZ_OP_INVALID
	SYSZ_OP_REG     = C.SYSZ_OP_REG
	SYSZ_OP_ACREG   = C.SYSZ_OP_ACREG
	SYSZ_OP_IMM     = C.SYSZ_OP_IMM
	SYSZ_OP_MEM     = C.SYSZ_OP_MEM
)

const (
	SYSZ_REG_INVALID = C.SYSZ_REG_INVALID
	SYSZ_REG_0       = C.SYSZ_REG_0
	SYSZ_REG_1       = C.SYSZ_REG_1
	SYSZ_REG_2       = C.SYSZ_REG_2
	SYSZ_REG_3       = C.SYSZ_REG_3
	SYSZ_REG_4       = C.SYSZ_REG_4
	SYSZ_REG_5       = C.SYSZ_REG_5
	SYSZ_REG_6       = C.SYSZ_REG_6
	SYSZ_REG_7       = C.SYSZ_REG_7
	SYSZ_REG_8       = C.SYSZ_REG_8
	SYSZ_REG_9       = C.SYSZ_REG_9
	SYSZ_REG_10      = C.SYSZ_REG_10
	SYSZ_REG_11      = C.SYSZ_REG_11
	SYSZ_REG_12      = C.SYSZ_REG_12
	SYSZ_REG_13      = C.SYSZ_REG_13
	SYSZ_REG_14      = C.SYSZ_REG_14
	SYSZ_REG_15      = C.SYSZ_REG_15
	SYSZ_REG_CC      = C.SYSZ_REG_CC
	SYSZ_REG_F0      = C.SYSZ_REG_F0
	SYSZ_REG_F1      = C.SYSZ_REG_F1
	SYSZ_REG_F2      = C.SYSZ_REG_F2
	SYSZ_REG_F3      = C.SYSZ_REG_F3
	SYSZ_REG_F4      = C.SYSZ_REG_F4
	SYSZ_REG_F5      = C.SYSZ_REG_F5
	SYSZ_REG_F6      = C.SYSZ_REG_F6
	SYSZ_REG_F7      = C.SYSZ_REG_F7
	SYSZ_REG_F8      = C.SYSZ_REG_F8
	SYSZ_REG_F9      = C.SYSZ_REG_F9
	SYSZ_REG_F10     = C.SYSZ_REG_F10
	SYSZ_REG_F11     = C.SYSZ_REG_F11
	SYSZ_REG_F12     = C.SYSZ_REG_F12
	SYSZ_REG_F13     = C.SYSZ_REG_F13
	SYSZ_REG_F14     = C.SYSZ_REG_F14
	SYSZ_REG_F15     = C.SYSZ_REG_F15
	SYSZ_REG_R0L     = C.SYSZ_REG_R0L
	SYSZ_REG_MAX     = C.SYSZ_REG_MAX
)

const (
	// Instructions
	SYSZ_INS_INVALID  = C.SYSZ_INS_INVALID
	SYSZ_INS_A        = C.SYSZ_INS_A
	SYSZ_INS_ADB      = C.SYSZ_INS_ADB
	SYSZ_INS_ADBR     = C.SYSZ_INS_ADBR
	SYSZ_INS_AEB      = C.SYSZ_INS_AEB
	SYSZ_INS_AEBR     = C.SYSZ_INS_AEBR
	SYSZ_INS_AFI      = C.SYSZ_INS_AFI
	SYSZ_INS_AG       = C.SYSZ_INS_AG
	SYSZ_INS_AGF      = C.SYSZ_INS_AGF
	SYSZ_INS_AGFI     = C.SYSZ_INS_AGFI
	SYSZ_INS_AGFR     = C.SYSZ_INS_AGFR
	SYSZ_INS_AGHI     = C.SYSZ_INS_AGHI
	SYSZ_INS_AGHIK    = C.SYSZ_INS_AGHIK
	SYSZ_INS_AGR      = C.SYSZ_INS_AGR
	SYSZ_INS_AGRK     = C.SYSZ_INS_AGRK
	SYSZ_INS_AGSI     = C.SYSZ_INS_AGSI
	SYSZ_INS_AH       = C.SYSZ_INS_AH
	SYSZ_INS_AHI      = C.SYSZ_INS_AHI
	SYSZ_INS_AHIK     = C.SYSZ_INS_AHIK
	SYSZ_INS_AHY      = C.SYSZ_INS_AHY
	SYSZ_INS_AIH      = C.SYSZ_INS_AIH
	SYSZ_INS_AL       = C.SYSZ_INS_AL
	SYSZ_INS_ALC      = C.SYSZ_INS_ALC
	SYSZ_INS_ALCG     = C.SYSZ_INS_ALCG
	SYSZ_INS_ALCGR    = C.SYSZ_INS_ALCGR
	SYSZ_INS_ALCR     = C.SYSZ_INS_ALCR
	SYSZ_INS_ALFI     = C.SYSZ_INS_ALFI
	SYSZ_INS_ALG      = C.SYSZ_INS_ALG
	SYSZ_INS_ALGF     = C.SYSZ_INS_ALGF
	SYSZ_INS_ALGFI    = C.SYSZ_INS_ALGFI
	SYSZ_INS_ALGFR    = C.SYSZ_INS_ALGFR
	SYSZ_INS_ALGHSIK  = C.SYSZ_INS_ALGHSIK
	SYSZ_INS_ALGR     = C.SYSZ_INS_ALGR
	SYSZ_INS_ALGRK    = C.SYSZ_INS_ALGRK
	SYSZ_INS_ALHSIK   = C.SYSZ_INS_ALHSIK
	SYSZ_INS_ALR      = C.SYSZ_INS_ALR
	SYSZ_INS_ALRK     = C.SYSZ_INS_ALRK
	SYSZ_INS_ALY      = C.SYSZ_INS_ALY
	SYSZ_INS_AR       = C.SYSZ_INS_AR
	SYSZ_INS_ARK      = C.SYSZ_INS_ARK
	SYSZ_INS_ASI      = C.SYSZ_INS_ASI
	SYSZ_INS_AXBR     = C.SYSZ_INS_AXBR
	SYSZ_INS_AY       = C.SYSZ_INS_AY
	SYSZ_INS_BCR      = C.SYSZ_INS_BCR
	SYSZ_INS_BRC      = C.SYSZ_INS_BRC
	SYSZ_INS_BRCL     = C.SYSZ_INS_BRCL
	SYSZ_INS_CGIJ     = C.SYSZ_INS_CGIJ
	SYSZ_INS_CGRJ     = C.SYSZ_INS_CGRJ
	SYSZ_INS_CIJ      = C.SYSZ_INS_CIJ
	SYSZ_INS_CLGIJ    = C.SYSZ_INS_CLGIJ
	SYSZ_INS_CLGRJ    = C.SYSZ_INS_CLGRJ
	SYSZ_INS_CLIJ     = C.SYSZ_INS_CLIJ
	SYSZ_INS_CLRJ     = C.SYSZ_INS_CLRJ
	SYSZ_INS_CRJ      = C.SYSZ_INS_CRJ
	SYSZ_INS_BER      = C.SYSZ_INS_BER
	SYSZ_INS_JE       = C.SYSZ_INS_JE
	SYSZ_INS_JGE      = C.SYSZ_INS_JGE
	SYSZ_INS_LOCE     = C.SYSZ_INS_LOCE
	SYSZ_INS_LOCGE    = C.SYSZ_INS_LOCGE
	SYSZ_INS_LOCGRE   = C.SYSZ_INS_LOCGRE
	SYSZ_INS_LOCRE    = C.SYSZ_INS_LOCRE
	SYSZ_INS_STOCE    = C.SYSZ_INS_STOCE
	SYSZ_INS_STOCGE   = C.SYSZ_INS_STOCGE
	SYSZ_INS_BHR      = C.SYSZ_INS_BHR
	SYSZ_INS_BHER     = C.SYSZ_INS_BHER
	SYSZ_INS_JHE      = C.SYSZ_INS_JHE
	SYSZ_INS_JGHE     = C.SYSZ_INS_JGHE
	SYSZ_INS_LOCHE    = C.SYSZ_INS_LOCHE
	SYSZ_INS_LOCGHE   = C.SYSZ_INS_LOCGHE
	SYSZ_INS_LOCGRHE  = C.SYSZ_INS_LOCGRHE
	SYSZ_INS_LOCRHE   = C.SYSZ_INS_LOCRHE
	SYSZ_INS_STOCHE   = C.SYSZ_INS_STOCHE
	SYSZ_INS_STOCGHE  = C.SYSZ_INS_STOCGHE
	SYSZ_INS_JH       = C.SYSZ_INS_JH
	SYSZ_INS_JGH      = C.SYSZ_INS_JGH
	SYSZ_INS_LOCH     = C.SYSZ_INS_LOCH
	SYSZ_INS_LOCGH    = C.SYSZ_INS_LOCGH
	SYSZ_INS_LOCGRH   = C.SYSZ_INS_LOCGRH
	SYSZ_INS_LOCRH    = C.SYSZ_INS_LOCRH
	SYSZ_INS_STOCH    = C.SYSZ_INS_STOCH
	SYSZ_INS_STOCGH   = C.SYSZ_INS_STOCGH
	SYSZ_INS_CGIJNLH  = C.SYSZ_INS_CGIJNLH
	SYSZ_INS_CGRJNLH  = C.SYSZ_INS_CGRJNLH
	SYSZ_INS_CIJNLH   = C.SYSZ_INS_CIJNLH
	SYSZ_INS_CLGIJNLH = C.SYSZ_INS_CLGIJNLH
	SYSZ_INS_CLGRJNLH = C.SYSZ_INS_CLGRJNLH
	SYSZ_INS_CLIJNLH  = C.SYSZ_INS_CLIJNLH
	SYSZ_INS_CLRJNLH  = C.SYSZ_INS_CLRJNLH
	SYSZ_INS_CRJNLH   = C.SYSZ_INS_CRJNLH
	SYSZ_INS_CGIJE    = C.SYSZ_INS_CGIJE
	SYSZ_INS_CGRJE    = C.SYSZ_INS_CGRJE
	SYSZ_INS_CIJE     = C.SYSZ_INS_CIJE
	SYSZ_INS_CLGIJE   = C.SYSZ_INS_CLGIJE
	SYSZ_INS_CLGRJE   = C.SYSZ_INS_CLGRJE
	SYSZ_INS_CLIJE    = C.SYSZ_INS_CLIJE
	SYSZ_INS_CLRJE    = C.SYSZ_INS_CLRJE
	SYSZ_INS_CRJE     = C.SYSZ_INS_CRJE
	SYSZ_INS_CGIJNLE  = C.SYSZ_INS_CGIJNLE
	SYSZ_INS_CGRJNLE  = C.SYSZ_INS_CGRJNLE
	SYSZ_INS_CIJNLE   = C.SYSZ_INS_CIJNLE
	SYSZ_INS_CLGIJNLE = C.SYSZ_INS_CLGIJNLE
	SYSZ_INS_CLGRJNLE = C.SYSZ_INS_CLGRJNLE
	SYSZ_INS_CLIJNLE  = C.SYSZ_INS_CLIJNLE
	SYSZ_INS_CLRJNLE  = C.SYSZ_INS_CLRJNLE
	SYSZ_INS_CRJNLE   = C.SYSZ_INS_CRJNLE
	SYSZ_INS_CGIJH    = C.SYSZ_INS_CGIJH
	SYSZ_INS_CGRJH    = C.SYSZ_INS_CGRJH
	SYSZ_INS_CIJH     = C.SYSZ_INS_CIJH
	SYSZ_INS_CLGIJH   = C.SYSZ_INS_CLGIJH
	SYSZ_INS_CLGRJH   = C.SYSZ_INS_CLGRJH
	SYSZ_INS_CLIJH    = C.SYSZ_INS_CLIJH
	SYSZ_INS_CLRJH    = C.SYSZ_INS_CLRJH
	SYSZ_INS_CRJH     = C.SYSZ_INS_CRJH
	SYSZ_INS_CGIJNL   = C.SYSZ_INS_CGIJNL
	SYSZ_INS_CGRJNL   = C.SYSZ_INS_CGRJNL
	SYSZ_INS_CIJNL    = C.SYSZ_INS_CIJNL
	SYSZ_INS_CLGIJNL  = C.SYSZ_INS_CLGIJNL
	SYSZ_INS_CLGRJNL  = C.SYSZ_INS_CLGRJNL
	SYSZ_INS_CLIJNL   = C.SYSZ_INS_CLIJNL
	SYSZ_INS_CLRJNL   = C.SYSZ_INS_CLRJNL
	SYSZ_INS_CRJNL    = C.SYSZ_INS_CRJNL
	SYSZ_INS_CGIJHE   = C.SYSZ_INS_CGIJHE
	SYSZ_INS_CGRJHE   = C.SYSZ_INS_CGRJHE
	SYSZ_INS_CIJHE    = C.SYSZ_INS_CIJHE
	SYSZ_INS_CLGIJHE  = C.SYSZ_INS_CLGIJHE
	SYSZ_INS_CLGRJHE  = C.SYSZ_INS_CLGRJHE
	SYSZ_INS_CLIJHE   = C.SYSZ_INS_CLIJHE
	SYSZ_INS_CLRJHE   = C.SYSZ_INS_CLRJHE
	SYSZ_INS_CRJHE    = C.SYSZ_INS_CRJHE
	SYSZ_INS_CGIJNHE  = C.SYSZ_INS_CGIJNHE
	SYSZ_INS_CGRJNHE  = C.SYSZ_INS_CGRJNHE
	SYSZ_INS_CIJNHE   = C.SYSZ_INS_CIJNHE
	SYSZ_INS_CLGIJNHE = C.SYSZ_INS_CLGIJNHE
	SYSZ_INS_CLGRJNHE = C.SYSZ_INS_CLGRJNHE
	SYSZ_INS_CLIJNHE  = C.SYSZ_INS_CLIJNHE
	SYSZ_INS_CLRJNHE  = C.SYSZ_INS_CLRJNHE
	SYSZ_INS_CRJNHE   = C.SYSZ_INS_CRJNHE
	SYSZ_INS_CGIJL    = C.SYSZ_INS_CGIJL
	SYSZ_INS_CGRJL    = C.SYSZ_INS_CGRJL
	SYSZ_INS_CIJL     = C.SYSZ_INS_CIJL
	SYSZ_INS_CLGIJL   = C.SYSZ_INS_CLGIJL
	SYSZ_INS_CLGRJL   = C.SYSZ_INS_CLGRJL
	SYSZ_INS_CLIJL    = C.SYSZ_INS_CLIJL
	SYSZ_INS_CLRJL    = C.SYSZ_INS_CLRJL
	SYSZ_INS_CRJL     = C.SYSZ_INS_CRJL
	SYSZ_INS_CGIJNH   = C.SYSZ_INS_CGIJNH
	SYSZ_INS_CGRJNH   = C.SYSZ_INS_CGRJNH
	SYSZ_INS_CIJNH    = C.SYSZ_INS_CIJNH
	SYSZ_INS_CLGIJNH  = C.SYSZ_INS_CLGIJNH
	SYSZ_INS_CLGRJNH  = C.SYSZ_INS_CLGRJNH
	SYSZ_INS_CLIJNH   = C.SYSZ_INS_CLIJNH
	SYSZ_INS_CLRJNH   = C.SYSZ_INS_CLRJNH
	SYSZ_INS_CRJNH    = C.SYSZ_INS_CRJNH
	SYSZ_INS_CGIJLE   = C.SYSZ_INS_CGIJLE
	SYSZ_INS_CGRJLE   = C.SYSZ_INS_CGRJLE
	SYSZ_INS_CIJLE    = C.SYSZ_INS_CIJLE
	SYSZ_INS_CLGIJLE  = C.SYSZ_INS_CLGIJLE
	SYSZ_INS_CLGRJLE  = C.SYSZ_INS_CLGRJLE
	SYSZ_INS_CLIJLE   = C.SYSZ_INS_CLIJLE
	SYSZ_INS_CLRJLE   = C.SYSZ_INS_CLRJLE
	SYSZ_INS_CRJLE    = C.SYSZ_INS_CRJLE
	SYSZ_INS_CGIJNE   = C.SYSZ_INS_CGIJNE
	SYSZ_INS_CGRJNE   = C.SYSZ_INS_CGRJNE
	SYSZ_INS_CIJNE    = C.SYSZ_INS_CIJNE
	SYSZ_INS_CLGIJNE  = C.SYSZ_INS_CLGIJNE
	SYSZ_INS_CLGRJNE  = C.SYSZ_INS_CLGRJNE
	SYSZ_INS_CLIJNE   = C.SYSZ_INS_CLIJNE
	SYSZ_INS_CLRJNE   = C.SYSZ_INS_CLRJNE
	SYSZ_INS_CRJNE    = C.SYSZ_INS_CRJNE
	SYSZ_INS_CGIJLH   = C.SYSZ_INS_CGIJLH
	SYSZ_INS_CGRJLH   = C.SYSZ_INS_CGRJLH
	SYSZ_INS_CIJLH    = C.SYSZ_INS_CIJLH
	SYSZ_INS_CLGIJLH  = C.SYSZ_INS_CLGIJLH
	SYSZ_INS_CLGRJLH  = C.SYSZ_INS_CLGRJLH
	SYSZ_INS_CLIJLH   = C.SYSZ_INS_CLIJLH
	SYSZ_INS_CLRJLH   = C.SYSZ_INS_CLRJLH
	SYSZ_INS_CRJLH    = C.SYSZ_INS_CRJLH
	SYSZ_INS_BLR      = C.SYSZ_INS_BLR
	SYSZ_INS_BLER     = C.SYSZ_INS_BLER
	SYSZ_INS_JLE      = C.SYSZ_INS_JLE
	SYSZ_INS_JGLE     = C.SYSZ_INS_JGLE
	SYSZ_INS_LOCLE    = C.SYSZ_INS_LOCLE
	SYSZ_INS_LOCGLE   = C.SYSZ_INS_LOCGLE
	SYSZ_INS_LOCGRLE  = C.SYSZ_INS_LOCGRLE
	SYSZ_INS_LOCRLE   = C.SYSZ_INS_LOCRLE
	SYSZ_INS_STOCLE   = C.SYSZ_INS_STOCLE
	SYSZ_INS_STOCGLE  = C.SYSZ_INS_STOCGLE
	SYSZ_INS_BLHR     = C.SYSZ_INS_BLHR
	SYSZ_INS_JLH      = C.SYSZ_INS_JLH
	SYSZ_INS_JGLH     = C.SYSZ_INS_JGLH
	SYSZ_INS_LOCLH    = C.SYSZ_INS_LOCLH
	SYSZ_INS_LOCGLH   = C.SYSZ_INS_LOCGLH
	SYSZ_INS_LOCGRLH  = C.SYSZ_INS_LOCGRLH
	SYSZ_INS_LOCRLH   = C.SYSZ_INS_LOCRLH
	SYSZ_INS_STOCLH   = C.SYSZ_INS_STOCLH
	SYSZ_INS_STOCGLH  = C.SYSZ_INS_STOCGLH
	SYSZ_INS_JL       = C.SYSZ_INS_JL
	SYSZ_INS_JGL      = C.SYSZ_INS_JGL
	SYSZ_INS_LOCL     = C.SYSZ_INS_LOCL
	SYSZ_INS_LOCGL    = C.SYSZ_INS_LOCGL
	SYSZ_INS_LOCGRL   = C.SYSZ_INS_LOCGRL
	SYSZ_INS_LOCRL    = C.SYSZ_INS_LOCRL
	SYSZ_INS_LOC      = C.SYSZ_INS_LOC
	SYSZ_INS_LOCG     = C.SYSZ_INS_LOCG
	SYSZ_INS_LOCGR    = C.SYSZ_INS_LOCGR
	SYSZ_INS_LOCR     = C.SYSZ_INS_LOCR
	SYSZ_INS_STOCL    = C.SYSZ_INS_STOCL
	SYSZ_INS_STOCGL   = C.SYSZ_INS_STOCGL
	SYSZ_INS_BNER     = C.SYSZ_INS_BNER
	SYSZ_INS_JNE      = C.SYSZ_INS_JNE
	SYSZ_INS_JGNE     = C.SYSZ_INS_JGNE
	SYSZ_INS_LOCNE    = C.SYSZ_INS_LOCNE
	SYSZ_INS_LOCGNE   = C.SYSZ_INS_LOCGNE
	SYSZ_INS_LOCGRNE  = C.SYSZ_INS_LOCGRNE
	SYSZ_INS_LOCRNE   = C.SYSZ_INS_LOCRNE
	SYSZ_INS_STOCNE   = C.SYSZ_INS_STOCNE
	SYSZ_INS_STOCGNE  = C.SYSZ_INS_STOCGNE
	SYSZ_INS_BNHR     = C.SYSZ_INS_BNHR
	SYSZ_INS_BNHER    = C.SYSZ_INS_BNHER
	SYSZ_INS_JNHE     = C.SYSZ_INS_JNHE
	SYSZ_INS_JGNHE    = C.SYSZ_INS_JGNHE
	SYSZ_INS_LOCNHE   = C.SYSZ_INS_LOCNHE
	SYSZ_INS_LOCGNHE  = C.SYSZ_INS_LOCGNHE
	SYSZ_INS_LOCGRNHE = C.SYSZ_INS_LOCGRNHE
	SYSZ_INS_LOCRNHE  = C.SYSZ_INS_LOCRNHE
	SYSZ_INS_STOCNHE  = C.SYSZ_INS_STOCNHE
	SYSZ_INS_STOCGNHE = C.SYSZ_INS_STOCGNHE
	SYSZ_INS_JNH      = C.SYSZ_INS_JNH
	SYSZ_INS_JGNH     = C.SYSZ_INS_JGNH
	SYSZ_INS_LOCNH    = C.SYSZ_INS_LOCNH
	SYSZ_INS_LOCGNH   = C.SYSZ_INS_LOCGNH
	SYSZ_INS_LOCGRNH  = C.SYSZ_INS_LOCGRNH
	SYSZ_INS_LOCRNH   = C.SYSZ_INS_LOCRNH
	SYSZ_INS_STOCNH   = C.SYSZ_INS_STOCNH
	SYSZ_INS_STOCGNH  = C.SYSZ_INS_STOCGNH
	SYSZ_INS_BNLR     = C.SYSZ_INS_BNLR
	SYSZ_INS_BNLER    = C.SYSZ_INS_BNLER
	SYSZ_INS_JNLE     = C.SYSZ_INS_JNLE
	SYSZ_INS_JGNLE    = C.SYSZ_INS_JGNLE
	SYSZ_INS_LOCNLE   = C.SYSZ_INS_LOCNLE
	SYSZ_INS_LOCGNLE  = C.SYSZ_INS_LOCGNLE
	SYSZ_INS_LOCGRNLE = C.SYSZ_INS_LOCGRNLE
	SYSZ_INS_LOCRNLE  = C.SYSZ_INS_LOCRNLE
	SYSZ_INS_STOCNLE  = C.SYSZ_INS_STOCNLE
	SYSZ_INS_STOCGNLE = C.SYSZ_INS_STOCGNLE
	SYSZ_INS_BNLHR    = C.SYSZ_INS_BNLHR
	SYSZ_INS_JNLH     = C.SYSZ_INS_JNLH
	SYSZ_INS_JGNLH    = C.SYSZ_INS_JGNLH
	SYSZ_INS_LOCNLH   = C.SYSZ_INS_LOCNLH
	SYSZ_INS_LOCGNLH  = C.SYSZ_INS_LOCGNLH
	SYSZ_INS_LOCGRNLH = C.SYSZ_INS_LOCGRNLH
	SYSZ_INS_LOCRNLH  = C.SYSZ_INS_LOCRNLH
	SYSZ_INS_STOCNLH  = C.SYSZ_INS_STOCNLH
	SYSZ_INS_STOCGNLH = C.SYSZ_INS_STOCGNLH
	SYSZ_INS_JNL      = C.SYSZ_INS_JNL
	SYSZ_INS_JGNL     = C.SYSZ_INS_JGNL
	SYSZ_INS_LOCNL    = C.SYSZ_INS_LOCNL
	SYSZ_INS_LOCGNL   = C.SYSZ_INS_LOCGNL
	SYSZ_INS_LOCGRNL  = C.SYSZ_INS_LOCGRNL
	SYSZ_INS_LOCRNL   = C.SYSZ_INS_LOCRNL
	SYSZ_INS_STOCNL   = C.SYSZ_INS_STOCNL
	SYSZ_INS_STOCGNL  = C.SYSZ_INS_STOCGNL
	SYSZ_INS_BNOR     = C.SYSZ_INS_BNOR
	SYSZ_INS_JNO      = C.SYSZ_INS_JNO
	SYSZ_INS_JGNO     = C.SYSZ_INS_JGNO
	SYSZ_INS_LOCNO    = C.SYSZ_INS_LOCNO
	SYSZ_INS_LOCGNO   = C.SYSZ_INS_LOCGNO
	SYSZ_INS_LOCGRNO  = C.SYSZ_INS_LOCGRNO
	SYSZ_INS_LOCRNO   = C.SYSZ_INS_LOCRNO
	SYSZ_INS_STOCNO   = C.SYSZ_INS_STOCNO
	SYSZ_INS_STOCGNO  = C.SYSZ_INS_STOCGNO
	SYSZ_INS_BOR      = C.SYSZ_INS_BOR
	SYSZ_INS_JO       = C.SYSZ_INS_JO
	SYSZ_INS_JGO      = C.SYSZ_INS_JGO
	SYSZ_INS_LOCO     = C.SYSZ_INS_LOCO
	SYSZ_INS_LOCGO    = C.SYSZ_INS_LOCGO
	SYSZ_INS_LOCGRO   = C.SYSZ_INS_LOCGRO
	SYSZ_INS_LOCRO    = C.SYSZ_INS_LOCRO
	SYSZ_INS_STOCO    = C.SYSZ_INS_STOCO
	SYSZ_INS_STOCGO   = C.SYSZ_INS_STOCGO
	SYSZ_INS_STOC     = C.SYSZ_INS_STOC
	SYSZ_INS_STOCG    = C.SYSZ_INS_STOCG
	SYSZ_INS_BASR     = C.SYSZ_INS_BASR
	SYSZ_INS_BR       = C.SYSZ_INS_BR
	SYSZ_INS_BRAS     = C.SYSZ_INS_BRAS
	SYSZ_INS_BRASL    = C.SYSZ_INS_BRASL
	SYSZ_INS_J        = C.SYSZ_INS_J
	SYSZ_INS_JG       = C.SYSZ_INS_JG
	SYSZ_INS_BRCT     = C.SYSZ_INS_BRCT
	SYSZ_INS_BRCTG    = C.SYSZ_INS_BRCTG
	SYSZ_INS_C        = C.SYSZ_INS_C
	SYSZ_INS_CDB      = C.SYSZ_INS_CDB
	SYSZ_INS_CDBR     = C.SYSZ_INS_CDBR
	SYSZ_INS_CDFBR    = C.SYSZ_INS_CDFBR
	SYSZ_INS_CDGBR    = C.SYSZ_INS_CDGBR
	SYSZ_INS_CDLFBR   = C.SYSZ_INS_CDLFBR
	SYSZ_INS_CDLGBR   = C.SYSZ_INS_CDLGBR
	SYSZ_INS_CEB      = C.SYSZ_INS_CEB
	SYSZ_INS_CEBR     = C.SYSZ_INS_CEBR
	SYSZ_INS_CEFBR    = C.SYSZ_INS_CEFBR
	SYSZ_INS_CEGBR    = C.SYSZ_INS_CEGBR
	SYSZ_INS_CELFBR   = C.SYSZ_INS_CELFBR
	SYSZ_INS_CELGBR   = C.SYSZ_INS_CELGBR
	SYSZ_INS_CFDBR    = C.SYSZ_INS_CFDBR
	SYSZ_INS_CFEBR    = C.SYSZ_INS_CFEBR
	SYSZ_INS_CFI      = C.SYSZ_INS_CFI
	SYSZ_INS_CFXBR    = C.SYSZ_INS_CFXBR
	SYSZ_INS_CG       = C.SYSZ_INS_CG
	SYSZ_INS_CGDBR    = C.SYSZ_INS_CGDBR
	SYSZ_INS_CGEBR    = C.SYSZ_INS_CGEBR
	SYSZ_INS_CGF      = C.SYSZ_INS_CGF
	SYSZ_INS_CGFI     = C.SYSZ_INS_CGFI
	SYSZ_INS_CGFR     = C.SYSZ_INS_CGFR
	SYSZ_INS_CGFRL    = C.SYSZ_INS_CGFRL
	SYSZ_INS_CGH      = C.SYSZ_INS_CGH
	SYSZ_INS_CGHI     = C.SYSZ_INS_CGHI
	SYSZ_INS_CGHRL    = C.SYSZ_INS_CGHRL
	SYSZ_INS_CGHSI    = C.SYSZ_INS_CGHSI
	SYSZ_INS_CGR      = C.SYSZ_INS_CGR
	SYSZ_INS_CGRL     = C.SYSZ_INS_CGRL
	SYSZ_INS_CGXBR    = C.SYSZ_INS_CGXBR
	SYSZ_INS_CH       = C.SYSZ_INS_CH
	SYSZ_INS_CHF      = C.SYSZ_INS_CHF
	SYSZ_INS_CHHSI    = C.SYSZ_INS_CHHSI
	SYSZ_INS_CHI      = C.SYSZ_INS_CHI
	SYSZ_INS_CHRL     = C.SYSZ_INS_CHRL
	SYSZ_INS_CHSI     = C.SYSZ_INS_CHSI
	SYSZ_INS_CHY      = C.SYSZ_INS_CHY
	SYSZ_INS_CIH      = C.SYSZ_INS_CIH
	SYSZ_INS_CL       = C.SYSZ_INS_CL
	SYSZ_INS_CLC      = C.SYSZ_INS_CLC
	SYSZ_INS_CLFDBR   = C.SYSZ_INS_CLFDBR
	SYSZ_INS_CLFEBR   = C.SYSZ_INS_CLFEBR
	SYSZ_INS_CLFHSI   = C.SYSZ_INS_CLFHSI
	SYSZ_INS_CLFI     = C.SYSZ_INS_CLFI
	SYSZ_INS_CLFXBR   = C.SYSZ_INS_CLFXBR
	SYSZ_INS_CLG      = C.SYSZ_INS_CLG
	SYSZ_INS_CLGDBR   = C.SYSZ_INS_CLGDBR
	SYSZ_INS_CLGEBR   = C.SYSZ_INS_CLGEBR
	SYSZ_INS_CLGF     = C.SYSZ_INS_CLGF
	SYSZ_INS_CLGFI    = C.SYSZ_INS_CLGFI
	SYSZ_INS_CLGFR    = C.SYSZ_INS_CLGFR
	SYSZ_INS_CLGFRL   = C.SYSZ_INS_CLGFRL
	SYSZ_INS_CLGHRL   = C.SYSZ_INS_CLGHRL
	SYSZ_INS_CLGHSI   = C.SYSZ_INS_CLGHSI
	SYSZ_INS_CLGR     = C.SYSZ_INS_CLGR
	SYSZ_INS_CLGRL    = C.SYSZ_INS_CLGRL
	SYSZ_INS_CLGXBR   = C.SYSZ_INS_CLGXBR
	SYSZ_INS_CLHF     = C.SYSZ_INS_CLHF
	SYSZ_INS_CLHHSI   = C.SYSZ_INS_CLHHSI
	SYSZ_INS_CLHRL    = C.SYSZ_INS_CLHRL
	SYSZ_INS_CLI      = C.SYSZ_INS_CLI
	SYSZ_INS_CLIH     = C.SYSZ_INS_CLIH
	SYSZ_INS_CLIY     = C.SYSZ_INS_CLIY
	SYSZ_INS_CLR      = C.SYSZ_INS_CLR
	SYSZ_INS_CLRL     = C.SYSZ_INS_CLRL
	SYSZ_INS_CLST     = C.SYSZ_INS_CLST
	SYSZ_INS_CLY      = C.SYSZ_INS_CLY
	SYSZ_INS_CPSDR    = C.SYSZ_INS_CPSDR
	SYSZ_INS_CR       = C.SYSZ_INS_CR
	SYSZ_INS_CRL      = C.SYSZ_INS_CRL
	SYSZ_INS_CS       = C.SYSZ_INS_CS
	SYSZ_INS_CSG      = C.SYSZ_INS_CSG
	SYSZ_INS_CSY      = C.SYSZ_INS_CSY
	SYSZ_INS_CXBR     = C.SYSZ_INS_CXBR
	SYSZ_INS_CXFBR    = C.SYSZ_INS_CXFBR
	SYSZ_INS_CXGBR    = C.SYSZ_INS_CXGBR
	SYSZ_INS_CXLFBR   = C.SYSZ_INS_CXLFBR
	SYSZ_INS_CXLGBR   = C.SYSZ_INS_CXLGBR
	SYSZ_INS_CY       = C.SYSZ_INS_CY
	SYSZ_INS_DDB      = C.SYSZ_INS_DDB
	SYSZ_INS_DDBR     = C.SYSZ_INS_DDBR
	SYSZ_INS_DEB      = C.SYSZ_INS_DEB
	SYSZ_INS_DEBR     = C.SYSZ_INS_DEBR
	SYSZ_INS_DL       = C.SYSZ_INS_DL
	SYSZ_INS_DLG      = C.SYSZ_INS_DLG
	SYSZ_INS_DLGR     = C.SYSZ_INS_DLGR
	SYSZ_INS_DLR      = C.SYSZ_INS_DLR
	SYSZ_INS_DSG      = C.SYSZ_INS_DSG
	SYSZ_INS_DSGF     = C.SYSZ_INS_DSGF
	SYSZ_INS_DSGFR    = C.SYSZ_INS_DSGFR
	SYSZ_INS_DSGR     = C.SYSZ_INS_DSGR
	SYSZ_INS_DXBR     = C.SYSZ_INS_DXBR
	SYSZ_INS_EAR      = C.SYSZ_INS_EAR
	SYSZ_INS_FIDBR    = C.SYSZ_INS_FIDBR
	SYSZ_INS_FIDBRA   = C.SYSZ_INS_FIDBRA
	SYSZ_INS_FIEBR    = C.SYSZ_INS_FIEBR
	SYSZ_INS_FIEBRA   = C.SYSZ_INS_FIEBRA
	SYSZ_INS_FIXBR    = C.SYSZ_INS_FIXBR
	SYSZ_INS_FIXBRA   = C.SYSZ_INS_FIXBRA
	SYSZ_INS_FLOGR    = C.SYSZ_INS_FLOGR
	SYSZ_INS_IC       = C.SYSZ_INS_IC
	SYSZ_INS_ICY      = C.SYSZ_INS_ICY
	SYSZ_INS_IIHF     = C.SYSZ_INS_IIHF
	SYSZ_INS_IIHH     = C.SYSZ_INS_IIHH
	SYSZ_INS_IIHL     = C.SYSZ_INS_IIHL
	SYSZ_INS_IILF     = C.SYSZ_INS_IILF
	SYSZ_INS_IILH     = C.SYSZ_INS_IILH
	SYSZ_INS_IILL     = C.SYSZ_INS_IILL
	SYSZ_INS_IPM      = C.SYSZ_INS_IPM
	SYSZ_INS_L        = C.SYSZ_INS_L
	SYSZ_INS_LA       = C.SYSZ_INS_LA
	SYSZ_INS_LAA      = C.SYSZ_INS_LAA
	SYSZ_INS_LAAG     = C.SYSZ_INS_LAAG
	SYSZ_INS_LAAL     = C.SYSZ_INS_LAAL
	SYSZ_INS_LAALG    = C.SYSZ_INS_LAALG
	SYSZ_INS_LAN      = C.SYSZ_INS_LAN
	SYSZ_INS_LANG     = C.SYSZ_INS_LANG
	SYSZ_INS_LAO      = C.SYSZ_INS_LAO
	SYSZ_INS_LAOG     = C.SYSZ_INS_LAOG
	SYSZ_INS_LARL     = C.SYSZ_INS_LARL
	SYSZ_INS_LAX      = C.SYSZ_INS_LAX
	SYSZ_INS_LAXG     = C.SYSZ_INS_LAXG
	SYSZ_INS_LAY      = C.SYSZ_INS_LAY
	SYSZ_INS_LB       = C.SYSZ_INS_LB
	SYSZ_INS_LBH      = C.SYSZ_INS_LBH
	SYSZ_INS_LBR      = C.SYSZ_INS_LBR
	SYSZ_INS_LCDBR    = C.SYSZ_INS_LCDBR
	SYSZ_INS_LCEBR    = C.SYSZ_INS_LCEBR
	SYSZ_INS_LCGFR    = C.SYSZ_INS_LCGFR
	SYSZ_INS_LCGR     = C.SYSZ_INS_LCGR
	SYSZ_INS_LCR      = C.SYSZ_INS_LCR
	SYSZ_INS_LCXBR    = C.SYSZ_INS_LCXBR
	SYSZ_INS_LD       = C.SYSZ_INS_LD
	SYSZ_INS_LDEB     = C.SYSZ_INS_LDEB
	SYSZ_INS_LDEBR    = C.SYSZ_INS_LDEBR
	SYSZ_INS_LDGR     = C.SYSZ_INS_LDGR
	SYSZ_INS_LDR      = C.SYSZ_INS_LDR
	SYSZ_INS_LDXBR    = C.SYSZ_INS_LDXBR
	SYSZ_INS_LDY      = C.SYSZ_INS_LDY
	SYSZ_INS_LE       = C.SYSZ_INS_LE
	SYSZ_INS_LEDBR    = C.SYSZ_INS_LEDBR
	SYSZ_INS_LER      = C.SYSZ_INS_LER
	SYSZ_INS_LEXBR    = C.SYSZ_INS_LEXBR
	SYSZ_INS_LEY      = C.SYSZ_INS_LEY
	SYSZ_INS_LFH      = C.SYSZ_INS_LFH
	SYSZ_INS_LG       = C.SYSZ_INS_LG
	SYSZ_INS_LGB      = C.SYSZ_INS_LGB
	SYSZ_INS_LGBR     = C.SYSZ_INS_LGBR
	SYSZ_INS_LGDR     = C.SYSZ_INS_LGDR
	SYSZ_INS_LGF      = C.SYSZ_INS_LGF
	SYSZ_INS_LGFI     = C.SYSZ_INS_LGFI
	SYSZ_INS_LGFR     = C.SYSZ_INS_LGFR
	SYSZ_INS_LGFRL    = C.SYSZ_INS_LGFRL
	SYSZ_INS_LGH      = C.SYSZ_INS_LGH
	SYSZ_INS_LGHI     = C.SYSZ_INS_LGHI
	SYSZ_INS_LGHR     = C.SYSZ_INS_LGHR
	SYSZ_INS_LGHRL    = C.SYSZ_INS_LGHRL
	SYSZ_INS_LGR      = C.SYSZ_INS_LGR
	SYSZ_INS_LGRL     = C.SYSZ_INS_LGRL
	SYSZ_INS_LH       = C.SYSZ_INS_LH
	SYSZ_INS_LHH      = C.SYSZ_INS_LHH
	SYSZ_INS_LHI      = C.SYSZ_INS_LHI
	SYSZ_INS_LHR      = C.SYSZ_INS_LHR
	SYSZ_INS_LHRL     = C.SYSZ_INS_LHRL
	SYSZ_INS_LHY      = C.SYSZ_INS_LHY
	SYSZ_INS_LLC      = C.SYSZ_INS_LLC
	SYSZ_INS_LLCH     = C.SYSZ_INS_LLCH
	SYSZ_INS_LLCR     = C.SYSZ_INS_LLCR
	SYSZ_INS_LLGC     = C.SYSZ_INS_LLGC
	SYSZ_INS_LLGCR    = C.SYSZ_INS_LLGCR
	SYSZ_INS_LLGF     = C.SYSZ_INS_LLGF
	SYSZ_INS_LLGFR    = C.SYSZ_INS_LLGFR
	SYSZ_INS_LLGFRL   = C.SYSZ_INS_LLGFRL
	SYSZ_INS_LLGH     = C.SYSZ_INS_LLGH
	SYSZ_INS_LLGHR    = C.SYSZ_INS_LLGHR
	SYSZ_INS_LLGHRL   = C.SYSZ_INS_LLGHRL
	SYSZ_INS_LLH      = C.SYSZ_INS_LLH
	SYSZ_INS_LLHH     = C.SYSZ_INS_LLHH
	SYSZ_INS_LLHR     = C.SYSZ_INS_LLHR
	SYSZ_INS_LLHRL    = C.SYSZ_INS_LLHRL
	SYSZ_INS_LLIHF    = C.SYSZ_INS_LLIHF
	SYSZ_INS_LLIHH    = C.SYSZ_INS_LLIHH
	SYSZ_INS_LLIHL    = C.SYSZ_INS_LLIHL
	SYSZ_INS_LLILF    = C.SYSZ_INS_LLILF
	SYSZ_INS_LLILH    = C.SYSZ_INS_LLILH
	SYSZ_INS_LLILL    = C.SYSZ_INS_LLILL
	SYSZ_INS_LMG      = C.SYSZ_INS_LMG
	SYSZ_INS_LNDBR    = C.SYSZ_INS_LNDBR
	SYSZ_INS_LNEBR    = C.SYSZ_INS_LNEBR
	SYSZ_INS_LNGFR    = C.SYSZ_INS_LNGFR
	SYSZ_INS_LNGR     = C.SYSZ_INS_LNGR
	SYSZ_INS_LNR      = C.SYSZ_INS_LNR
	SYSZ_INS_LNXBR    = C.SYSZ_INS_LNXBR
	SYSZ_INS_LPDBR    = C.SYSZ_INS_LPDBR
	SYSZ_INS_LPEBR    = C.SYSZ_INS_LPEBR
	SYSZ_INS_LPGFR    = C.SYSZ_INS_LPGFR
	SYSZ_INS_LPGR     = C.SYSZ_INS_LPGR
	SYSZ_INS_LPR      = C.SYSZ_INS_LPR
	SYSZ_INS_LPXBR    = C.SYSZ_INS_LPXBR
	SYSZ_INS_LR       = C.SYSZ_INS_LR
	SYSZ_INS_LRL      = C.SYSZ_INS_LRL
	SYSZ_INS_LRV      = C.SYSZ_INS_LRV
	SYSZ_INS_LRVG     = C.SYSZ_INS_LRVG
	SYSZ_INS_LRVGR    = C.SYSZ_INS_LRVGR
	SYSZ_INS_LRVR     = C.SYSZ_INS_LRVR
	SYSZ_INS_LT       = C.SYSZ_INS_LT
	SYSZ_INS_LTDBR    = C.SYSZ_INS_LTDBR
	SYSZ_INS_LTEBR    = C.SYSZ_INS_LTEBR
	SYSZ_INS_LTG      = C.SYSZ_INS_LTG
	SYSZ_INS_LTGF     = C.SYSZ_INS_LTGF
	SYSZ_INS_LTGFR    = C.SYSZ_INS_LTGFR
	SYSZ_INS_LTGR     = C.SYSZ_INS_LTGR
	SYSZ_INS_LTR      = C.SYSZ_INS_LTR
	SYSZ_INS_LTXBR    = C.SYSZ_INS_LTXBR
	SYSZ_INS_LXDB     = C.SYSZ_INS_LXDB
	SYSZ_INS_LXDBR    = C.SYSZ_INS_LXDBR
	SYSZ_INS_LXEB     = C.SYSZ_INS_LXEB
	SYSZ_INS_LXEBR    = C.SYSZ_INS_LXEBR
	SYSZ_INS_LXR      = C.SYSZ_INS_LXR
	SYSZ_INS_LY       = C.SYSZ_INS_LY
	SYSZ_INS_LZDR     = C.SYSZ_INS_LZDR
	SYSZ_INS_LZER     = C.SYSZ_INS_LZER
	SYSZ_INS_LZXR     = C.SYSZ_INS_LZXR
	SYSZ_INS_MADB     = C.SYSZ_INS_MADB
	SYSZ_INS_MADBR    = C.SYSZ_INS_MADBR
	SYSZ_INS_MAEB     = C.SYSZ_INS_MAEB
	SYSZ_INS_MAEBR    = C.SYSZ_INS_MAEBR
	SYSZ_INS_MDB      = C.SYSZ_INS_MDB
	SYSZ_INS_MDBR     = C.SYSZ_INS_MDBR
	SYSZ_INS_MDEB     = C.SYSZ_INS_MDEB
	SYSZ_INS_MDEBR    = C.SYSZ_INS_MDEBR
	SYSZ_INS_MEEB     = C.SYSZ_INS_MEEB
	SYSZ_INS_MEEBR    = C.SYSZ_INS_MEEBR
	SYSZ_INS_MGHI     = C.SYSZ_INS_MGHI
	SYSZ_INS_MH       = C.SYSZ_INS_MH
	SYSZ_INS_MHI      = C.SYSZ_INS_MHI
	SYSZ_INS_MHY      = C.SYSZ_INS_MHY
	SYSZ_INS_MLG      = C.SYSZ_INS_MLG
	SYSZ_INS_MLGR     = C.SYSZ_INS_MLGR
	SYSZ_INS_MS       = C.SYSZ_INS_MS
	SYSZ_INS_MSDB     = C.SYSZ_INS_MSDB
	SYSZ_INS_MSDBR    = C.SYSZ_INS_MSDBR
	SYSZ_INS_MSEB     = C.SYSZ_INS_MSEB
	SYSZ_INS_MSEBR    = C.SYSZ_INS_MSEBR
	SYSZ_INS_MSFI     = C.SYSZ_INS_MSFI
	SYSZ_INS_MSG      = C.SYSZ_INS_MSG
	SYSZ_INS_MSGF     = C.SYSZ_INS_MSGF
	SYSZ_INS_MSGFI    = C.SYSZ_INS_MSGFI
	SYSZ_INS_MSGFR    = C.SYSZ_INS_MSGFR
	SYSZ_INS_MSGR     = C.SYSZ_INS_MSGR
	SYSZ_INS_MSR      = C.SYSZ_INS_MSR
	SYSZ_INS_MSY      = C.SYSZ_INS_MSY
	SYSZ_INS_MVC      = C.SYSZ_INS_MVC
	SYSZ_INS_MVGHI    = C.SYSZ_INS_MVGHI
	SYSZ_INS_MVHHI    = C.SYSZ_INS_MVHHI
	SYSZ_INS_MVHI     = C.SYSZ_INS_MVHI
	SYSZ_INS_MVI      = C.SYSZ_INS_MVI
	SYSZ_INS_MVIY     = C.SYSZ_INS_MVIY
	SYSZ_INS_MVST     = C.SYSZ_INS_MVST
	SYSZ_INS_MXBR     = C.SYSZ_INS_MXBR
	SYSZ_INS_MXDB     = C.SYSZ_INS_MXDB
	SYSZ_INS_MXDBR    = C.SYSZ_INS_MXDBR
	SYSZ_INS_N        = C.SYSZ_INS_N
	SYSZ_INS_NC       = C.SYSZ_INS_NC
	SYSZ_INS_NG       = C.SYSZ_INS_NG
	SYSZ_INS_NGR      = C.SYSZ_INS_NGR
	SYSZ_INS_NGRK     = C.SYSZ_INS_NGRK
	SYSZ_INS_NI       = C.SYSZ_INS_NI
	SYSZ_INS_NIHF     = C.SYSZ_INS_NIHF
	SYSZ_INS_NIHH     = C.SYSZ_INS_NIHH
	SYSZ_INS_NIHL     = C.SYSZ_INS_NIHL
	SYSZ_INS_NILF     = C.SYSZ_INS_NILF
	SYSZ_INS_NILH     = C.SYSZ_INS_NILH
	SYSZ_INS_NILL     = C.SYSZ_INS_NILL
	SYSZ_INS_NIY      = C.SYSZ_INS_NIY
	SYSZ_INS_NR       = C.SYSZ_INS_NR
	SYSZ_INS_NRK      = C.SYSZ_INS_NRK
	SYSZ_INS_NY       = C.SYSZ_INS_NY
	SYSZ_INS_O        = C.SYSZ_INS_O
	SYSZ_INS_OC       = C.SYSZ_INS_OC
	SYSZ_INS_OG       = C.SYSZ_INS_OG
	SYSZ_INS_OGR      = C.SYSZ_INS_OGR
	SYSZ_INS_OGRK     = C.SYSZ_INS_OGRK
	SYSZ_INS_OI       = C.SYSZ_INS_OI
	SYSZ_INS_OIHF     = C.SYSZ_INS_OIHF
	SYSZ_INS_OIHH     = C.SYSZ_INS_OIHH
	SYSZ_INS_OIHL     = C.SYSZ_INS_OIHL
	SYSZ_INS_OILF     = C.SYSZ_INS_OILF
	SYSZ_INS_OILH     = C.SYSZ_INS_OILH
	SYSZ_INS_OILL     = C.SYSZ_INS_OILL
	SYSZ_INS_OIY      = C.SYSZ_INS_OIY
	SYSZ_INS_OR       = C.SYSZ_INS_OR
	SYSZ_INS_ORK      = C.SYSZ_INS_ORK
	SYSZ_INS_OY       = C.SYSZ_INS_OY
	SYSZ_INS_PFD      = C.SYSZ_INS_PFD
	SYSZ_INS_PFDRL    = C.SYSZ_INS_PFDRL
	SYSZ_INS_RISBG    = C.SYSZ_INS_RISBG
	SYSZ_INS_RISBHG   = C.SYSZ_INS_RISBHG
	SYSZ_INS_RISBLG   = C.SYSZ_INS_RISBLG
	SYSZ_INS_RLL      = C.SYSZ_INS_RLL
	SYSZ_INS_RLLG     = C.SYSZ_INS_RLLG
	SYSZ_INS_RNSBG    = C.SYSZ_INS_RNSBG
	SYSZ_INS_ROSBG    = C.SYSZ_INS_ROSBG
	SYSZ_INS_RXSBG    = C.SYSZ_INS_RXSBG
	SYSZ_INS_S        = C.SYSZ_INS_S
	SYSZ_INS_SDB      = C.SYSZ_INS_SDB
	SYSZ_INS_SDBR     = C.SYSZ_INS_SDBR
	SYSZ_INS_SEB      = C.SYSZ_INS_SEB
	SYSZ_INS_SEBR     = C.SYSZ_INS_SEBR
	SYSZ_INS_SG       = C.SYSZ_INS_SG
	SYSZ_INS_SGF      = C.SYSZ_INS_SGF
	SYSZ_INS_SGFR     = C.SYSZ_INS_SGFR
	SYSZ_INS_SGR      = C.SYSZ_INS_SGR
	SYSZ_INS_SGRK     = C.SYSZ_INS_SGRK
	SYSZ_INS_SH       = C.SYSZ_INS_SH
	SYSZ_INS_SHY      = C.SYSZ_INS_SHY
	SYSZ_INS_SL       = C.SYSZ_INS_SL
	SYSZ_INS_SLB      = C.SYSZ_INS_SLB
	SYSZ_INS_SLBG     = C.SYSZ_INS_SLBG
	SYSZ_INS_SLBR     = C.SYSZ_INS_SLBR
	SYSZ_INS_SLFI     = C.SYSZ_INS_SLFI
	SYSZ_INS_SLG      = C.SYSZ_INS_SLG
	SYSZ_INS_SLBGR    = C.SYSZ_INS_SLBGR
	SYSZ_INS_SLGF     = C.SYSZ_INS_SLGF
	SYSZ_INS_SLGFI    = C.SYSZ_INS_SLGFI
	SYSZ_INS_SLGFR    = C.SYSZ_INS_SLGFR
	SYSZ_INS_SLGR     = C.SYSZ_INS_SLGR
	SYSZ_INS_SLGRK    = C.SYSZ_INS_SLGRK
	SYSZ_INS_SLL      = C.SYSZ_INS_SLL
	SYSZ_INS_SLLG     = C.SYSZ_INS_SLLG
	SYSZ_INS_SLLK     = C.SYSZ_INS_SLLK
	SYSZ_INS_SLR      = C.SYSZ_INS_SLR
	SYSZ_INS_SLRK     = C.SYSZ_INS_SLRK
	SYSZ_INS_SLY      = C.SYSZ_INS_SLY
	SYSZ_INS_SQDB     = C.SYSZ_INS_SQDB
	SYSZ_INS_SQDBR    = C.SYSZ_INS_SQDBR
	SYSZ_INS_SQEB     = C.SYSZ_INS_SQEB
	SYSZ_INS_SQEBR    = C.SYSZ_INS_SQEBR
	SYSZ_INS_SQXBR    = C.SYSZ_INS_SQXBR
	SYSZ_INS_SR       = C.SYSZ_INS_SR
	SYSZ_INS_SRA      = C.SYSZ_INS_SRA
	SYSZ_INS_SRAG     = C.SYSZ_INS_SRAG
	SYSZ_INS_SRAK     = C.SYSZ_INS_SRAK
	SYSZ_INS_SRK      = C.SYSZ_INS_SRK
	SYSZ_INS_SRL      = C.SYSZ_INS_SRL
	SYSZ_INS_SRLG     = C.SYSZ_INS_SRLG
	SYSZ_INS_SRLK     = C.SYSZ_INS_SRLK
	SYSZ_INS_SRST     = C.SYSZ_INS_SRST
	SYSZ_INS_ST       = C.SYSZ_INS_ST
	SYSZ_INS_STC      = C.SYSZ_INS_STC
	SYSZ_INS_STCH     = C.SYSZ_INS_STCH
	SYSZ_INS_STCY     = C.SYSZ_INS_STCY
	SYSZ_INS_STD      = C.SYSZ_INS_STD
	SYSZ_INS_STDY     = C.SYSZ_INS_STDY
	SYSZ_INS_STE      = C.SYSZ_INS_STE
	SYSZ_INS_STEY     = C.SYSZ_INS_STEY
	SYSZ_INS_STFH     = C.SYSZ_INS_STFH
	SYSZ_INS_STG      = C.SYSZ_INS_STG
	SYSZ_INS_STGRL    = C.SYSZ_INS_STGRL
	SYSZ_INS_STH      = C.SYSZ_INS_STH
	SYSZ_INS_STHH     = C.SYSZ_INS_STHH
	SYSZ_INS_STHRL    = C.SYSZ_INS_STHRL
	SYSZ_INS_STHY     = C.SYSZ_INS_STHY
	SYSZ_INS_STMG     = C.SYSZ_INS_STMG
	SYSZ_INS_STRL     = C.SYSZ_INS_STRL
	SYSZ_INS_STRV     = C.SYSZ_INS_STRV
	SYSZ_INS_STRVG    = C.SYSZ_INS_STRVG
	SYSZ_INS_STY      = C.SYSZ_INS_STY
	SYSZ_INS_SXBR     = C.SYSZ_INS_SXBR
	SYSZ_INS_SY       = C.SYSZ_INS_SY
	SYSZ_INS_TM       = C.SYSZ_INS_TM
	SYSZ_INS_TMHH     = C.SYSZ_INS_TMHH
	SYSZ_INS_TMHL     = C.SYSZ_INS_TMHL
	SYSZ_INS_TMLH     = C.SYSZ_INS_TMLH
	SYSZ_INS_TMLL     = C.SYSZ_INS_TMLL
	SYSZ_INS_TMY      = C.SYSZ_INS_TMY
	SYSZ_INS_X        = C.SYSZ_INS_X
	SYSZ_INS_XC       = C.SYSZ_INS_XC
	SYSZ_INS_XG       = C.SYSZ_INS_XG
	SYSZ_INS_XGR      = C.SYSZ_INS_XGR
	SYSZ_INS_XGRK     = C.SYSZ_INS_XGRK
	SYSZ_INS_XI       = C.SYSZ_INS_XI
	SYSZ_INS_XIHF     = C.SYSZ_INS_XIHF
	SYSZ_INS_XILF     = C.SYSZ_INS_XILF
	SYSZ_INS_XIY      = C.SYSZ_INS_XIY
	SYSZ_INS_XR       = C.SYSZ_INS_XR
	SYSZ_INS_XRK      = C.SYSZ_INS_XRK
	SYSZ_INS_XY       = C.SYSZ_INS_XY
	SYSZ_INS_MAX      = C.SYSZ_INS_MAX
)

const (
	// Groups
	SYSZ_GRP_INVALID                   = C.SYSZ_GRP_INVALID
	SYSZ_GRP_FEATUREDISTINCTOPS        = C.SYSZ_GRP_FEATUREDISTINCTOPS
	SYSZ_GRP_FEATUREFPEXTENSION        = C.SYSZ_GRP_FEATUREFPEXTENSION
	SYSZ_GRP_FEATUREHIGHWORD           = C.SYSZ_GRP_FEATUREHIGHWORD
	SYSZ_GRP_FEATUREINTERLOCKEDACCESS1 = C.SYSZ_GRP_FEATUREINTERLOCKEDACCESS1
	SYSZ_GRP_FEATURELOADSTOREONCOND    = C.SYSZ_GRP_FEATURELOADSTOREONCOND
	SYSZ_GRP_JUMP                      = C.SYSZ_GRP_JUMP
	SYSZ_GRP_MAX                       = C.SYSZ_GRP_MAX
)
