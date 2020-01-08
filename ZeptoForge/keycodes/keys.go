package keycodes

type Key int

const (
	/* Printable keys */
	ZF_KeySpace        Key = 32
	ZF_KeyApostrophe   Key = 39 /* ' */
	ZF_KeyComma        Key = 44 /* , */
	ZF_KeyMinus        Key = 45 /* - */
	ZF_KeyPeriod       Key = 46 /* . */
	ZF_KeySlash        Key = 47 /* / */
	ZF_Key0            Key = 48
	ZF_Key1            Key = 49
	ZF_Key2            Key = 50
	ZF_Key3            Key = 51
	ZF_Key4            Key = 52
	ZF_Key5            Key = 53
	ZF_Key6            Key = 54
	ZF_Key7            Key = 55
	ZF_Key8            Key = 56
	ZF_Key9            Key = 57
	ZF_KeySemicolon    Key = 59 /* ; */
	ZF_KeyEqual        Key = 61 /*Key = */
	ZF_KeyA            Key = 65
	ZF_KeyB            Key = 66
	ZF_KeyC            Key = 67
	ZF_KeyD            Key = 68
	ZF_KeyE            Key = 69
	ZF_KeyF            Key = 70
	ZF_KeyG            Key = 71
	ZF_KeyH            Key = 72
	ZF_KeyI            Key = 73
	ZF_KeyJ            Key = 74
	ZF_KeyK            Key = 75
	ZF_KeyL            Key = 76
	ZF_KeyM            Key = 77
	ZF_KeyN            Key = 78
	ZF_KeyO            Key = 79
	ZF_KeyP            Key = 80
	ZF_KeyQ            Key = 81
	ZF_KeyR            Key = 82
	ZF_KeyS            Key = 83
	ZF_KeyT            Key = 84
	ZF_KeyU            Key = 85
	ZF_KeyV            Key = 86
	ZF_KeyW            Key = 87
	ZF_KeyX            Key = 88
	ZF_KeyY            Key = 89
	ZF_KeyZ            Key = 90
	ZF_KeyLeftBracket  Key = 91  /* [ */
	ZF_KeyBackslash    Key = 92  /* \ */
	ZF_KeyRightBracket Key = 93  /* ] */
	ZF_KeyGraveAccent  Key = 96  /* ` */
	ZF_KeyWorld1       Key = 161 /* non-US #1 */
	ZF_KeyWorld2       Key = 162 /* non-US #2 */

	/* Function keys */
	ZF_KeyEscape       Key = 256
	ZF_KeyEnter        Key = 257
	ZF_KeyTab          Key = 258
	ZF_KeyBackspace    Key = 259
	ZF_KeyInsert       Key = 260
	ZF_KeyDelete       Key = 261
	ZF_KeyRight        Key = 262
	ZF_KeyLeft         Key = 263
	ZF_KeyDown         Key = 264
	ZF_KeyUp           Key = 265
	ZF_KeyPageUp       Key = 266
	ZF_KeyPageDown     Key = 267
	ZF_KeyHome         Key = 268
	ZF_KeyEnd          Key = 269
	ZF_KeyCapsLock     Key = 280
	ZF_KeyScrollLock   Key = 281
	ZF_KeyNumLock      Key = 282
	ZF_KeyPrintScreen  Key = 283
	ZF_KeyPause        Key = 284
	ZF_KeyF1           Key = 290
	ZF_KeyF2           Key = 291
	ZF_KeyF3           Key = 292
	ZF_KeyF4           Key = 293
	ZF_KeyF5           Key = 294
	ZF_KeyF6           Key = 295
	ZF_KeyF7           Key = 296
	ZF_KeyF8           Key = 297
	ZF_KeyF9           Key = 298
	ZF_KeyF10          Key = 299
	ZF_KeyF11          Key = 300
	ZF_KeyF12          Key = 301
	ZF_KeyF13          Key = 302
	ZF_KeyF14          Key = 303
	ZF_KeyF15          Key = 304
	ZF_KeyF16          Key = 305
	ZF_KeyF17          Key = 306
	ZF_KeyF18          Key = 307
	ZF_KeyF19          Key = 308
	ZF_KeyF20          Key = 309
	ZF_KeyF21          Key = 310
	ZF_KeyF22          Key = 311
	ZF_KeyF23          Key = 312
	ZF_KeyF24          Key = 313
	ZF_KeyF25          Key = 314
	ZF_KeyKP0          Key = 320
	ZF_KeyKP1          Key = 321
	ZF_KeyKP2          Key = 322
	ZF_KeyKP3          Key = 323
	ZF_KeyKP4          Key = 324
	ZF_KeyKP5          Key = 325
	ZF_KeyKP6          Key = 326
	ZF_KeyKP7          Key = 327
	ZF_KeyKP8          Key = 328
	ZF_KeyKP9          Key = 329
	ZF_KeyKPDecimal    Key = 330
	ZF_KeyKPDivide     Key = 331
	ZF_KeyKPMultiply   Key = 332
	ZF_KeyKPSubtract   Key = 333
	ZF_KeyKPAdd        Key = 334
	ZF_KeyKPEnter      Key = 335
	ZF_KeyKPEqual      Key = 336
	ZF_KeyLeftShift    Key = 340
	ZF_KeyLeftControl  Key = 341
	ZF_KeyLeftAlt      Key = 342
	ZF_KeyLeftSuper    Key = 343
	ZF_KeyRightShift   Key = 344
	ZF_KeyRightControl Key = 345
	ZF_KeyRightAlt     Key = 346
	ZF_KeyRightSuper   Key = 347
	ZF_KeyMenu         Key = 348

	ZF_KeyLast Key = ZF_KeyMenu
)