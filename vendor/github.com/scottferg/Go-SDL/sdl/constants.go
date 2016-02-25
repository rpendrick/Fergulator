package sdl

// #cgo pkg-config: sdl
// #include <SDL.h>
import "C"

const (
	// init flags

	INIT_AUDIO       = C.SDL_INIT_AUDIO
	INIT_VIDEO       = C.SDL_INIT_VIDEO
	INIT_CDROM       = C.SDL_INIT_CDROM
	INIT_TIMER       = C.SDL_INIT_TIMER
	INIT_JOYSTICK    = C.SDL_INIT_JOYSTICK
	INIT_NOPARACHUTE = C.SDL_INIT_NOPARACHUTE
	INIT_EVENTTHREAD = C.SDL_INIT_EVENTTHREAD
	INIT_EVERYTHING  = C.SDL_INIT_EVERYTHING

	// application states

	APPMOUSEFOCUS = C.SDL_APPMOUSEFOCUS
	APPINPUTFOCUS = C.SDL_APPINPUTFOCUS
	APPACTIVE     = C.SDL_APPACTIVE

	// setvideo flags

	SWSURFACE    = C.SDL_SWSURFACE
	HWSURFACE    = C.SDL_HWSURFACE
	ASYNCBLIT    = C.SDL_ASYNCBLIT
	ANYFORMAT    = C.SDL_ANYFORMAT
	HWPALETTE    = C.SDL_HWPALETTE
	DOUBLEBUF    = C.SDL_DOUBLEBUF
	FULLSCREEN   = C.SDL_FULLSCREEN
	OPENGL       = C.SDL_OPENGL
	OPENGLBLIT   = C.SDL_OPENGLBLIT
	RESIZABLE    = C.SDL_RESIZABLE
	NOFRAME      = C.SDL_NOFRAME
	HWACCEL      = C.SDL_HWACCEL
	SRCCOLORKEY  = C.SDL_SRCCOLORKEY
	RLEACCELOK   = C.SDL_RLEACCELOK
	RLEACCEL     = C.SDL_RLEACCEL
	SRCALPHA     = C.SDL_SRCALPHA
	PREALLOC     = C.SDL_PREALLOC
	YV12_OVERLAY = C.SDL_YV12_OVERLAY
	IYUV_OVERLAY = C.SDL_IYUV_OVERLAY
	YUY2_OVERLAY = C.SDL_YUY2_OVERLAY
	UYVY_OVERLAY = C.SDL_UYVY_OVERLAY
	YVYU_OVERLAY = C.SDL_YVYU_OVERLAY
	LOGPAL       = C.SDL_LOGPAL
	PHYSPAL      = C.SDL_PHYSPAL

	// More setvideo flags: GLattr enumeration

	GL_RED_SIZE           = C.SDL_GL_RED_SIZE
	GL_GREEN_SIZE         = C.SDL_GL_GREEN_SIZE
	GL_BLUE_SIZE          = C.SDL_GL_BLUE_SIZE
	GL_ALPHA_SIZE         = C.SDL_GL_ALPHA_SIZE
	GL_BUFFER_SIZE        = C.SDL_GL_BUFFER_SIZE
	GL_DOUBLEBUFFER       = C.SDL_GL_DOUBLEBUFFER
	GL_DEPTH_SIZE         = C.SDL_GL_DEPTH_SIZE
	GL_STENCIL_SIZE       = C.SDL_GL_STENCIL_SIZE
	GL_ACCUM_RED_SIZE     = C.SDL_GL_ACCUM_RED_SIZE
	GL_ACCUM_GREEN_SIZE   = C.SDL_GL_ACCUM_GREEN_SIZE
	GL_ACCUM_BLUE_SIZE    = C.SDL_GL_ACCUM_BLUE_SIZE
	GL_ACCUM_ALPHA_SIZE   = C.SDL_GL_ACCUM_ALPHA_SIZE
	GL_STEREO             = C.SDL_GL_STEREO
	GL_MULTISAMPLEBUFFERS = C.SDL_GL_MULTISAMPLEBUFFERS
	GL_MULTISAMPLESAMPLES = C.SDL_GL_MULTISAMPLESAMPLES
	GL_ACCELERATED_VISUAL = C.SDL_GL_ACCELERATED_VISUAL
	GL_SWAP_CONTROL       = C.SDL_GL_SWAP_CONTROL

	// event types

	NOEVENT         = C.SDL_NOEVENT
	ACTIVEEVENT     = C.SDL_ACTIVEEVENT
	KEYDOWN         = C.SDL_KEYDOWN
	KEYUP           = C.SDL_KEYUP
	MOUSEMOTION     = C.SDL_MOUSEMOTION
	MOUSEBUTTONDOWN = C.SDL_MOUSEBUTTONDOWN
	MOUSEBUTTONUP   = C.SDL_MOUSEBUTTONUP
	JOYAXISMOTION   = C.SDL_JOYAXISMOTION
	JOYBALLMOTION   = C.SDL_JOYBALLMOTION
	JOYHATMOTION    = C.SDL_JOYHATMOTION
	JOYBUTTONDOWN   = C.SDL_JOYBUTTONDOWN
	JOYBUTTONUP     = C.SDL_JOYBUTTONUP
	QUIT            = C.SDL_QUIT
	SYSWMEVENT      = C.SDL_SYSWMEVENT
	EVENT_RESERVEDA = C.SDL_EVENT_RESERVEDA
	EVENT_RESERVEDB = C.SDL_EVENT_RESERVEDB
	VIDEORESIZE     = C.SDL_VIDEORESIZE
	VIDEOEXPOSE     = C.SDL_VIDEOEXPOSE
	EVENT_RESERVED2 = C.SDL_EVENT_RESERVED2
	EVENT_RESERVED3 = C.SDL_EVENT_RESERVED3
	EVENT_RESERVED4 = C.SDL_EVENT_RESERVED4
	EVENT_RESERVED5 = C.SDL_EVENT_RESERVED5
	EVENT_RESERVED6 = C.SDL_EVENT_RESERVED6
	EVENT_RESERVED7 = C.SDL_EVENT_RESERVED7

	USEREVENT = C.SDL_USEREVENT

	NUMEVENTS = C.SDL_NUMEVENTS

	// event masks

	ACTIVEEVENTMASK     = C.SDL_ACTIVEEVENTMASK
	KEYDOWNMASK         = C.SDL_KEYDOWNMASK
	KEYUPMASK           = C.SDL_KEYUPMASK
	KEYEVENTMASK        = C.SDL_KEYEVENTMASK
	MOUSEMOTIONMASK     = C.SDL_MOUSEMOTIONMASK
	MOUSEBUTTONDOWNMASK = C.SDL_MOUSEBUTTONDOWNMASK
	MOUSEBUTTONUPMASK   = C.SDL_MOUSEBUTTONUPMASK
	MOUSEEVENTMASK      = C.SDL_MOUSEEVENTMASK
	JOYAXISMOTIONMASK   = C.SDL_JOYAXISMOTIONMASK
	JOYBALLMOTIONMASK   = C.SDL_JOYBALLMOTIONMASK
	JOYHATMOTIONMASK    = C.SDL_JOYHATMOTIONMASK
	JOYBUTTONDOWNMASK   = C.SDL_JOYBUTTONDOWNMASK
	JOYBUTTONUPMASK     = C.SDL_JOYBUTTONUPMASK
	JOYEVENTMASK        = C.SDL_JOYEVENTMASK
	VIDEORESIZEMASK     = C.SDL_VIDEORESIZEMASK
	VIDEOEXPOSEMASK     = C.SDL_VIDEOEXPOSEMASK
	QUITMASK            = C.SDL_QUITMASK
	SYSWMEVENTMASK      = C.SDL_SYSWMEVENTMASK

	// event state

	QUERY   = C.SDL_QUERY
	DISABLE = C.SDL_DISABLE
	ENABLE  = C.SDL_ENABLE

	// keys
	K_UNKNOWN      = C.SDLK_UNKNOWN
	K_FIRST        = C.SDLK_FIRST
	K_BACKSPACE    = C.SDLK_BACKSPACE
	K_TAB          = C.SDLK_TAB
	K_CLEAR        = C.SDLK_CLEAR
	K_RETURN       = C.SDLK_RETURN
	K_PAUSE        = C.SDLK_PAUSE
	K_ESCAPE       = C.SDLK_ESCAPE
	K_SPACE        = C.SDLK_SPACE
	K_EXCLAIM      = C.SDLK_EXCLAIM
	K_QUOTEDBL     = C.SDLK_QUOTEDBL
	K_HASH         = C.SDLK_HASH
	K_DOLLAR       = C.SDLK_DOLLAR
	K_AMPERSAND    = C.SDLK_AMPERSAND
	K_QUOTE        = C.SDLK_QUOTE
	K_LEFTPAREN    = C.SDLK_LEFTPAREN
	K_RIGHTPAREN   = C.SDLK_RIGHTPAREN
	K_ASTERISK     = C.SDLK_ASTERISK
	K_PLUS         = C.SDLK_PLUS
	K_COMMA        = C.SDLK_COMMA
	K_MINUS        = C.SDLK_MINUS
	K_PERIOD       = C.SDLK_PERIOD
	K_SLASH        = C.SDLK_SLASH
	K_0            = C.SDLK_0
	K_1            = C.SDLK_1
	K_2            = C.SDLK_2
	K_3            = C.SDLK_3
	K_4            = C.SDLK_4
	K_5            = C.SDLK_5
	K_6            = C.SDLK_6
	K_7            = C.SDLK_7
	K_8            = C.SDLK_8
	K_9            = C.SDLK_9
	K_COLON        = C.SDLK_COLON
	K_SEMICOLON    = C.SDLK_SEMICOLON
	K_LESS         = C.SDLK_LESS
	K_EQUALS       = C.SDLK_EQUALS
	K_GREATER      = C.SDLK_GREATER
	K_QUESTION     = C.SDLK_QUESTION
	K_AT           = C.SDLK_AT
	K_LEFTBRACKET  = C.SDLK_LEFTBRACKET
	K_BACKSLASH    = C.SDLK_BACKSLASH
	K_RIGHTBRACKET = C.SDLK_RIGHTBRACKET
	K_CARET        = C.SDLK_CARET
	K_UNDERSCORE   = C.SDLK_UNDERSCORE
	K_BACKQUOTE    = C.SDLK_BACKQUOTE
	K_a            = C.SDLK_a
	K_b            = C.SDLK_b
	K_c            = C.SDLK_c
	K_d            = C.SDLK_d
	K_e            = C.SDLK_e
	K_f            = C.SDLK_f
	K_g            = C.SDLK_g
	K_h            = C.SDLK_h
	K_i            = C.SDLK_i
	K_j            = C.SDLK_j
	K_k            = C.SDLK_k
	K_l            = C.SDLK_l
	K_m            = C.SDLK_m
	K_n            = C.SDLK_n
	K_o            = C.SDLK_o
	K_p            = C.SDLK_p
	K_q            = C.SDLK_q
	K_r            = C.SDLK_r
	K_s            = C.SDLK_s
	K_t            = C.SDLK_t
	K_u            = C.SDLK_u
	K_v            = C.SDLK_v
	K_w            = C.SDLK_w
	K_x            = C.SDLK_x
	K_y            = C.SDLK_y
	K_z            = C.SDLK_z
	K_DELETE       = C.SDLK_DELETE
	K_WORLD_0      = C.SDLK_WORLD_0
	K_WORLD_1      = C.SDLK_WORLD_1
	K_WORLD_2      = C.SDLK_WORLD_2
	K_WORLD_3      = C.SDLK_WORLD_3
	K_WORLD_4      = C.SDLK_WORLD_4
	K_WORLD_5      = C.SDLK_WORLD_5
	K_WORLD_6      = C.SDLK_WORLD_6
	K_WORLD_7      = C.SDLK_WORLD_7
	K_WORLD_8      = C.SDLK_WORLD_8
	K_WORLD_9      = C.SDLK_WORLD_9
	K_WORLD_10     = C.SDLK_WORLD_10
	K_WORLD_11     = C.SDLK_WORLD_11
	K_WORLD_12     = C.SDLK_WORLD_12
	K_WORLD_13     = C.SDLK_WORLD_13
	K_WORLD_14     = C.SDLK_WORLD_14
	K_WORLD_15     = C.SDLK_WORLD_15
	K_WORLD_16     = C.SDLK_WORLD_16
	K_WORLD_17     = C.SDLK_WORLD_17
	K_WORLD_18     = C.SDLK_WORLD_18
	K_WORLD_19     = C.SDLK_WORLD_19
	K_WORLD_20     = C.SDLK_WORLD_20
	K_WORLD_21     = C.SDLK_WORLD_21
	K_WORLD_22     = C.SDLK_WORLD_22
	K_WORLD_23     = C.SDLK_WORLD_23
	K_WORLD_24     = C.SDLK_WORLD_24
	K_WORLD_25     = C.SDLK_WORLD_25
	K_WORLD_26     = C.SDLK_WORLD_26
	K_WORLD_27     = C.SDLK_WORLD_27
	K_WORLD_28     = C.SDLK_WORLD_28
	K_WORLD_29     = C.SDLK_WORLD_29
	K_WORLD_30     = C.SDLK_WORLD_30
	K_WORLD_31     = C.SDLK_WORLD_31
	K_WORLD_32     = C.SDLK_WORLD_32
	K_WORLD_33     = C.SDLK_WORLD_33
	K_WORLD_34     = C.SDLK_WORLD_34
	K_WORLD_35     = C.SDLK_WORLD_35
	K_WORLD_36     = C.SDLK_WORLD_36
	K_WORLD_37     = C.SDLK_WORLD_37
	K_WORLD_38     = C.SDLK_WORLD_38
	K_WORLD_39     = C.SDLK_WORLD_39
	K_WORLD_40     = C.SDLK_WORLD_40
	K_WORLD_41     = C.SDLK_WORLD_41
	K_WORLD_42     = C.SDLK_WORLD_42
	K_WORLD_43     = C.SDLK_WORLD_43
	K_WORLD_44     = C.SDLK_WORLD_44
	K_WORLD_45     = C.SDLK_WORLD_45
	K_WORLD_46     = C.SDLK_WORLD_46
	K_WORLD_47     = C.SDLK_WORLD_47
	K_WORLD_48     = C.SDLK_WORLD_48
	K_WORLD_49     = C.SDLK_WORLD_49
	K_WORLD_50     = C.SDLK_WORLD_50
	K_WORLD_51     = C.SDLK_WORLD_51
	K_WORLD_52     = C.SDLK_WORLD_52
	K_WORLD_53     = C.SDLK_WORLD_53
	K_WORLD_54     = C.SDLK_WORLD_54
	K_WORLD_55     = C.SDLK_WORLD_55
	K_WORLD_56     = C.SDLK_WORLD_56
	K_WORLD_57     = C.SDLK_WORLD_57
	K_WORLD_58     = C.SDLK_WORLD_58
	K_WORLD_59     = C.SDLK_WORLD_59
	K_WORLD_60     = C.SDLK_WORLD_60
	K_WORLD_61     = C.SDLK_WORLD_61
	K_WORLD_62     = C.SDLK_WORLD_62
	K_WORLD_63     = C.SDLK_WORLD_63
	K_WORLD_64     = C.SDLK_WORLD_64
	K_WORLD_65     = C.SDLK_WORLD_65
	K_WORLD_66     = C.SDLK_WORLD_66
	K_WORLD_67     = C.SDLK_WORLD_67
	K_WORLD_68     = C.SDLK_WORLD_68
	K_WORLD_69     = C.SDLK_WORLD_69
	K_WORLD_70     = C.SDLK_WORLD_70
	K_WORLD_71     = C.SDLK_WORLD_71
	K_WORLD_72     = C.SDLK_WORLD_72
	K_WORLD_73     = C.SDLK_WORLD_73
	K_WORLD_74     = C.SDLK_WORLD_74
	K_WORLD_75     = C.SDLK_WORLD_75
	K_WORLD_76     = C.SDLK_WORLD_76
	K_WORLD_77     = C.SDLK_WORLD_77
	K_WORLD_78     = C.SDLK_WORLD_78
	K_WORLD_79     = C.SDLK_WORLD_79
	K_WORLD_80     = C.SDLK_WORLD_80
	K_WORLD_81     = C.SDLK_WORLD_81
	K_WORLD_82     = C.SDLK_WORLD_82
	K_WORLD_83     = C.SDLK_WORLD_83
	K_WORLD_84     = C.SDLK_WORLD_84
	K_WORLD_85     = C.SDLK_WORLD_85
	K_WORLD_86     = C.SDLK_WORLD_86
	K_WORLD_87     = C.SDLK_WORLD_87
	K_WORLD_88     = C.SDLK_WORLD_88
	K_WORLD_89     = C.SDLK_WORLD_89
	K_WORLD_90     = C.SDLK_WORLD_90
	K_WORLD_91     = C.SDLK_WORLD_91
	K_WORLD_92     = C.SDLK_WORLD_92
	K_WORLD_93     = C.SDLK_WORLD_93
	K_WORLD_94     = C.SDLK_WORLD_94
	K_WORLD_95     = C.SDLK_WORLD_95
	K_KP0          = C.SDLK_KP0
	K_KP1          = C.SDLK_KP1
	K_KP2          = C.SDLK_KP2
	K_KP3          = C.SDLK_KP3
	K_KP4          = C.SDLK_KP4
	K_KP5          = C.SDLK_KP5
	K_KP6          = C.SDLK_KP6
	K_KP7          = C.SDLK_KP7
	K_KP8          = C.SDLK_KP8
	K_KP9          = C.SDLK_KP9
	K_KP_PERIOD    = C.SDLK_KP_PERIOD
	K_KP_DIVIDE    = C.SDLK_KP_DIVIDE
	K_KP_MULTIPLY  = C.SDLK_KP_MULTIPLY
	K_KP_MINUS     = C.SDLK_KP_MINUS
	K_KP_PLUS      = C.SDLK_KP_PLUS
	K_KP_ENTER     = C.SDLK_KP_ENTER
	K_KP_EQUALS    = C.SDLK_KP_EQUALS
	K_UP           = C.SDLK_UP
	K_DOWN         = C.SDLK_DOWN
	K_RIGHT        = C.SDLK_RIGHT
	K_LEFT         = C.SDLK_LEFT
	K_INSERT       = C.SDLK_INSERT
	K_HOME         = C.SDLK_HOME
	K_END          = C.SDLK_END
	K_PAGEUP       = C.SDLK_PAGEUP
	K_PAGEDOWN     = C.SDLK_PAGEDOWN
	K_F1           = C.SDLK_F1
	K_F2           = C.SDLK_F2
	K_F3           = C.SDLK_F3
	K_F4           = C.SDLK_F4
	K_F5           = C.SDLK_F5
	K_F6           = C.SDLK_F6
	K_F7           = C.SDLK_F7
	K_F8           = C.SDLK_F8
	K_F9           = C.SDLK_F9
	K_F10          = C.SDLK_F10
	K_F11          = C.SDLK_F11
	K_F12          = C.SDLK_F12
	K_F13          = C.SDLK_F13
	K_F14          = C.SDLK_F14
	K_F15          = C.SDLK_F15
	K_NUMLOCK      = C.SDLK_NUMLOCK
	K_CAPSLOCK     = C.SDLK_CAPSLOCK
	K_SCROLLOCK    = C.SDLK_SCROLLOCK
	K_RSHIFT       = C.SDLK_RSHIFT
	K_LSHIFT       = C.SDLK_LSHIFT
	K_RCTRL        = C.SDLK_RCTRL
	K_LCTRL        = C.SDLK_LCTRL
	K_RALT         = C.SDLK_RALT
	K_LALT         = C.SDLK_LALT
	K_RMETA        = C.SDLK_RMETA
	K_LMETA        = C.SDLK_LMETA
	K_LSUPER       = C.SDLK_LSUPER
	K_RSUPER       = C.SDLK_RSUPER
	K_MODE         = C.SDLK_MODE
	K_COMPOSE      = C.SDLK_COMPOSE
	K_HELP         = C.SDLK_HELP
	K_PRINT        = C.SDLK_PRINT
	K_SYSREQ       = C.SDLK_SYSREQ
	K_BREAK        = C.SDLK_BREAK
	K_MENU         = C.SDLK_MENU
	K_POWER        = C.SDLK_POWER
	K_EURO         = C.SDLK_EURO
	K_UNDO         = C.SDLK_UNDO

	// key mods

	KMOD_NONE     = C.KMOD_NONE
	KMOD_LSHIFT   = C.KMOD_LSHIFT
	KMOD_RSHIFT   = C.KMOD_RSHIFT
	KMOD_LCTRL    = C.KMOD_LCTRL
	KMOD_RCTRL    = C.KMOD_RCTRL
	KMOD_LALT     = C.KMOD_LALT
	KMOD_RALT     = C.KMOD_RALT
	KMOD_LMETA    = C.KMOD_LMETA
	KMOD_RMETA    = C.KMOD_RMETA
	KMOD_NUM      = C.KMOD_NUM
	KMOD_CAPS     = C.KMOD_CAPS
	KMOD_MODE     = C.KMOD_MODE
	KMOD_RESERVED = C.KMOD_RESERVED

	// hat states

	HAT_CENTERED  = C.SDL_HAT_CENTERED
	HAT_UP        = C.SDL_HAT_UP
	HAT_RIGHT     = C.SDL_HAT_RIGHT
	HAT_DOWN      = C.SDL_HAT_DOWN
	HAT_LEFT      = C.SDL_HAT_LEFT
	HAT_RIGHTUP   = C.SDL_HAT_RIGHTUP
	HAT_RIGHTDOWN = C.SDL_HAT_RIGHTDOWN
	HAT_LEFTUP    = C.SDL_HAT_LEFTUP
	HAT_LEFTDOWN  = C.SDL_HAT_LEFTDOWN

	// keyboard/mouse state

	RELEASED = C.SDL_RELEASED
	PRESSED  = C.SDL_PRESSED

	// mouse button constants

	BUTTON_LEFT      = C.SDL_BUTTON_LEFT
	BUTTON_MIDDLE    = C.SDL_BUTTON_MIDDLE
	BUTTON_RIGHT     = C.SDL_BUTTON_RIGHT
	BUTTON_WHEELUP   = C.SDL_BUTTON_WHEELUP
	BUTTON_WHEELDOWN = C.SDL_BUTTON_WHEELDOWN
	BUTTON_X1        = C.SDL_BUTTON_X1
	BUTTON_X2        = C.SDL_BUTTON_X2

	// mouse button masks

	BUTTON_LMASK         = 1 << (BUTTON_LEFT - 1)
	BUTTON_MMASK         = 1 << (BUTTON_MIDDLE - 1)
	BUTTON_RMASK         = 1 << (BUTTON_RIGHT - 1)
	BUTTON_WHEELUPMASK   = 1 << (BUTTON_WHEELUP - 1)
	BUTTON_WHEELDOWNMASK = 1 << (BUTTON_WHEELDOWN - 1)
	BUTTON_X1MASK        = 1 << (BUTTON_X1 - 1)
	BUTTON_X2MASK        = 1 << (BUTTON_X2 - 1)
)