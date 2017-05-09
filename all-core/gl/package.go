// Copyright (c) 2010 Khronos Group.
// This material may be distributed subject to the terms and conditions
// set forth in the Open Publication License, v 1.0, 8 June 1999.
// http://opencontent.org/openpub/.
//
// Copyright (c) 1991-2006 Silicon Graphics, Inc.
// This document is licensed under the SGI Free Software B License.
// For details, see http://oss.sgi.com/projects/FreeB.

// Package gl implements Go bindings to OpenGL.
//
// This package was automatically generated using Glow:
//  http://github.com/go-gl/glow
//
// Generated based on the OpenGL XML specification:
//  SVN revision 27695
package gl

// #cgo darwin  LDFLAGS: -framework OpenGL
// #cgo linux   LDFLAGS: -lGL
// #cgo windows LDFLAGS: -lopengl32
// #if defined(_WIN32) && !defined(APIENTRY) && !defined(__CYGWIN__) && !defined(__SCITECH_SNAP__)
// #ifndef WIN32_LEAN_AND_MEAN
// #define WIN32_LEAN_AND_MEAN 1
// #endif
// #include <windows.h>
// #endif
// #ifndef APIENTRY
// #define APIENTRY
// #endif
// #ifndef APIENTRYP
// #define APIENTRYP APIENTRY *
// #endif
// #ifndef GLAPI
// #define GLAPI extern
// #endif
// #include <stddef.h>
// #ifndef GLEXT_64_TYPES_DEFINED
// /* This code block is duplicated in glxext.h, so must be protected */
// #define GLEXT_64_TYPES_DEFINED
// /* Define int32_t, int64_t, and uint64_t types for UST/MSC */
// /* (as used in the GL_EXT_timer_query extension). */
// #if defined(__STDC_VERSION__) && __STDC_VERSION__ >= 199901L
// #include <inttypes.h>
// #elif defined(__sun__) || defined(__digital__)
// #include <inttypes.h>
// #if defined(__STDC__)
// #if defined(__arch64__) || defined(_LP64)
// typedef long int int64_t;
// typedef unsigned long int uint64_t;
// #else
// typedef long long int int64_t;
// typedef unsigned long long int uint64_t;
// #endif /* __arch64__ */
// #endif /* __STDC__ */
// #elif defined( __VMS ) || defined(__sgi)
// #include <inttypes.h>
// #elif defined(__SCO__) || defined(__USLC__)
// #include <stdint.h>
// #elif defined(__UNIXOS2__) || defined(__SOL64__)
// typedef long int int32_t;
// typedef long long int int64_t;
// typedef unsigned long long int uint64_t;
// #elif defined(_WIN32) && defined(__GNUC__)
// #include <stdint.h>
// #elif defined(_WIN32)
// typedef __int32 int32_t;
// typedef __int64 int64_t;
// typedef unsigned __int64 uint64_t;
// #else
// /* Fallback if nothing above works */
// #include <inttypes.h>
// #endif
// #endif
// typedef unsigned int GLenum;
// typedef unsigned char GLboolean;
// typedef unsigned int GLbitfield;
// typedef signed char GLbyte;
// typedef short GLshort;
// typedef int GLint;
// typedef unsigned char GLubyte;
// typedef unsigned short GLushort;
// typedef unsigned int GLuint;
// typedef int GLsizei;
// typedef float GLfloat;
// typedef double GLdouble;
// typedef char GLchar;
// typedef ptrdiff_t GLintptr;
// typedef ptrdiff_t GLsizeiptr;
// typedef int64_t GLint64;
// typedef uint64_t GLuint64;
// typedef uint64_t GLuint64EXT;
// typedef struct __GLsync *GLsync;
// struct _cl_context;
// struct _cl_event;
// typedef void (APIENTRY *GLDEBUGPROC)(GLenum source,GLenum type,GLuint id,GLenum severity,GLsizei length,const GLchar *message,const void *userParam);
// typedef void (APIENTRY *GLDEBUGPROCARB)(GLenum source,GLenum type,GLuint id,GLenum severity,GLsizei length,const GLchar *message,const void *userParam);
// typedef void (APIENTRY *GLDEBUGPROCKHR)(GLenum source,GLenum type,GLuint id,GLenum severity,GLsizei length,const GLchar *message,const void *userParam);
// extern void glowDebugCallback_glcoreall(GLenum source, GLenum type, GLuint id, GLenum severity, GLsizei length, const GLchar* message, const void* userParam);
// static void APIENTRY glowCDebugCallback(GLenum source, GLenum type, GLuint id, GLenum severity, GLsizei length, const GLchar* message, const void* userParam) {
//   glowDebugCallback_glcoreall(source, type, id, severity, length, message, userParam);
// }
// typedef void  (APIENTRYP GPACCUM)(GLenum  op, GLfloat  value);
// typedef void  (APIENTRYP GPACTIVESHADERPROGRAM)(GLuint  pipeline, GLuint  program);
// typedef void  (APIENTRYP GPACTIVETEXTURE)(GLenum  texture);
// typedef void  (APIENTRYP GPALPHAFUNC)(GLenum  func, GLfloat  ref);
// typedef GLboolean  (APIENTRYP GPARETEXTURESRESIDENT)(GLsizei  n, const GLuint * textures, GLboolean * residences);
// typedef void  (APIENTRYP GPARRAYELEMENT)(GLint  i);
// typedef void  (APIENTRYP GPATTACHSHADER)(GLuint  program, GLuint  shader);
// typedef void  (APIENTRYP GPBEGIN)(GLenum  mode);
// typedef void  (APIENTRYP GPBEGINCONDITIONALRENDER)(GLuint  id, GLenum  mode);
// typedef void  (APIENTRYP GPBEGINQUERY)(GLenum  target, GLuint  id);
// typedef void  (APIENTRYP GPBEGINQUERYINDEXED)(GLenum  target, GLuint  index, GLuint  id);
// typedef void  (APIENTRYP GPBEGINTRANSFORMFEEDBACK)(GLenum  primitiveMode);
// typedef void  (APIENTRYP GPBINDATTRIBLOCATION)(GLuint  program, GLuint  index, const GLchar * name);
// typedef void  (APIENTRYP GPBINDBUFFER)(GLenum  target, GLuint  buffer);
// typedef void  (APIENTRYP GPBINDBUFFERBASE)(GLenum  target, GLuint  index, GLuint  buffer);
// typedef void  (APIENTRYP GPBINDBUFFERRANGE)(GLenum  target, GLuint  index, GLuint  buffer, GLintptr  offset, GLsizeiptr  size);
// typedef void  (APIENTRYP GPBINDBUFFERSBASE)(GLenum  target, GLuint  first, GLsizei  count, const GLuint * buffers);
// typedef void  (APIENTRYP GPBINDBUFFERSRANGE)(GLenum  target, GLuint  first, GLsizei  count, const GLuint * buffers, const GLintptr * offsets, const GLsizeiptr * sizes);
// typedef void  (APIENTRYP GPBINDFRAGDATALOCATION)(GLuint  program, GLuint  color, const GLchar * name);
// typedef void  (APIENTRYP GPBINDFRAGDATALOCATIONINDEXED)(GLuint  program, GLuint  colorNumber, GLuint  index, const GLchar * name);
// typedef void  (APIENTRYP GPBINDFRAMEBUFFER)(GLenum  target, GLuint  framebuffer);
// typedef void  (APIENTRYP GPBINDIMAGETEXTURE)(GLuint  unit, GLuint  texture, GLint  level, GLboolean  layered, GLint  layer, GLenum  access, GLenum  format);
// typedef void  (APIENTRYP GPBINDIMAGETEXTURES)(GLuint  first, GLsizei  count, const GLuint * textures);
// typedef void  (APIENTRYP GPBINDPROGRAMPIPELINE)(GLuint  pipeline);
// typedef void  (APIENTRYP GPBINDRENDERBUFFER)(GLenum  target, GLuint  renderbuffer);
// typedef void  (APIENTRYP GPBINDSAMPLER)(GLuint  unit, GLuint  sampler);
// typedef void  (APIENTRYP GPBINDSAMPLERS)(GLuint  first, GLsizei  count, const GLuint * samplers);
// typedef void  (APIENTRYP GPBINDTEXTURE)(GLenum  target, GLuint  texture);
// typedef void  (APIENTRYP GPBINDTEXTUREUNIT)(GLuint  unit, GLuint  texture);
// typedef void  (APIENTRYP GPBINDTEXTURES)(GLuint  first, GLsizei  count, const GLuint * textures);
// typedef void  (APIENTRYP GPBINDTRANSFORMFEEDBACK)(GLenum  target, GLuint  id);
// typedef void  (APIENTRYP GPBINDVERTEXARRAY)(GLuint  array);
// typedef void  (APIENTRYP GPBINDVERTEXBUFFER)(GLuint  bindingindex, GLuint  buffer, GLintptr  offset, GLsizei  stride);
// typedef void  (APIENTRYP GPBINDVERTEXBUFFERS)(GLuint  first, GLsizei  count, const GLuint * buffers, const GLintptr * offsets, const GLsizei * strides);
// typedef void  (APIENTRYP GPBITMAP)(GLsizei  width, GLsizei  height, GLfloat  xorig, GLfloat  yorig, GLfloat  xmove, GLfloat  ymove, const GLubyte * bitmap);
// typedef void  (APIENTRYP GPBLENDCOLOR)(GLfloat  red, GLfloat  green, GLfloat  blue, GLfloat  alpha);
// typedef void  (APIENTRYP GPBLENDEQUATION)(GLenum  mode);
// typedef void  (APIENTRYP GPBLENDEQUATIONSEPARATE)(GLenum  modeRGB, GLenum  modeAlpha);
// typedef void  (APIENTRYP GPBLENDEQUATIONSEPARATEI)(GLuint  buf, GLenum  modeRGB, GLenum  modeAlpha);
// typedef void  (APIENTRYP GPBLENDEQUATIONSEPARATEIARB)(GLuint  buf, GLenum  modeRGB, GLenum  modeAlpha);
// typedef void  (APIENTRYP GPBLENDEQUATIONI)(GLuint  buf, GLenum  mode);
// typedef void  (APIENTRYP GPBLENDEQUATIONIARB)(GLuint  buf, GLenum  mode);
// typedef void  (APIENTRYP GPBLENDFUNC)(GLenum  sfactor, GLenum  dfactor);
// typedef void  (APIENTRYP GPBLENDFUNCSEPARATE)(GLenum  sfactorRGB, GLenum  dfactorRGB, GLenum  sfactorAlpha, GLenum  dfactorAlpha);
// typedef void  (APIENTRYP GPBLENDFUNCSEPARATEI)(GLuint  buf, GLenum  srcRGB, GLenum  dstRGB, GLenum  srcAlpha, GLenum  dstAlpha);
// typedef void  (APIENTRYP GPBLENDFUNCSEPARATEIARB)(GLuint  buf, GLenum  srcRGB, GLenum  dstRGB, GLenum  srcAlpha, GLenum  dstAlpha);
// typedef void  (APIENTRYP GPBLENDFUNCI)(GLuint  buf, GLenum  src, GLenum  dst);
// typedef void  (APIENTRYP GPBLENDFUNCIARB)(GLuint  buf, GLenum  src, GLenum  dst);
// typedef void  (APIENTRYP GPBLITFRAMEBUFFER)(GLint  srcX0, GLint  srcY0, GLint  srcX1, GLint  srcY1, GLint  dstX0, GLint  dstY0, GLint  dstX1, GLint  dstY1, GLbitfield  mask, GLenum  filter);
// typedef void  (APIENTRYP GPBLITNAMEDFRAMEBUFFER)(GLuint  readFramebuffer, GLuint  drawFramebuffer, GLint  srcX0, GLint  srcY0, GLint  srcX1, GLint  srcY1, GLint  dstX0, GLint  dstY0, GLint  dstX1, GLint  dstY1, GLbitfield  mask, GLenum  filter);
// typedef void  (APIENTRYP GPBUFFERDATA)(GLenum  target, GLsizeiptr  size, const void * data, GLenum  usage);
// typedef void  (APIENTRYP GPBUFFERPAGECOMMITMENTARB)(GLenum  target, GLintptr  offset, GLsizei  size, GLboolean  commit);
// typedef void  (APIENTRYP GPBUFFERSTORAGE)(GLenum  target, GLsizeiptr  size, const void * data, GLbitfield  flags);
// typedef void  (APIENTRYP GPBUFFERSUBDATA)(GLenum  target, GLintptr  offset, GLsizeiptr  size, const void * data);
// typedef void  (APIENTRYP GPCALLLIST)(GLuint  list);
// typedef void  (APIENTRYP GPCALLLISTS)(GLsizei  n, GLenum  type, const void * lists);
// typedef GLenum  (APIENTRYP GPCHECKFRAMEBUFFERSTATUS)(GLenum  target);
// typedef GLenum  (APIENTRYP GPCHECKNAMEDFRAMEBUFFERSTATUS)(GLuint  framebuffer, GLenum  target);
// typedef void  (APIENTRYP GPCLAMPCOLOR)(GLenum  target, GLenum  clamp);
// typedef void  (APIENTRYP GPCLEAR)(GLbitfield  mask);
// typedef void  (APIENTRYP GPCLEARACCUM)(GLfloat  red, GLfloat  green, GLfloat  blue, GLfloat  alpha);
// typedef void  (APIENTRYP GPCLEARBUFFERDATA)(GLenum  target, GLenum  internalformat, GLenum  format, GLenum  type, const void * data);
// typedef void  (APIENTRYP GPCLEARBUFFERSUBDATA)(GLenum  target, GLenum  internalformat, GLintptr  offset, GLsizeiptr  size, GLenum  format, GLenum  type, const void * data);
// typedef void  (APIENTRYP GPCLEARBUFFERFI)(GLenum  buffer, GLint  drawbuffer, GLfloat  depth, GLint  stencil);
// typedef void  (APIENTRYP GPCLEARBUFFERFV)(GLenum  buffer, GLint  drawbuffer, const GLfloat * value);
// typedef void  (APIENTRYP GPCLEARBUFFERIV)(GLenum  buffer, GLint  drawbuffer, const GLint * value);
// typedef void  (APIENTRYP GPCLEARBUFFERUIV)(GLenum  buffer, GLint  drawbuffer, const GLuint * value);
// typedef void  (APIENTRYP GPCLEARCOLOR)(GLfloat  red, GLfloat  green, GLfloat  blue, GLfloat  alpha);
// typedef void  (APIENTRYP GPCLEARDEPTH)(GLdouble  depth);
// typedef void  (APIENTRYP GPCLEARDEPTHF)(GLfloat  d);
// typedef void  (APIENTRYP GPCLEARINDEX)(GLfloat  c);
// typedef void  (APIENTRYP GPCLEARNAMEDBUFFERDATA)(GLuint  buffer, GLenum  internalformat, GLenum  format, GLenum  type, const void * data);
// typedef void  (APIENTRYP GPCLEARNAMEDBUFFERSUBDATA)(GLuint  buffer, GLenum  internalformat, GLintptr  offset, GLsizei  size, GLenum  format, GLenum  type, const void * data);
// typedef void  (APIENTRYP GPCLEARNAMEDFRAMEBUFFERFI)(GLuint  framebuffer, GLenum  buffer, const GLfloat  depth, GLint  stencil);
// typedef void  (APIENTRYP GPCLEARNAMEDFRAMEBUFFERFV)(GLuint  framebuffer, GLenum  buffer, GLint  drawbuffer, const GLfloat * value);
// typedef void  (APIENTRYP GPCLEARNAMEDFRAMEBUFFERIV)(GLuint  framebuffer, GLenum  buffer, GLint  drawbuffer, const GLint * value);
// typedef void  (APIENTRYP GPCLEARNAMEDFRAMEBUFFERUIV)(GLuint  framebuffer, GLenum  buffer, GLint  drawbuffer, const GLuint * value);
// typedef void  (APIENTRYP GPCLEARSTENCIL)(GLint  s);
// typedef void  (APIENTRYP GPCLEARTEXIMAGE)(GLuint  texture, GLint  level, GLenum  format, GLenum  type, const void * data);
// typedef void  (APIENTRYP GPCLEARTEXSUBIMAGE)(GLuint  texture, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth, GLenum  format, GLenum  type, const void * data);
// typedef void  (APIENTRYP GPCLIENTACTIVETEXTURE)(GLenum  texture);
// typedef GLenum  (APIENTRYP GPCLIENTWAITSYNC)(GLsync  sync, GLbitfield  flags, GLuint64  timeout);
// typedef void  (APIENTRYP GPCLIPCONTROL)(GLenum  origin, GLenum  depth);
// typedef void  (APIENTRYP GPCLIPPLANE)(GLenum  plane, const GLdouble * equation);
// typedef void  (APIENTRYP GPCOLOR3B)(GLbyte  red, GLbyte  green, GLbyte  blue);
// typedef void  (APIENTRYP GPCOLOR3BV)(const GLbyte * v);
// typedef void  (APIENTRYP GPCOLOR3D)(GLdouble  red, GLdouble  green, GLdouble  blue);
// typedef void  (APIENTRYP GPCOLOR3DV)(const GLdouble * v);
// typedef void  (APIENTRYP GPCOLOR3F)(GLfloat  red, GLfloat  green, GLfloat  blue);
// typedef void  (APIENTRYP GPCOLOR3FV)(const GLfloat * v);
// typedef void  (APIENTRYP GPCOLOR3I)(GLint  red, GLint  green, GLint  blue);
// typedef void  (APIENTRYP GPCOLOR3IV)(const GLint * v);
// typedef void  (APIENTRYP GPCOLOR3S)(GLshort  red, GLshort  green, GLshort  blue);
// typedef void  (APIENTRYP GPCOLOR3SV)(const GLshort * v);
// typedef void  (APIENTRYP GPCOLOR3UB)(GLubyte  red, GLubyte  green, GLubyte  blue);
// typedef void  (APIENTRYP GPCOLOR3UBV)(const GLubyte * v);
// typedef void  (APIENTRYP GPCOLOR3UI)(GLuint  red, GLuint  green, GLuint  blue);
// typedef void  (APIENTRYP GPCOLOR3UIV)(const GLuint * v);
// typedef void  (APIENTRYP GPCOLOR3US)(GLushort  red, GLushort  green, GLushort  blue);
// typedef void  (APIENTRYP GPCOLOR3USV)(const GLushort * v);
// typedef void  (APIENTRYP GPCOLOR4B)(GLbyte  red, GLbyte  green, GLbyte  blue, GLbyte  alpha);
// typedef void  (APIENTRYP GPCOLOR4BV)(const GLbyte * v);
// typedef void  (APIENTRYP GPCOLOR4D)(GLdouble  red, GLdouble  green, GLdouble  blue, GLdouble  alpha);
// typedef void  (APIENTRYP GPCOLOR4DV)(const GLdouble * v);
// typedef void  (APIENTRYP GPCOLOR4F)(GLfloat  red, GLfloat  green, GLfloat  blue, GLfloat  alpha);
// typedef void  (APIENTRYP GPCOLOR4FV)(const GLfloat * v);
// typedef void  (APIENTRYP GPCOLOR4I)(GLint  red, GLint  green, GLint  blue, GLint  alpha);
// typedef void  (APIENTRYP GPCOLOR4IV)(const GLint * v);
// typedef void  (APIENTRYP GPCOLOR4S)(GLshort  red, GLshort  green, GLshort  blue, GLshort  alpha);
// typedef void  (APIENTRYP GPCOLOR4SV)(const GLshort * v);
// typedef void  (APIENTRYP GPCOLOR4UB)(GLubyte  red, GLubyte  green, GLubyte  blue, GLubyte  alpha);
// typedef void  (APIENTRYP GPCOLOR4UBV)(const GLubyte * v);
// typedef void  (APIENTRYP GPCOLOR4UI)(GLuint  red, GLuint  green, GLuint  blue, GLuint  alpha);
// typedef void  (APIENTRYP GPCOLOR4UIV)(const GLuint * v);
// typedef void  (APIENTRYP GPCOLOR4US)(GLushort  red, GLushort  green, GLushort  blue, GLushort  alpha);
// typedef void  (APIENTRYP GPCOLOR4USV)(const GLushort * v);
// typedef void  (APIENTRYP GPCOLORMASK)(GLboolean  red, GLboolean  green, GLboolean  blue, GLboolean  alpha);
// typedef void  (APIENTRYP GPCOLORMASKI)(GLuint  index, GLboolean  r, GLboolean  g, GLboolean  b, GLboolean  a);
// typedef void  (APIENTRYP GPCOLORMATERIAL)(GLenum  face, GLenum  mode);
// typedef void  (APIENTRYP GPCOLORPOINTER)(GLint  size, GLenum  type, GLsizei  stride, const void * pointer);
// typedef void  (APIENTRYP GPCOMPILESHADER)(GLuint  shader);
// typedef void  (APIENTRYP GPCOMPILESHADERINCLUDEARB)(GLuint  shader, GLsizei  count, const GLchar *const* path, const GLint * length);
// typedef void  (APIENTRYP GPCOMPRESSEDTEXIMAGE1D)(GLenum  target, GLint  level, GLenum  internalformat, GLsizei  width, GLint  border, GLsizei  imageSize, const void * data);
// typedef void  (APIENTRYP GPCOMPRESSEDTEXIMAGE2D)(GLenum  target, GLint  level, GLenum  internalformat, GLsizei  width, GLsizei  height, GLint  border, GLsizei  imageSize, const void * data);
// typedef void  (APIENTRYP GPCOMPRESSEDTEXIMAGE3D)(GLenum  target, GLint  level, GLenum  internalformat, GLsizei  width, GLsizei  height, GLsizei  depth, GLint  border, GLsizei  imageSize, const void * data);
// typedef void  (APIENTRYP GPCOMPRESSEDTEXSUBIMAGE1D)(GLenum  target, GLint  level, GLint  xoffset, GLsizei  width, GLenum  format, GLsizei  imageSize, const void * data);
// typedef void  (APIENTRYP GPCOMPRESSEDTEXSUBIMAGE2D)(GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLsizei  width, GLsizei  height, GLenum  format, GLsizei  imageSize, const void * data);
// typedef void  (APIENTRYP GPCOMPRESSEDTEXSUBIMAGE3D)(GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth, GLenum  format, GLsizei  imageSize, const void * data);
// typedef void  (APIENTRYP GPCOMPRESSEDTEXTURESUBIMAGE1D)(GLuint  texture, GLint  level, GLint  xoffset, GLsizei  width, GLenum  format, GLsizei  imageSize, const void * data);
// typedef void  (APIENTRYP GPCOMPRESSEDTEXTURESUBIMAGE2D)(GLuint  texture, GLint  level, GLint  xoffset, GLint  yoffset, GLsizei  width, GLsizei  height, GLenum  format, GLsizei  imageSize, const void * data);
// typedef void  (APIENTRYP GPCOMPRESSEDTEXTURESUBIMAGE3D)(GLuint  texture, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth, GLenum  format, GLsizei  imageSize, const void * data);
// typedef void  (APIENTRYP GPCOPYBUFFERSUBDATA)(GLenum  readTarget, GLenum  writeTarget, GLintptr  readOffset, GLintptr  writeOffset, GLsizeiptr  size);
// typedef void  (APIENTRYP GPCOPYIMAGESUBDATA)(GLuint  srcName, GLenum  srcTarget, GLint  srcLevel, GLint  srcX, GLint  srcY, GLint  srcZ, GLuint  dstName, GLenum  dstTarget, GLint  dstLevel, GLint  dstX, GLint  dstY, GLint  dstZ, GLsizei  srcWidth, GLsizei  srcHeight, GLsizei  srcDepth);
// typedef void  (APIENTRYP GPCOPYNAMEDBUFFERSUBDATA)(GLuint  readBuffer, GLuint  writeBuffer, GLintptr  readOffset, GLintptr  writeOffset, GLsizei  size);
// typedef void  (APIENTRYP GPCOPYPIXELS)(GLint  x, GLint  y, GLsizei  width, GLsizei  height, GLenum  type);
// typedef void  (APIENTRYP GPCOPYTEXIMAGE1D)(GLenum  target, GLint  level, GLenum  internalformat, GLint  x, GLint  y, GLsizei  width, GLint  border);
// typedef void  (APIENTRYP GPCOPYTEXIMAGE2D)(GLenum  target, GLint  level, GLenum  internalformat, GLint  x, GLint  y, GLsizei  width, GLsizei  height, GLint  border);
// typedef void  (APIENTRYP GPCOPYTEXSUBIMAGE1D)(GLenum  target, GLint  level, GLint  xoffset, GLint  x, GLint  y, GLsizei  width);
// typedef void  (APIENTRYP GPCOPYTEXSUBIMAGE2D)(GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLint  x, GLint  y, GLsizei  width, GLsizei  height);
// typedef void  (APIENTRYP GPCOPYTEXSUBIMAGE3D)(GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLint  x, GLint  y, GLsizei  width, GLsizei  height);
// typedef void  (APIENTRYP GPCOPYTEXTURESUBIMAGE1D)(GLuint  texture, GLint  level, GLint  xoffset, GLint  x, GLint  y, GLsizei  width);
// typedef void  (APIENTRYP GPCOPYTEXTURESUBIMAGE2D)(GLuint  texture, GLint  level, GLint  xoffset, GLint  yoffset, GLint  x, GLint  y, GLsizei  width, GLsizei  height);
// typedef void  (APIENTRYP GPCOPYTEXTURESUBIMAGE3D)(GLuint  texture, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLint  x, GLint  y, GLsizei  width, GLsizei  height);
// typedef void  (APIENTRYP GPCREATEBUFFERS)(GLsizei  n, GLuint * buffers);
// typedef void  (APIENTRYP GPCREATEFRAMEBUFFERS)(GLsizei  n, GLuint * framebuffers);
// typedef GLuint  (APIENTRYP GPCREATEPROGRAM)();
// typedef void  (APIENTRYP GPCREATEPROGRAMPIPELINES)(GLsizei  n, GLuint * pipelines);
// typedef void  (APIENTRYP GPCREATEQUERIES)(GLenum  target, GLsizei  n, GLuint * ids);
// typedef void  (APIENTRYP GPCREATERENDERBUFFERS)(GLsizei  n, GLuint * renderbuffers);
// typedef void  (APIENTRYP GPCREATESAMPLERS)(GLsizei  n, GLuint * samplers);
// typedef GLuint  (APIENTRYP GPCREATESHADER)(GLenum  type);
// typedef GLuint  (APIENTRYP GPCREATESHADERPROGRAMV)(GLenum  type, GLsizei  count, const GLchar *const* strings);
// typedef GLsync  (APIENTRYP GPCREATESYNCFROMCLEVENTARB)(struct _cl_context * context, struct _cl_event * event, GLbitfield  flags);
// typedef void  (APIENTRYP GPCREATETEXTURES)(GLenum  target, GLsizei  n, GLuint * textures);
// typedef void  (APIENTRYP GPCREATETRANSFORMFEEDBACKS)(GLsizei  n, GLuint * ids);
// typedef void  (APIENTRYP GPCREATEVERTEXARRAYS)(GLsizei  n, GLuint * arrays);
// typedef void  (APIENTRYP GPCULLFACE)(GLenum  mode);
// typedef void  (APIENTRYP GPDEBUGMESSAGECALLBACK)(GLDEBUGPROC  callback, const void * userParam);
// typedef void  (APIENTRYP GPDEBUGMESSAGECALLBACKARB)(GLDEBUGPROCARB  callback, const void * userParam);
// typedef void  (APIENTRYP GPDEBUGMESSAGECALLBACKKHR)(GLDEBUGPROCKHR  callback, const void * userParam);
// typedef void  (APIENTRYP GPDEBUGMESSAGECONTROL)(GLenum  source, GLenum  type, GLenum  severity, GLsizei  count, const GLuint * ids, GLboolean  enabled);
// typedef void  (APIENTRYP GPDEBUGMESSAGECONTROLARB)(GLenum  source, GLenum  type, GLenum  severity, GLsizei  count, const GLuint * ids, GLboolean  enabled);
// typedef void  (APIENTRYP GPDEBUGMESSAGECONTROLKHR)(GLenum  source, GLenum  type, GLenum  severity, GLsizei  count, const GLuint * ids, GLboolean  enabled);
// typedef void  (APIENTRYP GPDEBUGMESSAGEINSERT)(GLenum  source, GLenum  type, GLuint  id, GLenum  severity, GLsizei  length, const GLchar * buf);
// typedef void  (APIENTRYP GPDEBUGMESSAGEINSERTARB)(GLenum  source, GLenum  type, GLuint  id, GLenum  severity, GLsizei  length, const GLchar * buf);
// typedef void  (APIENTRYP GPDEBUGMESSAGEINSERTKHR)(GLenum  source, GLenum  type, GLuint  id, GLenum  severity, GLsizei  length, const GLchar * buf);
// typedef void  (APIENTRYP GPDELETEBUFFERS)(GLsizei  n, const GLuint * buffers);
// typedef void  (APIENTRYP GPDELETEFRAMEBUFFERS)(GLsizei  n, const GLuint * framebuffers);
// typedef void  (APIENTRYP GPDELETELISTS)(GLuint  list, GLsizei  range);
// typedef void  (APIENTRYP GPDELETENAMEDSTRINGARB)(GLint  namelen, const GLchar * name);
// typedef void  (APIENTRYP GPDELETEPROGRAM)(GLuint  program);
// typedef void  (APIENTRYP GPDELETEPROGRAMPIPELINES)(GLsizei  n, const GLuint * pipelines);
// typedef void  (APIENTRYP GPDELETEQUERIES)(GLsizei  n, const GLuint * ids);
// typedef void  (APIENTRYP GPDELETERENDERBUFFERS)(GLsizei  n, const GLuint * renderbuffers);
// typedef void  (APIENTRYP GPDELETESAMPLERS)(GLsizei  count, const GLuint * samplers);
// typedef void  (APIENTRYP GPDELETESHADER)(GLuint  shader);
// typedef void  (APIENTRYP GPDELETESYNC)(GLsync  sync);
// typedef void  (APIENTRYP GPDELETETEXTURES)(GLsizei  n, const GLuint * textures);
// typedef void  (APIENTRYP GPDELETETRANSFORMFEEDBACKS)(GLsizei  n, const GLuint * ids);
// typedef void  (APIENTRYP GPDELETEVERTEXARRAYS)(GLsizei  n, const GLuint * arrays);
// typedef void  (APIENTRYP GPDEPTHFUNC)(GLenum  func);
// typedef void  (APIENTRYP GPDEPTHMASK)(GLboolean  flag);
// typedef void  (APIENTRYP GPDEPTHRANGE)(GLdouble  xnear, GLdouble  xfar);
// typedef void  (APIENTRYP GPDEPTHRANGEARRAYV)(GLuint  first, GLsizei  count, const GLdouble * v);
// typedef void  (APIENTRYP GPDEPTHRANGEINDEXED)(GLuint  index, GLdouble  n, GLdouble  f);
// typedef void  (APIENTRYP GPDEPTHRANGEF)(GLfloat  n, GLfloat  f);
// typedef void  (APIENTRYP GPDETACHSHADER)(GLuint  program, GLuint  shader);
// typedef void  (APIENTRYP GPDISABLE)(GLenum  cap);
// typedef void  (APIENTRYP GPDISABLECLIENTSTATE)(GLenum  array);
// typedef void  (APIENTRYP GPDISABLEVERTEXARRAYATTRIB)(GLuint  vaobj, GLuint  index);
// typedef void  (APIENTRYP GPDISABLEVERTEXATTRIBARRAY)(GLuint  index);
// typedef void  (APIENTRYP GPDISABLEI)(GLenum  target, GLuint  index);
// typedef void  (APIENTRYP GPDISPATCHCOMPUTE)(GLuint  num_groups_x, GLuint  num_groups_y, GLuint  num_groups_z);
// typedef void  (APIENTRYP GPDISPATCHCOMPUTEGROUPSIZEARB)(GLuint  num_groups_x, GLuint  num_groups_y, GLuint  num_groups_z, GLuint  group_size_x, GLuint  group_size_y, GLuint  group_size_z);
// typedef void  (APIENTRYP GPDISPATCHCOMPUTEINDIRECT)(GLintptr  indirect);
// typedef void  (APIENTRYP GPDRAWARRAYS)(GLenum  mode, GLint  first, GLsizei  count);
// typedef void  (APIENTRYP GPDRAWARRAYSINDIRECT)(GLenum  mode, const void * indirect);
// typedef void  (APIENTRYP GPDRAWARRAYSINSTANCED)(GLenum  mode, GLint  first, GLsizei  count, GLsizei  instancecount);
// typedef void  (APIENTRYP GPDRAWARRAYSINSTANCEDBASEINSTANCE)(GLenum  mode, GLint  first, GLsizei  count, GLsizei  instancecount, GLuint  baseinstance);
// typedef void  (APIENTRYP GPDRAWBUFFER)(GLenum  buf);
// typedef void  (APIENTRYP GPDRAWBUFFERS)(GLsizei  n, const GLenum * bufs);
// typedef void  (APIENTRYP GPDRAWELEMENTS)(GLenum  mode, GLsizei  count, GLenum  type, const void * indices);
// typedef void  (APIENTRYP GPDRAWELEMENTSBASEVERTEX)(GLenum  mode, GLsizei  count, GLenum  type, const void * indices, GLint  basevertex);
// typedef void  (APIENTRYP GPDRAWELEMENTSINDIRECT)(GLenum  mode, GLenum  type, const void * indirect);
// typedef void  (APIENTRYP GPDRAWELEMENTSINSTANCED)(GLenum  mode, GLsizei  count, GLenum  type, const void * indices, GLsizei  instancecount);
// typedef void  (APIENTRYP GPDRAWELEMENTSINSTANCEDBASEINSTANCE)(GLenum  mode, GLsizei  count, GLenum  type, const void * indices, GLsizei  instancecount, GLuint  baseinstance);
// typedef void  (APIENTRYP GPDRAWELEMENTSINSTANCEDBASEVERTEX)(GLenum  mode, GLsizei  count, GLenum  type, const void * indices, GLsizei  instancecount, GLint  basevertex);
// typedef void  (APIENTRYP GPDRAWELEMENTSINSTANCEDBASEVERTEXBASEINSTANCE)(GLenum  mode, GLsizei  count, GLenum  type, const void * indices, GLsizei  instancecount, GLint  basevertex, GLuint  baseinstance);
// typedef void  (APIENTRYP GPDRAWPIXELS)(GLsizei  width, GLsizei  height, GLenum  format, GLenum  type, const void * pixels);
// typedef void  (APIENTRYP GPDRAWRANGEELEMENTS)(GLenum  mode, GLuint  start, GLuint  end, GLsizei  count, GLenum  type, const void * indices);
// typedef void  (APIENTRYP GPDRAWRANGEELEMENTSBASEVERTEX)(GLenum  mode, GLuint  start, GLuint  end, GLsizei  count, GLenum  type, const void * indices, GLint  basevertex);
// typedef void  (APIENTRYP GPDRAWTRANSFORMFEEDBACK)(GLenum  mode, GLuint  id);
// typedef void  (APIENTRYP GPDRAWTRANSFORMFEEDBACKINSTANCED)(GLenum  mode, GLuint  id, GLsizei  instancecount);
// typedef void  (APIENTRYP GPDRAWTRANSFORMFEEDBACKSTREAM)(GLenum  mode, GLuint  id, GLuint  stream);
// typedef void  (APIENTRYP GPDRAWTRANSFORMFEEDBACKSTREAMINSTANCED)(GLenum  mode, GLuint  id, GLuint  stream, GLsizei  instancecount);
// typedef void  (APIENTRYP GPEDGEFLAG)(GLboolean  flag);
// typedef void  (APIENTRYP GPEDGEFLAGPOINTER)(GLsizei  stride, const void * pointer);
// typedef void  (APIENTRYP GPEDGEFLAGV)(const GLboolean * flag);
// typedef void  (APIENTRYP GPENABLE)(GLenum  cap);
// typedef void  (APIENTRYP GPENABLECLIENTSTATE)(GLenum  array);
// typedef void  (APIENTRYP GPENABLEVERTEXARRAYATTRIB)(GLuint  vaobj, GLuint  index);
// typedef void  (APIENTRYP GPENABLEVERTEXATTRIBARRAY)(GLuint  index);
// typedef void  (APIENTRYP GPENABLEI)(GLenum  target, GLuint  index);
// typedef void  (APIENTRYP GPEND)();
// typedef void  (APIENTRYP GPENDCONDITIONALRENDER)();
// typedef void  (APIENTRYP GPENDLIST)();
// typedef void  (APIENTRYP GPENDQUERY)(GLenum  target);
// typedef void  (APIENTRYP GPENDQUERYINDEXED)(GLenum  target, GLuint  index);
// typedef void  (APIENTRYP GPENDTRANSFORMFEEDBACK)();
// typedef void  (APIENTRYP GPEVALCOORD1D)(GLdouble  u);
// typedef void  (APIENTRYP GPEVALCOORD1DV)(const GLdouble * u);
// typedef void  (APIENTRYP GPEVALCOORD1F)(GLfloat  u);
// typedef void  (APIENTRYP GPEVALCOORD1FV)(const GLfloat * u);
// typedef void  (APIENTRYP GPEVALCOORD2D)(GLdouble  u, GLdouble  v);
// typedef void  (APIENTRYP GPEVALCOORD2DV)(const GLdouble * u);
// typedef void  (APIENTRYP GPEVALCOORD2F)(GLfloat  u, GLfloat  v);
// typedef void  (APIENTRYP GPEVALCOORD2FV)(const GLfloat * u);
// typedef void  (APIENTRYP GPEVALMESH1)(GLenum  mode, GLint  i1, GLint  i2);
// typedef void  (APIENTRYP GPEVALMESH2)(GLenum  mode, GLint  i1, GLint  i2, GLint  j1, GLint  j2);
// typedef void  (APIENTRYP GPEVALPOINT1)(GLint  i);
// typedef void  (APIENTRYP GPEVALPOINT2)(GLint  i, GLint  j);
// typedef void  (APIENTRYP GPFEEDBACKBUFFER)(GLsizei  size, GLenum  type, GLfloat * buffer);
// typedef GLsync  (APIENTRYP GPFENCESYNC)(GLenum  condition, GLbitfield  flags);
// typedef void  (APIENTRYP GPFINISH)();
// typedef void  (APIENTRYP GPFLUSH)();
// typedef void  (APIENTRYP GPFLUSHMAPPEDBUFFERRANGE)(GLenum  target, GLintptr  offset, GLsizeiptr  length);
// typedef void  (APIENTRYP GPFLUSHMAPPEDNAMEDBUFFERRANGE)(GLuint  buffer, GLintptr  offset, GLsizei  length);
// typedef void  (APIENTRYP GPFOGCOORDPOINTER)(GLenum  type, GLsizei  stride, const void * pointer);
// typedef void  (APIENTRYP GPFOGCOORDD)(GLdouble  coord);
// typedef void  (APIENTRYP GPFOGCOORDDV)(const GLdouble * coord);
// typedef void  (APIENTRYP GPFOGCOORDF)(GLfloat  coord);
// typedef void  (APIENTRYP GPFOGCOORDFV)(const GLfloat * coord);
// typedef void  (APIENTRYP GPFOGF)(GLenum  pname, GLfloat  param);
// typedef void  (APIENTRYP GPFOGFV)(GLenum  pname, const GLfloat * params);
// typedef void  (APIENTRYP GPFOGI)(GLenum  pname, GLint  param);
// typedef void  (APIENTRYP GPFOGIV)(GLenum  pname, const GLint * params);
// typedef void  (APIENTRYP GPFRAMEBUFFERPARAMETERI)(GLenum  target, GLenum  pname, GLint  param);
// typedef void  (APIENTRYP GPFRAMEBUFFERRENDERBUFFER)(GLenum  target, GLenum  attachment, GLenum  renderbuffertarget, GLuint  renderbuffer);
// typedef void  (APIENTRYP GPFRAMEBUFFERTEXTURE)(GLenum  target, GLenum  attachment, GLuint  texture, GLint  level);
// typedef void  (APIENTRYP GPFRAMEBUFFERTEXTURE1D)(GLenum  target, GLenum  attachment, GLenum  textarget, GLuint  texture, GLint  level);
// typedef void  (APIENTRYP GPFRAMEBUFFERTEXTURE2D)(GLenum  target, GLenum  attachment, GLenum  textarget, GLuint  texture, GLint  level);
// typedef void  (APIENTRYP GPFRAMEBUFFERTEXTURE3D)(GLenum  target, GLenum  attachment, GLenum  textarget, GLuint  texture, GLint  level, GLint  zoffset);
// typedef void  (APIENTRYP GPFRAMEBUFFERTEXTURELAYER)(GLenum  target, GLenum  attachment, GLuint  texture, GLint  level, GLint  layer);
// typedef void  (APIENTRYP GPFRONTFACE)(GLenum  mode);
// typedef void  (APIENTRYP GPFRUSTUM)(GLdouble  left, GLdouble  right, GLdouble  bottom, GLdouble  top, GLdouble  zNear, GLdouble  zFar);
// typedef void  (APIENTRYP GPGENBUFFERS)(GLsizei  n, GLuint * buffers);
// typedef void  (APIENTRYP GPGENFRAMEBUFFERS)(GLsizei  n, GLuint * framebuffers);
// typedef GLuint  (APIENTRYP GPGENLISTS)(GLsizei  range);
// typedef void  (APIENTRYP GPGENPROGRAMPIPELINES)(GLsizei  n, GLuint * pipelines);
// typedef void  (APIENTRYP GPGENQUERIES)(GLsizei  n, GLuint * ids);
// typedef void  (APIENTRYP GPGENRENDERBUFFERS)(GLsizei  n, GLuint * renderbuffers);
// typedef void  (APIENTRYP GPGENSAMPLERS)(GLsizei  count, GLuint * samplers);
// typedef void  (APIENTRYP GPGENTEXTURES)(GLsizei  n, GLuint * textures);
// typedef void  (APIENTRYP GPGENTRANSFORMFEEDBACKS)(GLsizei  n, GLuint * ids);
// typedef void  (APIENTRYP GPGENVERTEXARRAYS)(GLsizei  n, GLuint * arrays);
// typedef void  (APIENTRYP GPGENERATEMIPMAP)(GLenum  target);
// typedef void  (APIENTRYP GPGENERATETEXTUREMIPMAP)(GLuint  texture);
// typedef void  (APIENTRYP GPGETACTIVEATOMICCOUNTERBUFFERIV)(GLuint  program, GLuint  bufferIndex, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETACTIVEATTRIB)(GLuint  program, GLuint  index, GLsizei  bufSize, GLsizei * length, GLint * size, GLenum * type, GLchar * name);
// typedef void  (APIENTRYP GPGETACTIVESUBROUTINENAME)(GLuint  program, GLenum  shadertype, GLuint  index, GLsizei  bufsize, GLsizei * length, GLchar * name);
// typedef void  (APIENTRYP GPGETACTIVESUBROUTINEUNIFORMNAME)(GLuint  program, GLenum  shadertype, GLuint  index, GLsizei  bufsize, GLsizei * length, GLchar * name);
// typedef void  (APIENTRYP GPGETACTIVESUBROUTINEUNIFORMIV)(GLuint  program, GLenum  shadertype, GLuint  index, GLenum  pname, GLint * values);
// typedef void  (APIENTRYP GPGETACTIVEUNIFORM)(GLuint  program, GLuint  index, GLsizei  bufSize, GLsizei * length, GLint * size, GLenum * type, GLchar * name);
// typedef void  (APIENTRYP GPGETACTIVEUNIFORMBLOCKNAME)(GLuint  program, GLuint  uniformBlockIndex, GLsizei  bufSize, GLsizei * length, GLchar * uniformBlockName);
// typedef void  (APIENTRYP GPGETACTIVEUNIFORMBLOCKIV)(GLuint  program, GLuint  uniformBlockIndex, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETACTIVEUNIFORMNAME)(GLuint  program, GLuint  uniformIndex, GLsizei  bufSize, GLsizei * length, GLchar * uniformName);
// typedef void  (APIENTRYP GPGETACTIVEUNIFORMSIV)(GLuint  program, GLsizei  uniformCount, const GLuint * uniformIndices, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETATTACHEDSHADERS)(GLuint  program, GLsizei  maxCount, GLsizei * count, GLuint * shaders);
// typedef GLint  (APIENTRYP GPGETATTRIBLOCATION)(GLuint  program, const GLchar * name);
// typedef void  (APIENTRYP GPGETBOOLEANI_V)(GLenum  target, GLuint  index, GLboolean * data);
// typedef void  (APIENTRYP GPGETBOOLEANV)(GLenum  pname, GLboolean * data);
// typedef void  (APIENTRYP GPGETBUFFERPARAMETERI64V)(GLenum  target, GLenum  pname, GLint64 * params);
// typedef void  (APIENTRYP GPGETBUFFERPARAMETERIV)(GLenum  target, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETBUFFERPOINTERV)(GLenum  target, GLenum  pname, void ** params);
// typedef void  (APIENTRYP GPGETBUFFERSUBDATA)(GLenum  target, GLintptr  offset, GLsizeiptr  size, void * data);
// typedef void  (APIENTRYP GPGETCLIPPLANE)(GLenum  plane, GLdouble * equation);
// typedef void  (APIENTRYP GPGETCOMPRESSEDTEXIMAGE)(GLenum  target, GLint  level, void * img);
// typedef void  (APIENTRYP GPGETCOMPRESSEDTEXTUREIMAGE)(GLuint  texture, GLint  level, GLsizei  bufSize, void * pixels);
// typedef void  (APIENTRYP GPGETCOMPRESSEDTEXTURESUBIMAGE)(GLuint  texture, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth, GLsizei  bufSize, void * pixels);
// typedef GLuint  (APIENTRYP GPGETDEBUGMESSAGELOG)(GLuint  count, GLsizei  bufSize, GLenum * sources, GLenum * types, GLuint * ids, GLenum * severities, GLsizei * lengths, GLchar * messageLog);
// typedef GLuint  (APIENTRYP GPGETDEBUGMESSAGELOGARB)(GLuint  count, GLsizei  bufSize, GLenum * sources, GLenum * types, GLuint * ids, GLenum * severities, GLsizei * lengths, GLchar * messageLog);
// typedef GLuint  (APIENTRYP GPGETDEBUGMESSAGELOGKHR)(GLuint  count, GLsizei  bufSize, GLenum * sources, GLenum * types, GLuint * ids, GLenum * severities, GLsizei * lengths, GLchar * messageLog);
// typedef void  (APIENTRYP GPGETDOUBLEI_V)(GLenum  target, GLuint  index, GLdouble * data);
// typedef void  (APIENTRYP GPGETDOUBLEV)(GLenum  pname, GLdouble * data);
// typedef GLenum  (APIENTRYP GPGETERROR)();
// typedef void  (APIENTRYP GPGETFLOATI_V)(GLenum  target, GLuint  index, GLfloat * data);
// typedef void  (APIENTRYP GPGETFLOATV)(GLenum  pname, GLfloat * data);
// typedef GLint  (APIENTRYP GPGETFRAGDATAINDEX)(GLuint  program, const GLchar * name);
// typedef GLint  (APIENTRYP GPGETFRAGDATALOCATION)(GLuint  program, const GLchar * name);
// typedef void  (APIENTRYP GPGETFRAMEBUFFERATTACHMENTPARAMETERIV)(GLenum  target, GLenum  attachment, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETFRAMEBUFFERPARAMETERIV)(GLenum  target, GLenum  pname, GLint * params);
// typedef GLenum  (APIENTRYP GPGETGRAPHICSRESETSTATUS)();
// typedef GLenum  (APIENTRYP GPGETGRAPHICSRESETSTATUSARB)();
// typedef GLenum  (APIENTRYP GPGETGRAPHICSRESETSTATUSKHR)();
// typedef GLuint64  (APIENTRYP GPGETIMAGEHANDLEARB)(GLuint  texture, GLint  level, GLboolean  layered, GLint  layer, GLenum  format);
// typedef void  (APIENTRYP GPGETINTEGER64I_V)(GLenum  target, GLuint  index, GLint64 * data);
// typedef void  (APIENTRYP GPGETINTEGER64V)(GLenum  pname, GLint64 * data);
// typedef void  (APIENTRYP GPGETINTEGERI_V)(GLenum  target, GLuint  index, GLint * data);
// typedef void  (APIENTRYP GPGETINTEGERV)(GLenum  pname, GLint * data);
// typedef void  (APIENTRYP GPGETINTERNALFORMATI64V)(GLenum  target, GLenum  internalformat, GLenum  pname, GLsizei  bufSize, GLint64 * params);
// typedef void  (APIENTRYP GPGETINTERNALFORMATIV)(GLenum  target, GLenum  internalformat, GLenum  pname, GLsizei  bufSize, GLint * params);
// typedef void  (APIENTRYP GPGETLIGHTFV)(GLenum  light, GLenum  pname, GLfloat * params);
// typedef void  (APIENTRYP GPGETLIGHTIV)(GLenum  light, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETMAPDV)(GLenum  target, GLenum  query, GLdouble * v);
// typedef void  (APIENTRYP GPGETMAPFV)(GLenum  target, GLenum  query, GLfloat * v);
// typedef void  (APIENTRYP GPGETMAPIV)(GLenum  target, GLenum  query, GLint * v);
// typedef void  (APIENTRYP GPGETMATERIALFV)(GLenum  face, GLenum  pname, GLfloat * params);
// typedef void  (APIENTRYP GPGETMATERIALIV)(GLenum  face, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETMULTISAMPLEFV)(GLenum  pname, GLuint  index, GLfloat * val);
// typedef void  (APIENTRYP GPGETNAMEDBUFFERPARAMETERI64V)(GLuint  buffer, GLenum  pname, GLint64 * params);
// typedef void  (APIENTRYP GPGETNAMEDBUFFERPARAMETERIV)(GLuint  buffer, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETNAMEDBUFFERPOINTERV)(GLuint  buffer, GLenum  pname, void ** params);
// typedef void  (APIENTRYP GPGETNAMEDBUFFERSUBDATA)(GLuint  buffer, GLintptr  offset, GLsizei  size, void * data);
// typedef void  (APIENTRYP GPGETNAMEDFRAMEBUFFERATTACHMENTPARAMETERIV)(GLuint  framebuffer, GLenum  attachment, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETNAMEDFRAMEBUFFERPARAMETERIV)(GLuint  framebuffer, GLenum  pname, GLint * param);
// typedef void  (APIENTRYP GPGETNAMEDRENDERBUFFERPARAMETERIV)(GLuint  renderbuffer, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETNAMEDSTRINGARB)(GLint  namelen, const GLchar * name, GLsizei  bufSize, GLint * stringlen, GLchar * string);
// typedef void  (APIENTRYP GPGETNAMEDSTRINGIVARB)(GLint  namelen, const GLchar * name, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETOBJECTLABEL)(GLenum  identifier, GLuint  name, GLsizei  bufSize, GLsizei * length, GLchar * label);
// typedef void  (APIENTRYP GPGETOBJECTLABELKHR)(GLenum  identifier, GLuint  name, GLsizei  bufSize, GLsizei * length, GLchar * label);
// typedef void  (APIENTRYP GPGETOBJECTPTRLABEL)(const void * ptr, GLsizei  bufSize, GLsizei * length, GLchar * label);
// typedef void  (APIENTRYP GPGETOBJECTPTRLABELKHR)(const void * ptr, GLsizei  bufSize, GLsizei * length, GLchar * label);
// typedef void  (APIENTRYP GPGETPIXELMAPFV)(GLenum  map, GLfloat * values);
// typedef void  (APIENTRYP GPGETPIXELMAPUIV)(GLenum  map, GLuint * values);
// typedef void  (APIENTRYP GPGETPIXELMAPUSV)(GLenum  map, GLushort * values);
// typedef void  (APIENTRYP GPGETPOINTERV)(GLenum  pname, void ** params);
// typedef void  (APIENTRYP GPGETPOINTERVKHR)(GLenum  pname, void ** params);
// typedef void  (APIENTRYP GPGETPOLYGONSTIPPLE)(GLubyte * mask);
// typedef void  (APIENTRYP GPGETPROGRAMBINARY)(GLuint  program, GLsizei  bufSize, GLsizei * length, GLenum * binaryFormat, void * binary);
// typedef void  (APIENTRYP GPGETPROGRAMINFOLOG)(GLuint  program, GLsizei  bufSize, GLsizei * length, GLchar * infoLog);
// typedef void  (APIENTRYP GPGETPROGRAMINTERFACEIV)(GLuint  program, GLenum  programInterface, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETPROGRAMPIPELINEINFOLOG)(GLuint  pipeline, GLsizei  bufSize, GLsizei * length, GLchar * infoLog);
// typedef void  (APIENTRYP GPGETPROGRAMPIPELINEIV)(GLuint  pipeline, GLenum  pname, GLint * params);
// typedef GLuint  (APIENTRYP GPGETPROGRAMRESOURCEINDEX)(GLuint  program, GLenum  programInterface, const GLchar * name);
// typedef GLint  (APIENTRYP GPGETPROGRAMRESOURCELOCATION)(GLuint  program, GLenum  programInterface, const GLchar * name);
// typedef GLint  (APIENTRYP GPGETPROGRAMRESOURCELOCATIONINDEX)(GLuint  program, GLenum  programInterface, const GLchar * name);
// typedef void  (APIENTRYP GPGETPROGRAMRESOURCENAME)(GLuint  program, GLenum  programInterface, GLuint  index, GLsizei  bufSize, GLsizei * length, GLchar * name);
// typedef void  (APIENTRYP GPGETPROGRAMRESOURCEIV)(GLuint  program, GLenum  programInterface, GLuint  index, GLsizei  propCount, const GLenum * props, GLsizei  bufSize, GLsizei * length, GLint * params);
// typedef void  (APIENTRYP GPGETPROGRAMSTAGEIV)(GLuint  program, GLenum  shadertype, GLenum  pname, GLint * values);
// typedef void  (APIENTRYP GPGETPROGRAMIV)(GLuint  program, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETQUERYINDEXEDIV)(GLenum  target, GLuint  index, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETQUERYOBJECTI64V)(GLuint  id, GLenum  pname, GLint64 * params);
// typedef void  (APIENTRYP GPGETQUERYOBJECTIV)(GLuint  id, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETQUERYOBJECTUI64V)(GLuint  id, GLenum  pname, GLuint64 * params);
// typedef void  (APIENTRYP GPGETQUERYOBJECTUIV)(GLuint  id, GLenum  pname, GLuint * params);
// typedef void  (APIENTRYP GPGETQUERYIV)(GLenum  target, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETRENDERBUFFERPARAMETERIV)(GLenum  target, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETSAMPLERPARAMETERIIV)(GLuint  sampler, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETSAMPLERPARAMETERIUIV)(GLuint  sampler, GLenum  pname, GLuint * params);
// typedef void  (APIENTRYP GPGETSAMPLERPARAMETERFV)(GLuint  sampler, GLenum  pname, GLfloat * params);
// typedef void  (APIENTRYP GPGETSAMPLERPARAMETERIV)(GLuint  sampler, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETSHADERINFOLOG)(GLuint  shader, GLsizei  bufSize, GLsizei * length, GLchar * infoLog);
// typedef void  (APIENTRYP GPGETSHADERPRECISIONFORMAT)(GLenum  shadertype, GLenum  precisiontype, GLint * range, GLint * precision);
// typedef void  (APIENTRYP GPGETSHADERSOURCE)(GLuint  shader, GLsizei  bufSize, GLsizei * length, GLchar * source);
// typedef void  (APIENTRYP GPGETSHADERIV)(GLuint  shader, GLenum  pname, GLint * params);
// typedef const GLubyte * (APIENTRYP GPGETSTRING)(GLenum  name);
// typedef const GLubyte * (APIENTRYP GPGETSTRINGI)(GLenum  name, GLuint  index);
// typedef GLuint  (APIENTRYP GPGETSUBROUTINEINDEX)(GLuint  program, GLenum  shadertype, const GLchar * name);
// typedef GLint  (APIENTRYP GPGETSUBROUTINEUNIFORMLOCATION)(GLuint  program, GLenum  shadertype, const GLchar * name);
// typedef void  (APIENTRYP GPGETSYNCIV)(GLsync  sync, GLenum  pname, GLsizei  bufSize, GLsizei * length, GLint * values);
// typedef void  (APIENTRYP GPGETTEXENVFV)(GLenum  target, GLenum  pname, GLfloat * params);
// typedef void  (APIENTRYP GPGETTEXENVIV)(GLenum  target, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETTEXGENDV)(GLenum  coord, GLenum  pname, GLdouble * params);
// typedef void  (APIENTRYP GPGETTEXGENFV)(GLenum  coord, GLenum  pname, GLfloat * params);
// typedef void  (APIENTRYP GPGETTEXGENIV)(GLenum  coord, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETTEXIMAGE)(GLenum  target, GLint  level, GLenum  format, GLenum  type, void * pixels);
// typedef void  (APIENTRYP GPGETTEXLEVELPARAMETERFV)(GLenum  target, GLint  level, GLenum  pname, GLfloat * params);
// typedef void  (APIENTRYP GPGETTEXLEVELPARAMETERIV)(GLenum  target, GLint  level, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETTEXPARAMETERIIV)(GLenum  target, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETTEXPARAMETERIUIV)(GLenum  target, GLenum  pname, GLuint * params);
// typedef void  (APIENTRYP GPGETTEXPARAMETERFV)(GLenum  target, GLenum  pname, GLfloat * params);
// typedef void  (APIENTRYP GPGETTEXPARAMETERIV)(GLenum  target, GLenum  pname, GLint * params);
// typedef GLuint64  (APIENTRYP GPGETTEXTUREHANDLEARB)(GLuint  texture);
// typedef void  (APIENTRYP GPGETTEXTUREIMAGE)(GLuint  texture, GLint  level, GLenum  format, GLenum  type, GLsizei  bufSize, void * pixels);
// typedef void  (APIENTRYP GPGETTEXTURELEVELPARAMETERFV)(GLuint  texture, GLint  level, GLenum  pname, GLfloat * params);
// typedef void  (APIENTRYP GPGETTEXTURELEVELPARAMETERIV)(GLuint  texture, GLint  level, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETTEXTUREPARAMETERIIV)(GLuint  texture, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETTEXTUREPARAMETERIUIV)(GLuint  texture, GLenum  pname, GLuint * params);
// typedef void  (APIENTRYP GPGETTEXTUREPARAMETERFV)(GLuint  texture, GLenum  pname, GLfloat * params);
// typedef void  (APIENTRYP GPGETTEXTUREPARAMETERIV)(GLuint  texture, GLenum  pname, GLint * params);
// typedef GLuint64  (APIENTRYP GPGETTEXTURESAMPLERHANDLEARB)(GLuint  texture, GLuint  sampler);
// typedef void  (APIENTRYP GPGETTEXTURESUBIMAGE)(GLuint  texture, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth, GLenum  format, GLenum  type, GLsizei  bufSize, void * pixels);
// typedef void  (APIENTRYP GPGETTRANSFORMFEEDBACKVARYING)(GLuint  program, GLuint  index, GLsizei  bufSize, GLsizei * length, GLsizei * size, GLenum * type, GLchar * name);
// typedef void  (APIENTRYP GPGETTRANSFORMFEEDBACKI64_V)(GLuint  xfb, GLenum  pname, GLuint  index, GLint64 * param);
// typedef void  (APIENTRYP GPGETTRANSFORMFEEDBACKI_V)(GLuint  xfb, GLenum  pname, GLuint  index, GLint * param);
// typedef void  (APIENTRYP GPGETTRANSFORMFEEDBACKIV)(GLuint  xfb, GLenum  pname, GLint * param);
// typedef GLuint  (APIENTRYP GPGETUNIFORMBLOCKINDEX)(GLuint  program, const GLchar * uniformBlockName);
// typedef void  (APIENTRYP GPGETUNIFORMINDICES)(GLuint  program, GLsizei  uniformCount, const GLchar *const* uniformNames, GLuint * uniformIndices);
// typedef GLint  (APIENTRYP GPGETUNIFORMLOCATION)(GLuint  program, const GLchar * name);
// typedef void  (APIENTRYP GPGETUNIFORMSUBROUTINEUIV)(GLenum  shadertype, GLint  location, GLuint * params);
// typedef void  (APIENTRYP GPGETUNIFORMDV)(GLuint  program, GLint  location, GLdouble * params);
// typedef void  (APIENTRYP GPGETUNIFORMFV)(GLuint  program, GLint  location, GLfloat * params);
// typedef void  (APIENTRYP GPGETUNIFORMIV)(GLuint  program, GLint  location, GLint * params);
// typedef void  (APIENTRYP GPGETUNIFORMUIV)(GLuint  program, GLint  location, GLuint * params);
// typedef void  (APIENTRYP GPGETVERTEXARRAYINDEXED64IV)(GLuint  vaobj, GLuint  index, GLenum  pname, GLint64 * param);
// typedef void  (APIENTRYP GPGETVERTEXARRAYINDEXEDIV)(GLuint  vaobj, GLuint  index, GLenum  pname, GLint * param);
// typedef void  (APIENTRYP GPGETVERTEXARRAYIV)(GLuint  vaobj, GLenum  pname, GLint * param);
// typedef void  (APIENTRYP GPGETVERTEXATTRIBIIV)(GLuint  index, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETVERTEXATTRIBIUIV)(GLuint  index, GLenum  pname, GLuint * params);
// typedef void  (APIENTRYP GPGETVERTEXATTRIBLDV)(GLuint  index, GLenum  pname, GLdouble * params);
// typedef void  (APIENTRYP GPGETVERTEXATTRIBLUI64VARB)(GLuint  index, GLenum  pname, GLuint64EXT * params);
// typedef void  (APIENTRYP GPGETVERTEXATTRIBPOINTERV)(GLuint  index, GLenum  pname, void ** pointer);
// typedef void  (APIENTRYP GPGETVERTEXATTRIBDV)(GLuint  index, GLenum  pname, GLdouble * params);
// typedef void  (APIENTRYP GPGETVERTEXATTRIBFV)(GLuint  index, GLenum  pname, GLfloat * params);
// typedef void  (APIENTRYP GPGETVERTEXATTRIBIV)(GLuint  index, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETNCOMPRESSEDTEXIMAGE)(GLenum  target, GLint  lod, GLsizei  bufSize, void * pixels);
// typedef void  (APIENTRYP GPGETNCOMPRESSEDTEXIMAGEARB)(GLenum  target, GLint  lod, GLsizei  bufSize, void * img);
// typedef void  (APIENTRYP GPGETNTEXIMAGE)(GLenum  target, GLint  level, GLenum  format, GLenum  type, GLsizei  bufSize, void * pixels);
// typedef void  (APIENTRYP GPGETNTEXIMAGEARB)(GLenum  target, GLint  level, GLenum  format, GLenum  type, GLsizei  bufSize, void * img);
// typedef void  (APIENTRYP GPGETNUNIFORMDV)(GLuint  program, GLint  location, GLsizei  bufSize, GLdouble * params);
// typedef void  (APIENTRYP GPGETNUNIFORMDVARB)(GLuint  program, GLint  location, GLsizei  bufSize, GLdouble * params);
// typedef void  (APIENTRYP GPGETNUNIFORMFV)(GLuint  program, GLint  location, GLsizei  bufSize, GLfloat * params);
// typedef void  (APIENTRYP GPGETNUNIFORMFVARB)(GLuint  program, GLint  location, GLsizei  bufSize, GLfloat * params);
// typedef void  (APIENTRYP GPGETNUNIFORMFVKHR)(GLuint  program, GLint  location, GLsizei  bufSize, GLfloat * params);
// typedef void  (APIENTRYP GPGETNUNIFORMIV)(GLuint  program, GLint  location, GLsizei  bufSize, GLint * params);
// typedef void  (APIENTRYP GPGETNUNIFORMIVARB)(GLuint  program, GLint  location, GLsizei  bufSize, GLint * params);
// typedef void  (APIENTRYP GPGETNUNIFORMIVKHR)(GLuint  program, GLint  location, GLsizei  bufSize, GLint * params);
// typedef void  (APIENTRYP GPGETNUNIFORMUIV)(GLuint  program, GLint  location, GLsizei  bufSize, GLuint * params);
// typedef void  (APIENTRYP GPGETNUNIFORMUIVARB)(GLuint  program, GLint  location, GLsizei  bufSize, GLuint * params);
// typedef void  (APIENTRYP GPGETNUNIFORMUIVKHR)(GLuint  program, GLint  location, GLsizei  bufSize, GLuint * params);
// typedef void  (APIENTRYP GPHINT)(GLenum  target, GLenum  mode);
// typedef void  (APIENTRYP GPINDEXMASK)(GLuint  mask);
// typedef void  (APIENTRYP GPINDEXPOINTER)(GLenum  type, GLsizei  stride, const void * pointer);
// typedef void  (APIENTRYP GPINDEXD)(GLdouble  c);
// typedef void  (APIENTRYP GPINDEXDV)(const GLdouble * c);
// typedef void  (APIENTRYP GPINDEXF)(GLfloat  c);
// typedef void  (APIENTRYP GPINDEXFV)(const GLfloat * c);
// typedef void  (APIENTRYP GPINDEXI)(GLint  c);
// typedef void  (APIENTRYP GPINDEXIV)(const GLint * c);
// typedef void  (APIENTRYP GPINDEXS)(GLshort  c);
// typedef void  (APIENTRYP GPINDEXSV)(const GLshort * c);
// typedef void  (APIENTRYP GPINDEXUB)(GLubyte  c);
// typedef void  (APIENTRYP GPINDEXUBV)(const GLubyte * c);
// typedef void  (APIENTRYP GPINITNAMES)();
// typedef void  (APIENTRYP GPINTERLEAVEDARRAYS)(GLenum  format, GLsizei  stride, const void * pointer);
// typedef void  (APIENTRYP GPINVALIDATEBUFFERDATA)(GLuint  buffer);
// typedef void  (APIENTRYP GPINVALIDATEBUFFERSUBDATA)(GLuint  buffer, GLintptr  offset, GLsizeiptr  length);
// typedef void  (APIENTRYP GPINVALIDATEFRAMEBUFFER)(GLenum  target, GLsizei  numAttachments, const GLenum * attachments);
// typedef void  (APIENTRYP GPINVALIDATENAMEDFRAMEBUFFERDATA)(GLuint  framebuffer, GLsizei  numAttachments, const GLenum * attachments);
// typedef void  (APIENTRYP GPINVALIDATENAMEDFRAMEBUFFERSUBDATA)(GLuint  framebuffer, GLsizei  numAttachments, const GLenum * attachments, GLint  x, GLint  y, GLsizei  width, GLsizei  height);
// typedef void  (APIENTRYP GPINVALIDATESUBFRAMEBUFFER)(GLenum  target, GLsizei  numAttachments, const GLenum * attachments, GLint  x, GLint  y, GLsizei  width, GLsizei  height);
// typedef void  (APIENTRYP GPINVALIDATETEXIMAGE)(GLuint  texture, GLint  level);
// typedef void  (APIENTRYP GPINVALIDATETEXSUBIMAGE)(GLuint  texture, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth);
// typedef GLboolean  (APIENTRYP GPISBUFFER)(GLuint  buffer);
// typedef GLboolean  (APIENTRYP GPISENABLED)(GLenum  cap);
// typedef GLboolean  (APIENTRYP GPISENABLEDI)(GLenum  target, GLuint  index);
// typedef GLboolean  (APIENTRYP GPISFRAMEBUFFER)(GLuint  framebuffer);
// typedef GLboolean  (APIENTRYP GPISIMAGEHANDLERESIDENTARB)(GLuint64  handle);
// typedef GLboolean  (APIENTRYP GPISLIST)(GLuint  list);
// typedef GLboolean  (APIENTRYP GPISNAMEDSTRINGARB)(GLint  namelen, const GLchar * name);
// typedef GLboolean  (APIENTRYP GPISPROGRAM)(GLuint  program);
// typedef GLboolean  (APIENTRYP GPISPROGRAMPIPELINE)(GLuint  pipeline);
// typedef GLboolean  (APIENTRYP GPISQUERY)(GLuint  id);
// typedef GLboolean  (APIENTRYP GPISRENDERBUFFER)(GLuint  renderbuffer);
// typedef GLboolean  (APIENTRYP GPISSAMPLER)(GLuint  sampler);
// typedef GLboolean  (APIENTRYP GPISSHADER)(GLuint  shader);
// typedef GLboolean  (APIENTRYP GPISSYNC)(GLsync  sync);
// typedef GLboolean  (APIENTRYP GPISTEXTURE)(GLuint  texture);
// typedef GLboolean  (APIENTRYP GPISTEXTUREHANDLERESIDENTARB)(GLuint64  handle);
// typedef GLboolean  (APIENTRYP GPISTRANSFORMFEEDBACK)(GLuint  id);
// typedef GLboolean  (APIENTRYP GPISVERTEXARRAY)(GLuint  array);
// typedef void  (APIENTRYP GPLIGHTMODELF)(GLenum  pname, GLfloat  param);
// typedef void  (APIENTRYP GPLIGHTMODELFV)(GLenum  pname, const GLfloat * params);
// typedef void  (APIENTRYP GPLIGHTMODELI)(GLenum  pname, GLint  param);
// typedef void  (APIENTRYP GPLIGHTMODELIV)(GLenum  pname, const GLint * params);
// typedef void  (APIENTRYP GPLIGHTF)(GLenum  light, GLenum  pname, GLfloat  param);
// typedef void  (APIENTRYP GPLIGHTFV)(GLenum  light, GLenum  pname, const GLfloat * params);
// typedef void  (APIENTRYP GPLIGHTI)(GLenum  light, GLenum  pname, GLint  param);
// typedef void  (APIENTRYP GPLIGHTIV)(GLenum  light, GLenum  pname, const GLint * params);
// typedef void  (APIENTRYP GPLINESTIPPLE)(GLint  factor, GLushort  pattern);
// typedef void  (APIENTRYP GPLINEWIDTH)(GLfloat  width);
// typedef void  (APIENTRYP GPLINKPROGRAM)(GLuint  program);
// typedef void  (APIENTRYP GPLISTBASE)(GLuint  base);
// typedef void  (APIENTRYP GPLOADIDENTITY)();
// typedef void  (APIENTRYP GPLOADMATRIXD)(const GLdouble * m);
// typedef void  (APIENTRYP GPLOADMATRIXF)(const GLfloat * m);
// typedef void  (APIENTRYP GPLOADNAME)(GLuint  name);
// typedef void  (APIENTRYP GPLOADTRANSPOSEMATRIXD)(const GLdouble * m);
// typedef void  (APIENTRYP GPLOADTRANSPOSEMATRIXF)(const GLfloat * m);
// typedef void  (APIENTRYP GPLOGICOP)(GLenum  opcode);
// typedef void  (APIENTRYP GPMAKEIMAGEHANDLENONRESIDENTARB)(GLuint64  handle);
// typedef void  (APIENTRYP GPMAKEIMAGEHANDLERESIDENTARB)(GLuint64  handle, GLenum  access);
// typedef void  (APIENTRYP GPMAKETEXTUREHANDLENONRESIDENTARB)(GLuint64  handle);
// typedef void  (APIENTRYP GPMAKETEXTUREHANDLERESIDENTARB)(GLuint64  handle);
// typedef void  (APIENTRYP GPMAP1D)(GLenum  target, GLdouble  u1, GLdouble  u2, GLint  stride, GLint  order, const GLdouble * points);
// typedef void  (APIENTRYP GPMAP1F)(GLenum  target, GLfloat  u1, GLfloat  u2, GLint  stride, GLint  order, const GLfloat * points);
// typedef void  (APIENTRYP GPMAP2D)(GLenum  target, GLdouble  u1, GLdouble  u2, GLint  ustride, GLint  uorder, GLdouble  v1, GLdouble  v2, GLint  vstride, GLint  vorder, const GLdouble * points);
// typedef void  (APIENTRYP GPMAP2F)(GLenum  target, GLfloat  u1, GLfloat  u2, GLint  ustride, GLint  uorder, GLfloat  v1, GLfloat  v2, GLint  vstride, GLint  vorder, const GLfloat * points);
// typedef void * (APIENTRYP GPMAPBUFFER)(GLenum  target, GLenum  access);
// typedef void * (APIENTRYP GPMAPBUFFERRANGE)(GLenum  target, GLintptr  offset, GLsizeiptr  length, GLbitfield  access);
// typedef void  (APIENTRYP GPMAPGRID1D)(GLint  un, GLdouble  u1, GLdouble  u2);
// typedef void  (APIENTRYP GPMAPGRID1F)(GLint  un, GLfloat  u1, GLfloat  u2);
// typedef void  (APIENTRYP GPMAPGRID2D)(GLint  un, GLdouble  u1, GLdouble  u2, GLint  vn, GLdouble  v1, GLdouble  v2);
// typedef void  (APIENTRYP GPMAPGRID2F)(GLint  un, GLfloat  u1, GLfloat  u2, GLint  vn, GLfloat  v1, GLfloat  v2);
// typedef void * (APIENTRYP GPMAPNAMEDBUFFER)(GLuint  buffer, GLenum  access);
// typedef void * (APIENTRYP GPMAPNAMEDBUFFERRANGE)(GLuint  buffer, GLintptr  offset, GLsizei  length, GLbitfield  access);
// typedef void  (APIENTRYP GPMATERIALF)(GLenum  face, GLenum  pname, GLfloat  param);
// typedef void  (APIENTRYP GPMATERIALFV)(GLenum  face, GLenum  pname, const GLfloat * params);
// typedef void  (APIENTRYP GPMATERIALI)(GLenum  face, GLenum  pname, GLint  param);
// typedef void  (APIENTRYP GPMATERIALIV)(GLenum  face, GLenum  pname, const GLint * params);
// typedef void  (APIENTRYP GPMATRIXMODE)(GLenum  mode);
// typedef void  (APIENTRYP GPMEMORYBARRIER)(GLbitfield  barriers);
// typedef void  (APIENTRYP GPMEMORYBARRIERBYREGION)(GLbitfield  barriers);
// typedef void  (APIENTRYP GPMINSAMPLESHADING)(GLfloat  value);
// typedef void  (APIENTRYP GPMINSAMPLESHADINGARB)(GLfloat  value);
// typedef void  (APIENTRYP GPMULTMATRIXD)(const GLdouble * m);
// typedef void  (APIENTRYP GPMULTMATRIXF)(const GLfloat * m);
// typedef void  (APIENTRYP GPMULTTRANSPOSEMATRIXD)(const GLdouble * m);
// typedef void  (APIENTRYP GPMULTTRANSPOSEMATRIXF)(const GLfloat * m);
// typedef void  (APIENTRYP GPMULTIDRAWARRAYS)(GLenum  mode, const GLint * first, const GLsizei * count, GLsizei  drawcount);
// typedef void  (APIENTRYP GPMULTIDRAWARRAYSINDIRECT)(GLenum  mode, const void * indirect, GLsizei  drawcount, GLsizei  stride);
// typedef void  (APIENTRYP GPMULTIDRAWARRAYSINDIRECTCOUNTARB)(GLenum  mode, GLintptr  indirect, GLintptr  drawcount, GLsizei  maxdrawcount, GLsizei  stride);
// typedef void  (APIENTRYP GPMULTIDRAWELEMENTS)(GLenum  mode, const GLsizei * count, GLenum  type, const void *const* indices, GLsizei  drawcount);
// typedef void  (APIENTRYP GPMULTIDRAWELEMENTSBASEVERTEX)(GLenum  mode, const GLsizei * count, GLenum  type, const void *const* indices, GLsizei  drawcount, const GLint * basevertex);
// typedef void  (APIENTRYP GPMULTIDRAWELEMENTSINDIRECT)(GLenum  mode, GLenum  type, const void * indirect, GLsizei  drawcount, GLsizei  stride);
// typedef void  (APIENTRYP GPMULTIDRAWELEMENTSINDIRECTCOUNTARB)(GLenum  mode, GLenum  type, GLintptr  indirect, GLintptr  drawcount, GLsizei  maxdrawcount, GLsizei  stride);
// typedef void  (APIENTRYP GPMULTITEXCOORD1D)(GLenum  target, GLdouble  s);
// typedef void  (APIENTRYP GPMULTITEXCOORD1DV)(GLenum  target, const GLdouble * v);
// typedef void  (APIENTRYP GPMULTITEXCOORD1F)(GLenum  target, GLfloat  s);
// typedef void  (APIENTRYP GPMULTITEXCOORD1FV)(GLenum  target, const GLfloat * v);
// typedef void  (APIENTRYP GPMULTITEXCOORD1I)(GLenum  target, GLint  s);
// typedef void  (APIENTRYP GPMULTITEXCOORD1IV)(GLenum  target, const GLint * v);
// typedef void  (APIENTRYP GPMULTITEXCOORD1S)(GLenum  target, GLshort  s);
// typedef void  (APIENTRYP GPMULTITEXCOORD1SV)(GLenum  target, const GLshort * v);
// typedef void  (APIENTRYP GPMULTITEXCOORD2D)(GLenum  target, GLdouble  s, GLdouble  t);
// typedef void  (APIENTRYP GPMULTITEXCOORD2DV)(GLenum  target, const GLdouble * v);
// typedef void  (APIENTRYP GPMULTITEXCOORD2F)(GLenum  target, GLfloat  s, GLfloat  t);
// typedef void  (APIENTRYP GPMULTITEXCOORD2FV)(GLenum  target, const GLfloat * v);
// typedef void  (APIENTRYP GPMULTITEXCOORD2I)(GLenum  target, GLint  s, GLint  t);
// typedef void  (APIENTRYP GPMULTITEXCOORD2IV)(GLenum  target, const GLint * v);
// typedef void  (APIENTRYP GPMULTITEXCOORD2S)(GLenum  target, GLshort  s, GLshort  t);
// typedef void  (APIENTRYP GPMULTITEXCOORD2SV)(GLenum  target, const GLshort * v);
// typedef void  (APIENTRYP GPMULTITEXCOORD3D)(GLenum  target, GLdouble  s, GLdouble  t, GLdouble  r);
// typedef void  (APIENTRYP GPMULTITEXCOORD3DV)(GLenum  target, const GLdouble * v);
// typedef void  (APIENTRYP GPMULTITEXCOORD3F)(GLenum  target, GLfloat  s, GLfloat  t, GLfloat  r);
// typedef void  (APIENTRYP GPMULTITEXCOORD3FV)(GLenum  target, const GLfloat * v);
// typedef void  (APIENTRYP GPMULTITEXCOORD3I)(GLenum  target, GLint  s, GLint  t, GLint  r);
// typedef void  (APIENTRYP GPMULTITEXCOORD3IV)(GLenum  target, const GLint * v);
// typedef void  (APIENTRYP GPMULTITEXCOORD3S)(GLenum  target, GLshort  s, GLshort  t, GLshort  r);
// typedef void  (APIENTRYP GPMULTITEXCOORD3SV)(GLenum  target, const GLshort * v);
// typedef void  (APIENTRYP GPMULTITEXCOORD4D)(GLenum  target, GLdouble  s, GLdouble  t, GLdouble  r, GLdouble  q);
// typedef void  (APIENTRYP GPMULTITEXCOORD4DV)(GLenum  target, const GLdouble * v);
// typedef void  (APIENTRYP GPMULTITEXCOORD4F)(GLenum  target, GLfloat  s, GLfloat  t, GLfloat  r, GLfloat  q);
// typedef void  (APIENTRYP GPMULTITEXCOORD4FV)(GLenum  target, const GLfloat * v);
// typedef void  (APIENTRYP GPMULTITEXCOORD4I)(GLenum  target, GLint  s, GLint  t, GLint  r, GLint  q);
// typedef void  (APIENTRYP GPMULTITEXCOORD4IV)(GLenum  target, const GLint * v);
// typedef void  (APIENTRYP GPMULTITEXCOORD4S)(GLenum  target, GLshort  s, GLshort  t, GLshort  r, GLshort  q);
// typedef void  (APIENTRYP GPMULTITEXCOORD4SV)(GLenum  target, const GLshort * v);
// typedef void  (APIENTRYP GPNAMEDBUFFERDATA)(GLuint  buffer, GLsizei  size, const void * data, GLenum  usage);
// typedef void  (APIENTRYP GPNAMEDBUFFERPAGECOMMITMENTARB)(GLuint  buffer, GLintptr  offset, GLsizei  size, GLboolean  commit);
// typedef void  (APIENTRYP GPNAMEDBUFFERPAGECOMMITMENTEXT)(GLuint  buffer, GLintptr  offset, GLsizei  size, GLboolean  commit);
// typedef void  (APIENTRYP GPNAMEDBUFFERSTORAGE)(GLuint  buffer, GLsizei  size, const void * data, GLbitfield  flags);
// typedef void  (APIENTRYP GPNAMEDBUFFERSUBDATA)(GLuint  buffer, GLintptr  offset, GLsizei  size, const void * data);
// typedef void  (APIENTRYP GPNAMEDFRAMEBUFFERDRAWBUFFER)(GLuint  framebuffer, GLenum  buf);
// typedef void  (APIENTRYP GPNAMEDFRAMEBUFFERDRAWBUFFERS)(GLuint  framebuffer, GLsizei  n, const GLenum * bufs);
// typedef void  (APIENTRYP GPNAMEDFRAMEBUFFERPARAMETERI)(GLuint  framebuffer, GLenum  pname, GLint  param);
// typedef void  (APIENTRYP GPNAMEDFRAMEBUFFERREADBUFFER)(GLuint  framebuffer, GLenum  src);
// typedef void  (APIENTRYP GPNAMEDFRAMEBUFFERRENDERBUFFER)(GLuint  framebuffer, GLenum  attachment, GLenum  renderbuffertarget, GLuint  renderbuffer);
// typedef void  (APIENTRYP GPNAMEDFRAMEBUFFERTEXTURE)(GLuint  framebuffer, GLenum  attachment, GLuint  texture, GLint  level);
// typedef void  (APIENTRYP GPNAMEDFRAMEBUFFERTEXTURELAYER)(GLuint  framebuffer, GLenum  attachment, GLuint  texture, GLint  level, GLint  layer);
// typedef void  (APIENTRYP GPNAMEDRENDERBUFFERSTORAGE)(GLuint  renderbuffer, GLenum  internalformat, GLsizei  width, GLsizei  height);
// typedef void  (APIENTRYP GPNAMEDRENDERBUFFERSTORAGEMULTISAMPLE)(GLuint  renderbuffer, GLsizei  samples, GLenum  internalformat, GLsizei  width, GLsizei  height);
// typedef void  (APIENTRYP GPNAMEDSTRINGARB)(GLenum  type, GLint  namelen, const GLchar * name, GLint  stringlen, const GLchar * string);
// typedef void  (APIENTRYP GPNEWLIST)(GLuint  list, GLenum  mode);
// typedef void  (APIENTRYP GPNORMAL3B)(GLbyte  nx, GLbyte  ny, GLbyte  nz);
// typedef void  (APIENTRYP GPNORMAL3BV)(const GLbyte * v);
// typedef void  (APIENTRYP GPNORMAL3D)(GLdouble  nx, GLdouble  ny, GLdouble  nz);
// typedef void  (APIENTRYP GPNORMAL3DV)(const GLdouble * v);
// typedef void  (APIENTRYP GPNORMAL3F)(GLfloat  nx, GLfloat  ny, GLfloat  nz);
// typedef void  (APIENTRYP GPNORMAL3FV)(const GLfloat * v);
// typedef void  (APIENTRYP GPNORMAL3I)(GLint  nx, GLint  ny, GLint  nz);
// typedef void  (APIENTRYP GPNORMAL3IV)(const GLint * v);
// typedef void  (APIENTRYP GPNORMAL3S)(GLshort  nx, GLshort  ny, GLshort  nz);
// typedef void  (APIENTRYP GPNORMAL3SV)(const GLshort * v);
// typedef void  (APIENTRYP GPNORMALPOINTER)(GLenum  type, GLsizei  stride, const void * pointer);
// typedef void  (APIENTRYP GPOBJECTLABEL)(GLenum  identifier, GLuint  name, GLsizei  length, const GLchar * label);
// typedef void  (APIENTRYP GPOBJECTLABELKHR)(GLenum  identifier, GLuint  name, GLsizei  length, const GLchar * label);
// typedef void  (APIENTRYP GPOBJECTPTRLABEL)(const void * ptr, GLsizei  length, const GLchar * label);
// typedef void  (APIENTRYP GPOBJECTPTRLABELKHR)(const void * ptr, GLsizei  length, const GLchar * label);
// typedef void  (APIENTRYP GPORTHO)(GLdouble  left, GLdouble  right, GLdouble  bottom, GLdouble  top, GLdouble  zNear, GLdouble  zFar);
// typedef void  (APIENTRYP GPPASSTHROUGH)(GLfloat  token);
// typedef void  (APIENTRYP GPPATCHPARAMETERFV)(GLenum  pname, const GLfloat * values);
// typedef void  (APIENTRYP GPPATCHPARAMETERI)(GLenum  pname, GLint  value);
// typedef void  (APIENTRYP GPPAUSETRANSFORMFEEDBACK)();
// typedef void  (APIENTRYP GPPIXELMAPFV)(GLenum  map, GLsizei  mapsize, const GLfloat * values);
// typedef void  (APIENTRYP GPPIXELMAPUIV)(GLenum  map, GLsizei  mapsize, const GLuint * values);
// typedef void  (APIENTRYP GPPIXELMAPUSV)(GLenum  map, GLsizei  mapsize, const GLushort * values);
// typedef void  (APIENTRYP GPPIXELSTOREF)(GLenum  pname, GLfloat  param);
// typedef void  (APIENTRYP GPPIXELSTOREI)(GLenum  pname, GLint  param);
// typedef void  (APIENTRYP GPPIXELTRANSFERF)(GLenum  pname, GLfloat  param);
// typedef void  (APIENTRYP GPPIXELTRANSFERI)(GLenum  pname, GLint  param);
// typedef void  (APIENTRYP GPPIXELZOOM)(GLfloat  xfactor, GLfloat  yfactor);
// typedef void  (APIENTRYP GPPOINTPARAMETERF)(GLenum  pname, GLfloat  param);
// typedef void  (APIENTRYP GPPOINTPARAMETERFV)(GLenum  pname, const GLfloat * params);
// typedef void  (APIENTRYP GPPOINTPARAMETERI)(GLenum  pname, GLint  param);
// typedef void  (APIENTRYP GPPOINTPARAMETERIV)(GLenum  pname, const GLint * params);
// typedef void  (APIENTRYP GPPOINTSIZE)(GLfloat  size);
// typedef void  (APIENTRYP GPPOLYGONMODE)(GLenum  face, GLenum  mode);
// typedef void  (APIENTRYP GPPOLYGONOFFSET)(GLfloat  factor, GLfloat  units);
// typedef void  (APIENTRYP GPPOLYGONSTIPPLE)(const GLubyte * mask);
// typedef void  (APIENTRYP GPPOPATTRIB)();
// typedef void  (APIENTRYP GPPOPCLIENTATTRIB)();
// typedef void  (APIENTRYP GPPOPDEBUGGROUP)();
// typedef void  (APIENTRYP GPPOPDEBUGGROUPKHR)();
// typedef void  (APIENTRYP GPPOPMATRIX)();
// typedef void  (APIENTRYP GPPOPNAME)();
// typedef void  (APIENTRYP GPPRIMITIVERESTARTINDEX)(GLuint  index);
// typedef void  (APIENTRYP GPPRIORITIZETEXTURES)(GLsizei  n, const GLuint * textures, const GLfloat * priorities);
// typedef void  (APIENTRYP GPPROGRAMBINARY)(GLuint  program, GLenum  binaryFormat, const void * binary, GLsizei  length);
// typedef void  (APIENTRYP GPPROGRAMPARAMETERI)(GLuint  program, GLenum  pname, GLint  value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM1D)(GLuint  program, GLint  location, GLdouble  v0);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM1DV)(GLuint  program, GLint  location, GLsizei  count, const GLdouble * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM1F)(GLuint  program, GLint  location, GLfloat  v0);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM1FV)(GLuint  program, GLint  location, GLsizei  count, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM1I)(GLuint  program, GLint  location, GLint  v0);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM1IV)(GLuint  program, GLint  location, GLsizei  count, const GLint * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM1UI)(GLuint  program, GLint  location, GLuint  v0);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM1UIV)(GLuint  program, GLint  location, GLsizei  count, const GLuint * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM2D)(GLuint  program, GLint  location, GLdouble  v0, GLdouble  v1);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM2DV)(GLuint  program, GLint  location, GLsizei  count, const GLdouble * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM2F)(GLuint  program, GLint  location, GLfloat  v0, GLfloat  v1);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM2FV)(GLuint  program, GLint  location, GLsizei  count, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM2I)(GLuint  program, GLint  location, GLint  v0, GLint  v1);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM2IV)(GLuint  program, GLint  location, GLsizei  count, const GLint * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM2UI)(GLuint  program, GLint  location, GLuint  v0, GLuint  v1);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM2UIV)(GLuint  program, GLint  location, GLsizei  count, const GLuint * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM3D)(GLuint  program, GLint  location, GLdouble  v0, GLdouble  v1, GLdouble  v2);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM3DV)(GLuint  program, GLint  location, GLsizei  count, const GLdouble * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM3F)(GLuint  program, GLint  location, GLfloat  v0, GLfloat  v1, GLfloat  v2);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM3FV)(GLuint  program, GLint  location, GLsizei  count, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM3I)(GLuint  program, GLint  location, GLint  v0, GLint  v1, GLint  v2);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM3IV)(GLuint  program, GLint  location, GLsizei  count, const GLint * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM3UI)(GLuint  program, GLint  location, GLuint  v0, GLuint  v1, GLuint  v2);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM3UIV)(GLuint  program, GLint  location, GLsizei  count, const GLuint * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM4D)(GLuint  program, GLint  location, GLdouble  v0, GLdouble  v1, GLdouble  v2, GLdouble  v3);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM4DV)(GLuint  program, GLint  location, GLsizei  count, const GLdouble * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM4F)(GLuint  program, GLint  location, GLfloat  v0, GLfloat  v1, GLfloat  v2, GLfloat  v3);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM4FV)(GLuint  program, GLint  location, GLsizei  count, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM4I)(GLuint  program, GLint  location, GLint  v0, GLint  v1, GLint  v2, GLint  v3);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM4IV)(GLuint  program, GLint  location, GLsizei  count, const GLint * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM4UI)(GLuint  program, GLint  location, GLuint  v0, GLuint  v1, GLuint  v2, GLuint  v3);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM4UIV)(GLuint  program, GLint  location, GLsizei  count, const GLuint * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMHANDLEUI64ARB)(GLuint  program, GLint  location, GLuint64  value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMHANDLEUI64VARB)(GLuint  program, GLint  location, GLsizei  count, const GLuint64 * values);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX2DV)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX2FV)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX2X3DV)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX2X3FV)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX2X4DV)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX2X4FV)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX3DV)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX3FV)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX3X2DV)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX3X2FV)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX3X4DV)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX3X4FV)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX4DV)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX4FV)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX4X2DV)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX4X2FV)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX4X3DV)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX4X3FV)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPPROVOKINGVERTEX)(GLenum  mode);
// typedef void  (APIENTRYP GPPUSHATTRIB)(GLbitfield  mask);
// typedef void  (APIENTRYP GPPUSHCLIENTATTRIB)(GLbitfield  mask);
// typedef void  (APIENTRYP GPPUSHDEBUGGROUP)(GLenum  source, GLuint  id, GLsizei  length, const GLchar * message);
// typedef void  (APIENTRYP GPPUSHDEBUGGROUPKHR)(GLenum  source, GLuint  id, GLsizei  length, const GLchar * message);
// typedef void  (APIENTRYP GPPUSHMATRIX)();
// typedef void  (APIENTRYP GPPUSHNAME)(GLuint  name);
// typedef void  (APIENTRYP GPQUERYCOUNTER)(GLuint  id, GLenum  target);
// typedef void  (APIENTRYP GPRASTERPOS2D)(GLdouble  x, GLdouble  y);
// typedef void  (APIENTRYP GPRASTERPOS2DV)(const GLdouble * v);
// typedef void  (APIENTRYP GPRASTERPOS2F)(GLfloat  x, GLfloat  y);
// typedef void  (APIENTRYP GPRASTERPOS2FV)(const GLfloat * v);
// typedef void  (APIENTRYP GPRASTERPOS2I)(GLint  x, GLint  y);
// typedef void  (APIENTRYP GPRASTERPOS2IV)(const GLint * v);
// typedef void  (APIENTRYP GPRASTERPOS2S)(GLshort  x, GLshort  y);
// typedef void  (APIENTRYP GPRASTERPOS2SV)(const GLshort * v);
// typedef void  (APIENTRYP GPRASTERPOS3D)(GLdouble  x, GLdouble  y, GLdouble  z);
// typedef void  (APIENTRYP GPRASTERPOS3DV)(const GLdouble * v);
// typedef void  (APIENTRYP GPRASTERPOS3F)(GLfloat  x, GLfloat  y, GLfloat  z);
// typedef void  (APIENTRYP GPRASTERPOS3FV)(const GLfloat * v);
// typedef void  (APIENTRYP GPRASTERPOS3I)(GLint  x, GLint  y, GLint  z);
// typedef void  (APIENTRYP GPRASTERPOS3IV)(const GLint * v);
// typedef void  (APIENTRYP GPRASTERPOS3S)(GLshort  x, GLshort  y, GLshort  z);
// typedef void  (APIENTRYP GPRASTERPOS3SV)(const GLshort * v);
// typedef void  (APIENTRYP GPRASTERPOS4D)(GLdouble  x, GLdouble  y, GLdouble  z, GLdouble  w);
// typedef void  (APIENTRYP GPRASTERPOS4DV)(const GLdouble * v);
// typedef void  (APIENTRYP GPRASTERPOS4F)(GLfloat  x, GLfloat  y, GLfloat  z, GLfloat  w);
// typedef void  (APIENTRYP GPRASTERPOS4FV)(const GLfloat * v);
// typedef void  (APIENTRYP GPRASTERPOS4I)(GLint  x, GLint  y, GLint  z, GLint  w);
// typedef void  (APIENTRYP GPRASTERPOS4IV)(const GLint * v);
// typedef void  (APIENTRYP GPRASTERPOS4S)(GLshort  x, GLshort  y, GLshort  z, GLshort  w);
// typedef void  (APIENTRYP GPRASTERPOS4SV)(const GLshort * v);
// typedef void  (APIENTRYP GPREADBUFFER)(GLenum  src);
// typedef void  (APIENTRYP GPREADPIXELS)(GLint  x, GLint  y, GLsizei  width, GLsizei  height, GLenum  format, GLenum  type, void * pixels);
// typedef void  (APIENTRYP GPREADNPIXELS)(GLint  x, GLint  y, GLsizei  width, GLsizei  height, GLenum  format, GLenum  type, GLsizei  bufSize, void * data);
// typedef void  (APIENTRYP GPREADNPIXELSARB)(GLint  x, GLint  y, GLsizei  width, GLsizei  height, GLenum  format, GLenum  type, GLsizei  bufSize, void * data);
// typedef void  (APIENTRYP GPREADNPIXELSKHR)(GLint  x, GLint  y, GLsizei  width, GLsizei  height, GLenum  format, GLenum  type, GLsizei  bufSize, void * data);
// typedef void  (APIENTRYP GPRECTD)(GLdouble  x1, GLdouble  y1, GLdouble  x2, GLdouble  y2);
// typedef void  (APIENTRYP GPRECTDV)(const GLdouble * v1, const GLdouble * v2);
// typedef void  (APIENTRYP GPRECTF)(GLfloat  x1, GLfloat  y1, GLfloat  x2, GLfloat  y2);
// typedef void  (APIENTRYP GPRECTFV)(const GLfloat * v1, const GLfloat * v2);
// typedef void  (APIENTRYP GPRECTI)(GLint  x1, GLint  y1, GLint  x2, GLint  y2);
// typedef void  (APIENTRYP GPRECTIV)(const GLint * v1, const GLint * v2);
// typedef void  (APIENTRYP GPRECTS)(GLshort  x1, GLshort  y1, GLshort  x2, GLshort  y2);
// typedef void  (APIENTRYP GPRECTSV)(const GLshort * v1, const GLshort * v2);
// typedef void  (APIENTRYP GPRELEASESHADERCOMPILER)();
// typedef GLint  (APIENTRYP GPRENDERMODE)(GLenum  mode);
// typedef void  (APIENTRYP GPRENDERBUFFERSTORAGE)(GLenum  target, GLenum  internalformat, GLsizei  width, GLsizei  height);
// typedef void  (APIENTRYP GPRENDERBUFFERSTORAGEMULTISAMPLE)(GLenum  target, GLsizei  samples, GLenum  internalformat, GLsizei  width, GLsizei  height);
// typedef void  (APIENTRYP GPRESUMETRANSFORMFEEDBACK)();
// typedef void  (APIENTRYP GPROTATED)(GLdouble  angle, GLdouble  x, GLdouble  y, GLdouble  z);
// typedef void  (APIENTRYP GPROTATEF)(GLfloat  angle, GLfloat  x, GLfloat  y, GLfloat  z);
// typedef void  (APIENTRYP GPSAMPLECOVERAGE)(GLfloat  value, GLboolean  invert);
// typedef void  (APIENTRYP GPSAMPLEMASKI)(GLuint  maskNumber, GLbitfield  mask);
// typedef void  (APIENTRYP GPSAMPLERPARAMETERIIV)(GLuint  sampler, GLenum  pname, const GLint * param);
// typedef void  (APIENTRYP GPSAMPLERPARAMETERIUIV)(GLuint  sampler, GLenum  pname, const GLuint * param);
// typedef void  (APIENTRYP GPSAMPLERPARAMETERF)(GLuint  sampler, GLenum  pname, GLfloat  param);
// typedef void  (APIENTRYP GPSAMPLERPARAMETERFV)(GLuint  sampler, GLenum  pname, const GLfloat * param);
// typedef void  (APIENTRYP GPSAMPLERPARAMETERI)(GLuint  sampler, GLenum  pname, GLint  param);
// typedef void  (APIENTRYP GPSAMPLERPARAMETERIV)(GLuint  sampler, GLenum  pname, const GLint * param);
// typedef void  (APIENTRYP GPSCALED)(GLdouble  x, GLdouble  y, GLdouble  z);
// typedef void  (APIENTRYP GPSCALEF)(GLfloat  x, GLfloat  y, GLfloat  z);
// typedef void  (APIENTRYP GPSCISSOR)(GLint  x, GLint  y, GLsizei  width, GLsizei  height);
// typedef void  (APIENTRYP GPSCISSORARRAYV)(GLuint  first, GLsizei  count, const GLint * v);
// typedef void  (APIENTRYP GPSCISSORINDEXED)(GLuint  index, GLint  left, GLint  bottom, GLsizei  width, GLsizei  height);
// typedef void  (APIENTRYP GPSCISSORINDEXEDV)(GLuint  index, const GLint * v);
// typedef void  (APIENTRYP GPSECONDARYCOLOR3B)(GLbyte  red, GLbyte  green, GLbyte  blue);
// typedef void  (APIENTRYP GPSECONDARYCOLOR3BV)(const GLbyte * v);
// typedef void  (APIENTRYP GPSECONDARYCOLOR3D)(GLdouble  red, GLdouble  green, GLdouble  blue);
// typedef void  (APIENTRYP GPSECONDARYCOLOR3DV)(const GLdouble * v);
// typedef void  (APIENTRYP GPSECONDARYCOLOR3F)(GLfloat  red, GLfloat  green, GLfloat  blue);
// typedef void  (APIENTRYP GPSECONDARYCOLOR3FV)(const GLfloat * v);
// typedef void  (APIENTRYP GPSECONDARYCOLOR3I)(GLint  red, GLint  green, GLint  blue);
// typedef void  (APIENTRYP GPSECONDARYCOLOR3IV)(const GLint * v);
// typedef void  (APIENTRYP GPSECONDARYCOLOR3S)(GLshort  red, GLshort  green, GLshort  blue);
// typedef void  (APIENTRYP GPSECONDARYCOLOR3SV)(const GLshort * v);
// typedef void  (APIENTRYP GPSECONDARYCOLOR3UB)(GLubyte  red, GLubyte  green, GLubyte  blue);
// typedef void  (APIENTRYP GPSECONDARYCOLOR3UBV)(const GLubyte * v);
// typedef void  (APIENTRYP GPSECONDARYCOLOR3UI)(GLuint  red, GLuint  green, GLuint  blue);
// typedef void  (APIENTRYP GPSECONDARYCOLOR3UIV)(const GLuint * v);
// typedef void  (APIENTRYP GPSECONDARYCOLOR3US)(GLushort  red, GLushort  green, GLushort  blue);
// typedef void  (APIENTRYP GPSECONDARYCOLOR3USV)(const GLushort * v);
// typedef void  (APIENTRYP GPSECONDARYCOLORPOINTER)(GLint  size, GLenum  type, GLsizei  stride, const void * pointer);
// typedef void  (APIENTRYP GPSELECTBUFFER)(GLsizei  size, GLuint * buffer);
// typedef void  (APIENTRYP GPSHADEMODEL)(GLenum  mode);
// typedef void  (APIENTRYP GPSHADERBINARY)(GLsizei  count, const GLuint * shaders, GLenum  binaryformat, const void * binary, GLsizei  length);
// typedef void  (APIENTRYP GPSHADERSOURCE)(GLuint  shader, GLsizei  count, const GLchar *const* string, const GLint * length);
// typedef void  (APIENTRYP GPSHADERSTORAGEBLOCKBINDING)(GLuint  program, GLuint  storageBlockIndex, GLuint  storageBlockBinding);
// typedef void  (APIENTRYP GPSTENCILFUNC)(GLenum  func, GLint  ref, GLuint  mask);
// typedef void  (APIENTRYP GPSTENCILFUNCSEPARATE)(GLenum  face, GLenum  func, GLint  ref, GLuint  mask);
// typedef void  (APIENTRYP GPSTENCILMASK)(GLuint  mask);
// typedef void  (APIENTRYP GPSTENCILMASKSEPARATE)(GLenum  face, GLuint  mask);
// typedef void  (APIENTRYP GPSTENCILOP)(GLenum  fail, GLenum  zfail, GLenum  zpass);
// typedef void  (APIENTRYP GPSTENCILOPSEPARATE)(GLenum  face, GLenum  sfail, GLenum  dpfail, GLenum  dppass);
// typedef void  (APIENTRYP GPTEXBUFFER)(GLenum  target, GLenum  internalformat, GLuint  buffer);
// typedef void  (APIENTRYP GPTEXBUFFERRANGE)(GLenum  target, GLenum  internalformat, GLuint  buffer, GLintptr  offset, GLsizeiptr  size);
// typedef void  (APIENTRYP GPTEXCOORD1D)(GLdouble  s);
// typedef void  (APIENTRYP GPTEXCOORD1DV)(const GLdouble * v);
// typedef void  (APIENTRYP GPTEXCOORD1F)(GLfloat  s);
// typedef void  (APIENTRYP GPTEXCOORD1FV)(const GLfloat * v);
// typedef void  (APIENTRYP GPTEXCOORD1I)(GLint  s);
// typedef void  (APIENTRYP GPTEXCOORD1IV)(const GLint * v);
// typedef void  (APIENTRYP GPTEXCOORD1S)(GLshort  s);
// typedef void  (APIENTRYP GPTEXCOORD1SV)(const GLshort * v);
// typedef void  (APIENTRYP GPTEXCOORD2D)(GLdouble  s, GLdouble  t);
// typedef void  (APIENTRYP GPTEXCOORD2DV)(const GLdouble * v);
// typedef void  (APIENTRYP GPTEXCOORD2F)(GLfloat  s, GLfloat  t);
// typedef void  (APIENTRYP GPTEXCOORD2FV)(const GLfloat * v);
// typedef void  (APIENTRYP GPTEXCOORD2I)(GLint  s, GLint  t);
// typedef void  (APIENTRYP GPTEXCOORD2IV)(const GLint * v);
// typedef void  (APIENTRYP GPTEXCOORD2S)(GLshort  s, GLshort  t);
// typedef void  (APIENTRYP GPTEXCOORD2SV)(const GLshort * v);
// typedef void  (APIENTRYP GPTEXCOORD3D)(GLdouble  s, GLdouble  t, GLdouble  r);
// typedef void  (APIENTRYP GPTEXCOORD3DV)(const GLdouble * v);
// typedef void  (APIENTRYP GPTEXCOORD3F)(GLfloat  s, GLfloat  t, GLfloat  r);
// typedef void  (APIENTRYP GPTEXCOORD3FV)(const GLfloat * v);
// typedef void  (APIENTRYP GPTEXCOORD3I)(GLint  s, GLint  t, GLint  r);
// typedef void  (APIENTRYP GPTEXCOORD3IV)(const GLint * v);
// typedef void  (APIENTRYP GPTEXCOORD3S)(GLshort  s, GLshort  t, GLshort  r);
// typedef void  (APIENTRYP GPTEXCOORD3SV)(const GLshort * v);
// typedef void  (APIENTRYP GPTEXCOORD4D)(GLdouble  s, GLdouble  t, GLdouble  r, GLdouble  q);
// typedef void  (APIENTRYP GPTEXCOORD4DV)(const GLdouble * v);
// typedef void  (APIENTRYP GPTEXCOORD4F)(GLfloat  s, GLfloat  t, GLfloat  r, GLfloat  q);
// typedef void  (APIENTRYP GPTEXCOORD4FV)(const GLfloat * v);
// typedef void  (APIENTRYP GPTEXCOORD4I)(GLint  s, GLint  t, GLint  r, GLint  q);
// typedef void  (APIENTRYP GPTEXCOORD4IV)(const GLint * v);
// typedef void  (APIENTRYP GPTEXCOORD4S)(GLshort  s, GLshort  t, GLshort  r, GLshort  q);
// typedef void  (APIENTRYP GPTEXCOORD4SV)(const GLshort * v);
// typedef void  (APIENTRYP GPTEXCOORDPOINTER)(GLint  size, GLenum  type, GLsizei  stride, const void * pointer);
// typedef void  (APIENTRYP GPTEXENVF)(GLenum  target, GLenum  pname, GLfloat  param);
// typedef void  (APIENTRYP GPTEXENVFV)(GLenum  target, GLenum  pname, const GLfloat * params);
// typedef void  (APIENTRYP GPTEXENVI)(GLenum  target, GLenum  pname, GLint  param);
// typedef void  (APIENTRYP GPTEXENVIV)(GLenum  target, GLenum  pname, const GLint * params);
// typedef void  (APIENTRYP GPTEXGEND)(GLenum  coord, GLenum  pname, GLdouble  param);
// typedef void  (APIENTRYP GPTEXGENDV)(GLenum  coord, GLenum  pname, const GLdouble * params);
// typedef void  (APIENTRYP GPTEXGENF)(GLenum  coord, GLenum  pname, GLfloat  param);
// typedef void  (APIENTRYP GPTEXGENFV)(GLenum  coord, GLenum  pname, const GLfloat * params);
// typedef void  (APIENTRYP GPTEXGENI)(GLenum  coord, GLenum  pname, GLint  param);
// typedef void  (APIENTRYP GPTEXGENIV)(GLenum  coord, GLenum  pname, const GLint * params);
// typedef void  (APIENTRYP GPTEXIMAGE1D)(GLenum  target, GLint  level, GLint  internalformat, GLsizei  width, GLint  border, GLenum  format, GLenum  type, const void * pixels);
// typedef void  (APIENTRYP GPTEXIMAGE2D)(GLenum  target, GLint  level, GLint  internalformat, GLsizei  width, GLsizei  height, GLint  border, GLenum  format, GLenum  type, const void * pixels);
// typedef void  (APIENTRYP GPTEXIMAGE2DMULTISAMPLE)(GLenum  target, GLsizei  samples, GLenum  internalformat, GLsizei  width, GLsizei  height, GLboolean  fixedsamplelocations);
// typedef void  (APIENTRYP GPTEXIMAGE3D)(GLenum  target, GLint  level, GLint  internalformat, GLsizei  width, GLsizei  height, GLsizei  depth, GLint  border, GLenum  format, GLenum  type, const void * pixels);
// typedef void  (APIENTRYP GPTEXIMAGE3DMULTISAMPLE)(GLenum  target, GLsizei  samples, GLenum  internalformat, GLsizei  width, GLsizei  height, GLsizei  depth, GLboolean  fixedsamplelocations);
// typedef void  (APIENTRYP GPTEXPAGECOMMITMENTARB)(GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth, GLboolean  resident);
// typedef void  (APIENTRYP GPTEXPARAMETERIIV)(GLenum  target, GLenum  pname, const GLint * params);
// typedef void  (APIENTRYP GPTEXPARAMETERIUIV)(GLenum  target, GLenum  pname, const GLuint * params);
// typedef void  (APIENTRYP GPTEXPARAMETERF)(GLenum  target, GLenum  pname, GLfloat  param);
// typedef void  (APIENTRYP GPTEXPARAMETERFV)(GLenum  target, GLenum  pname, const GLfloat * params);
// typedef void  (APIENTRYP GPTEXPARAMETERI)(GLenum  target, GLenum  pname, GLint  param);
// typedef void  (APIENTRYP GPTEXPARAMETERIV)(GLenum  target, GLenum  pname, const GLint * params);
// typedef void  (APIENTRYP GPTEXSTORAGE1D)(GLenum  target, GLsizei  levels, GLenum  internalformat, GLsizei  width);
// typedef void  (APIENTRYP GPTEXSTORAGE2D)(GLenum  target, GLsizei  levels, GLenum  internalformat, GLsizei  width, GLsizei  height);
// typedef void  (APIENTRYP GPTEXSTORAGE2DMULTISAMPLE)(GLenum  target, GLsizei  samples, GLenum  internalformat, GLsizei  width, GLsizei  height, GLboolean  fixedsamplelocations);
// typedef void  (APIENTRYP GPTEXSTORAGE3D)(GLenum  target, GLsizei  levels, GLenum  internalformat, GLsizei  width, GLsizei  height, GLsizei  depth);
// typedef void  (APIENTRYP GPTEXSTORAGE3DMULTISAMPLE)(GLenum  target, GLsizei  samples, GLenum  internalformat, GLsizei  width, GLsizei  height, GLsizei  depth, GLboolean  fixedsamplelocations);
// typedef void  (APIENTRYP GPTEXSUBIMAGE1D)(GLenum  target, GLint  level, GLint  xoffset, GLsizei  width, GLenum  format, GLenum  type, const void * pixels);
// typedef void  (APIENTRYP GPTEXSUBIMAGE2D)(GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLsizei  width, GLsizei  height, GLenum  format, GLenum  type, const void * pixels);
// typedef void  (APIENTRYP GPTEXSUBIMAGE3D)(GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth, GLenum  format, GLenum  type, const void * pixels);
// typedef void  (APIENTRYP GPTEXTUREBARRIER)();
// typedef void  (APIENTRYP GPTEXTUREBUFFER)(GLuint  texture, GLenum  internalformat, GLuint  buffer);
// typedef void  (APIENTRYP GPTEXTUREBUFFERRANGE)(GLuint  texture, GLenum  internalformat, GLuint  buffer, GLintptr  offset, GLsizei  size);
// typedef void  (APIENTRYP GPTEXTUREPARAMETERIIV)(GLuint  texture, GLenum  pname, const GLint * params);
// typedef void  (APIENTRYP GPTEXTUREPARAMETERIUIV)(GLuint  texture, GLenum  pname, const GLuint * params);
// typedef void  (APIENTRYP GPTEXTUREPARAMETERF)(GLuint  texture, GLenum  pname, GLfloat  param);
// typedef void  (APIENTRYP GPTEXTUREPARAMETERFV)(GLuint  texture, GLenum  pname, const GLfloat * param);
// typedef void  (APIENTRYP GPTEXTUREPARAMETERI)(GLuint  texture, GLenum  pname, GLint  param);
// typedef void  (APIENTRYP GPTEXTUREPARAMETERIV)(GLuint  texture, GLenum  pname, const GLint * param);
// typedef void  (APIENTRYP GPTEXTURESTORAGE1D)(GLuint  texture, GLsizei  levels, GLenum  internalformat, GLsizei  width);
// typedef void  (APIENTRYP GPTEXTURESTORAGE2D)(GLuint  texture, GLsizei  levels, GLenum  internalformat, GLsizei  width, GLsizei  height);
// typedef void  (APIENTRYP GPTEXTURESTORAGE2DMULTISAMPLE)(GLuint  texture, GLsizei  samples, GLenum  internalformat, GLsizei  width, GLsizei  height, GLboolean  fixedsamplelocations);
// typedef void  (APIENTRYP GPTEXTURESTORAGE3D)(GLuint  texture, GLsizei  levels, GLenum  internalformat, GLsizei  width, GLsizei  height, GLsizei  depth);
// typedef void  (APIENTRYP GPTEXTURESTORAGE3DMULTISAMPLE)(GLuint  texture, GLsizei  samples, GLenum  internalformat, GLsizei  width, GLsizei  height, GLsizei  depth, GLboolean  fixedsamplelocations);
// typedef void  (APIENTRYP GPTEXTURESUBIMAGE1D)(GLuint  texture, GLint  level, GLint  xoffset, GLsizei  width, GLenum  format, GLenum  type, const void * pixels);
// typedef void  (APIENTRYP GPTEXTURESUBIMAGE2D)(GLuint  texture, GLint  level, GLint  xoffset, GLint  yoffset, GLsizei  width, GLsizei  height, GLenum  format, GLenum  type, const void * pixels);
// typedef void  (APIENTRYP GPTEXTURESUBIMAGE3D)(GLuint  texture, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth, GLenum  format, GLenum  type, const void * pixels);
// typedef void  (APIENTRYP GPTEXTUREVIEW)(GLuint  texture, GLenum  target, GLuint  origtexture, GLenum  internalformat, GLuint  minlevel, GLuint  numlevels, GLuint  minlayer, GLuint  numlayers);
// typedef void  (APIENTRYP GPTRANSFORMFEEDBACKBUFFERBASE)(GLuint  xfb, GLuint  index, GLuint  buffer);
// typedef void  (APIENTRYP GPTRANSFORMFEEDBACKBUFFERRANGE)(GLuint  xfb, GLuint  index, GLuint  buffer, GLintptr  offset, GLsizei  size);
// typedef void  (APIENTRYP GPTRANSFORMFEEDBACKVARYINGS)(GLuint  program, GLsizei  count, const GLchar *const* varyings, GLenum  bufferMode);
// typedef void  (APIENTRYP GPTRANSLATED)(GLdouble  x, GLdouble  y, GLdouble  z);
// typedef void  (APIENTRYP GPTRANSLATEF)(GLfloat  x, GLfloat  y, GLfloat  z);
// typedef void  (APIENTRYP GPUNIFORM1D)(GLint  location, GLdouble  x);
// typedef void  (APIENTRYP GPUNIFORM1DV)(GLint  location, GLsizei  count, const GLdouble * value);
// typedef void  (APIENTRYP GPUNIFORM1F)(GLint  location, GLfloat  v0);
// typedef void  (APIENTRYP GPUNIFORM1FV)(GLint  location, GLsizei  count, const GLfloat * value);
// typedef void  (APIENTRYP GPUNIFORM1I)(GLint  location, GLint  v0);
// typedef void  (APIENTRYP GPUNIFORM1IV)(GLint  location, GLsizei  count, const GLint * value);
// typedef void  (APIENTRYP GPUNIFORM1UI)(GLint  location, GLuint  v0);
// typedef void  (APIENTRYP GPUNIFORM1UIV)(GLint  location, GLsizei  count, const GLuint * value);
// typedef void  (APIENTRYP GPUNIFORM2D)(GLint  location, GLdouble  x, GLdouble  y);
// typedef void  (APIENTRYP GPUNIFORM2DV)(GLint  location, GLsizei  count, const GLdouble * value);
// typedef void  (APIENTRYP GPUNIFORM2F)(GLint  location, GLfloat  v0, GLfloat  v1);
// typedef void  (APIENTRYP GPUNIFORM2FV)(GLint  location, GLsizei  count, const GLfloat * value);
// typedef void  (APIENTRYP GPUNIFORM2I)(GLint  location, GLint  v0, GLint  v1);
// typedef void  (APIENTRYP GPUNIFORM2IV)(GLint  location, GLsizei  count, const GLint * value);
// typedef void  (APIENTRYP GPUNIFORM2UI)(GLint  location, GLuint  v0, GLuint  v1);
// typedef void  (APIENTRYP GPUNIFORM2UIV)(GLint  location, GLsizei  count, const GLuint * value);
// typedef void  (APIENTRYP GPUNIFORM3D)(GLint  location, GLdouble  x, GLdouble  y, GLdouble  z);
// typedef void  (APIENTRYP GPUNIFORM3DV)(GLint  location, GLsizei  count, const GLdouble * value);
// typedef void  (APIENTRYP GPUNIFORM3F)(GLint  location, GLfloat  v0, GLfloat  v1, GLfloat  v2);
// typedef void  (APIENTRYP GPUNIFORM3FV)(GLint  location, GLsizei  count, const GLfloat * value);
// typedef void  (APIENTRYP GPUNIFORM3I)(GLint  location, GLint  v0, GLint  v1, GLint  v2);
// typedef void  (APIENTRYP GPUNIFORM3IV)(GLint  location, GLsizei  count, const GLint * value);
// typedef void  (APIENTRYP GPUNIFORM3UI)(GLint  location, GLuint  v0, GLuint  v1, GLuint  v2);
// typedef void  (APIENTRYP GPUNIFORM3UIV)(GLint  location, GLsizei  count, const GLuint * value);
// typedef void  (APIENTRYP GPUNIFORM4D)(GLint  location, GLdouble  x, GLdouble  y, GLdouble  z, GLdouble  w);
// typedef void  (APIENTRYP GPUNIFORM4DV)(GLint  location, GLsizei  count, const GLdouble * value);
// typedef void  (APIENTRYP GPUNIFORM4F)(GLint  location, GLfloat  v0, GLfloat  v1, GLfloat  v2, GLfloat  v3);
// typedef void  (APIENTRYP GPUNIFORM4FV)(GLint  location, GLsizei  count, const GLfloat * value);
// typedef void  (APIENTRYP GPUNIFORM4I)(GLint  location, GLint  v0, GLint  v1, GLint  v2, GLint  v3);
// typedef void  (APIENTRYP GPUNIFORM4IV)(GLint  location, GLsizei  count, const GLint * value);
// typedef void  (APIENTRYP GPUNIFORM4UI)(GLint  location, GLuint  v0, GLuint  v1, GLuint  v2, GLuint  v3);
// typedef void  (APIENTRYP GPUNIFORM4UIV)(GLint  location, GLsizei  count, const GLuint * value);
// typedef void  (APIENTRYP GPUNIFORMBLOCKBINDING)(GLuint  program, GLuint  uniformBlockIndex, GLuint  uniformBlockBinding);
// typedef void  (APIENTRYP GPUNIFORMHANDLEUI64ARB)(GLint  location, GLuint64  value);
// typedef void  (APIENTRYP GPUNIFORMHANDLEUI64VARB)(GLint  location, GLsizei  count, const GLuint64 * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX2DV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX2FV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX2X3DV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX2X3FV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX2X4DV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX2X4FV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX3DV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX3FV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX3X2DV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX3X2FV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX3X4DV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX3X4FV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX4DV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX4FV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX4X2DV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX4X2FV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX4X3DV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX4X3FV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPUNIFORMSUBROUTINESUIV)(GLenum  shadertype, GLsizei  count, const GLuint * indices);
// typedef GLboolean  (APIENTRYP GPUNMAPBUFFER)(GLenum  target);
// typedef GLboolean  (APIENTRYP GPUNMAPNAMEDBUFFER)(GLuint  buffer);
// typedef void  (APIENTRYP GPUSEPROGRAM)(GLuint  program);
// typedef void  (APIENTRYP GPUSEPROGRAMSTAGES)(GLuint  pipeline, GLbitfield  stages, GLuint  program);
// typedef void  (APIENTRYP GPVALIDATEPROGRAM)(GLuint  program);
// typedef void  (APIENTRYP GPVALIDATEPROGRAMPIPELINE)(GLuint  pipeline);
// typedef void  (APIENTRYP GPVERTEX2D)(GLdouble  x, GLdouble  y);
// typedef void  (APIENTRYP GPVERTEX2DV)(const GLdouble * v);
// typedef void  (APIENTRYP GPVERTEX2F)(GLfloat  x, GLfloat  y);
// typedef void  (APIENTRYP GPVERTEX2FV)(const GLfloat * v);
// typedef void  (APIENTRYP GPVERTEX2I)(GLint  x, GLint  y);
// typedef void  (APIENTRYP GPVERTEX2IV)(const GLint * v);
// typedef void  (APIENTRYP GPVERTEX2S)(GLshort  x, GLshort  y);
// typedef void  (APIENTRYP GPVERTEX2SV)(const GLshort * v);
// typedef void  (APIENTRYP GPVERTEX3D)(GLdouble  x, GLdouble  y, GLdouble  z);
// typedef void  (APIENTRYP GPVERTEX3DV)(const GLdouble * v);
// typedef void  (APIENTRYP GPVERTEX3F)(GLfloat  x, GLfloat  y, GLfloat  z);
// typedef void  (APIENTRYP GPVERTEX3FV)(const GLfloat * v);
// typedef void  (APIENTRYP GPVERTEX3I)(GLint  x, GLint  y, GLint  z);
// typedef void  (APIENTRYP GPVERTEX3IV)(const GLint * v);
// typedef void  (APIENTRYP GPVERTEX3S)(GLshort  x, GLshort  y, GLshort  z);
// typedef void  (APIENTRYP GPVERTEX3SV)(const GLshort * v);
// typedef void  (APIENTRYP GPVERTEX4D)(GLdouble  x, GLdouble  y, GLdouble  z, GLdouble  w);
// typedef void  (APIENTRYP GPVERTEX4DV)(const GLdouble * v);
// typedef void  (APIENTRYP GPVERTEX4F)(GLfloat  x, GLfloat  y, GLfloat  z, GLfloat  w);
// typedef void  (APIENTRYP GPVERTEX4FV)(const GLfloat * v);
// typedef void  (APIENTRYP GPVERTEX4I)(GLint  x, GLint  y, GLint  z, GLint  w);
// typedef void  (APIENTRYP GPVERTEX4IV)(const GLint * v);
// typedef void  (APIENTRYP GPVERTEX4S)(GLshort  x, GLshort  y, GLshort  z, GLshort  w);
// typedef void  (APIENTRYP GPVERTEX4SV)(const GLshort * v);
// typedef void  (APIENTRYP GPVERTEXARRAYATTRIBBINDING)(GLuint  vaobj, GLuint  attribindex, GLuint  bindingindex);
// typedef void  (APIENTRYP GPVERTEXARRAYATTRIBFORMAT)(GLuint  vaobj, GLuint  attribindex, GLint  size, GLenum  type, GLboolean  normalized, GLuint  relativeoffset);
// typedef void  (APIENTRYP GPVERTEXARRAYATTRIBIFORMAT)(GLuint  vaobj, GLuint  attribindex, GLint  size, GLenum  type, GLuint  relativeoffset);
// typedef void  (APIENTRYP GPVERTEXARRAYATTRIBLFORMAT)(GLuint  vaobj, GLuint  attribindex, GLint  size, GLenum  type, GLuint  relativeoffset);
// typedef void  (APIENTRYP GPVERTEXARRAYBINDINGDIVISOR)(GLuint  vaobj, GLuint  bindingindex, GLuint  divisor);
// typedef void  (APIENTRYP GPVERTEXARRAYELEMENTBUFFER)(GLuint  vaobj, GLuint  buffer);
// typedef void  (APIENTRYP GPVERTEXARRAYVERTEXBUFFER)(GLuint  vaobj, GLuint  bindingindex, GLuint  buffer, GLintptr  offset, GLsizei  stride);
// typedef void  (APIENTRYP GPVERTEXARRAYVERTEXBUFFERS)(GLuint  vaobj, GLuint  first, GLsizei  count, const GLuint * buffers, const GLintptr * offsets, const GLsizei * strides);
// typedef void  (APIENTRYP GPVERTEXATTRIB1D)(GLuint  index, GLdouble  x);
// typedef void  (APIENTRYP GPVERTEXATTRIB1DV)(GLuint  index, const GLdouble * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB1F)(GLuint  index, GLfloat  x);
// typedef void  (APIENTRYP GPVERTEXATTRIB1FV)(GLuint  index, const GLfloat * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB1S)(GLuint  index, GLshort  x);
// typedef void  (APIENTRYP GPVERTEXATTRIB1SV)(GLuint  index, const GLshort * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB2D)(GLuint  index, GLdouble  x, GLdouble  y);
// typedef void  (APIENTRYP GPVERTEXATTRIB2DV)(GLuint  index, const GLdouble * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB2F)(GLuint  index, GLfloat  x, GLfloat  y);
// typedef void  (APIENTRYP GPVERTEXATTRIB2FV)(GLuint  index, const GLfloat * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB2S)(GLuint  index, GLshort  x, GLshort  y);
// typedef void  (APIENTRYP GPVERTEXATTRIB2SV)(GLuint  index, const GLshort * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB3D)(GLuint  index, GLdouble  x, GLdouble  y, GLdouble  z);
// typedef void  (APIENTRYP GPVERTEXATTRIB3DV)(GLuint  index, const GLdouble * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB3F)(GLuint  index, GLfloat  x, GLfloat  y, GLfloat  z);
// typedef void  (APIENTRYP GPVERTEXATTRIB3FV)(GLuint  index, const GLfloat * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB3S)(GLuint  index, GLshort  x, GLshort  y, GLshort  z);
// typedef void  (APIENTRYP GPVERTEXATTRIB3SV)(GLuint  index, const GLshort * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB4NBV)(GLuint  index, const GLbyte * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB4NIV)(GLuint  index, const GLint * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB4NSV)(GLuint  index, const GLshort * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB4NUB)(GLuint  index, GLubyte  x, GLubyte  y, GLubyte  z, GLubyte  w);
// typedef void  (APIENTRYP GPVERTEXATTRIB4NUBV)(GLuint  index, const GLubyte * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB4NUIV)(GLuint  index, const GLuint * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB4NUSV)(GLuint  index, const GLushort * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB4BV)(GLuint  index, const GLbyte * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB4D)(GLuint  index, GLdouble  x, GLdouble  y, GLdouble  z, GLdouble  w);
// typedef void  (APIENTRYP GPVERTEXATTRIB4DV)(GLuint  index, const GLdouble * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB4F)(GLuint  index, GLfloat  x, GLfloat  y, GLfloat  z, GLfloat  w);
// typedef void  (APIENTRYP GPVERTEXATTRIB4FV)(GLuint  index, const GLfloat * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB4IV)(GLuint  index, const GLint * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB4S)(GLuint  index, GLshort  x, GLshort  y, GLshort  z, GLshort  w);
// typedef void  (APIENTRYP GPVERTEXATTRIB4SV)(GLuint  index, const GLshort * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB4UBV)(GLuint  index, const GLubyte * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB4UIV)(GLuint  index, const GLuint * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB4USV)(GLuint  index, const GLushort * v);
// typedef void  (APIENTRYP GPVERTEXATTRIBBINDING)(GLuint  attribindex, GLuint  bindingindex);
// typedef void  (APIENTRYP GPVERTEXATTRIBDIVISOR)(GLuint  index, GLuint  divisor);
// typedef void  (APIENTRYP GPVERTEXATTRIBFORMAT)(GLuint  attribindex, GLint  size, GLenum  type, GLboolean  normalized, GLuint  relativeoffset);
// typedef void  (APIENTRYP GPVERTEXATTRIBI1I)(GLuint  index, GLint  x);
// typedef void  (APIENTRYP GPVERTEXATTRIBI1IV)(GLuint  index, const GLint * v);
// typedef void  (APIENTRYP GPVERTEXATTRIBI1UI)(GLuint  index, GLuint  x);
// typedef void  (APIENTRYP GPVERTEXATTRIBI1UIV)(GLuint  index, const GLuint * v);
// typedef void  (APIENTRYP GPVERTEXATTRIBI2I)(GLuint  index, GLint  x, GLint  y);
// typedef void  (APIENTRYP GPVERTEXATTRIBI2IV)(GLuint  index, const GLint * v);
// typedef void  (APIENTRYP GPVERTEXATTRIBI2UI)(GLuint  index, GLuint  x, GLuint  y);
// typedef void  (APIENTRYP GPVERTEXATTRIBI2UIV)(GLuint  index, const GLuint * v);
// typedef void  (APIENTRYP GPVERTEXATTRIBI3I)(GLuint  index, GLint  x, GLint  y, GLint  z);
// typedef void  (APIENTRYP GPVERTEXATTRIBI3IV)(GLuint  index, const GLint * v);
// typedef void  (APIENTRYP GPVERTEXATTRIBI3UI)(GLuint  index, GLuint  x, GLuint  y, GLuint  z);
// typedef void  (APIENTRYP GPVERTEXATTRIBI3UIV)(GLuint  index, const GLuint * v);
// typedef void  (APIENTRYP GPVERTEXATTRIBI4BV)(GLuint  index, const GLbyte * v);
// typedef void  (APIENTRYP GPVERTEXATTRIBI4I)(GLuint  index, GLint  x, GLint  y, GLint  z, GLint  w);
// typedef void  (APIENTRYP GPVERTEXATTRIBI4IV)(GLuint  index, const GLint * v);
// typedef void  (APIENTRYP GPVERTEXATTRIBI4SV)(GLuint  index, const GLshort * v);
// typedef void  (APIENTRYP GPVERTEXATTRIBI4UBV)(GLuint  index, const GLubyte * v);
// typedef void  (APIENTRYP GPVERTEXATTRIBI4UI)(GLuint  index, GLuint  x, GLuint  y, GLuint  z, GLuint  w);
// typedef void  (APIENTRYP GPVERTEXATTRIBI4UIV)(GLuint  index, const GLuint * v);
// typedef void  (APIENTRYP GPVERTEXATTRIBI4USV)(GLuint  index, const GLushort * v);
// typedef void  (APIENTRYP GPVERTEXATTRIBIFORMAT)(GLuint  attribindex, GLint  size, GLenum  type, GLuint  relativeoffset);
// typedef void  (APIENTRYP GPVERTEXATTRIBIPOINTER)(GLuint  index, GLint  size, GLenum  type, GLsizei  stride, const void * pointer);
// typedef void  (APIENTRYP GPVERTEXATTRIBL1D)(GLuint  index, GLdouble  x);
// typedef void  (APIENTRYP GPVERTEXATTRIBL1DV)(GLuint  index, const GLdouble * v);
// typedef void  (APIENTRYP GPVERTEXATTRIBL1UI64ARB)(GLuint  index, GLuint64EXT  x);
// typedef void  (APIENTRYP GPVERTEXATTRIBL1UI64VARB)(GLuint  index, const GLuint64EXT * v);
// typedef void  (APIENTRYP GPVERTEXATTRIBL2D)(GLuint  index, GLdouble  x, GLdouble  y);
// typedef void  (APIENTRYP GPVERTEXATTRIBL2DV)(GLuint  index, const GLdouble * v);
// typedef void  (APIENTRYP GPVERTEXATTRIBL3D)(GLuint  index, GLdouble  x, GLdouble  y, GLdouble  z);
// typedef void  (APIENTRYP GPVERTEXATTRIBL3DV)(GLuint  index, const GLdouble * v);
// typedef void  (APIENTRYP GPVERTEXATTRIBL4D)(GLuint  index, GLdouble  x, GLdouble  y, GLdouble  z, GLdouble  w);
// typedef void  (APIENTRYP GPVERTEXATTRIBL4DV)(GLuint  index, const GLdouble * v);
// typedef void  (APIENTRYP GPVERTEXATTRIBLFORMAT)(GLuint  attribindex, GLint  size, GLenum  type, GLuint  relativeoffset);
// typedef void  (APIENTRYP GPVERTEXATTRIBLPOINTER)(GLuint  index, GLint  size, GLenum  type, GLsizei  stride, const void * pointer);
// typedef void  (APIENTRYP GPVERTEXATTRIBP1UI)(GLuint  index, GLenum  type, GLboolean  normalized, GLuint  value);
// typedef void  (APIENTRYP GPVERTEXATTRIBP1UIV)(GLuint  index, GLenum  type, GLboolean  normalized, const GLuint * value);
// typedef void  (APIENTRYP GPVERTEXATTRIBP2UI)(GLuint  index, GLenum  type, GLboolean  normalized, GLuint  value);
// typedef void  (APIENTRYP GPVERTEXATTRIBP2UIV)(GLuint  index, GLenum  type, GLboolean  normalized, const GLuint * value);
// typedef void  (APIENTRYP GPVERTEXATTRIBP3UI)(GLuint  index, GLenum  type, GLboolean  normalized, GLuint  value);
// typedef void  (APIENTRYP GPVERTEXATTRIBP3UIV)(GLuint  index, GLenum  type, GLboolean  normalized, const GLuint * value);
// typedef void  (APIENTRYP GPVERTEXATTRIBP4UI)(GLuint  index, GLenum  type, GLboolean  normalized, GLuint  value);
// typedef void  (APIENTRYP GPVERTEXATTRIBP4UIV)(GLuint  index, GLenum  type, GLboolean  normalized, const GLuint * value);
// typedef void  (APIENTRYP GPVERTEXATTRIBPOINTER)(GLuint  index, GLint  size, GLenum  type, GLboolean  normalized, GLsizei  stride, const void * pointer);
// typedef void  (APIENTRYP GPVERTEXBINDINGDIVISOR)(GLuint  bindingindex, GLuint  divisor);
// typedef void  (APIENTRYP GPVERTEXPOINTER)(GLint  size, GLenum  type, GLsizei  stride, const void * pointer);
// typedef void  (APIENTRYP GPVIEWPORT)(GLint  x, GLint  y, GLsizei  width, GLsizei  height);
// typedef void  (APIENTRYP GPVIEWPORTARRAYV)(GLuint  first, GLsizei  count, const GLfloat * v);
// typedef void  (APIENTRYP GPVIEWPORTINDEXEDF)(GLuint  index, GLfloat  x, GLfloat  y, GLfloat  w, GLfloat  h);
// typedef void  (APIENTRYP GPVIEWPORTINDEXEDFV)(GLuint  index, const GLfloat * v);
// typedef void  (APIENTRYP GPWAITSYNC)(GLsync  sync, GLbitfield  flags, GLuint64  timeout);
// typedef void  (APIENTRYP GPWINDOWPOS2D)(GLdouble  x, GLdouble  y);
// typedef void  (APIENTRYP GPWINDOWPOS2DV)(const GLdouble * v);
// typedef void  (APIENTRYP GPWINDOWPOS2F)(GLfloat  x, GLfloat  y);
// typedef void  (APIENTRYP GPWINDOWPOS2FV)(const GLfloat * v);
// typedef void  (APIENTRYP GPWINDOWPOS2I)(GLint  x, GLint  y);
// typedef void  (APIENTRYP GPWINDOWPOS2IV)(const GLint * v);
// typedef void  (APIENTRYP GPWINDOWPOS2S)(GLshort  x, GLshort  y);
// typedef void  (APIENTRYP GPWINDOWPOS2SV)(const GLshort * v);
// typedef void  (APIENTRYP GPWINDOWPOS3D)(GLdouble  x, GLdouble  y, GLdouble  z);
// typedef void  (APIENTRYP GPWINDOWPOS3DV)(const GLdouble * v);
// typedef void  (APIENTRYP GPWINDOWPOS3F)(GLfloat  x, GLfloat  y, GLfloat  z);
// typedef void  (APIENTRYP GPWINDOWPOS3FV)(const GLfloat * v);
// typedef void  (APIENTRYP GPWINDOWPOS3I)(GLint  x, GLint  y, GLint  z);
// typedef void  (APIENTRYP GPWINDOWPOS3IV)(const GLint * v);
// typedef void  (APIENTRYP GPWINDOWPOS3S)(GLshort  x, GLshort  y, GLshort  z);
// typedef void  (APIENTRYP GPWINDOWPOS3SV)(const GLshort * v);
// static void  glowAccum(GPACCUM fnptr, GLenum  op, GLfloat  value) {
//   (*fnptr)(op, value);
// }
// static void  glowActiveShaderProgram(GPACTIVESHADERPROGRAM fnptr, GLuint  pipeline, GLuint  program) {
//   (*fnptr)(pipeline, program);
// }
// static void  glowActiveTexture(GPACTIVETEXTURE fnptr, GLenum  texture) {
//   (*fnptr)(texture);
// }
// static void  glowAlphaFunc(GPALPHAFUNC fnptr, GLenum  func, GLfloat  ref) {
//   (*fnptr)(func, ref);
// }
// static GLboolean  glowAreTexturesResident(GPARETEXTURESRESIDENT fnptr, GLsizei  n, const GLuint * textures, GLboolean * residences) {
//   return (*fnptr)(n, textures, residences);
// }
// static void  glowArrayElement(GPARRAYELEMENT fnptr, GLint  i) {
//   (*fnptr)(i);
// }
// static void  glowAttachShader(GPATTACHSHADER fnptr, GLuint  program, GLuint  shader) {
//   (*fnptr)(program, shader);
// }
// static void  glowBegin(GPBEGIN fnptr, GLenum  mode) {
//   (*fnptr)(mode);
// }
// static void  glowBeginConditionalRender(GPBEGINCONDITIONALRENDER fnptr, GLuint  id, GLenum  mode) {
//   (*fnptr)(id, mode);
// }
// static void  glowBeginQuery(GPBEGINQUERY fnptr, GLenum  target, GLuint  id) {
//   (*fnptr)(target, id);
// }
// static void  glowBeginQueryIndexed(GPBEGINQUERYINDEXED fnptr, GLenum  target, GLuint  index, GLuint  id) {
//   (*fnptr)(target, index, id);
// }
// static void  glowBeginTransformFeedback(GPBEGINTRANSFORMFEEDBACK fnptr, GLenum  primitiveMode) {
//   (*fnptr)(primitiveMode);
// }
// static void  glowBindAttribLocation(GPBINDATTRIBLOCATION fnptr, GLuint  program, GLuint  index, const GLchar * name) {
//   (*fnptr)(program, index, name);
// }
// static void  glowBindBuffer(GPBINDBUFFER fnptr, GLenum  target, GLuint  buffer) {
//   (*fnptr)(target, buffer);
// }
// static void  glowBindBufferBase(GPBINDBUFFERBASE fnptr, GLenum  target, GLuint  index, GLuint  buffer) {
//   (*fnptr)(target, index, buffer);
// }
// static void  glowBindBufferRange(GPBINDBUFFERRANGE fnptr, GLenum  target, GLuint  index, GLuint  buffer, GLintptr  offset, GLsizeiptr  size) {
//   (*fnptr)(target, index, buffer, offset, size);
// }
// static void  glowBindBuffersBase(GPBINDBUFFERSBASE fnptr, GLenum  target, GLuint  first, GLsizei  count, const GLuint * buffers) {
//   (*fnptr)(target, first, count, buffers);
// }
// static void  glowBindBuffersRange(GPBINDBUFFERSRANGE fnptr, GLenum  target, GLuint  first, GLsizei  count, const GLuint * buffers, const GLintptr * offsets, const GLsizeiptr * sizes) {
//   (*fnptr)(target, first, count, buffers, offsets, sizes);
// }
// static void  glowBindFragDataLocation(GPBINDFRAGDATALOCATION fnptr, GLuint  program, GLuint  color, const GLchar * name) {
//   (*fnptr)(program, color, name);
// }
// static void  glowBindFragDataLocationIndexed(GPBINDFRAGDATALOCATIONINDEXED fnptr, GLuint  program, GLuint  colorNumber, GLuint  index, const GLchar * name) {
//   (*fnptr)(program, colorNumber, index, name);
// }
// static void  glowBindFramebuffer(GPBINDFRAMEBUFFER fnptr, GLenum  target, GLuint  framebuffer) {
//   (*fnptr)(target, framebuffer);
// }
// static void  glowBindImageTexture(GPBINDIMAGETEXTURE fnptr, GLuint  unit, GLuint  texture, GLint  level, GLboolean  layered, GLint  layer, GLenum  access, GLenum  format) {
//   (*fnptr)(unit, texture, level, layered, layer, access, format);
// }
// static void  glowBindImageTextures(GPBINDIMAGETEXTURES fnptr, GLuint  first, GLsizei  count, const GLuint * textures) {
//   (*fnptr)(first, count, textures);
// }
// static void  glowBindProgramPipeline(GPBINDPROGRAMPIPELINE fnptr, GLuint  pipeline) {
//   (*fnptr)(pipeline);
// }
// static void  glowBindRenderbuffer(GPBINDRENDERBUFFER fnptr, GLenum  target, GLuint  renderbuffer) {
//   (*fnptr)(target, renderbuffer);
// }
// static void  glowBindSampler(GPBINDSAMPLER fnptr, GLuint  unit, GLuint  sampler) {
//   (*fnptr)(unit, sampler);
// }
// static void  glowBindSamplers(GPBINDSAMPLERS fnptr, GLuint  first, GLsizei  count, const GLuint * samplers) {
//   (*fnptr)(first, count, samplers);
// }
// static void  glowBindTexture(GPBINDTEXTURE fnptr, GLenum  target, GLuint  texture) {
//   (*fnptr)(target, texture);
// }
// static void  glowBindTextureUnit(GPBINDTEXTUREUNIT fnptr, GLuint  unit, GLuint  texture) {
//   (*fnptr)(unit, texture);
// }
// static void  glowBindTextures(GPBINDTEXTURES fnptr, GLuint  first, GLsizei  count, const GLuint * textures) {
//   (*fnptr)(first, count, textures);
// }
// static void  glowBindTransformFeedback(GPBINDTRANSFORMFEEDBACK fnptr, GLenum  target, GLuint  id) {
//   (*fnptr)(target, id);
// }
// static void  glowBindVertexArray(GPBINDVERTEXARRAY fnptr, GLuint  array) {
//   (*fnptr)(array);
// }
// static void  glowBindVertexBuffer(GPBINDVERTEXBUFFER fnptr, GLuint  bindingindex, GLuint  buffer, GLintptr  offset, GLsizei  stride) {
//   (*fnptr)(bindingindex, buffer, offset, stride);
// }
// static void  glowBindVertexBuffers(GPBINDVERTEXBUFFERS fnptr, GLuint  first, GLsizei  count, const GLuint * buffers, const GLintptr * offsets, const GLsizei * strides) {
//   (*fnptr)(first, count, buffers, offsets, strides);
// }
// static void  glowBitmap(GPBITMAP fnptr, GLsizei  width, GLsizei  height, GLfloat  xorig, GLfloat  yorig, GLfloat  xmove, GLfloat  ymove, const GLubyte * bitmap) {
//   (*fnptr)(width, height, xorig, yorig, xmove, ymove, bitmap);
// }
// static void  glowBlendColor(GPBLENDCOLOR fnptr, GLfloat  red, GLfloat  green, GLfloat  blue, GLfloat  alpha) {
//   (*fnptr)(red, green, blue, alpha);
// }
// static void  glowBlendEquation(GPBLENDEQUATION fnptr, GLenum  mode) {
//   (*fnptr)(mode);
// }
// static void  glowBlendEquationSeparate(GPBLENDEQUATIONSEPARATE fnptr, GLenum  modeRGB, GLenum  modeAlpha) {
//   (*fnptr)(modeRGB, modeAlpha);
// }
// static void  glowBlendEquationSeparatei(GPBLENDEQUATIONSEPARATEI fnptr, GLuint  buf, GLenum  modeRGB, GLenum  modeAlpha) {
//   (*fnptr)(buf, modeRGB, modeAlpha);
// }
// static void  glowBlendEquationSeparateiARB(GPBLENDEQUATIONSEPARATEIARB fnptr, GLuint  buf, GLenum  modeRGB, GLenum  modeAlpha) {
//   (*fnptr)(buf, modeRGB, modeAlpha);
// }
// static void  glowBlendEquationi(GPBLENDEQUATIONI fnptr, GLuint  buf, GLenum  mode) {
//   (*fnptr)(buf, mode);
// }
// static void  glowBlendEquationiARB(GPBLENDEQUATIONIARB fnptr, GLuint  buf, GLenum  mode) {
//   (*fnptr)(buf, mode);
// }
// static void  glowBlendFunc(GPBLENDFUNC fnptr, GLenum  sfactor, GLenum  dfactor) {
//   (*fnptr)(sfactor, dfactor);
// }
// static void  glowBlendFuncSeparate(GPBLENDFUNCSEPARATE fnptr, GLenum  sfactorRGB, GLenum  dfactorRGB, GLenum  sfactorAlpha, GLenum  dfactorAlpha) {
//   (*fnptr)(sfactorRGB, dfactorRGB, sfactorAlpha, dfactorAlpha);
// }
// static void  glowBlendFuncSeparatei(GPBLENDFUNCSEPARATEI fnptr, GLuint  buf, GLenum  srcRGB, GLenum  dstRGB, GLenum  srcAlpha, GLenum  dstAlpha) {
//   (*fnptr)(buf, srcRGB, dstRGB, srcAlpha, dstAlpha);
// }
// static void  glowBlendFuncSeparateiARB(GPBLENDFUNCSEPARATEIARB fnptr, GLuint  buf, GLenum  srcRGB, GLenum  dstRGB, GLenum  srcAlpha, GLenum  dstAlpha) {
//   (*fnptr)(buf, srcRGB, dstRGB, srcAlpha, dstAlpha);
// }
// static void  glowBlendFunci(GPBLENDFUNCI fnptr, GLuint  buf, GLenum  src, GLenum  dst) {
//   (*fnptr)(buf, src, dst);
// }
// static void  glowBlendFunciARB(GPBLENDFUNCIARB fnptr, GLuint  buf, GLenum  src, GLenum  dst) {
//   (*fnptr)(buf, src, dst);
// }
// static void  glowBlitFramebuffer(GPBLITFRAMEBUFFER fnptr, GLint  srcX0, GLint  srcY0, GLint  srcX1, GLint  srcY1, GLint  dstX0, GLint  dstY0, GLint  dstX1, GLint  dstY1, GLbitfield  mask, GLenum  filter) {
//   (*fnptr)(srcX0, srcY0, srcX1, srcY1, dstX0, dstY0, dstX1, dstY1, mask, filter);
// }
// static void  glowBlitNamedFramebuffer(GPBLITNAMEDFRAMEBUFFER fnptr, GLuint  readFramebuffer, GLuint  drawFramebuffer, GLint  srcX0, GLint  srcY0, GLint  srcX1, GLint  srcY1, GLint  dstX0, GLint  dstY0, GLint  dstX1, GLint  dstY1, GLbitfield  mask, GLenum  filter) {
//   (*fnptr)(readFramebuffer, drawFramebuffer, srcX0, srcY0, srcX1, srcY1, dstX0, dstY0, dstX1, dstY1, mask, filter);
// }
// static void  glowBufferData(GPBUFFERDATA fnptr, GLenum  target, GLsizeiptr  size, const void * data, GLenum  usage) {
//   (*fnptr)(target, size, data, usage);
// }
// static void  glowBufferPageCommitmentARB(GPBUFFERPAGECOMMITMENTARB fnptr, GLenum  target, GLintptr  offset, GLsizei  size, GLboolean  commit) {
//   (*fnptr)(target, offset, size, commit);
// }
// static void  glowBufferStorage(GPBUFFERSTORAGE fnptr, GLenum  target, GLsizeiptr  size, const void * data, GLbitfield  flags) {
//   (*fnptr)(target, size, data, flags);
// }
// static void  glowBufferSubData(GPBUFFERSUBDATA fnptr, GLenum  target, GLintptr  offset, GLsizeiptr  size, const void * data) {
//   (*fnptr)(target, offset, size, data);
// }
// static void  glowCallList(GPCALLLIST fnptr, GLuint  list) {
//   (*fnptr)(list);
// }
// static void  glowCallLists(GPCALLLISTS fnptr, GLsizei  n, GLenum  type, const void * lists) {
//   (*fnptr)(n, type, lists);
// }
// static GLenum  glowCheckFramebufferStatus(GPCHECKFRAMEBUFFERSTATUS fnptr, GLenum  target) {
//   return (*fnptr)(target);
// }
// static GLenum  glowCheckNamedFramebufferStatus(GPCHECKNAMEDFRAMEBUFFERSTATUS fnptr, GLuint  framebuffer, GLenum  target) {
//   return (*fnptr)(framebuffer, target);
// }
// static void  glowClampColor(GPCLAMPCOLOR fnptr, GLenum  target, GLenum  clamp) {
//   (*fnptr)(target, clamp);
// }
// static void  glowClear(GPCLEAR fnptr, GLbitfield  mask) {
//   (*fnptr)(mask);
// }
// static void  glowClearAccum(GPCLEARACCUM fnptr, GLfloat  red, GLfloat  green, GLfloat  blue, GLfloat  alpha) {
//   (*fnptr)(red, green, blue, alpha);
// }
// static void  glowClearBufferData(GPCLEARBUFFERDATA fnptr, GLenum  target, GLenum  internalformat, GLenum  format, GLenum  type, const void * data) {
//   (*fnptr)(target, internalformat, format, type, data);
// }
// static void  glowClearBufferSubData(GPCLEARBUFFERSUBDATA fnptr, GLenum  target, GLenum  internalformat, GLintptr  offset, GLsizeiptr  size, GLenum  format, GLenum  type, const void * data) {
//   (*fnptr)(target, internalformat, offset, size, format, type, data);
// }
// static void  glowClearBufferfi(GPCLEARBUFFERFI fnptr, GLenum  buffer, GLint  drawbuffer, GLfloat  depth, GLint  stencil) {
//   (*fnptr)(buffer, drawbuffer, depth, stencil);
// }
// static void  glowClearBufferfv(GPCLEARBUFFERFV fnptr, GLenum  buffer, GLint  drawbuffer, const GLfloat * value) {
//   (*fnptr)(buffer, drawbuffer, value);
// }
// static void  glowClearBufferiv(GPCLEARBUFFERIV fnptr, GLenum  buffer, GLint  drawbuffer, const GLint * value) {
//   (*fnptr)(buffer, drawbuffer, value);
// }
// static void  glowClearBufferuiv(GPCLEARBUFFERUIV fnptr, GLenum  buffer, GLint  drawbuffer, const GLuint * value) {
//   (*fnptr)(buffer, drawbuffer, value);
// }
// static void  glowClearColor(GPCLEARCOLOR fnptr, GLfloat  red, GLfloat  green, GLfloat  blue, GLfloat  alpha) {
//   (*fnptr)(red, green, blue, alpha);
// }
// static void  glowClearDepth(GPCLEARDEPTH fnptr, GLdouble  depth) {
//   (*fnptr)(depth);
// }
// static void  glowClearDepthf(GPCLEARDEPTHF fnptr, GLfloat  d) {
//   (*fnptr)(d);
// }
// static void  glowClearIndex(GPCLEARINDEX fnptr, GLfloat  c) {
//   (*fnptr)(c);
// }
// static void  glowClearNamedBufferData(GPCLEARNAMEDBUFFERDATA fnptr, GLuint  buffer, GLenum  internalformat, GLenum  format, GLenum  type, const void * data) {
//   (*fnptr)(buffer, internalformat, format, type, data);
// }
// static void  glowClearNamedBufferSubData(GPCLEARNAMEDBUFFERSUBDATA fnptr, GLuint  buffer, GLenum  internalformat, GLintptr  offset, GLsizei  size, GLenum  format, GLenum  type, const void * data) {
//   (*fnptr)(buffer, internalformat, offset, size, format, type, data);
// }
// static void  glowClearNamedFramebufferfi(GPCLEARNAMEDFRAMEBUFFERFI fnptr, GLuint  framebuffer, GLenum  buffer, const GLfloat  depth, GLint  stencil) {
//   (*fnptr)(framebuffer, buffer, depth, stencil);
// }
// static void  glowClearNamedFramebufferfv(GPCLEARNAMEDFRAMEBUFFERFV fnptr, GLuint  framebuffer, GLenum  buffer, GLint  drawbuffer, const GLfloat * value) {
//   (*fnptr)(framebuffer, buffer, drawbuffer, value);
// }
// static void  glowClearNamedFramebufferiv(GPCLEARNAMEDFRAMEBUFFERIV fnptr, GLuint  framebuffer, GLenum  buffer, GLint  drawbuffer, const GLint * value) {
//   (*fnptr)(framebuffer, buffer, drawbuffer, value);
// }
// static void  glowClearNamedFramebufferuiv(GPCLEARNAMEDFRAMEBUFFERUIV fnptr, GLuint  framebuffer, GLenum  buffer, GLint  drawbuffer, const GLuint * value) {
//   (*fnptr)(framebuffer, buffer, drawbuffer, value);
// }
// static void  glowClearStencil(GPCLEARSTENCIL fnptr, GLint  s) {
//   (*fnptr)(s);
// }
// static void  glowClearTexImage(GPCLEARTEXIMAGE fnptr, GLuint  texture, GLint  level, GLenum  format, GLenum  type, const void * data) {
//   (*fnptr)(texture, level, format, type, data);
// }
// static void  glowClearTexSubImage(GPCLEARTEXSUBIMAGE fnptr, GLuint  texture, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth, GLenum  format, GLenum  type, const void * data) {
//   (*fnptr)(texture, level, xoffset, yoffset, zoffset, width, height, depth, format, type, data);
// }
// static void  glowClientActiveTexture(GPCLIENTACTIVETEXTURE fnptr, GLenum  texture) {
//   (*fnptr)(texture);
// }
// static GLenum  glowClientWaitSync(GPCLIENTWAITSYNC fnptr, GLsync  sync, GLbitfield  flags, GLuint64  timeout) {
//   return (*fnptr)(sync, flags, timeout);
// }
// static void  glowClipControl(GPCLIPCONTROL fnptr, GLenum  origin, GLenum  depth) {
//   (*fnptr)(origin, depth);
// }
// static void  glowClipPlane(GPCLIPPLANE fnptr, GLenum  plane, const GLdouble * equation) {
//   (*fnptr)(plane, equation);
// }
// static void  glowColor3b(GPCOLOR3B fnptr, GLbyte  red, GLbyte  green, GLbyte  blue) {
//   (*fnptr)(red, green, blue);
// }
// static void  glowColor3bv(GPCOLOR3BV fnptr, const GLbyte * v) {
//   (*fnptr)(v);
// }
// static void  glowColor3d(GPCOLOR3D fnptr, GLdouble  red, GLdouble  green, GLdouble  blue) {
//   (*fnptr)(red, green, blue);
// }
// static void  glowColor3dv(GPCOLOR3DV fnptr, const GLdouble * v) {
//   (*fnptr)(v);
// }
// static void  glowColor3f(GPCOLOR3F fnptr, GLfloat  red, GLfloat  green, GLfloat  blue) {
//   (*fnptr)(red, green, blue);
// }
// static void  glowColor3fv(GPCOLOR3FV fnptr, const GLfloat * v) {
//   (*fnptr)(v);
// }
// static void  glowColor3i(GPCOLOR3I fnptr, GLint  red, GLint  green, GLint  blue) {
//   (*fnptr)(red, green, blue);
// }
// static void  glowColor3iv(GPCOLOR3IV fnptr, const GLint * v) {
//   (*fnptr)(v);
// }
// static void  glowColor3s(GPCOLOR3S fnptr, GLshort  red, GLshort  green, GLshort  blue) {
//   (*fnptr)(red, green, blue);
// }
// static void  glowColor3sv(GPCOLOR3SV fnptr, const GLshort * v) {
//   (*fnptr)(v);
// }
// static void  glowColor3ub(GPCOLOR3UB fnptr, GLubyte  red, GLubyte  green, GLubyte  blue) {
//   (*fnptr)(red, green, blue);
// }
// static void  glowColor3ubv(GPCOLOR3UBV fnptr, const GLubyte * v) {
//   (*fnptr)(v);
// }
// static void  glowColor3ui(GPCOLOR3UI fnptr, GLuint  red, GLuint  green, GLuint  blue) {
//   (*fnptr)(red, green, blue);
// }
// static void  glowColor3uiv(GPCOLOR3UIV fnptr, const GLuint * v) {
//   (*fnptr)(v);
// }
// static void  glowColor3us(GPCOLOR3US fnptr, GLushort  red, GLushort  green, GLushort  blue) {
//   (*fnptr)(red, green, blue);
// }
// static void  glowColor3usv(GPCOLOR3USV fnptr, const GLushort * v) {
//   (*fnptr)(v);
// }
// static void  glowColor4b(GPCOLOR4B fnptr, GLbyte  red, GLbyte  green, GLbyte  blue, GLbyte  alpha) {
//   (*fnptr)(red, green, blue, alpha);
// }
// static void  glowColor4bv(GPCOLOR4BV fnptr, const GLbyte * v) {
//   (*fnptr)(v);
// }
// static void  glowColor4d(GPCOLOR4D fnptr, GLdouble  red, GLdouble  green, GLdouble  blue, GLdouble  alpha) {
//   (*fnptr)(red, green, blue, alpha);
// }
// static void  glowColor4dv(GPCOLOR4DV fnptr, const GLdouble * v) {
//   (*fnptr)(v);
// }
// static void  glowColor4f(GPCOLOR4F fnptr, GLfloat  red, GLfloat  green, GLfloat  blue, GLfloat  alpha) {
//   (*fnptr)(red, green, blue, alpha);
// }
// static void  glowColor4fv(GPCOLOR4FV fnptr, const GLfloat * v) {
//   (*fnptr)(v);
// }
// static void  glowColor4i(GPCOLOR4I fnptr, GLint  red, GLint  green, GLint  blue, GLint  alpha) {
//   (*fnptr)(red, green, blue, alpha);
// }
// static void  glowColor4iv(GPCOLOR4IV fnptr, const GLint * v) {
//   (*fnptr)(v);
// }
// static void  glowColor4s(GPCOLOR4S fnptr, GLshort  red, GLshort  green, GLshort  blue, GLshort  alpha) {
//   (*fnptr)(red, green, blue, alpha);
// }
// static void  glowColor4sv(GPCOLOR4SV fnptr, const GLshort * v) {
//   (*fnptr)(v);
// }
// static void  glowColor4ub(GPCOLOR4UB fnptr, GLubyte  red, GLubyte  green, GLubyte  blue, GLubyte  alpha) {
//   (*fnptr)(red, green, blue, alpha);
// }
// static void  glowColor4ubv(GPCOLOR4UBV fnptr, const GLubyte * v) {
//   (*fnptr)(v);
// }
// static void  glowColor4ui(GPCOLOR4UI fnptr, GLuint  red, GLuint  green, GLuint  blue, GLuint  alpha) {
//   (*fnptr)(red, green, blue, alpha);
// }
// static void  glowColor4uiv(GPCOLOR4UIV fnptr, const GLuint * v) {
//   (*fnptr)(v);
// }
// static void  glowColor4us(GPCOLOR4US fnptr, GLushort  red, GLushort  green, GLushort  blue, GLushort  alpha) {
//   (*fnptr)(red, green, blue, alpha);
// }
// static void  glowColor4usv(GPCOLOR4USV fnptr, const GLushort * v) {
//   (*fnptr)(v);
// }
// static void  glowColorMask(GPCOLORMASK fnptr, GLboolean  red, GLboolean  green, GLboolean  blue, GLboolean  alpha) {
//   (*fnptr)(red, green, blue, alpha);
// }
// static void  glowColorMaski(GPCOLORMASKI fnptr, GLuint  index, GLboolean  r, GLboolean  g, GLboolean  b, GLboolean  a) {
//   (*fnptr)(index, r, g, b, a);
// }
// static void  glowColorMaterial(GPCOLORMATERIAL fnptr, GLenum  face, GLenum  mode) {
//   (*fnptr)(face, mode);
// }
// static void  glowColorPointer(GPCOLORPOINTER fnptr, GLint  size, GLenum  type, GLsizei  stride, const void * pointer) {
//   (*fnptr)(size, type, stride, pointer);
// }
// static void  glowCompileShader(GPCOMPILESHADER fnptr, GLuint  shader) {
//   (*fnptr)(shader);
// }
// static void  glowCompileShaderIncludeARB(GPCOMPILESHADERINCLUDEARB fnptr, GLuint  shader, GLsizei  count, const GLchar *const* path, const GLint * length) {
//   (*fnptr)(shader, count, path, length);
// }
// static void  glowCompressedTexImage1D(GPCOMPRESSEDTEXIMAGE1D fnptr, GLenum  target, GLint  level, GLenum  internalformat, GLsizei  width, GLint  border, GLsizei  imageSize, const void * data) {
//   (*fnptr)(target, level, internalformat, width, border, imageSize, data);
// }
// static void  glowCompressedTexImage2D(GPCOMPRESSEDTEXIMAGE2D fnptr, GLenum  target, GLint  level, GLenum  internalformat, GLsizei  width, GLsizei  height, GLint  border, GLsizei  imageSize, const void * data) {
//   (*fnptr)(target, level, internalformat, width, height, border, imageSize, data);
// }
// static void  glowCompressedTexImage3D(GPCOMPRESSEDTEXIMAGE3D fnptr, GLenum  target, GLint  level, GLenum  internalformat, GLsizei  width, GLsizei  height, GLsizei  depth, GLint  border, GLsizei  imageSize, const void * data) {
//   (*fnptr)(target, level, internalformat, width, height, depth, border, imageSize, data);
// }
// static void  glowCompressedTexSubImage1D(GPCOMPRESSEDTEXSUBIMAGE1D fnptr, GLenum  target, GLint  level, GLint  xoffset, GLsizei  width, GLenum  format, GLsizei  imageSize, const void * data) {
//   (*fnptr)(target, level, xoffset, width, format, imageSize, data);
// }
// static void  glowCompressedTexSubImage2D(GPCOMPRESSEDTEXSUBIMAGE2D fnptr, GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLsizei  width, GLsizei  height, GLenum  format, GLsizei  imageSize, const void * data) {
//   (*fnptr)(target, level, xoffset, yoffset, width, height, format, imageSize, data);
// }
// static void  glowCompressedTexSubImage3D(GPCOMPRESSEDTEXSUBIMAGE3D fnptr, GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth, GLenum  format, GLsizei  imageSize, const void * data) {
//   (*fnptr)(target, level, xoffset, yoffset, zoffset, width, height, depth, format, imageSize, data);
// }
// static void  glowCompressedTextureSubImage1D(GPCOMPRESSEDTEXTURESUBIMAGE1D fnptr, GLuint  texture, GLint  level, GLint  xoffset, GLsizei  width, GLenum  format, GLsizei  imageSize, const void * data) {
//   (*fnptr)(texture, level, xoffset, width, format, imageSize, data);
// }
// static void  glowCompressedTextureSubImage2D(GPCOMPRESSEDTEXTURESUBIMAGE2D fnptr, GLuint  texture, GLint  level, GLint  xoffset, GLint  yoffset, GLsizei  width, GLsizei  height, GLenum  format, GLsizei  imageSize, const void * data) {
//   (*fnptr)(texture, level, xoffset, yoffset, width, height, format, imageSize, data);
// }
// static void  glowCompressedTextureSubImage3D(GPCOMPRESSEDTEXTURESUBIMAGE3D fnptr, GLuint  texture, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth, GLenum  format, GLsizei  imageSize, const void * data) {
//   (*fnptr)(texture, level, xoffset, yoffset, zoffset, width, height, depth, format, imageSize, data);
// }
// static void  glowCopyBufferSubData(GPCOPYBUFFERSUBDATA fnptr, GLenum  readTarget, GLenum  writeTarget, GLintptr  readOffset, GLintptr  writeOffset, GLsizeiptr  size) {
//   (*fnptr)(readTarget, writeTarget, readOffset, writeOffset, size);
// }
// static void  glowCopyImageSubData(GPCOPYIMAGESUBDATA fnptr, GLuint  srcName, GLenum  srcTarget, GLint  srcLevel, GLint  srcX, GLint  srcY, GLint  srcZ, GLuint  dstName, GLenum  dstTarget, GLint  dstLevel, GLint  dstX, GLint  dstY, GLint  dstZ, GLsizei  srcWidth, GLsizei  srcHeight, GLsizei  srcDepth) {
//   (*fnptr)(srcName, srcTarget, srcLevel, srcX, srcY, srcZ, dstName, dstTarget, dstLevel, dstX, dstY, dstZ, srcWidth, srcHeight, srcDepth);
// }
// static void  glowCopyNamedBufferSubData(GPCOPYNAMEDBUFFERSUBDATA fnptr, GLuint  readBuffer, GLuint  writeBuffer, GLintptr  readOffset, GLintptr  writeOffset, GLsizei  size) {
//   (*fnptr)(readBuffer, writeBuffer, readOffset, writeOffset, size);
// }
// static void  glowCopyPixels(GPCOPYPIXELS fnptr, GLint  x, GLint  y, GLsizei  width, GLsizei  height, GLenum  type) {
//   (*fnptr)(x, y, width, height, type);
// }
// static void  glowCopyTexImage1D(GPCOPYTEXIMAGE1D fnptr, GLenum  target, GLint  level, GLenum  internalformat, GLint  x, GLint  y, GLsizei  width, GLint  border) {
//   (*fnptr)(target, level, internalformat, x, y, width, border);
// }
// static void  glowCopyTexImage2D(GPCOPYTEXIMAGE2D fnptr, GLenum  target, GLint  level, GLenum  internalformat, GLint  x, GLint  y, GLsizei  width, GLsizei  height, GLint  border) {
//   (*fnptr)(target, level, internalformat, x, y, width, height, border);
// }
// static void  glowCopyTexSubImage1D(GPCOPYTEXSUBIMAGE1D fnptr, GLenum  target, GLint  level, GLint  xoffset, GLint  x, GLint  y, GLsizei  width) {
//   (*fnptr)(target, level, xoffset, x, y, width);
// }
// static void  glowCopyTexSubImage2D(GPCOPYTEXSUBIMAGE2D fnptr, GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLint  x, GLint  y, GLsizei  width, GLsizei  height) {
//   (*fnptr)(target, level, xoffset, yoffset, x, y, width, height);
// }
// static void  glowCopyTexSubImage3D(GPCOPYTEXSUBIMAGE3D fnptr, GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLint  x, GLint  y, GLsizei  width, GLsizei  height) {
//   (*fnptr)(target, level, xoffset, yoffset, zoffset, x, y, width, height);
// }
// static void  glowCopyTextureSubImage1D(GPCOPYTEXTURESUBIMAGE1D fnptr, GLuint  texture, GLint  level, GLint  xoffset, GLint  x, GLint  y, GLsizei  width) {
//   (*fnptr)(texture, level, xoffset, x, y, width);
// }
// static void  glowCopyTextureSubImage2D(GPCOPYTEXTURESUBIMAGE2D fnptr, GLuint  texture, GLint  level, GLint  xoffset, GLint  yoffset, GLint  x, GLint  y, GLsizei  width, GLsizei  height) {
//   (*fnptr)(texture, level, xoffset, yoffset, x, y, width, height);
// }
// static void  glowCopyTextureSubImage3D(GPCOPYTEXTURESUBIMAGE3D fnptr, GLuint  texture, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLint  x, GLint  y, GLsizei  width, GLsizei  height) {
//   (*fnptr)(texture, level, xoffset, yoffset, zoffset, x, y, width, height);
// }
// static void  glowCreateBuffers(GPCREATEBUFFERS fnptr, GLsizei  n, GLuint * buffers) {
//   (*fnptr)(n, buffers);
// }
// static void  glowCreateFramebuffers(GPCREATEFRAMEBUFFERS fnptr, GLsizei  n, GLuint * framebuffers) {
//   (*fnptr)(n, framebuffers);
// }
// static GLuint  glowCreateProgram(GPCREATEPROGRAM fnptr) {
//   return (*fnptr)();
// }
// static void  glowCreateProgramPipelines(GPCREATEPROGRAMPIPELINES fnptr, GLsizei  n, GLuint * pipelines) {
//   (*fnptr)(n, pipelines);
// }
// static void  glowCreateQueries(GPCREATEQUERIES fnptr, GLenum  target, GLsizei  n, GLuint * ids) {
//   (*fnptr)(target, n, ids);
// }
// static void  glowCreateRenderbuffers(GPCREATERENDERBUFFERS fnptr, GLsizei  n, GLuint * renderbuffers) {
//   (*fnptr)(n, renderbuffers);
// }
// static void  glowCreateSamplers(GPCREATESAMPLERS fnptr, GLsizei  n, GLuint * samplers) {
//   (*fnptr)(n, samplers);
// }
// static GLuint  glowCreateShader(GPCREATESHADER fnptr, GLenum  type) {
//   return (*fnptr)(type);
// }
// static GLuint  glowCreateShaderProgramv(GPCREATESHADERPROGRAMV fnptr, GLenum  type, GLsizei  count, const GLchar *const* strings) {
//   return (*fnptr)(type, count, strings);
// }
// static GLsync  glowCreateSyncFromCLeventARB(GPCREATESYNCFROMCLEVENTARB fnptr, struct _cl_context * context, struct _cl_event * event, GLbitfield  flags) {
//   return (*fnptr)(context, event, flags);
// }
// static void  glowCreateTextures(GPCREATETEXTURES fnptr, GLenum  target, GLsizei  n, GLuint * textures) {
//   (*fnptr)(target, n, textures);
// }
// static void  glowCreateTransformFeedbacks(GPCREATETRANSFORMFEEDBACKS fnptr, GLsizei  n, GLuint * ids) {
//   (*fnptr)(n, ids);
// }
// static void  glowCreateVertexArrays(GPCREATEVERTEXARRAYS fnptr, GLsizei  n, GLuint * arrays) {
//   (*fnptr)(n, arrays);
// }
// static void  glowCullFace(GPCULLFACE fnptr, GLenum  mode) {
//   (*fnptr)(mode);
// }
// static void  glowDebugMessageCallback(GPDEBUGMESSAGECALLBACK fnptr, GLDEBUGPROC  callback, const void * userParam) {
//   (*fnptr)(glowCDebugCallback, userParam);
// }
// static void  glowDebugMessageCallbackARB(GPDEBUGMESSAGECALLBACKARB fnptr, GLDEBUGPROCARB  callback, const void * userParam) {
//   (*fnptr)(glowCDebugCallback, userParam);
// }
// static void  glowDebugMessageCallbackKHR(GPDEBUGMESSAGECALLBACKKHR fnptr, GLDEBUGPROCKHR  callback, const void * userParam) {
//   (*fnptr)(glowCDebugCallback, userParam);
// }
// static void  glowDebugMessageControl(GPDEBUGMESSAGECONTROL fnptr, GLenum  source, GLenum  type, GLenum  severity, GLsizei  count, const GLuint * ids, GLboolean  enabled) {
//   (*fnptr)(source, type, severity, count, ids, enabled);
// }
// static void  glowDebugMessageControlARB(GPDEBUGMESSAGECONTROLARB fnptr, GLenum  source, GLenum  type, GLenum  severity, GLsizei  count, const GLuint * ids, GLboolean  enabled) {
//   (*fnptr)(source, type, severity, count, ids, enabled);
// }
// static void  glowDebugMessageControlKHR(GPDEBUGMESSAGECONTROLKHR fnptr, GLenum  source, GLenum  type, GLenum  severity, GLsizei  count, const GLuint * ids, GLboolean  enabled) {
//   (*fnptr)(source, type, severity, count, ids, enabled);
// }
// static void  glowDebugMessageInsert(GPDEBUGMESSAGEINSERT fnptr, GLenum  source, GLenum  type, GLuint  id, GLenum  severity, GLsizei  length, const GLchar * buf) {
//   (*fnptr)(source, type, id, severity, length, buf);
// }
// static void  glowDebugMessageInsertARB(GPDEBUGMESSAGEINSERTARB fnptr, GLenum  source, GLenum  type, GLuint  id, GLenum  severity, GLsizei  length, const GLchar * buf) {
//   (*fnptr)(source, type, id, severity, length, buf);
// }
// static void  glowDebugMessageInsertKHR(GPDEBUGMESSAGEINSERTKHR fnptr, GLenum  source, GLenum  type, GLuint  id, GLenum  severity, GLsizei  length, const GLchar * buf) {
//   (*fnptr)(source, type, id, severity, length, buf);
// }
// static void  glowDeleteBuffers(GPDELETEBUFFERS fnptr, GLsizei  n, const GLuint * buffers) {
//   (*fnptr)(n, buffers);
// }
// static void  glowDeleteFramebuffers(GPDELETEFRAMEBUFFERS fnptr, GLsizei  n, const GLuint * framebuffers) {
//   (*fnptr)(n, framebuffers);
// }
// static void  glowDeleteLists(GPDELETELISTS fnptr, GLuint  list, GLsizei  range) {
//   (*fnptr)(list, range);
// }
// static void  glowDeleteNamedStringARB(GPDELETENAMEDSTRINGARB fnptr, GLint  namelen, const GLchar * name) {
//   (*fnptr)(namelen, name);
// }
// static void  glowDeleteProgram(GPDELETEPROGRAM fnptr, GLuint  program) {
//   (*fnptr)(program);
// }
// static void  glowDeleteProgramPipelines(GPDELETEPROGRAMPIPELINES fnptr, GLsizei  n, const GLuint * pipelines) {
//   (*fnptr)(n, pipelines);
// }
// static void  glowDeleteQueries(GPDELETEQUERIES fnptr, GLsizei  n, const GLuint * ids) {
//   (*fnptr)(n, ids);
// }
// static void  glowDeleteRenderbuffers(GPDELETERENDERBUFFERS fnptr, GLsizei  n, const GLuint * renderbuffers) {
//   (*fnptr)(n, renderbuffers);
// }
// static void  glowDeleteSamplers(GPDELETESAMPLERS fnptr, GLsizei  count, const GLuint * samplers) {
//   (*fnptr)(count, samplers);
// }
// static void  glowDeleteShader(GPDELETESHADER fnptr, GLuint  shader) {
//   (*fnptr)(shader);
// }
// static void  glowDeleteSync(GPDELETESYNC fnptr, GLsync  sync) {
//   (*fnptr)(sync);
// }
// static void  glowDeleteTextures(GPDELETETEXTURES fnptr, GLsizei  n, const GLuint * textures) {
//   (*fnptr)(n, textures);
// }
// static void  glowDeleteTransformFeedbacks(GPDELETETRANSFORMFEEDBACKS fnptr, GLsizei  n, const GLuint * ids) {
//   (*fnptr)(n, ids);
// }
// static void  glowDeleteVertexArrays(GPDELETEVERTEXARRAYS fnptr, GLsizei  n, const GLuint * arrays) {
//   (*fnptr)(n, arrays);
// }
// static void  glowDepthFunc(GPDEPTHFUNC fnptr, GLenum  func) {
//   (*fnptr)(func);
// }
// static void  glowDepthMask(GPDEPTHMASK fnptr, GLboolean  flag) {
//   (*fnptr)(flag);
// }
// static void  glowDepthRange(GPDEPTHRANGE fnptr, GLdouble  xnear, GLdouble  xfar) {
//   (*fnptr)(xnear, xfar);
// }
// static void  glowDepthRangeArrayv(GPDEPTHRANGEARRAYV fnptr, GLuint  first, GLsizei  count, const GLdouble * v) {
//   (*fnptr)(first, count, v);
// }
// static void  glowDepthRangeIndexed(GPDEPTHRANGEINDEXED fnptr, GLuint  index, GLdouble  n, GLdouble  f) {
//   (*fnptr)(index, n, f);
// }
// static void  glowDepthRangef(GPDEPTHRANGEF fnptr, GLfloat  n, GLfloat  f) {
//   (*fnptr)(n, f);
// }
// static void  glowDetachShader(GPDETACHSHADER fnptr, GLuint  program, GLuint  shader) {
//   (*fnptr)(program, shader);
// }
// static void  glowDisable(GPDISABLE fnptr, GLenum  cap) {
//   (*fnptr)(cap);
// }
// static void  glowDisableClientState(GPDISABLECLIENTSTATE fnptr, GLenum  array) {
//   (*fnptr)(array);
// }
// static void  glowDisableVertexArrayAttrib(GPDISABLEVERTEXARRAYATTRIB fnptr, GLuint  vaobj, GLuint  index) {
//   (*fnptr)(vaobj, index);
// }
// static void  glowDisableVertexAttribArray(GPDISABLEVERTEXATTRIBARRAY fnptr, GLuint  index) {
//   (*fnptr)(index);
// }
// static void  glowDisablei(GPDISABLEI fnptr, GLenum  target, GLuint  index) {
//   (*fnptr)(target, index);
// }
// static void  glowDispatchCompute(GPDISPATCHCOMPUTE fnptr, GLuint  num_groups_x, GLuint  num_groups_y, GLuint  num_groups_z) {
//   (*fnptr)(num_groups_x, num_groups_y, num_groups_z);
// }
// static void  glowDispatchComputeGroupSizeARB(GPDISPATCHCOMPUTEGROUPSIZEARB fnptr, GLuint  num_groups_x, GLuint  num_groups_y, GLuint  num_groups_z, GLuint  group_size_x, GLuint  group_size_y, GLuint  group_size_z) {
//   (*fnptr)(num_groups_x, num_groups_y, num_groups_z, group_size_x, group_size_y, group_size_z);
// }
// static void  glowDispatchComputeIndirect(GPDISPATCHCOMPUTEINDIRECT fnptr, GLintptr  indirect) {
//   (*fnptr)(indirect);
// }
// static void  glowDrawArrays(GPDRAWARRAYS fnptr, GLenum  mode, GLint  first, GLsizei  count) {
//   (*fnptr)(mode, first, count);
// }
// static void  glowDrawArraysIndirect(GPDRAWARRAYSINDIRECT fnptr, GLenum  mode, const void * indirect) {
//   (*fnptr)(mode, indirect);
// }
// static void  glowDrawArraysInstanced(GPDRAWARRAYSINSTANCED fnptr, GLenum  mode, GLint  first, GLsizei  count, GLsizei  instancecount) {
//   (*fnptr)(mode, first, count, instancecount);
// }
// static void  glowDrawArraysInstancedBaseInstance(GPDRAWARRAYSINSTANCEDBASEINSTANCE fnptr, GLenum  mode, GLint  first, GLsizei  count, GLsizei  instancecount, GLuint  baseinstance) {
//   (*fnptr)(mode, first, count, instancecount, baseinstance);
// }
// static void  glowDrawBuffer(GPDRAWBUFFER fnptr, GLenum  buf) {
//   (*fnptr)(buf);
// }
// static void  glowDrawBuffers(GPDRAWBUFFERS fnptr, GLsizei  n, const GLenum * bufs) {
//   (*fnptr)(n, bufs);
// }
// static void  glowDrawElements(GPDRAWELEMENTS fnptr, GLenum  mode, GLsizei  count, GLenum  type, const void * indices) {
//   (*fnptr)(mode, count, type, indices);
// }
// static void  glowDrawElementsBaseVertex(GPDRAWELEMENTSBASEVERTEX fnptr, GLenum  mode, GLsizei  count, GLenum  type, const void * indices, GLint  basevertex) {
//   (*fnptr)(mode, count, type, indices, basevertex);
// }
// static void  glowDrawElementsIndirect(GPDRAWELEMENTSINDIRECT fnptr, GLenum  mode, GLenum  type, const void * indirect) {
//   (*fnptr)(mode, type, indirect);
// }
// static void  glowDrawElementsInstanced(GPDRAWELEMENTSINSTANCED fnptr, GLenum  mode, GLsizei  count, GLenum  type, const void * indices, GLsizei  instancecount) {
//   (*fnptr)(mode, count, type, indices, instancecount);
// }
// static void  glowDrawElementsInstancedBaseInstance(GPDRAWELEMENTSINSTANCEDBASEINSTANCE fnptr, GLenum  mode, GLsizei  count, GLenum  type, const void * indices, GLsizei  instancecount, GLuint  baseinstance) {
//   (*fnptr)(mode, count, type, indices, instancecount, baseinstance);
// }
// static void  glowDrawElementsInstancedBaseVertex(GPDRAWELEMENTSINSTANCEDBASEVERTEX fnptr, GLenum  mode, GLsizei  count, GLenum  type, const void * indices, GLsizei  instancecount, GLint  basevertex) {
//   (*fnptr)(mode, count, type, indices, instancecount, basevertex);
// }
// static void  glowDrawElementsInstancedBaseVertexBaseInstance(GPDRAWELEMENTSINSTANCEDBASEVERTEXBASEINSTANCE fnptr, GLenum  mode, GLsizei  count, GLenum  type, const void * indices, GLsizei  instancecount, GLint  basevertex, GLuint  baseinstance) {
//   (*fnptr)(mode, count, type, indices, instancecount, basevertex, baseinstance);
// }
// static void  glowDrawPixels(GPDRAWPIXELS fnptr, GLsizei  width, GLsizei  height, GLenum  format, GLenum  type, const void * pixels) {
//   (*fnptr)(width, height, format, type, pixels);
// }
// static void  glowDrawRangeElements(GPDRAWRANGEELEMENTS fnptr, GLenum  mode, GLuint  start, GLuint  end, GLsizei  count, GLenum  type, const void * indices) {
//   (*fnptr)(mode, start, end, count, type, indices);
// }
// static void  glowDrawRangeElementsBaseVertex(GPDRAWRANGEELEMENTSBASEVERTEX fnptr, GLenum  mode, GLuint  start, GLuint  end, GLsizei  count, GLenum  type, const void * indices, GLint  basevertex) {
//   (*fnptr)(mode, start, end, count, type, indices, basevertex);
// }
// static void  glowDrawTransformFeedback(GPDRAWTRANSFORMFEEDBACK fnptr, GLenum  mode, GLuint  id) {
//   (*fnptr)(mode, id);
// }
// static void  glowDrawTransformFeedbackInstanced(GPDRAWTRANSFORMFEEDBACKINSTANCED fnptr, GLenum  mode, GLuint  id, GLsizei  instancecount) {
//   (*fnptr)(mode, id, instancecount);
// }
// static void  glowDrawTransformFeedbackStream(GPDRAWTRANSFORMFEEDBACKSTREAM fnptr, GLenum  mode, GLuint  id, GLuint  stream) {
//   (*fnptr)(mode, id, stream);
// }
// static void  glowDrawTransformFeedbackStreamInstanced(GPDRAWTRANSFORMFEEDBACKSTREAMINSTANCED fnptr, GLenum  mode, GLuint  id, GLuint  stream, GLsizei  instancecount) {
//   (*fnptr)(mode, id, stream, instancecount);
// }
// static void  glowEdgeFlag(GPEDGEFLAG fnptr, GLboolean  flag) {
//   (*fnptr)(flag);
// }
// static void  glowEdgeFlagPointer(GPEDGEFLAGPOINTER fnptr, GLsizei  stride, const void * pointer) {
//   (*fnptr)(stride, pointer);
// }
// static void  glowEdgeFlagv(GPEDGEFLAGV fnptr, const GLboolean * flag) {
//   (*fnptr)(flag);
// }
// static void  glowEnable(GPENABLE fnptr, GLenum  cap) {
//   (*fnptr)(cap);
// }
// static void  glowEnableClientState(GPENABLECLIENTSTATE fnptr, GLenum  array) {
//   (*fnptr)(array);
// }
// static void  glowEnableVertexArrayAttrib(GPENABLEVERTEXARRAYATTRIB fnptr, GLuint  vaobj, GLuint  index) {
//   (*fnptr)(vaobj, index);
// }
// static void  glowEnableVertexAttribArray(GPENABLEVERTEXATTRIBARRAY fnptr, GLuint  index) {
//   (*fnptr)(index);
// }
// static void  glowEnablei(GPENABLEI fnptr, GLenum  target, GLuint  index) {
//   (*fnptr)(target, index);
// }
// static void  glowEnd(GPEND fnptr) {
//   (*fnptr)();
// }
// static void  glowEndConditionalRender(GPENDCONDITIONALRENDER fnptr) {
//   (*fnptr)();
// }
// static void  glowEndList(GPENDLIST fnptr) {
//   (*fnptr)();
// }
// static void  glowEndQuery(GPENDQUERY fnptr, GLenum  target) {
//   (*fnptr)(target);
// }
// static void  glowEndQueryIndexed(GPENDQUERYINDEXED fnptr, GLenum  target, GLuint  index) {
//   (*fnptr)(target, index);
// }
// static void  glowEndTransformFeedback(GPENDTRANSFORMFEEDBACK fnptr) {
//   (*fnptr)();
// }
// static void  glowEvalCoord1d(GPEVALCOORD1D fnptr, GLdouble  u) {
//   (*fnptr)(u);
// }
// static void  glowEvalCoord1dv(GPEVALCOORD1DV fnptr, const GLdouble * u) {
//   (*fnptr)(u);
// }
// static void  glowEvalCoord1f(GPEVALCOORD1F fnptr, GLfloat  u) {
//   (*fnptr)(u);
// }
// static void  glowEvalCoord1fv(GPEVALCOORD1FV fnptr, const GLfloat * u) {
//   (*fnptr)(u);
// }
// static void  glowEvalCoord2d(GPEVALCOORD2D fnptr, GLdouble  u, GLdouble  v) {
//   (*fnptr)(u, v);
// }
// static void  glowEvalCoord2dv(GPEVALCOORD2DV fnptr, const GLdouble * u) {
//   (*fnptr)(u);
// }
// static void  glowEvalCoord2f(GPEVALCOORD2F fnptr, GLfloat  u, GLfloat  v) {
//   (*fnptr)(u, v);
// }
// static void  glowEvalCoord2fv(GPEVALCOORD2FV fnptr, const GLfloat * u) {
//   (*fnptr)(u);
// }
// static void  glowEvalMesh1(GPEVALMESH1 fnptr, GLenum  mode, GLint  i1, GLint  i2) {
//   (*fnptr)(mode, i1, i2);
// }
// static void  glowEvalMesh2(GPEVALMESH2 fnptr, GLenum  mode, GLint  i1, GLint  i2, GLint  j1, GLint  j2) {
//   (*fnptr)(mode, i1, i2, j1, j2);
// }
// static void  glowEvalPoint1(GPEVALPOINT1 fnptr, GLint  i) {
//   (*fnptr)(i);
// }
// static void  glowEvalPoint2(GPEVALPOINT2 fnptr, GLint  i, GLint  j) {
//   (*fnptr)(i, j);
// }
// static void  glowFeedbackBuffer(GPFEEDBACKBUFFER fnptr, GLsizei  size, GLenum  type, GLfloat * buffer) {
//   (*fnptr)(size, type, buffer);
// }
// static GLsync  glowFenceSync(GPFENCESYNC fnptr, GLenum  condition, GLbitfield  flags) {
//   return (*fnptr)(condition, flags);
// }
// static void  glowFinish(GPFINISH fnptr) {
//   (*fnptr)();
// }
// static void  glowFlush(GPFLUSH fnptr) {
//   (*fnptr)();
// }
// static void  glowFlushMappedBufferRange(GPFLUSHMAPPEDBUFFERRANGE fnptr, GLenum  target, GLintptr  offset, GLsizeiptr  length) {
//   (*fnptr)(target, offset, length);
// }
// static void  glowFlushMappedNamedBufferRange(GPFLUSHMAPPEDNAMEDBUFFERRANGE fnptr, GLuint  buffer, GLintptr  offset, GLsizei  length) {
//   (*fnptr)(buffer, offset, length);
// }
// static void  glowFogCoordPointer(GPFOGCOORDPOINTER fnptr, GLenum  type, GLsizei  stride, const void * pointer) {
//   (*fnptr)(type, stride, pointer);
// }
// static void  glowFogCoordd(GPFOGCOORDD fnptr, GLdouble  coord) {
//   (*fnptr)(coord);
// }
// static void  glowFogCoorddv(GPFOGCOORDDV fnptr, const GLdouble * coord) {
//   (*fnptr)(coord);
// }
// static void  glowFogCoordf(GPFOGCOORDF fnptr, GLfloat  coord) {
//   (*fnptr)(coord);
// }
// static void  glowFogCoordfv(GPFOGCOORDFV fnptr, const GLfloat * coord) {
//   (*fnptr)(coord);
// }
// static void  glowFogf(GPFOGF fnptr, GLenum  pname, GLfloat  param) {
//   (*fnptr)(pname, param);
// }
// static void  glowFogfv(GPFOGFV fnptr, GLenum  pname, const GLfloat * params) {
//   (*fnptr)(pname, params);
// }
// static void  glowFogi(GPFOGI fnptr, GLenum  pname, GLint  param) {
//   (*fnptr)(pname, param);
// }
// static void  glowFogiv(GPFOGIV fnptr, GLenum  pname, const GLint * params) {
//   (*fnptr)(pname, params);
// }
// static void  glowFramebufferParameteri(GPFRAMEBUFFERPARAMETERI fnptr, GLenum  target, GLenum  pname, GLint  param) {
//   (*fnptr)(target, pname, param);
// }
// static void  glowFramebufferRenderbuffer(GPFRAMEBUFFERRENDERBUFFER fnptr, GLenum  target, GLenum  attachment, GLenum  renderbuffertarget, GLuint  renderbuffer) {
//   (*fnptr)(target, attachment, renderbuffertarget, renderbuffer);
// }
// static void  glowFramebufferTexture(GPFRAMEBUFFERTEXTURE fnptr, GLenum  target, GLenum  attachment, GLuint  texture, GLint  level) {
//   (*fnptr)(target, attachment, texture, level);
// }
// static void  glowFramebufferTexture1D(GPFRAMEBUFFERTEXTURE1D fnptr, GLenum  target, GLenum  attachment, GLenum  textarget, GLuint  texture, GLint  level) {
//   (*fnptr)(target, attachment, textarget, texture, level);
// }
// static void  glowFramebufferTexture2D(GPFRAMEBUFFERTEXTURE2D fnptr, GLenum  target, GLenum  attachment, GLenum  textarget, GLuint  texture, GLint  level) {
//   (*fnptr)(target, attachment, textarget, texture, level);
// }
// static void  glowFramebufferTexture3D(GPFRAMEBUFFERTEXTURE3D fnptr, GLenum  target, GLenum  attachment, GLenum  textarget, GLuint  texture, GLint  level, GLint  zoffset) {
//   (*fnptr)(target, attachment, textarget, texture, level, zoffset);
// }
// static void  glowFramebufferTextureLayer(GPFRAMEBUFFERTEXTURELAYER fnptr, GLenum  target, GLenum  attachment, GLuint  texture, GLint  level, GLint  layer) {
//   (*fnptr)(target, attachment, texture, level, layer);
// }
// static void  glowFrontFace(GPFRONTFACE fnptr, GLenum  mode) {
//   (*fnptr)(mode);
// }
// static void  glowFrustum(GPFRUSTUM fnptr, GLdouble  left, GLdouble  right, GLdouble  bottom, GLdouble  top, GLdouble  zNear, GLdouble  zFar) {
//   (*fnptr)(left, right, bottom, top, zNear, zFar);
// }
// static void  glowGenBuffers(GPGENBUFFERS fnptr, GLsizei  n, GLuint * buffers) {
//   (*fnptr)(n, buffers);
// }
// static void  glowGenFramebuffers(GPGENFRAMEBUFFERS fnptr, GLsizei  n, GLuint * framebuffers) {
//   (*fnptr)(n, framebuffers);
// }
// static GLuint  glowGenLists(GPGENLISTS fnptr, GLsizei  range) {
//   return (*fnptr)(range);
// }
// static void  glowGenProgramPipelines(GPGENPROGRAMPIPELINES fnptr, GLsizei  n, GLuint * pipelines) {
//   (*fnptr)(n, pipelines);
// }
// static void  glowGenQueries(GPGENQUERIES fnptr, GLsizei  n, GLuint * ids) {
//   (*fnptr)(n, ids);
// }
// static void  glowGenRenderbuffers(GPGENRENDERBUFFERS fnptr, GLsizei  n, GLuint * renderbuffers) {
//   (*fnptr)(n, renderbuffers);
// }
// static void  glowGenSamplers(GPGENSAMPLERS fnptr, GLsizei  count, GLuint * samplers) {
//   (*fnptr)(count, samplers);
// }
// static void  glowGenTextures(GPGENTEXTURES fnptr, GLsizei  n, GLuint * textures) {
//   (*fnptr)(n, textures);
// }
// static void  glowGenTransformFeedbacks(GPGENTRANSFORMFEEDBACKS fnptr, GLsizei  n, GLuint * ids) {
//   (*fnptr)(n, ids);
// }
// static void  glowGenVertexArrays(GPGENVERTEXARRAYS fnptr, GLsizei  n, GLuint * arrays) {
//   (*fnptr)(n, arrays);
// }
// static void  glowGenerateMipmap(GPGENERATEMIPMAP fnptr, GLenum  target) {
//   (*fnptr)(target);
// }
// static void  glowGenerateTextureMipmap(GPGENERATETEXTUREMIPMAP fnptr, GLuint  texture) {
//   (*fnptr)(texture);
// }
// static void  glowGetActiveAtomicCounterBufferiv(GPGETACTIVEATOMICCOUNTERBUFFERIV fnptr, GLuint  program, GLuint  bufferIndex, GLenum  pname, GLint * params) {
//   (*fnptr)(program, bufferIndex, pname, params);
// }
// static void  glowGetActiveAttrib(GPGETACTIVEATTRIB fnptr, GLuint  program, GLuint  index, GLsizei  bufSize, GLsizei * length, GLint * size, GLenum * type, GLchar * name) {
//   (*fnptr)(program, index, bufSize, length, size, type, name);
// }
// static void  glowGetActiveSubroutineName(GPGETACTIVESUBROUTINENAME fnptr, GLuint  program, GLenum  shadertype, GLuint  index, GLsizei  bufsize, GLsizei * length, GLchar * name) {
//   (*fnptr)(program, shadertype, index, bufsize, length, name);
// }
// static void  glowGetActiveSubroutineUniformName(GPGETACTIVESUBROUTINEUNIFORMNAME fnptr, GLuint  program, GLenum  shadertype, GLuint  index, GLsizei  bufsize, GLsizei * length, GLchar * name) {
//   (*fnptr)(program, shadertype, index, bufsize, length, name);
// }
// static void  glowGetActiveSubroutineUniformiv(GPGETACTIVESUBROUTINEUNIFORMIV fnptr, GLuint  program, GLenum  shadertype, GLuint  index, GLenum  pname, GLint * values) {
//   (*fnptr)(program, shadertype, index, pname, values);
// }
// static void  glowGetActiveUniform(GPGETACTIVEUNIFORM fnptr, GLuint  program, GLuint  index, GLsizei  bufSize, GLsizei * length, GLint * size, GLenum * type, GLchar * name) {
//   (*fnptr)(program, index, bufSize, length, size, type, name);
// }
// static void  glowGetActiveUniformBlockName(GPGETACTIVEUNIFORMBLOCKNAME fnptr, GLuint  program, GLuint  uniformBlockIndex, GLsizei  bufSize, GLsizei * length, GLchar * uniformBlockName) {
//   (*fnptr)(program, uniformBlockIndex, bufSize, length, uniformBlockName);
// }
// static void  glowGetActiveUniformBlockiv(GPGETACTIVEUNIFORMBLOCKIV fnptr, GLuint  program, GLuint  uniformBlockIndex, GLenum  pname, GLint * params) {
//   (*fnptr)(program, uniformBlockIndex, pname, params);
// }
// static void  glowGetActiveUniformName(GPGETACTIVEUNIFORMNAME fnptr, GLuint  program, GLuint  uniformIndex, GLsizei  bufSize, GLsizei * length, GLchar * uniformName) {
//   (*fnptr)(program, uniformIndex, bufSize, length, uniformName);
// }
// static void  glowGetActiveUniformsiv(GPGETACTIVEUNIFORMSIV fnptr, GLuint  program, GLsizei  uniformCount, const GLuint * uniformIndices, GLenum  pname, GLint * params) {
//   (*fnptr)(program, uniformCount, uniformIndices, pname, params);
// }
// static void  glowGetAttachedShaders(GPGETATTACHEDSHADERS fnptr, GLuint  program, GLsizei  maxCount, GLsizei * count, GLuint * shaders) {
//   (*fnptr)(program, maxCount, count, shaders);
// }
// static GLint  glowGetAttribLocation(GPGETATTRIBLOCATION fnptr, GLuint  program, const GLchar * name) {
//   return (*fnptr)(program, name);
// }
// static void  glowGetBooleani_v(GPGETBOOLEANI_V fnptr, GLenum  target, GLuint  index, GLboolean * data) {
//   (*fnptr)(target, index, data);
// }
// static void  glowGetBooleanv(GPGETBOOLEANV fnptr, GLenum  pname, GLboolean * data) {
//   (*fnptr)(pname, data);
// }
// static void  glowGetBufferParameteri64v(GPGETBUFFERPARAMETERI64V fnptr, GLenum  target, GLenum  pname, GLint64 * params) {
//   (*fnptr)(target, pname, params);
// }
// static void  glowGetBufferParameteriv(GPGETBUFFERPARAMETERIV fnptr, GLenum  target, GLenum  pname, GLint * params) {
//   (*fnptr)(target, pname, params);
// }
// static void  glowGetBufferPointerv(GPGETBUFFERPOINTERV fnptr, GLenum  target, GLenum  pname, void ** params) {
//   (*fnptr)(target, pname, params);
// }
// static void  glowGetBufferSubData(GPGETBUFFERSUBDATA fnptr, GLenum  target, GLintptr  offset, GLsizeiptr  size, void * data) {
//   (*fnptr)(target, offset, size, data);
// }
// static void  glowGetClipPlane(GPGETCLIPPLANE fnptr, GLenum  plane, GLdouble * equation) {
//   (*fnptr)(plane, equation);
// }
// static void  glowGetCompressedTexImage(GPGETCOMPRESSEDTEXIMAGE fnptr, GLenum  target, GLint  level, void * img) {
//   (*fnptr)(target, level, img);
// }
// static void  glowGetCompressedTextureImage(GPGETCOMPRESSEDTEXTUREIMAGE fnptr, GLuint  texture, GLint  level, GLsizei  bufSize, void * pixels) {
//   (*fnptr)(texture, level, bufSize, pixels);
// }
// static void  glowGetCompressedTextureSubImage(GPGETCOMPRESSEDTEXTURESUBIMAGE fnptr, GLuint  texture, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth, GLsizei  bufSize, void * pixels) {
//   (*fnptr)(texture, level, xoffset, yoffset, zoffset, width, height, depth, bufSize, pixels);
// }
// static GLuint  glowGetDebugMessageLog(GPGETDEBUGMESSAGELOG fnptr, GLuint  count, GLsizei  bufSize, GLenum * sources, GLenum * types, GLuint * ids, GLenum * severities, GLsizei * lengths, GLchar * messageLog) {
//   return (*fnptr)(count, bufSize, sources, types, ids, severities, lengths, messageLog);
// }
// static GLuint  glowGetDebugMessageLogARB(GPGETDEBUGMESSAGELOGARB fnptr, GLuint  count, GLsizei  bufSize, GLenum * sources, GLenum * types, GLuint * ids, GLenum * severities, GLsizei * lengths, GLchar * messageLog) {
//   return (*fnptr)(count, bufSize, sources, types, ids, severities, lengths, messageLog);
// }
// static GLuint  glowGetDebugMessageLogKHR(GPGETDEBUGMESSAGELOGKHR fnptr, GLuint  count, GLsizei  bufSize, GLenum * sources, GLenum * types, GLuint * ids, GLenum * severities, GLsizei * lengths, GLchar * messageLog) {
//   return (*fnptr)(count, bufSize, sources, types, ids, severities, lengths, messageLog);
// }
// static void  glowGetDoublei_v(GPGETDOUBLEI_V fnptr, GLenum  target, GLuint  index, GLdouble * data) {
//   (*fnptr)(target, index, data);
// }
// static void  glowGetDoublev(GPGETDOUBLEV fnptr, GLenum  pname, GLdouble * data) {
//   (*fnptr)(pname, data);
// }
// static GLenum  glowGetError(GPGETERROR fnptr) {
//   return (*fnptr)();
// }
// static void  glowGetFloati_v(GPGETFLOATI_V fnptr, GLenum  target, GLuint  index, GLfloat * data) {
//   (*fnptr)(target, index, data);
// }
// static void  glowGetFloatv(GPGETFLOATV fnptr, GLenum  pname, GLfloat * data) {
//   (*fnptr)(pname, data);
// }
// static GLint  glowGetFragDataIndex(GPGETFRAGDATAINDEX fnptr, GLuint  program, const GLchar * name) {
//   return (*fnptr)(program, name);
// }
// static GLint  glowGetFragDataLocation(GPGETFRAGDATALOCATION fnptr, GLuint  program, const GLchar * name) {
//   return (*fnptr)(program, name);
// }
// static void  glowGetFramebufferAttachmentParameteriv(GPGETFRAMEBUFFERATTACHMENTPARAMETERIV fnptr, GLenum  target, GLenum  attachment, GLenum  pname, GLint * params) {
//   (*fnptr)(target, attachment, pname, params);
// }
// static void  glowGetFramebufferParameteriv(GPGETFRAMEBUFFERPARAMETERIV fnptr, GLenum  target, GLenum  pname, GLint * params) {
//   (*fnptr)(target, pname, params);
// }
// static GLenum  glowGetGraphicsResetStatus(GPGETGRAPHICSRESETSTATUS fnptr) {
//   return (*fnptr)();
// }
// static GLenum  glowGetGraphicsResetStatusARB(GPGETGRAPHICSRESETSTATUSARB fnptr) {
//   return (*fnptr)();
// }
// static GLenum  glowGetGraphicsResetStatusKHR(GPGETGRAPHICSRESETSTATUSKHR fnptr) {
//   return (*fnptr)();
// }
// static GLuint64  glowGetImageHandleARB(GPGETIMAGEHANDLEARB fnptr, GLuint  texture, GLint  level, GLboolean  layered, GLint  layer, GLenum  format) {
//   return (*fnptr)(texture, level, layered, layer, format);
// }
// static void  glowGetInteger64i_v(GPGETINTEGER64I_V fnptr, GLenum  target, GLuint  index, GLint64 * data) {
//   (*fnptr)(target, index, data);
// }
// static void  glowGetInteger64v(GPGETINTEGER64V fnptr, GLenum  pname, GLint64 * data) {
//   (*fnptr)(pname, data);
// }
// static void  glowGetIntegeri_v(GPGETINTEGERI_V fnptr, GLenum  target, GLuint  index, GLint * data) {
//   (*fnptr)(target, index, data);
// }
// static void  glowGetIntegerv(GPGETINTEGERV fnptr, GLenum  pname, GLint * data) {
//   (*fnptr)(pname, data);
// }
// static void  glowGetInternalformati64v(GPGETINTERNALFORMATI64V fnptr, GLenum  target, GLenum  internalformat, GLenum  pname, GLsizei  bufSize, GLint64 * params) {
//   (*fnptr)(target, internalformat, pname, bufSize, params);
// }
// static void  glowGetInternalformativ(GPGETINTERNALFORMATIV fnptr, GLenum  target, GLenum  internalformat, GLenum  pname, GLsizei  bufSize, GLint * params) {
//   (*fnptr)(target, internalformat, pname, bufSize, params);
// }
// static void  glowGetLightfv(GPGETLIGHTFV fnptr, GLenum  light, GLenum  pname, GLfloat * params) {
//   (*fnptr)(light, pname, params);
// }
// static void  glowGetLightiv(GPGETLIGHTIV fnptr, GLenum  light, GLenum  pname, GLint * params) {
//   (*fnptr)(light, pname, params);
// }
// static void  glowGetMapdv(GPGETMAPDV fnptr, GLenum  target, GLenum  query, GLdouble * v) {
//   (*fnptr)(target, query, v);
// }
// static void  glowGetMapfv(GPGETMAPFV fnptr, GLenum  target, GLenum  query, GLfloat * v) {
//   (*fnptr)(target, query, v);
// }
// static void  glowGetMapiv(GPGETMAPIV fnptr, GLenum  target, GLenum  query, GLint * v) {
//   (*fnptr)(target, query, v);
// }
// static void  glowGetMaterialfv(GPGETMATERIALFV fnptr, GLenum  face, GLenum  pname, GLfloat * params) {
//   (*fnptr)(face, pname, params);
// }
// static void  glowGetMaterialiv(GPGETMATERIALIV fnptr, GLenum  face, GLenum  pname, GLint * params) {
//   (*fnptr)(face, pname, params);
// }
// static void  glowGetMultisamplefv(GPGETMULTISAMPLEFV fnptr, GLenum  pname, GLuint  index, GLfloat * val) {
//   (*fnptr)(pname, index, val);
// }
// static void  glowGetNamedBufferParameteri64v(GPGETNAMEDBUFFERPARAMETERI64V fnptr, GLuint  buffer, GLenum  pname, GLint64 * params) {
//   (*fnptr)(buffer, pname, params);
// }
// static void  glowGetNamedBufferParameteriv(GPGETNAMEDBUFFERPARAMETERIV fnptr, GLuint  buffer, GLenum  pname, GLint * params) {
//   (*fnptr)(buffer, pname, params);
// }
// static void  glowGetNamedBufferPointerv(GPGETNAMEDBUFFERPOINTERV fnptr, GLuint  buffer, GLenum  pname, void ** params) {
//   (*fnptr)(buffer, pname, params);
// }
// static void  glowGetNamedBufferSubData(GPGETNAMEDBUFFERSUBDATA fnptr, GLuint  buffer, GLintptr  offset, GLsizei  size, void * data) {
//   (*fnptr)(buffer, offset, size, data);
// }
// static void  glowGetNamedFramebufferAttachmentParameteriv(GPGETNAMEDFRAMEBUFFERATTACHMENTPARAMETERIV fnptr, GLuint  framebuffer, GLenum  attachment, GLenum  pname, GLint * params) {
//   (*fnptr)(framebuffer, attachment, pname, params);
// }
// static void  glowGetNamedFramebufferParameteriv(GPGETNAMEDFRAMEBUFFERPARAMETERIV fnptr, GLuint  framebuffer, GLenum  pname, GLint * param) {
//   (*fnptr)(framebuffer, pname, param);
// }
// static void  glowGetNamedRenderbufferParameteriv(GPGETNAMEDRENDERBUFFERPARAMETERIV fnptr, GLuint  renderbuffer, GLenum  pname, GLint * params) {
//   (*fnptr)(renderbuffer, pname, params);
// }
// static void  glowGetNamedStringARB(GPGETNAMEDSTRINGARB fnptr, GLint  namelen, const GLchar * name, GLsizei  bufSize, GLint * stringlen, GLchar * string) {
//   (*fnptr)(namelen, name, bufSize, stringlen, string);
// }
// static void  glowGetNamedStringivARB(GPGETNAMEDSTRINGIVARB fnptr, GLint  namelen, const GLchar * name, GLenum  pname, GLint * params) {
//   (*fnptr)(namelen, name, pname, params);
// }
// static void  glowGetObjectLabel(GPGETOBJECTLABEL fnptr, GLenum  identifier, GLuint  name, GLsizei  bufSize, GLsizei * length, GLchar * label) {
//   (*fnptr)(identifier, name, bufSize, length, label);
// }
// static void  glowGetObjectLabelKHR(GPGETOBJECTLABELKHR fnptr, GLenum  identifier, GLuint  name, GLsizei  bufSize, GLsizei * length, GLchar * label) {
//   (*fnptr)(identifier, name, bufSize, length, label);
// }
// static void  glowGetObjectPtrLabel(GPGETOBJECTPTRLABEL fnptr, const void * ptr, GLsizei  bufSize, GLsizei * length, GLchar * label) {
//   (*fnptr)(ptr, bufSize, length, label);
// }
// static void  glowGetObjectPtrLabelKHR(GPGETOBJECTPTRLABELKHR fnptr, const void * ptr, GLsizei  bufSize, GLsizei * length, GLchar * label) {
//   (*fnptr)(ptr, bufSize, length, label);
// }
// static void  glowGetPixelMapfv(GPGETPIXELMAPFV fnptr, GLenum  map, GLfloat * values) {
//   (*fnptr)(map, values);
// }
// static void  glowGetPixelMapuiv(GPGETPIXELMAPUIV fnptr, GLenum  map, GLuint * values) {
//   (*fnptr)(map, values);
// }
// static void  glowGetPixelMapusv(GPGETPIXELMAPUSV fnptr, GLenum  map, GLushort * values) {
//   (*fnptr)(map, values);
// }
// static void  glowGetPointerv(GPGETPOINTERV fnptr, GLenum  pname, void ** params) {
//   (*fnptr)(pname, params);
// }
// static void  glowGetPointervKHR(GPGETPOINTERVKHR fnptr, GLenum  pname, void ** params) {
//   (*fnptr)(pname, params);
// }
// static void  glowGetPolygonStipple(GPGETPOLYGONSTIPPLE fnptr, GLubyte * mask) {
//   (*fnptr)(mask);
// }
// static void  glowGetProgramBinary(GPGETPROGRAMBINARY fnptr, GLuint  program, GLsizei  bufSize, GLsizei * length, GLenum * binaryFormat, void * binary) {
//   (*fnptr)(program, bufSize, length, binaryFormat, binary);
// }
// static void  glowGetProgramInfoLog(GPGETPROGRAMINFOLOG fnptr, GLuint  program, GLsizei  bufSize, GLsizei * length, GLchar * infoLog) {
//   (*fnptr)(program, bufSize, length, infoLog);
// }
// static void  glowGetProgramInterfaceiv(GPGETPROGRAMINTERFACEIV fnptr, GLuint  program, GLenum  programInterface, GLenum  pname, GLint * params) {
//   (*fnptr)(program, programInterface, pname, params);
// }
// static void  glowGetProgramPipelineInfoLog(GPGETPROGRAMPIPELINEINFOLOG fnptr, GLuint  pipeline, GLsizei  bufSize, GLsizei * length, GLchar * infoLog) {
//   (*fnptr)(pipeline, bufSize, length, infoLog);
// }
// static void  glowGetProgramPipelineiv(GPGETPROGRAMPIPELINEIV fnptr, GLuint  pipeline, GLenum  pname, GLint * params) {
//   (*fnptr)(pipeline, pname, params);
// }
// static GLuint  glowGetProgramResourceIndex(GPGETPROGRAMRESOURCEINDEX fnptr, GLuint  program, GLenum  programInterface, const GLchar * name) {
//   return (*fnptr)(program, programInterface, name);
// }
// static GLint  glowGetProgramResourceLocation(GPGETPROGRAMRESOURCELOCATION fnptr, GLuint  program, GLenum  programInterface, const GLchar * name) {
//   return (*fnptr)(program, programInterface, name);
// }
// static GLint  glowGetProgramResourceLocationIndex(GPGETPROGRAMRESOURCELOCATIONINDEX fnptr, GLuint  program, GLenum  programInterface, const GLchar * name) {
//   return (*fnptr)(program, programInterface, name);
// }
// static void  glowGetProgramResourceName(GPGETPROGRAMRESOURCENAME fnptr, GLuint  program, GLenum  programInterface, GLuint  index, GLsizei  bufSize, GLsizei * length, GLchar * name) {
//   (*fnptr)(program, programInterface, index, bufSize, length, name);
// }
// static void  glowGetProgramResourceiv(GPGETPROGRAMRESOURCEIV fnptr, GLuint  program, GLenum  programInterface, GLuint  index, GLsizei  propCount, const GLenum * props, GLsizei  bufSize, GLsizei * length, GLint * params) {
//   (*fnptr)(program, programInterface, index, propCount, props, bufSize, length, params);
// }
// static void  glowGetProgramStageiv(GPGETPROGRAMSTAGEIV fnptr, GLuint  program, GLenum  shadertype, GLenum  pname, GLint * values) {
//   (*fnptr)(program, shadertype, pname, values);
// }
// static void  glowGetProgramiv(GPGETPROGRAMIV fnptr, GLuint  program, GLenum  pname, GLint * params) {
//   (*fnptr)(program, pname, params);
// }
// static void  glowGetQueryIndexediv(GPGETQUERYINDEXEDIV fnptr, GLenum  target, GLuint  index, GLenum  pname, GLint * params) {
//   (*fnptr)(target, index, pname, params);
// }
// static void  glowGetQueryObjecti64v(GPGETQUERYOBJECTI64V fnptr, GLuint  id, GLenum  pname, GLint64 * params) {
//   (*fnptr)(id, pname, params);
// }
// static void  glowGetQueryObjectiv(GPGETQUERYOBJECTIV fnptr, GLuint  id, GLenum  pname, GLint * params) {
//   (*fnptr)(id, pname, params);
// }
// static void  glowGetQueryObjectui64v(GPGETQUERYOBJECTUI64V fnptr, GLuint  id, GLenum  pname, GLuint64 * params) {
//   (*fnptr)(id, pname, params);
// }
// static void  glowGetQueryObjectuiv(GPGETQUERYOBJECTUIV fnptr, GLuint  id, GLenum  pname, GLuint * params) {
//   (*fnptr)(id, pname, params);
// }
// static void  glowGetQueryiv(GPGETQUERYIV fnptr, GLenum  target, GLenum  pname, GLint * params) {
//   (*fnptr)(target, pname, params);
// }
// static void  glowGetRenderbufferParameteriv(GPGETRENDERBUFFERPARAMETERIV fnptr, GLenum  target, GLenum  pname, GLint * params) {
//   (*fnptr)(target, pname, params);
// }
// static void  glowGetSamplerParameterIiv(GPGETSAMPLERPARAMETERIIV fnptr, GLuint  sampler, GLenum  pname, GLint * params) {
//   (*fnptr)(sampler, pname, params);
// }
// static void  glowGetSamplerParameterIuiv(GPGETSAMPLERPARAMETERIUIV fnptr, GLuint  sampler, GLenum  pname, GLuint * params) {
//   (*fnptr)(sampler, pname, params);
// }
// static void  glowGetSamplerParameterfv(GPGETSAMPLERPARAMETERFV fnptr, GLuint  sampler, GLenum  pname, GLfloat * params) {
//   (*fnptr)(sampler, pname, params);
// }
// static void  glowGetSamplerParameteriv(GPGETSAMPLERPARAMETERIV fnptr, GLuint  sampler, GLenum  pname, GLint * params) {
//   (*fnptr)(sampler, pname, params);
// }
// static void  glowGetShaderInfoLog(GPGETSHADERINFOLOG fnptr, GLuint  shader, GLsizei  bufSize, GLsizei * length, GLchar * infoLog) {
//   (*fnptr)(shader, bufSize, length, infoLog);
// }
// static void  glowGetShaderPrecisionFormat(GPGETSHADERPRECISIONFORMAT fnptr, GLenum  shadertype, GLenum  precisiontype, GLint * range, GLint * precision) {
//   (*fnptr)(shadertype, precisiontype, range, precision);
// }
// static void  glowGetShaderSource(GPGETSHADERSOURCE fnptr, GLuint  shader, GLsizei  bufSize, GLsizei * length, GLchar * source) {
//   (*fnptr)(shader, bufSize, length, source);
// }
// static void  glowGetShaderiv(GPGETSHADERIV fnptr, GLuint  shader, GLenum  pname, GLint * params) {
//   (*fnptr)(shader, pname, params);
// }
// static const GLubyte * glowGetString(GPGETSTRING fnptr, GLenum  name) {
//   return (*fnptr)(name);
// }
// static const GLubyte * glowGetStringi(GPGETSTRINGI fnptr, GLenum  name, GLuint  index) {
//   return (*fnptr)(name, index);
// }
// static GLuint  glowGetSubroutineIndex(GPGETSUBROUTINEINDEX fnptr, GLuint  program, GLenum  shadertype, const GLchar * name) {
//   return (*fnptr)(program, shadertype, name);
// }
// static GLint  glowGetSubroutineUniformLocation(GPGETSUBROUTINEUNIFORMLOCATION fnptr, GLuint  program, GLenum  shadertype, const GLchar * name) {
//   return (*fnptr)(program, shadertype, name);
// }
// static void  glowGetSynciv(GPGETSYNCIV fnptr, GLsync  sync, GLenum  pname, GLsizei  bufSize, GLsizei * length, GLint * values) {
//   (*fnptr)(sync, pname, bufSize, length, values);
// }
// static void  glowGetTexEnvfv(GPGETTEXENVFV fnptr, GLenum  target, GLenum  pname, GLfloat * params) {
//   (*fnptr)(target, pname, params);
// }
// static void  glowGetTexEnviv(GPGETTEXENVIV fnptr, GLenum  target, GLenum  pname, GLint * params) {
//   (*fnptr)(target, pname, params);
// }
// static void  glowGetTexGendv(GPGETTEXGENDV fnptr, GLenum  coord, GLenum  pname, GLdouble * params) {
//   (*fnptr)(coord, pname, params);
// }
// static void  glowGetTexGenfv(GPGETTEXGENFV fnptr, GLenum  coord, GLenum  pname, GLfloat * params) {
//   (*fnptr)(coord, pname, params);
// }
// static void  glowGetTexGeniv(GPGETTEXGENIV fnptr, GLenum  coord, GLenum  pname, GLint * params) {
//   (*fnptr)(coord, pname, params);
// }
// static void  glowGetTexImage(GPGETTEXIMAGE fnptr, GLenum  target, GLint  level, GLenum  format, GLenum  type, void * pixels) {
//   (*fnptr)(target, level, format, type, pixels);
// }
// static void  glowGetTexLevelParameterfv(GPGETTEXLEVELPARAMETERFV fnptr, GLenum  target, GLint  level, GLenum  pname, GLfloat * params) {
//   (*fnptr)(target, level, pname, params);
// }
// static void  glowGetTexLevelParameteriv(GPGETTEXLEVELPARAMETERIV fnptr, GLenum  target, GLint  level, GLenum  pname, GLint * params) {
//   (*fnptr)(target, level, pname, params);
// }
// static void  glowGetTexParameterIiv(GPGETTEXPARAMETERIIV fnptr, GLenum  target, GLenum  pname, GLint * params) {
//   (*fnptr)(target, pname, params);
// }
// static void  glowGetTexParameterIuiv(GPGETTEXPARAMETERIUIV fnptr, GLenum  target, GLenum  pname, GLuint * params) {
//   (*fnptr)(target, pname, params);
// }
// static void  glowGetTexParameterfv(GPGETTEXPARAMETERFV fnptr, GLenum  target, GLenum  pname, GLfloat * params) {
//   (*fnptr)(target, pname, params);
// }
// static void  glowGetTexParameteriv(GPGETTEXPARAMETERIV fnptr, GLenum  target, GLenum  pname, GLint * params) {
//   (*fnptr)(target, pname, params);
// }
// static GLuint64  glowGetTextureHandleARB(GPGETTEXTUREHANDLEARB fnptr, GLuint  texture) {
//   return (*fnptr)(texture);
// }
// static void  glowGetTextureImage(GPGETTEXTUREIMAGE fnptr, GLuint  texture, GLint  level, GLenum  format, GLenum  type, GLsizei  bufSize, void * pixels) {
//   (*fnptr)(texture, level, format, type, bufSize, pixels);
// }
// static void  glowGetTextureLevelParameterfv(GPGETTEXTURELEVELPARAMETERFV fnptr, GLuint  texture, GLint  level, GLenum  pname, GLfloat * params) {
//   (*fnptr)(texture, level, pname, params);
// }
// static void  glowGetTextureLevelParameteriv(GPGETTEXTURELEVELPARAMETERIV fnptr, GLuint  texture, GLint  level, GLenum  pname, GLint * params) {
//   (*fnptr)(texture, level, pname, params);
// }
// static void  glowGetTextureParameterIiv(GPGETTEXTUREPARAMETERIIV fnptr, GLuint  texture, GLenum  pname, GLint * params) {
//   (*fnptr)(texture, pname, params);
// }
// static void  glowGetTextureParameterIuiv(GPGETTEXTUREPARAMETERIUIV fnptr, GLuint  texture, GLenum  pname, GLuint * params) {
//   (*fnptr)(texture, pname, params);
// }
// static void  glowGetTextureParameterfv(GPGETTEXTUREPARAMETERFV fnptr, GLuint  texture, GLenum  pname, GLfloat * params) {
//   (*fnptr)(texture, pname, params);
// }
// static void  glowGetTextureParameteriv(GPGETTEXTUREPARAMETERIV fnptr, GLuint  texture, GLenum  pname, GLint * params) {
//   (*fnptr)(texture, pname, params);
// }
// static GLuint64  glowGetTextureSamplerHandleARB(GPGETTEXTURESAMPLERHANDLEARB fnptr, GLuint  texture, GLuint  sampler) {
//   return (*fnptr)(texture, sampler);
// }
// static void  glowGetTextureSubImage(GPGETTEXTURESUBIMAGE fnptr, GLuint  texture, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth, GLenum  format, GLenum  type, GLsizei  bufSize, void * pixels) {
//   (*fnptr)(texture, level, xoffset, yoffset, zoffset, width, height, depth, format, type, bufSize, pixels);
// }
// static void  glowGetTransformFeedbackVarying(GPGETTRANSFORMFEEDBACKVARYING fnptr, GLuint  program, GLuint  index, GLsizei  bufSize, GLsizei * length, GLsizei * size, GLenum * type, GLchar * name) {
//   (*fnptr)(program, index, bufSize, length, size, type, name);
// }
// static void  glowGetTransformFeedbacki64_v(GPGETTRANSFORMFEEDBACKI64_V fnptr, GLuint  xfb, GLenum  pname, GLuint  index, GLint64 * param) {
//   (*fnptr)(xfb, pname, index, param);
// }
// static void  glowGetTransformFeedbacki_v(GPGETTRANSFORMFEEDBACKI_V fnptr, GLuint  xfb, GLenum  pname, GLuint  index, GLint * param) {
//   (*fnptr)(xfb, pname, index, param);
// }
// static void  glowGetTransformFeedbackiv(GPGETTRANSFORMFEEDBACKIV fnptr, GLuint  xfb, GLenum  pname, GLint * param) {
//   (*fnptr)(xfb, pname, param);
// }
// static GLuint  glowGetUniformBlockIndex(GPGETUNIFORMBLOCKINDEX fnptr, GLuint  program, const GLchar * uniformBlockName) {
//   return (*fnptr)(program, uniformBlockName);
// }
// static void  glowGetUniformIndices(GPGETUNIFORMINDICES fnptr, GLuint  program, GLsizei  uniformCount, const GLchar *const* uniformNames, GLuint * uniformIndices) {
//   (*fnptr)(program, uniformCount, uniformNames, uniformIndices);
// }
// static GLint  glowGetUniformLocation(GPGETUNIFORMLOCATION fnptr, GLuint  program, const GLchar * name) {
//   return (*fnptr)(program, name);
// }
// static void  glowGetUniformSubroutineuiv(GPGETUNIFORMSUBROUTINEUIV fnptr, GLenum  shadertype, GLint  location, GLuint * params) {
//   (*fnptr)(shadertype, location, params);
// }
// static void  glowGetUniformdv(GPGETUNIFORMDV fnptr, GLuint  program, GLint  location, GLdouble * params) {
//   (*fnptr)(program, location, params);
// }
// static void  glowGetUniformfv(GPGETUNIFORMFV fnptr, GLuint  program, GLint  location, GLfloat * params) {
//   (*fnptr)(program, location, params);
// }
// static void  glowGetUniformiv(GPGETUNIFORMIV fnptr, GLuint  program, GLint  location, GLint * params) {
//   (*fnptr)(program, location, params);
// }
// static void  glowGetUniformuiv(GPGETUNIFORMUIV fnptr, GLuint  program, GLint  location, GLuint * params) {
//   (*fnptr)(program, location, params);
// }
// static void  glowGetVertexArrayIndexed64iv(GPGETVERTEXARRAYINDEXED64IV fnptr, GLuint  vaobj, GLuint  index, GLenum  pname, GLint64 * param) {
//   (*fnptr)(vaobj, index, pname, param);
// }
// static void  glowGetVertexArrayIndexediv(GPGETVERTEXARRAYINDEXEDIV fnptr, GLuint  vaobj, GLuint  index, GLenum  pname, GLint * param) {
//   (*fnptr)(vaobj, index, pname, param);
// }
// static void  glowGetVertexArrayiv(GPGETVERTEXARRAYIV fnptr, GLuint  vaobj, GLenum  pname, GLint * param) {
//   (*fnptr)(vaobj, pname, param);
// }
// static void  glowGetVertexAttribIiv(GPGETVERTEXATTRIBIIV fnptr, GLuint  index, GLenum  pname, GLint * params) {
//   (*fnptr)(index, pname, params);
// }
// static void  glowGetVertexAttribIuiv(GPGETVERTEXATTRIBIUIV fnptr, GLuint  index, GLenum  pname, GLuint * params) {
//   (*fnptr)(index, pname, params);
// }
// static void  glowGetVertexAttribLdv(GPGETVERTEXATTRIBLDV fnptr, GLuint  index, GLenum  pname, GLdouble * params) {
//   (*fnptr)(index, pname, params);
// }
// static void  glowGetVertexAttribLui64vARB(GPGETVERTEXATTRIBLUI64VARB fnptr, GLuint  index, GLenum  pname, GLuint64EXT * params) {
//   (*fnptr)(index, pname, params);
// }
// static void  glowGetVertexAttribPointerv(GPGETVERTEXATTRIBPOINTERV fnptr, GLuint  index, GLenum  pname, void ** pointer) {
//   (*fnptr)(index, pname, pointer);
// }
// static void  glowGetVertexAttribdv(GPGETVERTEXATTRIBDV fnptr, GLuint  index, GLenum  pname, GLdouble * params) {
//   (*fnptr)(index, pname, params);
// }
// static void  glowGetVertexAttribfv(GPGETVERTEXATTRIBFV fnptr, GLuint  index, GLenum  pname, GLfloat * params) {
//   (*fnptr)(index, pname, params);
// }
// static void  glowGetVertexAttribiv(GPGETVERTEXATTRIBIV fnptr, GLuint  index, GLenum  pname, GLint * params) {
//   (*fnptr)(index, pname, params);
// }
// static void  glowGetnCompressedTexImage(GPGETNCOMPRESSEDTEXIMAGE fnptr, GLenum  target, GLint  lod, GLsizei  bufSize, void * pixels) {
//   (*fnptr)(target, lod, bufSize, pixels);
// }
// static void  glowGetnCompressedTexImageARB(GPGETNCOMPRESSEDTEXIMAGEARB fnptr, GLenum  target, GLint  lod, GLsizei  bufSize, void * img) {
//   (*fnptr)(target, lod, bufSize, img);
// }
// static void  glowGetnTexImage(GPGETNTEXIMAGE fnptr, GLenum  target, GLint  level, GLenum  format, GLenum  type, GLsizei  bufSize, void * pixels) {
//   (*fnptr)(target, level, format, type, bufSize, pixels);
// }
// static void  glowGetnTexImageARB(GPGETNTEXIMAGEARB fnptr, GLenum  target, GLint  level, GLenum  format, GLenum  type, GLsizei  bufSize, void * img) {
//   (*fnptr)(target, level, format, type, bufSize, img);
// }
// static void  glowGetnUniformdv(GPGETNUNIFORMDV fnptr, GLuint  program, GLint  location, GLsizei  bufSize, GLdouble * params) {
//   (*fnptr)(program, location, bufSize, params);
// }
// static void  glowGetnUniformdvARB(GPGETNUNIFORMDVARB fnptr, GLuint  program, GLint  location, GLsizei  bufSize, GLdouble * params) {
//   (*fnptr)(program, location, bufSize, params);
// }
// static void  glowGetnUniformfv(GPGETNUNIFORMFV fnptr, GLuint  program, GLint  location, GLsizei  bufSize, GLfloat * params) {
//   (*fnptr)(program, location, bufSize, params);
// }
// static void  glowGetnUniformfvARB(GPGETNUNIFORMFVARB fnptr, GLuint  program, GLint  location, GLsizei  bufSize, GLfloat * params) {
//   (*fnptr)(program, location, bufSize, params);
// }
// static void  glowGetnUniformfvKHR(GPGETNUNIFORMFVKHR fnptr, GLuint  program, GLint  location, GLsizei  bufSize, GLfloat * params) {
//   (*fnptr)(program, location, bufSize, params);
// }
// static void  glowGetnUniformiv(GPGETNUNIFORMIV fnptr, GLuint  program, GLint  location, GLsizei  bufSize, GLint * params) {
//   (*fnptr)(program, location, bufSize, params);
// }
// static void  glowGetnUniformivARB(GPGETNUNIFORMIVARB fnptr, GLuint  program, GLint  location, GLsizei  bufSize, GLint * params) {
//   (*fnptr)(program, location, bufSize, params);
// }
// static void  glowGetnUniformivKHR(GPGETNUNIFORMIVKHR fnptr, GLuint  program, GLint  location, GLsizei  bufSize, GLint * params) {
//   (*fnptr)(program, location, bufSize, params);
// }
// static void  glowGetnUniformuiv(GPGETNUNIFORMUIV fnptr, GLuint  program, GLint  location, GLsizei  bufSize, GLuint * params) {
//   (*fnptr)(program, location, bufSize, params);
// }
// static void  glowGetnUniformuivARB(GPGETNUNIFORMUIVARB fnptr, GLuint  program, GLint  location, GLsizei  bufSize, GLuint * params) {
//   (*fnptr)(program, location, bufSize, params);
// }
// static void  glowGetnUniformuivKHR(GPGETNUNIFORMUIVKHR fnptr, GLuint  program, GLint  location, GLsizei  bufSize, GLuint * params) {
//   (*fnptr)(program, location, bufSize, params);
// }
// static void  glowHint(GPHINT fnptr, GLenum  target, GLenum  mode) {
//   (*fnptr)(target, mode);
// }
// static void  glowIndexMask(GPINDEXMASK fnptr, GLuint  mask) {
//   (*fnptr)(mask);
// }
// static void  glowIndexPointer(GPINDEXPOINTER fnptr, GLenum  type, GLsizei  stride, const void * pointer) {
//   (*fnptr)(type, stride, pointer);
// }
// static void  glowIndexd(GPINDEXD fnptr, GLdouble  c) {
//   (*fnptr)(c);
// }
// static void  glowIndexdv(GPINDEXDV fnptr, const GLdouble * c) {
//   (*fnptr)(c);
// }
// static void  glowIndexf(GPINDEXF fnptr, GLfloat  c) {
//   (*fnptr)(c);
// }
// static void  glowIndexfv(GPINDEXFV fnptr, const GLfloat * c) {
//   (*fnptr)(c);
// }
// static void  glowIndexi(GPINDEXI fnptr, GLint  c) {
//   (*fnptr)(c);
// }
// static void  glowIndexiv(GPINDEXIV fnptr, const GLint * c) {
//   (*fnptr)(c);
// }
// static void  glowIndexs(GPINDEXS fnptr, GLshort  c) {
//   (*fnptr)(c);
// }
// static void  glowIndexsv(GPINDEXSV fnptr, const GLshort * c) {
//   (*fnptr)(c);
// }
// static void  glowIndexub(GPINDEXUB fnptr, GLubyte  c) {
//   (*fnptr)(c);
// }
// static void  glowIndexubv(GPINDEXUBV fnptr, const GLubyte * c) {
//   (*fnptr)(c);
// }
// static void  glowInitNames(GPINITNAMES fnptr) {
//   (*fnptr)();
// }
// static void  glowInterleavedArrays(GPINTERLEAVEDARRAYS fnptr, GLenum  format, GLsizei  stride, const void * pointer) {
//   (*fnptr)(format, stride, pointer);
// }
// static void  glowInvalidateBufferData(GPINVALIDATEBUFFERDATA fnptr, GLuint  buffer) {
//   (*fnptr)(buffer);
// }
// static void  glowInvalidateBufferSubData(GPINVALIDATEBUFFERSUBDATA fnptr, GLuint  buffer, GLintptr  offset, GLsizeiptr  length) {
//   (*fnptr)(buffer, offset, length);
// }
// static void  glowInvalidateFramebuffer(GPINVALIDATEFRAMEBUFFER fnptr, GLenum  target, GLsizei  numAttachments, const GLenum * attachments) {
//   (*fnptr)(target, numAttachments, attachments);
// }
// static void  glowInvalidateNamedFramebufferData(GPINVALIDATENAMEDFRAMEBUFFERDATA fnptr, GLuint  framebuffer, GLsizei  numAttachments, const GLenum * attachments) {
//   (*fnptr)(framebuffer, numAttachments, attachments);
// }
// static void  glowInvalidateNamedFramebufferSubData(GPINVALIDATENAMEDFRAMEBUFFERSUBDATA fnptr, GLuint  framebuffer, GLsizei  numAttachments, const GLenum * attachments, GLint  x, GLint  y, GLsizei  width, GLsizei  height) {
//   (*fnptr)(framebuffer, numAttachments, attachments, x, y, width, height);
// }
// static void  glowInvalidateSubFramebuffer(GPINVALIDATESUBFRAMEBUFFER fnptr, GLenum  target, GLsizei  numAttachments, const GLenum * attachments, GLint  x, GLint  y, GLsizei  width, GLsizei  height) {
//   (*fnptr)(target, numAttachments, attachments, x, y, width, height);
// }
// static void  glowInvalidateTexImage(GPINVALIDATETEXIMAGE fnptr, GLuint  texture, GLint  level) {
//   (*fnptr)(texture, level);
// }
// static void  glowInvalidateTexSubImage(GPINVALIDATETEXSUBIMAGE fnptr, GLuint  texture, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth) {
//   (*fnptr)(texture, level, xoffset, yoffset, zoffset, width, height, depth);
// }
// static GLboolean  glowIsBuffer(GPISBUFFER fnptr, GLuint  buffer) {
//   return (*fnptr)(buffer);
// }
// static GLboolean  glowIsEnabled(GPISENABLED fnptr, GLenum  cap) {
//   return (*fnptr)(cap);
// }
// static GLboolean  glowIsEnabledi(GPISENABLEDI fnptr, GLenum  target, GLuint  index) {
//   return (*fnptr)(target, index);
// }
// static GLboolean  glowIsFramebuffer(GPISFRAMEBUFFER fnptr, GLuint  framebuffer) {
//   return (*fnptr)(framebuffer);
// }
// static GLboolean  glowIsImageHandleResidentARB(GPISIMAGEHANDLERESIDENTARB fnptr, GLuint64  handle) {
//   return (*fnptr)(handle);
// }
// static GLboolean  glowIsList(GPISLIST fnptr, GLuint  list) {
//   return (*fnptr)(list);
// }
// static GLboolean  glowIsNamedStringARB(GPISNAMEDSTRINGARB fnptr, GLint  namelen, const GLchar * name) {
//   return (*fnptr)(namelen, name);
// }
// static GLboolean  glowIsProgram(GPISPROGRAM fnptr, GLuint  program) {
//   return (*fnptr)(program);
// }
// static GLboolean  glowIsProgramPipeline(GPISPROGRAMPIPELINE fnptr, GLuint  pipeline) {
//   return (*fnptr)(pipeline);
// }
// static GLboolean  glowIsQuery(GPISQUERY fnptr, GLuint  id) {
//   return (*fnptr)(id);
// }
// static GLboolean  glowIsRenderbuffer(GPISRENDERBUFFER fnptr, GLuint  renderbuffer) {
//   return (*fnptr)(renderbuffer);
// }
// static GLboolean  glowIsSampler(GPISSAMPLER fnptr, GLuint  sampler) {
//   return (*fnptr)(sampler);
// }
// static GLboolean  glowIsShader(GPISSHADER fnptr, GLuint  shader) {
//   return (*fnptr)(shader);
// }
// static GLboolean  glowIsSync(GPISSYNC fnptr, GLsync  sync) {
//   return (*fnptr)(sync);
// }
// static GLboolean  glowIsTexture(GPISTEXTURE fnptr, GLuint  texture) {
//   return (*fnptr)(texture);
// }
// static GLboolean  glowIsTextureHandleResidentARB(GPISTEXTUREHANDLERESIDENTARB fnptr, GLuint64  handle) {
//   return (*fnptr)(handle);
// }
// static GLboolean  glowIsTransformFeedback(GPISTRANSFORMFEEDBACK fnptr, GLuint  id) {
//   return (*fnptr)(id);
// }
// static GLboolean  glowIsVertexArray(GPISVERTEXARRAY fnptr, GLuint  array) {
//   return (*fnptr)(array);
// }
// static void  glowLightModelf(GPLIGHTMODELF fnptr, GLenum  pname, GLfloat  param) {
//   (*fnptr)(pname, param);
// }
// static void  glowLightModelfv(GPLIGHTMODELFV fnptr, GLenum  pname, const GLfloat * params) {
//   (*fnptr)(pname, params);
// }
// static void  glowLightModeli(GPLIGHTMODELI fnptr, GLenum  pname, GLint  param) {
//   (*fnptr)(pname, param);
// }
// static void  glowLightModeliv(GPLIGHTMODELIV fnptr, GLenum  pname, const GLint * params) {
//   (*fnptr)(pname, params);
// }
// static void  glowLightf(GPLIGHTF fnptr, GLenum  light, GLenum  pname, GLfloat  param) {
//   (*fnptr)(light, pname, param);
// }
// static void  glowLightfv(GPLIGHTFV fnptr, GLenum  light, GLenum  pname, const GLfloat * params) {
//   (*fnptr)(light, pname, params);
// }
// static void  glowLighti(GPLIGHTI fnptr, GLenum  light, GLenum  pname, GLint  param) {
//   (*fnptr)(light, pname, param);
// }
// static void  glowLightiv(GPLIGHTIV fnptr, GLenum  light, GLenum  pname, const GLint * params) {
//   (*fnptr)(light, pname, params);
// }
// static void  glowLineStipple(GPLINESTIPPLE fnptr, GLint  factor, GLushort  pattern) {
//   (*fnptr)(factor, pattern);
// }
// static void  glowLineWidth(GPLINEWIDTH fnptr, GLfloat  width) {
//   (*fnptr)(width);
// }
// static void  glowLinkProgram(GPLINKPROGRAM fnptr, GLuint  program) {
//   (*fnptr)(program);
// }
// static void  glowListBase(GPLISTBASE fnptr, GLuint  base) {
//   (*fnptr)(base);
// }
// static void  glowLoadIdentity(GPLOADIDENTITY fnptr) {
//   (*fnptr)();
// }
// static void  glowLoadMatrixd(GPLOADMATRIXD fnptr, const GLdouble * m) {
//   (*fnptr)(m);
// }
// static void  glowLoadMatrixf(GPLOADMATRIXF fnptr, const GLfloat * m) {
//   (*fnptr)(m);
// }
// static void  glowLoadName(GPLOADNAME fnptr, GLuint  name) {
//   (*fnptr)(name);
// }
// static void  glowLoadTransposeMatrixd(GPLOADTRANSPOSEMATRIXD fnptr, const GLdouble * m) {
//   (*fnptr)(m);
// }
// static void  glowLoadTransposeMatrixf(GPLOADTRANSPOSEMATRIXF fnptr, const GLfloat * m) {
//   (*fnptr)(m);
// }
// static void  glowLogicOp(GPLOGICOP fnptr, GLenum  opcode) {
//   (*fnptr)(opcode);
// }
// static void  glowMakeImageHandleNonResidentARB(GPMAKEIMAGEHANDLENONRESIDENTARB fnptr, GLuint64  handle) {
//   (*fnptr)(handle);
// }
// static void  glowMakeImageHandleResidentARB(GPMAKEIMAGEHANDLERESIDENTARB fnptr, GLuint64  handle, GLenum  access) {
//   (*fnptr)(handle, access);
// }
// static void  glowMakeTextureHandleNonResidentARB(GPMAKETEXTUREHANDLENONRESIDENTARB fnptr, GLuint64  handle) {
//   (*fnptr)(handle);
// }
// static void  glowMakeTextureHandleResidentARB(GPMAKETEXTUREHANDLERESIDENTARB fnptr, GLuint64  handle) {
//   (*fnptr)(handle);
// }
// static void  glowMap1d(GPMAP1D fnptr, GLenum  target, GLdouble  u1, GLdouble  u2, GLint  stride, GLint  order, const GLdouble * points) {
//   (*fnptr)(target, u1, u2, stride, order, points);
// }
// static void  glowMap1f(GPMAP1F fnptr, GLenum  target, GLfloat  u1, GLfloat  u2, GLint  stride, GLint  order, const GLfloat * points) {
//   (*fnptr)(target, u1, u2, stride, order, points);
// }
// static void  glowMap2d(GPMAP2D fnptr, GLenum  target, GLdouble  u1, GLdouble  u2, GLint  ustride, GLint  uorder, GLdouble  v1, GLdouble  v2, GLint  vstride, GLint  vorder, const GLdouble * points) {
//   (*fnptr)(target, u1, u2, ustride, uorder, v1, v2, vstride, vorder, points);
// }
// static void  glowMap2f(GPMAP2F fnptr, GLenum  target, GLfloat  u1, GLfloat  u2, GLint  ustride, GLint  uorder, GLfloat  v1, GLfloat  v2, GLint  vstride, GLint  vorder, const GLfloat * points) {
//   (*fnptr)(target, u1, u2, ustride, uorder, v1, v2, vstride, vorder, points);
// }
// static void * glowMapBuffer(GPMAPBUFFER fnptr, GLenum  target, GLenum  access) {
//   return (*fnptr)(target, access);
// }
// static void * glowMapBufferRange(GPMAPBUFFERRANGE fnptr, GLenum  target, GLintptr  offset, GLsizeiptr  length, GLbitfield  access) {
//   return (*fnptr)(target, offset, length, access);
// }
// static void  glowMapGrid1d(GPMAPGRID1D fnptr, GLint  un, GLdouble  u1, GLdouble  u2) {
//   (*fnptr)(un, u1, u2);
// }
// static void  glowMapGrid1f(GPMAPGRID1F fnptr, GLint  un, GLfloat  u1, GLfloat  u2) {
//   (*fnptr)(un, u1, u2);
// }
// static void  glowMapGrid2d(GPMAPGRID2D fnptr, GLint  un, GLdouble  u1, GLdouble  u2, GLint  vn, GLdouble  v1, GLdouble  v2) {
//   (*fnptr)(un, u1, u2, vn, v1, v2);
// }
// static void  glowMapGrid2f(GPMAPGRID2F fnptr, GLint  un, GLfloat  u1, GLfloat  u2, GLint  vn, GLfloat  v1, GLfloat  v2) {
//   (*fnptr)(un, u1, u2, vn, v1, v2);
// }
// static void * glowMapNamedBuffer(GPMAPNAMEDBUFFER fnptr, GLuint  buffer, GLenum  access) {
//   return (*fnptr)(buffer, access);
// }
// static void * glowMapNamedBufferRange(GPMAPNAMEDBUFFERRANGE fnptr, GLuint  buffer, GLintptr  offset, GLsizei  length, GLbitfield  access) {
//   return (*fnptr)(buffer, offset, length, access);
// }
// static void  glowMaterialf(GPMATERIALF fnptr, GLenum  face, GLenum  pname, GLfloat  param) {
//   (*fnptr)(face, pname, param);
// }
// static void  glowMaterialfv(GPMATERIALFV fnptr, GLenum  face, GLenum  pname, const GLfloat * params) {
//   (*fnptr)(face, pname, params);
// }
// static void  glowMateriali(GPMATERIALI fnptr, GLenum  face, GLenum  pname, GLint  param) {
//   (*fnptr)(face, pname, param);
// }
// static void  glowMaterialiv(GPMATERIALIV fnptr, GLenum  face, GLenum  pname, const GLint * params) {
//   (*fnptr)(face, pname, params);
// }
// static void  glowMatrixMode(GPMATRIXMODE fnptr, GLenum  mode) {
//   (*fnptr)(mode);
// }
// static void  glowMemoryBarrier(GPMEMORYBARRIER fnptr, GLbitfield  barriers) {
//   (*fnptr)(barriers);
// }
// static void  glowMemoryBarrierByRegion(GPMEMORYBARRIERBYREGION fnptr, GLbitfield  barriers) {
//   (*fnptr)(barriers);
// }
// static void  glowMinSampleShading(GPMINSAMPLESHADING fnptr, GLfloat  value) {
//   (*fnptr)(value);
// }
// static void  glowMinSampleShadingARB(GPMINSAMPLESHADINGARB fnptr, GLfloat  value) {
//   (*fnptr)(value);
// }
// static void  glowMultMatrixd(GPMULTMATRIXD fnptr, const GLdouble * m) {
//   (*fnptr)(m);
// }
// static void  glowMultMatrixf(GPMULTMATRIXF fnptr, const GLfloat * m) {
//   (*fnptr)(m);
// }
// static void  glowMultTransposeMatrixd(GPMULTTRANSPOSEMATRIXD fnptr, const GLdouble * m) {
//   (*fnptr)(m);
// }
// static void  glowMultTransposeMatrixf(GPMULTTRANSPOSEMATRIXF fnptr, const GLfloat * m) {
//   (*fnptr)(m);
// }
// static void  glowMultiDrawArrays(GPMULTIDRAWARRAYS fnptr, GLenum  mode, const GLint * first, const GLsizei * count, GLsizei  drawcount) {
//   (*fnptr)(mode, first, count, drawcount);
// }
// static void  glowMultiDrawArraysIndirect(GPMULTIDRAWARRAYSINDIRECT fnptr, GLenum  mode, const void * indirect, GLsizei  drawcount, GLsizei  stride) {
//   (*fnptr)(mode, indirect, drawcount, stride);
// }
// static void  glowMultiDrawArraysIndirectCountARB(GPMULTIDRAWARRAYSINDIRECTCOUNTARB fnptr, GLenum  mode, GLintptr  indirect, GLintptr  drawcount, GLsizei  maxdrawcount, GLsizei  stride) {
//   (*fnptr)(mode, indirect, drawcount, maxdrawcount, stride);
// }
// static void  glowMultiDrawElements(GPMULTIDRAWELEMENTS fnptr, GLenum  mode, const GLsizei * count, GLenum  type, const void *const* indices, GLsizei  drawcount) {
//   (*fnptr)(mode, count, type, indices, drawcount);
// }
// static void  glowMultiDrawElementsBaseVertex(GPMULTIDRAWELEMENTSBASEVERTEX fnptr, GLenum  mode, const GLsizei * count, GLenum  type, const void *const* indices, GLsizei  drawcount, const GLint * basevertex) {
//   (*fnptr)(mode, count, type, indices, drawcount, basevertex);
// }
// static void  glowMultiDrawElementsIndirect(GPMULTIDRAWELEMENTSINDIRECT fnptr, GLenum  mode, GLenum  type, const void * indirect, GLsizei  drawcount, GLsizei  stride) {
//   (*fnptr)(mode, type, indirect, drawcount, stride);
// }
// static void  glowMultiDrawElementsIndirectCountARB(GPMULTIDRAWELEMENTSINDIRECTCOUNTARB fnptr, GLenum  mode, GLenum  type, GLintptr  indirect, GLintptr  drawcount, GLsizei  maxdrawcount, GLsizei  stride) {
//   (*fnptr)(mode, type, indirect, drawcount, maxdrawcount, stride);
// }
// static void  glowMultiTexCoord1d(GPMULTITEXCOORD1D fnptr, GLenum  target, GLdouble  s) {
//   (*fnptr)(target, s);
// }
// static void  glowMultiTexCoord1dv(GPMULTITEXCOORD1DV fnptr, GLenum  target, const GLdouble * v) {
//   (*fnptr)(target, v);
// }
// static void  glowMultiTexCoord1f(GPMULTITEXCOORD1F fnptr, GLenum  target, GLfloat  s) {
//   (*fnptr)(target, s);
// }
// static void  glowMultiTexCoord1fv(GPMULTITEXCOORD1FV fnptr, GLenum  target, const GLfloat * v) {
//   (*fnptr)(target, v);
// }
// static void  glowMultiTexCoord1i(GPMULTITEXCOORD1I fnptr, GLenum  target, GLint  s) {
//   (*fnptr)(target, s);
// }
// static void  glowMultiTexCoord1iv(GPMULTITEXCOORD1IV fnptr, GLenum  target, const GLint * v) {
//   (*fnptr)(target, v);
// }
// static void  glowMultiTexCoord1s(GPMULTITEXCOORD1S fnptr, GLenum  target, GLshort  s) {
//   (*fnptr)(target, s);
// }
// static void  glowMultiTexCoord1sv(GPMULTITEXCOORD1SV fnptr, GLenum  target, const GLshort * v) {
//   (*fnptr)(target, v);
// }
// static void  glowMultiTexCoord2d(GPMULTITEXCOORD2D fnptr, GLenum  target, GLdouble  s, GLdouble  t) {
//   (*fnptr)(target, s, t);
// }
// static void  glowMultiTexCoord2dv(GPMULTITEXCOORD2DV fnptr, GLenum  target, const GLdouble * v) {
//   (*fnptr)(target, v);
// }
// static void  glowMultiTexCoord2f(GPMULTITEXCOORD2F fnptr, GLenum  target, GLfloat  s, GLfloat  t) {
//   (*fnptr)(target, s, t);
// }
// static void  glowMultiTexCoord2fv(GPMULTITEXCOORD2FV fnptr, GLenum  target, const GLfloat * v) {
//   (*fnptr)(target, v);
// }
// static void  glowMultiTexCoord2i(GPMULTITEXCOORD2I fnptr, GLenum  target, GLint  s, GLint  t) {
//   (*fnptr)(target, s, t);
// }
// static void  glowMultiTexCoord2iv(GPMULTITEXCOORD2IV fnptr, GLenum  target, const GLint * v) {
//   (*fnptr)(target, v);
// }
// static void  glowMultiTexCoord2s(GPMULTITEXCOORD2S fnptr, GLenum  target, GLshort  s, GLshort  t) {
//   (*fnptr)(target, s, t);
// }
// static void  glowMultiTexCoord2sv(GPMULTITEXCOORD2SV fnptr, GLenum  target, const GLshort * v) {
//   (*fnptr)(target, v);
// }
// static void  glowMultiTexCoord3d(GPMULTITEXCOORD3D fnptr, GLenum  target, GLdouble  s, GLdouble  t, GLdouble  r) {
//   (*fnptr)(target, s, t, r);
// }
// static void  glowMultiTexCoord3dv(GPMULTITEXCOORD3DV fnptr, GLenum  target, const GLdouble * v) {
//   (*fnptr)(target, v);
// }
// static void  glowMultiTexCoord3f(GPMULTITEXCOORD3F fnptr, GLenum  target, GLfloat  s, GLfloat  t, GLfloat  r) {
//   (*fnptr)(target, s, t, r);
// }
// static void  glowMultiTexCoord3fv(GPMULTITEXCOORD3FV fnptr, GLenum  target, const GLfloat * v) {
//   (*fnptr)(target, v);
// }
// static void  glowMultiTexCoord3i(GPMULTITEXCOORD3I fnptr, GLenum  target, GLint  s, GLint  t, GLint  r) {
//   (*fnptr)(target, s, t, r);
// }
// static void  glowMultiTexCoord3iv(GPMULTITEXCOORD3IV fnptr, GLenum  target, const GLint * v) {
//   (*fnptr)(target, v);
// }
// static void  glowMultiTexCoord3s(GPMULTITEXCOORD3S fnptr, GLenum  target, GLshort  s, GLshort  t, GLshort  r) {
//   (*fnptr)(target, s, t, r);
// }
// static void  glowMultiTexCoord3sv(GPMULTITEXCOORD3SV fnptr, GLenum  target, const GLshort * v) {
//   (*fnptr)(target, v);
// }
// static void  glowMultiTexCoord4d(GPMULTITEXCOORD4D fnptr, GLenum  target, GLdouble  s, GLdouble  t, GLdouble  r, GLdouble  q) {
//   (*fnptr)(target, s, t, r, q);
// }
// static void  glowMultiTexCoord4dv(GPMULTITEXCOORD4DV fnptr, GLenum  target, const GLdouble * v) {
//   (*fnptr)(target, v);
// }
// static void  glowMultiTexCoord4f(GPMULTITEXCOORD4F fnptr, GLenum  target, GLfloat  s, GLfloat  t, GLfloat  r, GLfloat  q) {
//   (*fnptr)(target, s, t, r, q);
// }
// static void  glowMultiTexCoord4fv(GPMULTITEXCOORD4FV fnptr, GLenum  target, const GLfloat * v) {
//   (*fnptr)(target, v);
// }
// static void  glowMultiTexCoord4i(GPMULTITEXCOORD4I fnptr, GLenum  target, GLint  s, GLint  t, GLint  r, GLint  q) {
//   (*fnptr)(target, s, t, r, q);
// }
// static void  glowMultiTexCoord4iv(GPMULTITEXCOORD4IV fnptr, GLenum  target, const GLint * v) {
//   (*fnptr)(target, v);
// }
// static void  glowMultiTexCoord4s(GPMULTITEXCOORD4S fnptr, GLenum  target, GLshort  s, GLshort  t, GLshort  r, GLshort  q) {
//   (*fnptr)(target, s, t, r, q);
// }
// static void  glowMultiTexCoord4sv(GPMULTITEXCOORD4SV fnptr, GLenum  target, const GLshort * v) {
//   (*fnptr)(target, v);
// }
// static void  glowNamedBufferData(GPNAMEDBUFFERDATA fnptr, GLuint  buffer, GLsizei  size, const void * data, GLenum  usage) {
//   (*fnptr)(buffer, size, data, usage);
// }
// static void  glowNamedBufferPageCommitmentARB(GPNAMEDBUFFERPAGECOMMITMENTARB fnptr, GLuint  buffer, GLintptr  offset, GLsizei  size, GLboolean  commit) {
//   (*fnptr)(buffer, offset, size, commit);
// }
// static void  glowNamedBufferPageCommitmentEXT(GPNAMEDBUFFERPAGECOMMITMENTEXT fnptr, GLuint  buffer, GLintptr  offset, GLsizei  size, GLboolean  commit) {
//   (*fnptr)(buffer, offset, size, commit);
// }
// static void  glowNamedBufferStorage(GPNAMEDBUFFERSTORAGE fnptr, GLuint  buffer, GLsizei  size, const void * data, GLbitfield  flags) {
//   (*fnptr)(buffer, size, data, flags);
// }
// static void  glowNamedBufferSubData(GPNAMEDBUFFERSUBDATA fnptr, GLuint  buffer, GLintptr  offset, GLsizei  size, const void * data) {
//   (*fnptr)(buffer, offset, size, data);
// }
// static void  glowNamedFramebufferDrawBuffer(GPNAMEDFRAMEBUFFERDRAWBUFFER fnptr, GLuint  framebuffer, GLenum  buf) {
//   (*fnptr)(framebuffer, buf);
// }
// static void  glowNamedFramebufferDrawBuffers(GPNAMEDFRAMEBUFFERDRAWBUFFERS fnptr, GLuint  framebuffer, GLsizei  n, const GLenum * bufs) {
//   (*fnptr)(framebuffer, n, bufs);
// }
// static void  glowNamedFramebufferParameteri(GPNAMEDFRAMEBUFFERPARAMETERI fnptr, GLuint  framebuffer, GLenum  pname, GLint  param) {
//   (*fnptr)(framebuffer, pname, param);
// }
// static void  glowNamedFramebufferReadBuffer(GPNAMEDFRAMEBUFFERREADBUFFER fnptr, GLuint  framebuffer, GLenum  src) {
//   (*fnptr)(framebuffer, src);
// }
// static void  glowNamedFramebufferRenderbuffer(GPNAMEDFRAMEBUFFERRENDERBUFFER fnptr, GLuint  framebuffer, GLenum  attachment, GLenum  renderbuffertarget, GLuint  renderbuffer) {
//   (*fnptr)(framebuffer, attachment, renderbuffertarget, renderbuffer);
// }
// static void  glowNamedFramebufferTexture(GPNAMEDFRAMEBUFFERTEXTURE fnptr, GLuint  framebuffer, GLenum  attachment, GLuint  texture, GLint  level) {
//   (*fnptr)(framebuffer, attachment, texture, level);
// }
// static void  glowNamedFramebufferTextureLayer(GPNAMEDFRAMEBUFFERTEXTURELAYER fnptr, GLuint  framebuffer, GLenum  attachment, GLuint  texture, GLint  level, GLint  layer) {
//   (*fnptr)(framebuffer, attachment, texture, level, layer);
// }
// static void  glowNamedRenderbufferStorage(GPNAMEDRENDERBUFFERSTORAGE fnptr, GLuint  renderbuffer, GLenum  internalformat, GLsizei  width, GLsizei  height) {
//   (*fnptr)(renderbuffer, internalformat, width, height);
// }
// static void  glowNamedRenderbufferStorageMultisample(GPNAMEDRENDERBUFFERSTORAGEMULTISAMPLE fnptr, GLuint  renderbuffer, GLsizei  samples, GLenum  internalformat, GLsizei  width, GLsizei  height) {
//   (*fnptr)(renderbuffer, samples, internalformat, width, height);
// }
// static void  glowNamedStringARB(GPNAMEDSTRINGARB fnptr, GLenum  type, GLint  namelen, const GLchar * name, GLint  stringlen, const GLchar * string) {
//   (*fnptr)(type, namelen, name, stringlen, string);
// }
// static void  glowNewList(GPNEWLIST fnptr, GLuint  list, GLenum  mode) {
//   (*fnptr)(list, mode);
// }
// static void  glowNormal3b(GPNORMAL3B fnptr, GLbyte  nx, GLbyte  ny, GLbyte  nz) {
//   (*fnptr)(nx, ny, nz);
// }
// static void  glowNormal3bv(GPNORMAL3BV fnptr, const GLbyte * v) {
//   (*fnptr)(v);
// }
// static void  glowNormal3d(GPNORMAL3D fnptr, GLdouble  nx, GLdouble  ny, GLdouble  nz) {
//   (*fnptr)(nx, ny, nz);
// }
// static void  glowNormal3dv(GPNORMAL3DV fnptr, const GLdouble * v) {
//   (*fnptr)(v);
// }
// static void  glowNormal3f(GPNORMAL3F fnptr, GLfloat  nx, GLfloat  ny, GLfloat  nz) {
//   (*fnptr)(nx, ny, nz);
// }
// static void  glowNormal3fv(GPNORMAL3FV fnptr, const GLfloat * v) {
//   (*fnptr)(v);
// }
// static void  glowNormal3i(GPNORMAL3I fnptr, GLint  nx, GLint  ny, GLint  nz) {
//   (*fnptr)(nx, ny, nz);
// }
// static void  glowNormal3iv(GPNORMAL3IV fnptr, const GLint * v) {
//   (*fnptr)(v);
// }
// static void  glowNormal3s(GPNORMAL3S fnptr, GLshort  nx, GLshort  ny, GLshort  nz) {
//   (*fnptr)(nx, ny, nz);
// }
// static void  glowNormal3sv(GPNORMAL3SV fnptr, const GLshort * v) {
//   (*fnptr)(v);
// }
// static void  glowNormalPointer(GPNORMALPOINTER fnptr, GLenum  type, GLsizei  stride, const void * pointer) {
//   (*fnptr)(type, stride, pointer);
// }
// static void  glowObjectLabel(GPOBJECTLABEL fnptr, GLenum  identifier, GLuint  name, GLsizei  length, const GLchar * label) {
//   (*fnptr)(identifier, name, length, label);
// }
// static void  glowObjectLabelKHR(GPOBJECTLABELKHR fnptr, GLenum  identifier, GLuint  name, GLsizei  length, const GLchar * label) {
//   (*fnptr)(identifier, name, length, label);
// }
// static void  glowObjectPtrLabel(GPOBJECTPTRLABEL fnptr, const void * ptr, GLsizei  length, const GLchar * label) {
//   (*fnptr)(ptr, length, label);
// }
// static void  glowObjectPtrLabelKHR(GPOBJECTPTRLABELKHR fnptr, const void * ptr, GLsizei  length, const GLchar * label) {
//   (*fnptr)(ptr, length, label);
// }
// static void  glowOrtho(GPORTHO fnptr, GLdouble  left, GLdouble  right, GLdouble  bottom, GLdouble  top, GLdouble  zNear, GLdouble  zFar) {
//   (*fnptr)(left, right, bottom, top, zNear, zFar);
// }
// static void  glowPassThrough(GPPASSTHROUGH fnptr, GLfloat  token) {
//   (*fnptr)(token);
// }
// static void  glowPatchParameterfv(GPPATCHPARAMETERFV fnptr, GLenum  pname, const GLfloat * values) {
//   (*fnptr)(pname, values);
// }
// static void  glowPatchParameteri(GPPATCHPARAMETERI fnptr, GLenum  pname, GLint  value) {
//   (*fnptr)(pname, value);
// }
// static void  glowPauseTransformFeedback(GPPAUSETRANSFORMFEEDBACK fnptr) {
//   (*fnptr)();
// }
// static void  glowPixelMapfv(GPPIXELMAPFV fnptr, GLenum  map, GLsizei  mapsize, const GLfloat * values) {
//   (*fnptr)(map, mapsize, values);
// }
// static void  glowPixelMapuiv(GPPIXELMAPUIV fnptr, GLenum  map, GLsizei  mapsize, const GLuint * values) {
//   (*fnptr)(map, mapsize, values);
// }
// static void  glowPixelMapusv(GPPIXELMAPUSV fnptr, GLenum  map, GLsizei  mapsize, const GLushort * values) {
//   (*fnptr)(map, mapsize, values);
// }
// static void  glowPixelStoref(GPPIXELSTOREF fnptr, GLenum  pname, GLfloat  param) {
//   (*fnptr)(pname, param);
// }
// static void  glowPixelStorei(GPPIXELSTOREI fnptr, GLenum  pname, GLint  param) {
//   (*fnptr)(pname, param);
// }
// static void  glowPixelTransferf(GPPIXELTRANSFERF fnptr, GLenum  pname, GLfloat  param) {
//   (*fnptr)(pname, param);
// }
// static void  glowPixelTransferi(GPPIXELTRANSFERI fnptr, GLenum  pname, GLint  param) {
//   (*fnptr)(pname, param);
// }
// static void  glowPixelZoom(GPPIXELZOOM fnptr, GLfloat  xfactor, GLfloat  yfactor) {
//   (*fnptr)(xfactor, yfactor);
// }
// static void  glowPointParameterf(GPPOINTPARAMETERF fnptr, GLenum  pname, GLfloat  param) {
//   (*fnptr)(pname, param);
// }
// static void  glowPointParameterfv(GPPOINTPARAMETERFV fnptr, GLenum  pname, const GLfloat * params) {
//   (*fnptr)(pname, params);
// }
// static void  glowPointParameteri(GPPOINTPARAMETERI fnptr, GLenum  pname, GLint  param) {
//   (*fnptr)(pname, param);
// }
// static void  glowPointParameteriv(GPPOINTPARAMETERIV fnptr, GLenum  pname, const GLint * params) {
//   (*fnptr)(pname, params);
// }
// static void  glowPointSize(GPPOINTSIZE fnptr, GLfloat  size) {
//   (*fnptr)(size);
// }
// static void  glowPolygonMode(GPPOLYGONMODE fnptr, GLenum  face, GLenum  mode) {
//   (*fnptr)(face, mode);
// }
// static void  glowPolygonOffset(GPPOLYGONOFFSET fnptr, GLfloat  factor, GLfloat  units) {
//   (*fnptr)(factor, units);
// }
// static void  glowPolygonStipple(GPPOLYGONSTIPPLE fnptr, const GLubyte * mask) {
//   (*fnptr)(mask);
// }
// static void  glowPopAttrib(GPPOPATTRIB fnptr) {
//   (*fnptr)();
// }
// static void  glowPopClientAttrib(GPPOPCLIENTATTRIB fnptr) {
//   (*fnptr)();
// }
// static void  glowPopDebugGroup(GPPOPDEBUGGROUP fnptr) {
//   (*fnptr)();
// }
// static void  glowPopDebugGroupKHR(GPPOPDEBUGGROUPKHR fnptr) {
//   (*fnptr)();
// }
// static void  glowPopMatrix(GPPOPMATRIX fnptr) {
//   (*fnptr)();
// }
// static void  glowPopName(GPPOPNAME fnptr) {
//   (*fnptr)();
// }
// static void  glowPrimitiveRestartIndex(GPPRIMITIVERESTARTINDEX fnptr, GLuint  index) {
//   (*fnptr)(index);
// }
// static void  glowPrioritizeTextures(GPPRIORITIZETEXTURES fnptr, GLsizei  n, const GLuint * textures, const GLfloat * priorities) {
//   (*fnptr)(n, textures, priorities);
// }
// static void  glowProgramBinary(GPPROGRAMBINARY fnptr, GLuint  program, GLenum  binaryFormat, const void * binary, GLsizei  length) {
//   (*fnptr)(program, binaryFormat, binary, length);
// }
// static void  glowProgramParameteri(GPPROGRAMPARAMETERI fnptr, GLuint  program, GLenum  pname, GLint  value) {
//   (*fnptr)(program, pname, value);
// }
// static void  glowProgramUniform1d(GPPROGRAMUNIFORM1D fnptr, GLuint  program, GLint  location, GLdouble  v0) {
//   (*fnptr)(program, location, v0);
// }
// static void  glowProgramUniform1dv(GPPROGRAMUNIFORM1DV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLdouble * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform1f(GPPROGRAMUNIFORM1F fnptr, GLuint  program, GLint  location, GLfloat  v0) {
//   (*fnptr)(program, location, v0);
// }
// static void  glowProgramUniform1fv(GPPROGRAMUNIFORM1FV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLfloat * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform1i(GPPROGRAMUNIFORM1I fnptr, GLuint  program, GLint  location, GLint  v0) {
//   (*fnptr)(program, location, v0);
// }
// static void  glowProgramUniform1iv(GPPROGRAMUNIFORM1IV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLint * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform1ui(GPPROGRAMUNIFORM1UI fnptr, GLuint  program, GLint  location, GLuint  v0) {
//   (*fnptr)(program, location, v0);
// }
// static void  glowProgramUniform1uiv(GPPROGRAMUNIFORM1UIV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLuint * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform2d(GPPROGRAMUNIFORM2D fnptr, GLuint  program, GLint  location, GLdouble  v0, GLdouble  v1) {
//   (*fnptr)(program, location, v0, v1);
// }
// static void  glowProgramUniform2dv(GPPROGRAMUNIFORM2DV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLdouble * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform2f(GPPROGRAMUNIFORM2F fnptr, GLuint  program, GLint  location, GLfloat  v0, GLfloat  v1) {
//   (*fnptr)(program, location, v0, v1);
// }
// static void  glowProgramUniform2fv(GPPROGRAMUNIFORM2FV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLfloat * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform2i(GPPROGRAMUNIFORM2I fnptr, GLuint  program, GLint  location, GLint  v0, GLint  v1) {
//   (*fnptr)(program, location, v0, v1);
// }
// static void  glowProgramUniform2iv(GPPROGRAMUNIFORM2IV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLint * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform2ui(GPPROGRAMUNIFORM2UI fnptr, GLuint  program, GLint  location, GLuint  v0, GLuint  v1) {
//   (*fnptr)(program, location, v0, v1);
// }
// static void  glowProgramUniform2uiv(GPPROGRAMUNIFORM2UIV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLuint * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform3d(GPPROGRAMUNIFORM3D fnptr, GLuint  program, GLint  location, GLdouble  v0, GLdouble  v1, GLdouble  v2) {
//   (*fnptr)(program, location, v0, v1, v2);
// }
// static void  glowProgramUniform3dv(GPPROGRAMUNIFORM3DV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLdouble * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform3f(GPPROGRAMUNIFORM3F fnptr, GLuint  program, GLint  location, GLfloat  v0, GLfloat  v1, GLfloat  v2) {
//   (*fnptr)(program, location, v0, v1, v2);
// }
// static void  glowProgramUniform3fv(GPPROGRAMUNIFORM3FV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLfloat * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform3i(GPPROGRAMUNIFORM3I fnptr, GLuint  program, GLint  location, GLint  v0, GLint  v1, GLint  v2) {
//   (*fnptr)(program, location, v0, v1, v2);
// }
// static void  glowProgramUniform3iv(GPPROGRAMUNIFORM3IV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLint * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform3ui(GPPROGRAMUNIFORM3UI fnptr, GLuint  program, GLint  location, GLuint  v0, GLuint  v1, GLuint  v2) {
//   (*fnptr)(program, location, v0, v1, v2);
// }
// static void  glowProgramUniform3uiv(GPPROGRAMUNIFORM3UIV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLuint * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform4d(GPPROGRAMUNIFORM4D fnptr, GLuint  program, GLint  location, GLdouble  v0, GLdouble  v1, GLdouble  v2, GLdouble  v3) {
//   (*fnptr)(program, location, v0, v1, v2, v3);
// }
// static void  glowProgramUniform4dv(GPPROGRAMUNIFORM4DV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLdouble * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform4f(GPPROGRAMUNIFORM4F fnptr, GLuint  program, GLint  location, GLfloat  v0, GLfloat  v1, GLfloat  v2, GLfloat  v3) {
//   (*fnptr)(program, location, v0, v1, v2, v3);
// }
// static void  glowProgramUniform4fv(GPPROGRAMUNIFORM4FV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLfloat * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform4i(GPPROGRAMUNIFORM4I fnptr, GLuint  program, GLint  location, GLint  v0, GLint  v1, GLint  v2, GLint  v3) {
//   (*fnptr)(program, location, v0, v1, v2, v3);
// }
// static void  glowProgramUniform4iv(GPPROGRAMUNIFORM4IV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLint * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform4ui(GPPROGRAMUNIFORM4UI fnptr, GLuint  program, GLint  location, GLuint  v0, GLuint  v1, GLuint  v2, GLuint  v3) {
//   (*fnptr)(program, location, v0, v1, v2, v3);
// }
// static void  glowProgramUniform4uiv(GPPROGRAMUNIFORM4UIV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLuint * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniformHandleui64ARB(GPPROGRAMUNIFORMHANDLEUI64ARB fnptr, GLuint  program, GLint  location, GLuint64  value) {
//   (*fnptr)(program, location, value);
// }
// static void  glowProgramUniformHandleui64vARB(GPPROGRAMUNIFORMHANDLEUI64VARB fnptr, GLuint  program, GLint  location, GLsizei  count, const GLuint64 * values) {
//   (*fnptr)(program, location, count, values);
// }
// static void  glowProgramUniformMatrix2dv(GPPROGRAMUNIFORMMATRIX2DV fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value) {
//   (*fnptr)(program, location, count, transpose, value);
// }
// static void  glowProgramUniformMatrix2fv(GPPROGRAMUNIFORMMATRIX2FV fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(program, location, count, transpose, value);
// }
// static void  glowProgramUniformMatrix2x3dv(GPPROGRAMUNIFORMMATRIX2X3DV fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value) {
//   (*fnptr)(program, location, count, transpose, value);
// }
// static void  glowProgramUniformMatrix2x3fv(GPPROGRAMUNIFORMMATRIX2X3FV fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(program, location, count, transpose, value);
// }
// static void  glowProgramUniformMatrix2x4dv(GPPROGRAMUNIFORMMATRIX2X4DV fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value) {
//   (*fnptr)(program, location, count, transpose, value);
// }
// static void  glowProgramUniformMatrix2x4fv(GPPROGRAMUNIFORMMATRIX2X4FV fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(program, location, count, transpose, value);
// }
// static void  glowProgramUniformMatrix3dv(GPPROGRAMUNIFORMMATRIX3DV fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value) {
//   (*fnptr)(program, location, count, transpose, value);
// }
// static void  glowProgramUniformMatrix3fv(GPPROGRAMUNIFORMMATRIX3FV fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(program, location, count, transpose, value);
// }
// static void  glowProgramUniformMatrix3x2dv(GPPROGRAMUNIFORMMATRIX3X2DV fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value) {
//   (*fnptr)(program, location, count, transpose, value);
// }
// static void  glowProgramUniformMatrix3x2fv(GPPROGRAMUNIFORMMATRIX3X2FV fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(program, location, count, transpose, value);
// }
// static void  glowProgramUniformMatrix3x4dv(GPPROGRAMUNIFORMMATRIX3X4DV fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value) {
//   (*fnptr)(program, location, count, transpose, value);
// }
// static void  glowProgramUniformMatrix3x4fv(GPPROGRAMUNIFORMMATRIX3X4FV fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(program, location, count, transpose, value);
// }
// static void  glowProgramUniformMatrix4dv(GPPROGRAMUNIFORMMATRIX4DV fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value) {
//   (*fnptr)(program, location, count, transpose, value);
// }
// static void  glowProgramUniformMatrix4fv(GPPROGRAMUNIFORMMATRIX4FV fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(program, location, count, transpose, value);
// }
// static void  glowProgramUniformMatrix4x2dv(GPPROGRAMUNIFORMMATRIX4X2DV fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value) {
//   (*fnptr)(program, location, count, transpose, value);
// }
// static void  glowProgramUniformMatrix4x2fv(GPPROGRAMUNIFORMMATRIX4X2FV fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(program, location, count, transpose, value);
// }
// static void  glowProgramUniformMatrix4x3dv(GPPROGRAMUNIFORMMATRIX4X3DV fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value) {
//   (*fnptr)(program, location, count, transpose, value);
// }
// static void  glowProgramUniformMatrix4x3fv(GPPROGRAMUNIFORMMATRIX4X3FV fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(program, location, count, transpose, value);
// }
// static void  glowProvokingVertex(GPPROVOKINGVERTEX fnptr, GLenum  mode) {
//   (*fnptr)(mode);
// }
// static void  glowPushAttrib(GPPUSHATTRIB fnptr, GLbitfield  mask) {
//   (*fnptr)(mask);
// }
// static void  glowPushClientAttrib(GPPUSHCLIENTATTRIB fnptr, GLbitfield  mask) {
//   (*fnptr)(mask);
// }
// static void  glowPushDebugGroup(GPPUSHDEBUGGROUP fnptr, GLenum  source, GLuint  id, GLsizei  length, const GLchar * message) {
//   (*fnptr)(source, id, length, message);
// }
// static void  glowPushDebugGroupKHR(GPPUSHDEBUGGROUPKHR fnptr, GLenum  source, GLuint  id, GLsizei  length, const GLchar * message) {
//   (*fnptr)(source, id, length, message);
// }
// static void  glowPushMatrix(GPPUSHMATRIX fnptr) {
//   (*fnptr)();
// }
// static void  glowPushName(GPPUSHNAME fnptr, GLuint  name) {
//   (*fnptr)(name);
// }
// static void  glowQueryCounter(GPQUERYCOUNTER fnptr, GLuint  id, GLenum  target) {
//   (*fnptr)(id, target);
// }
// static void  glowRasterPos2d(GPRASTERPOS2D fnptr, GLdouble  x, GLdouble  y) {
//   (*fnptr)(x, y);
// }
// static void  glowRasterPos2dv(GPRASTERPOS2DV fnptr, const GLdouble * v) {
//   (*fnptr)(v);
// }
// static void  glowRasterPos2f(GPRASTERPOS2F fnptr, GLfloat  x, GLfloat  y) {
//   (*fnptr)(x, y);
// }
// static void  glowRasterPos2fv(GPRASTERPOS2FV fnptr, const GLfloat * v) {
//   (*fnptr)(v);
// }
// static void  glowRasterPos2i(GPRASTERPOS2I fnptr, GLint  x, GLint  y) {
//   (*fnptr)(x, y);
// }
// static void  glowRasterPos2iv(GPRASTERPOS2IV fnptr, const GLint * v) {
//   (*fnptr)(v);
// }
// static void  glowRasterPos2s(GPRASTERPOS2S fnptr, GLshort  x, GLshort  y) {
//   (*fnptr)(x, y);
// }
// static void  glowRasterPos2sv(GPRASTERPOS2SV fnptr, const GLshort * v) {
//   (*fnptr)(v);
// }
// static void  glowRasterPos3d(GPRASTERPOS3D fnptr, GLdouble  x, GLdouble  y, GLdouble  z) {
//   (*fnptr)(x, y, z);
// }
// static void  glowRasterPos3dv(GPRASTERPOS3DV fnptr, const GLdouble * v) {
//   (*fnptr)(v);
// }
// static void  glowRasterPos3f(GPRASTERPOS3F fnptr, GLfloat  x, GLfloat  y, GLfloat  z) {
//   (*fnptr)(x, y, z);
// }
// static void  glowRasterPos3fv(GPRASTERPOS3FV fnptr, const GLfloat * v) {
//   (*fnptr)(v);
// }
// static void  glowRasterPos3i(GPRASTERPOS3I fnptr, GLint  x, GLint  y, GLint  z) {
//   (*fnptr)(x, y, z);
// }
// static void  glowRasterPos3iv(GPRASTERPOS3IV fnptr, const GLint * v) {
//   (*fnptr)(v);
// }
// static void  glowRasterPos3s(GPRASTERPOS3S fnptr, GLshort  x, GLshort  y, GLshort  z) {
//   (*fnptr)(x, y, z);
// }
// static void  glowRasterPos3sv(GPRASTERPOS3SV fnptr, const GLshort * v) {
//   (*fnptr)(v);
// }
// static void  glowRasterPos4d(GPRASTERPOS4D fnptr, GLdouble  x, GLdouble  y, GLdouble  z, GLdouble  w) {
//   (*fnptr)(x, y, z, w);
// }
// static void  glowRasterPos4dv(GPRASTERPOS4DV fnptr, const GLdouble * v) {
//   (*fnptr)(v);
// }
// static void  glowRasterPos4f(GPRASTERPOS4F fnptr, GLfloat  x, GLfloat  y, GLfloat  z, GLfloat  w) {
//   (*fnptr)(x, y, z, w);
// }
// static void  glowRasterPos4fv(GPRASTERPOS4FV fnptr, const GLfloat * v) {
//   (*fnptr)(v);
// }
// static void  glowRasterPos4i(GPRASTERPOS4I fnptr, GLint  x, GLint  y, GLint  z, GLint  w) {
//   (*fnptr)(x, y, z, w);
// }
// static void  glowRasterPos4iv(GPRASTERPOS4IV fnptr, const GLint * v) {
//   (*fnptr)(v);
// }
// static void  glowRasterPos4s(GPRASTERPOS4S fnptr, GLshort  x, GLshort  y, GLshort  z, GLshort  w) {
//   (*fnptr)(x, y, z, w);
// }
// static void  glowRasterPos4sv(GPRASTERPOS4SV fnptr, const GLshort * v) {
//   (*fnptr)(v);
// }
// static void  glowReadBuffer(GPREADBUFFER fnptr, GLenum  src) {
//   (*fnptr)(src);
// }
// static void  glowReadPixels(GPREADPIXELS fnptr, GLint  x, GLint  y, GLsizei  width, GLsizei  height, GLenum  format, GLenum  type, void * pixels) {
//   (*fnptr)(x, y, width, height, format, type, pixels);
// }
// static void  glowReadnPixels(GPREADNPIXELS fnptr, GLint  x, GLint  y, GLsizei  width, GLsizei  height, GLenum  format, GLenum  type, GLsizei  bufSize, void * data) {
//   (*fnptr)(x, y, width, height, format, type, bufSize, data);
// }
// static void  glowReadnPixelsARB(GPREADNPIXELSARB fnptr, GLint  x, GLint  y, GLsizei  width, GLsizei  height, GLenum  format, GLenum  type, GLsizei  bufSize, void * data) {
//   (*fnptr)(x, y, width, height, format, type, bufSize, data);
// }
// static void  glowReadnPixelsKHR(GPREADNPIXELSKHR fnptr, GLint  x, GLint  y, GLsizei  width, GLsizei  height, GLenum  format, GLenum  type, GLsizei  bufSize, void * data) {
//   (*fnptr)(x, y, width, height, format, type, bufSize, data);
// }
// static void  glowRectd(GPRECTD fnptr, GLdouble  x1, GLdouble  y1, GLdouble  x2, GLdouble  y2) {
//   (*fnptr)(x1, y1, x2, y2);
// }
// static void  glowRectdv(GPRECTDV fnptr, const GLdouble * v1, const GLdouble * v2) {
//   (*fnptr)(v1, v2);
// }
// static void  glowRectf(GPRECTF fnptr, GLfloat  x1, GLfloat  y1, GLfloat  x2, GLfloat  y2) {
//   (*fnptr)(x1, y1, x2, y2);
// }
// static void  glowRectfv(GPRECTFV fnptr, const GLfloat * v1, const GLfloat * v2) {
//   (*fnptr)(v1, v2);
// }
// static void  glowRecti(GPRECTI fnptr, GLint  x1, GLint  y1, GLint  x2, GLint  y2) {
//   (*fnptr)(x1, y1, x2, y2);
// }
// static void  glowRectiv(GPRECTIV fnptr, const GLint * v1, const GLint * v2) {
//   (*fnptr)(v1, v2);
// }
// static void  glowRects(GPRECTS fnptr, GLshort  x1, GLshort  y1, GLshort  x2, GLshort  y2) {
//   (*fnptr)(x1, y1, x2, y2);
// }
// static void  glowRectsv(GPRECTSV fnptr, const GLshort * v1, const GLshort * v2) {
//   (*fnptr)(v1, v2);
// }
// static void  glowReleaseShaderCompiler(GPRELEASESHADERCOMPILER fnptr) {
//   (*fnptr)();
// }
// static GLint  glowRenderMode(GPRENDERMODE fnptr, GLenum  mode) {
//   return (*fnptr)(mode);
// }
// static void  glowRenderbufferStorage(GPRENDERBUFFERSTORAGE fnptr, GLenum  target, GLenum  internalformat, GLsizei  width, GLsizei  height) {
//   (*fnptr)(target, internalformat, width, height);
// }
// static void  glowRenderbufferStorageMultisample(GPRENDERBUFFERSTORAGEMULTISAMPLE fnptr, GLenum  target, GLsizei  samples, GLenum  internalformat, GLsizei  width, GLsizei  height) {
//   (*fnptr)(target, samples, internalformat, width, height);
// }
// static void  glowResumeTransformFeedback(GPRESUMETRANSFORMFEEDBACK fnptr) {
//   (*fnptr)();
// }
// static void  glowRotated(GPROTATED fnptr, GLdouble  angle, GLdouble  x, GLdouble  y, GLdouble  z) {
//   (*fnptr)(angle, x, y, z);
// }
// static void  glowRotatef(GPROTATEF fnptr, GLfloat  angle, GLfloat  x, GLfloat  y, GLfloat  z) {
//   (*fnptr)(angle, x, y, z);
// }
// static void  glowSampleCoverage(GPSAMPLECOVERAGE fnptr, GLfloat  value, GLboolean  invert) {
//   (*fnptr)(value, invert);
// }
// static void  glowSampleMaski(GPSAMPLEMASKI fnptr, GLuint  maskNumber, GLbitfield  mask) {
//   (*fnptr)(maskNumber, mask);
// }
// static void  glowSamplerParameterIiv(GPSAMPLERPARAMETERIIV fnptr, GLuint  sampler, GLenum  pname, const GLint * param) {
//   (*fnptr)(sampler, pname, param);
// }
// static void  glowSamplerParameterIuiv(GPSAMPLERPARAMETERIUIV fnptr, GLuint  sampler, GLenum  pname, const GLuint * param) {
//   (*fnptr)(sampler, pname, param);
// }
// static void  glowSamplerParameterf(GPSAMPLERPARAMETERF fnptr, GLuint  sampler, GLenum  pname, GLfloat  param) {
//   (*fnptr)(sampler, pname, param);
// }
// static void  glowSamplerParameterfv(GPSAMPLERPARAMETERFV fnptr, GLuint  sampler, GLenum  pname, const GLfloat * param) {
//   (*fnptr)(sampler, pname, param);
// }
// static void  glowSamplerParameteri(GPSAMPLERPARAMETERI fnptr, GLuint  sampler, GLenum  pname, GLint  param) {
//   (*fnptr)(sampler, pname, param);
// }
// static void  glowSamplerParameteriv(GPSAMPLERPARAMETERIV fnptr, GLuint  sampler, GLenum  pname, const GLint * param) {
//   (*fnptr)(sampler, pname, param);
// }
// static void  glowScaled(GPSCALED fnptr, GLdouble  x, GLdouble  y, GLdouble  z) {
//   (*fnptr)(x, y, z);
// }
// static void  glowScalef(GPSCALEF fnptr, GLfloat  x, GLfloat  y, GLfloat  z) {
//   (*fnptr)(x, y, z);
// }
// static void  glowScissor(GPSCISSOR fnptr, GLint  x, GLint  y, GLsizei  width, GLsizei  height) {
//   (*fnptr)(x, y, width, height);
// }
// static void  glowScissorArrayv(GPSCISSORARRAYV fnptr, GLuint  first, GLsizei  count, const GLint * v) {
//   (*fnptr)(first, count, v);
// }
// static void  glowScissorIndexed(GPSCISSORINDEXED fnptr, GLuint  index, GLint  left, GLint  bottom, GLsizei  width, GLsizei  height) {
//   (*fnptr)(index, left, bottom, width, height);
// }
// static void  glowScissorIndexedv(GPSCISSORINDEXEDV fnptr, GLuint  index, const GLint * v) {
//   (*fnptr)(index, v);
// }
// static void  glowSecondaryColor3b(GPSECONDARYCOLOR3B fnptr, GLbyte  red, GLbyte  green, GLbyte  blue) {
//   (*fnptr)(red, green, blue);
// }
// static void  glowSecondaryColor3bv(GPSECONDARYCOLOR3BV fnptr, const GLbyte * v) {
//   (*fnptr)(v);
// }
// static void  glowSecondaryColor3d(GPSECONDARYCOLOR3D fnptr, GLdouble  red, GLdouble  green, GLdouble  blue) {
//   (*fnptr)(red, green, blue);
// }
// static void  glowSecondaryColor3dv(GPSECONDARYCOLOR3DV fnptr, const GLdouble * v) {
//   (*fnptr)(v);
// }
// static void  glowSecondaryColor3f(GPSECONDARYCOLOR3F fnptr, GLfloat  red, GLfloat  green, GLfloat  blue) {
//   (*fnptr)(red, green, blue);
// }
// static void  glowSecondaryColor3fv(GPSECONDARYCOLOR3FV fnptr, const GLfloat * v) {
//   (*fnptr)(v);
// }
// static void  glowSecondaryColor3i(GPSECONDARYCOLOR3I fnptr, GLint  red, GLint  green, GLint  blue) {
//   (*fnptr)(red, green, blue);
// }
// static void  glowSecondaryColor3iv(GPSECONDARYCOLOR3IV fnptr, const GLint * v) {
//   (*fnptr)(v);
// }
// static void  glowSecondaryColor3s(GPSECONDARYCOLOR3S fnptr, GLshort  red, GLshort  green, GLshort  blue) {
//   (*fnptr)(red, green, blue);
// }
// static void  glowSecondaryColor3sv(GPSECONDARYCOLOR3SV fnptr, const GLshort * v) {
//   (*fnptr)(v);
// }
// static void  glowSecondaryColor3ub(GPSECONDARYCOLOR3UB fnptr, GLubyte  red, GLubyte  green, GLubyte  blue) {
//   (*fnptr)(red, green, blue);
// }
// static void  glowSecondaryColor3ubv(GPSECONDARYCOLOR3UBV fnptr, const GLubyte * v) {
//   (*fnptr)(v);
// }
// static void  glowSecondaryColor3ui(GPSECONDARYCOLOR3UI fnptr, GLuint  red, GLuint  green, GLuint  blue) {
//   (*fnptr)(red, green, blue);
// }
// static void  glowSecondaryColor3uiv(GPSECONDARYCOLOR3UIV fnptr, const GLuint * v) {
//   (*fnptr)(v);
// }
// static void  glowSecondaryColor3us(GPSECONDARYCOLOR3US fnptr, GLushort  red, GLushort  green, GLushort  blue) {
//   (*fnptr)(red, green, blue);
// }
// static void  glowSecondaryColor3usv(GPSECONDARYCOLOR3USV fnptr, const GLushort * v) {
//   (*fnptr)(v);
// }
// static void  glowSecondaryColorPointer(GPSECONDARYCOLORPOINTER fnptr, GLint  size, GLenum  type, GLsizei  stride, const void * pointer) {
//   (*fnptr)(size, type, stride, pointer);
// }
// static void  glowSelectBuffer(GPSELECTBUFFER fnptr, GLsizei  size, GLuint * buffer) {
//   (*fnptr)(size, buffer);
// }
// static void  glowShadeModel(GPSHADEMODEL fnptr, GLenum  mode) {
//   (*fnptr)(mode);
// }
// static void  glowShaderBinary(GPSHADERBINARY fnptr, GLsizei  count, const GLuint * shaders, GLenum  binaryformat, const void * binary, GLsizei  length) {
//   (*fnptr)(count, shaders, binaryformat, binary, length);
// }
// static void  glowShaderSource(GPSHADERSOURCE fnptr, GLuint  shader, GLsizei  count, const GLchar *const* string, const GLint * length) {
//   (*fnptr)(shader, count, string, length);
// }
// static void  glowShaderStorageBlockBinding(GPSHADERSTORAGEBLOCKBINDING fnptr, GLuint  program, GLuint  storageBlockIndex, GLuint  storageBlockBinding) {
//   (*fnptr)(program, storageBlockIndex, storageBlockBinding);
// }
// static void  glowStencilFunc(GPSTENCILFUNC fnptr, GLenum  func, GLint  ref, GLuint  mask) {
//   (*fnptr)(func, ref, mask);
// }
// static void  glowStencilFuncSeparate(GPSTENCILFUNCSEPARATE fnptr, GLenum  face, GLenum  func, GLint  ref, GLuint  mask) {
//   (*fnptr)(face, func, ref, mask);
// }
// static void  glowStencilMask(GPSTENCILMASK fnptr, GLuint  mask) {
//   (*fnptr)(mask);
// }
// static void  glowStencilMaskSeparate(GPSTENCILMASKSEPARATE fnptr, GLenum  face, GLuint  mask) {
//   (*fnptr)(face, mask);
// }
// static void  glowStencilOp(GPSTENCILOP fnptr, GLenum  fail, GLenum  zfail, GLenum  zpass) {
//   (*fnptr)(fail, zfail, zpass);
// }
// static void  glowStencilOpSeparate(GPSTENCILOPSEPARATE fnptr, GLenum  face, GLenum  sfail, GLenum  dpfail, GLenum  dppass) {
//   (*fnptr)(face, sfail, dpfail, dppass);
// }
// static void  glowTexBuffer(GPTEXBUFFER fnptr, GLenum  target, GLenum  internalformat, GLuint  buffer) {
//   (*fnptr)(target, internalformat, buffer);
// }
// static void  glowTexBufferRange(GPTEXBUFFERRANGE fnptr, GLenum  target, GLenum  internalformat, GLuint  buffer, GLintptr  offset, GLsizeiptr  size) {
//   (*fnptr)(target, internalformat, buffer, offset, size);
// }
// static void  glowTexCoord1d(GPTEXCOORD1D fnptr, GLdouble  s) {
//   (*fnptr)(s);
// }
// static void  glowTexCoord1dv(GPTEXCOORD1DV fnptr, const GLdouble * v) {
//   (*fnptr)(v);
// }
// static void  glowTexCoord1f(GPTEXCOORD1F fnptr, GLfloat  s) {
//   (*fnptr)(s);
// }
// static void  glowTexCoord1fv(GPTEXCOORD1FV fnptr, const GLfloat * v) {
//   (*fnptr)(v);
// }
// static void  glowTexCoord1i(GPTEXCOORD1I fnptr, GLint  s) {
//   (*fnptr)(s);
// }
// static void  glowTexCoord1iv(GPTEXCOORD1IV fnptr, const GLint * v) {
//   (*fnptr)(v);
// }
// static void  glowTexCoord1s(GPTEXCOORD1S fnptr, GLshort  s) {
//   (*fnptr)(s);
// }
// static void  glowTexCoord1sv(GPTEXCOORD1SV fnptr, const GLshort * v) {
//   (*fnptr)(v);
// }
// static void  glowTexCoord2d(GPTEXCOORD2D fnptr, GLdouble  s, GLdouble  t) {
//   (*fnptr)(s, t);
// }
// static void  glowTexCoord2dv(GPTEXCOORD2DV fnptr, const GLdouble * v) {
//   (*fnptr)(v);
// }
// static void  glowTexCoord2f(GPTEXCOORD2F fnptr, GLfloat  s, GLfloat  t) {
//   (*fnptr)(s, t);
// }
// static void  glowTexCoord2fv(GPTEXCOORD2FV fnptr, const GLfloat * v) {
//   (*fnptr)(v);
// }
// static void  glowTexCoord2i(GPTEXCOORD2I fnptr, GLint  s, GLint  t) {
//   (*fnptr)(s, t);
// }
// static void  glowTexCoord2iv(GPTEXCOORD2IV fnptr, const GLint * v) {
//   (*fnptr)(v);
// }
// static void  glowTexCoord2s(GPTEXCOORD2S fnptr, GLshort  s, GLshort  t) {
//   (*fnptr)(s, t);
// }
// static void  glowTexCoord2sv(GPTEXCOORD2SV fnptr, const GLshort * v) {
//   (*fnptr)(v);
// }
// static void  glowTexCoord3d(GPTEXCOORD3D fnptr, GLdouble  s, GLdouble  t, GLdouble  r) {
//   (*fnptr)(s, t, r);
// }
// static void  glowTexCoord3dv(GPTEXCOORD3DV fnptr, const GLdouble * v) {
//   (*fnptr)(v);
// }
// static void  glowTexCoord3f(GPTEXCOORD3F fnptr, GLfloat  s, GLfloat  t, GLfloat  r) {
//   (*fnptr)(s, t, r);
// }
// static void  glowTexCoord3fv(GPTEXCOORD3FV fnptr, const GLfloat * v) {
//   (*fnptr)(v);
// }
// static void  glowTexCoord3i(GPTEXCOORD3I fnptr, GLint  s, GLint  t, GLint  r) {
//   (*fnptr)(s, t, r);
// }
// static void  glowTexCoord3iv(GPTEXCOORD3IV fnptr, const GLint * v) {
//   (*fnptr)(v);
// }
// static void  glowTexCoord3s(GPTEXCOORD3S fnptr, GLshort  s, GLshort  t, GLshort  r) {
//   (*fnptr)(s, t, r);
// }
// static void  glowTexCoord3sv(GPTEXCOORD3SV fnptr, const GLshort * v) {
//   (*fnptr)(v);
// }
// static void  glowTexCoord4d(GPTEXCOORD4D fnptr, GLdouble  s, GLdouble  t, GLdouble  r, GLdouble  q) {
//   (*fnptr)(s, t, r, q);
// }
// static void  glowTexCoord4dv(GPTEXCOORD4DV fnptr, const GLdouble * v) {
//   (*fnptr)(v);
// }
// static void  glowTexCoord4f(GPTEXCOORD4F fnptr, GLfloat  s, GLfloat  t, GLfloat  r, GLfloat  q) {
//   (*fnptr)(s, t, r, q);
// }
// static void  glowTexCoord4fv(GPTEXCOORD4FV fnptr, const GLfloat * v) {
//   (*fnptr)(v);
// }
// static void  glowTexCoord4i(GPTEXCOORD4I fnptr, GLint  s, GLint  t, GLint  r, GLint  q) {
//   (*fnptr)(s, t, r, q);
// }
// static void  glowTexCoord4iv(GPTEXCOORD4IV fnptr, const GLint * v) {
//   (*fnptr)(v);
// }
// static void  glowTexCoord4s(GPTEXCOORD4S fnptr, GLshort  s, GLshort  t, GLshort  r, GLshort  q) {
//   (*fnptr)(s, t, r, q);
// }
// static void  glowTexCoord4sv(GPTEXCOORD4SV fnptr, const GLshort * v) {
//   (*fnptr)(v);
// }
// static void  glowTexCoordPointer(GPTEXCOORDPOINTER fnptr, GLint  size, GLenum  type, GLsizei  stride, const void * pointer) {
//   (*fnptr)(size, type, stride, pointer);
// }
// static void  glowTexEnvf(GPTEXENVF fnptr, GLenum  target, GLenum  pname, GLfloat  param) {
//   (*fnptr)(target, pname, param);
// }
// static void  glowTexEnvfv(GPTEXENVFV fnptr, GLenum  target, GLenum  pname, const GLfloat * params) {
//   (*fnptr)(target, pname, params);
// }
// static void  glowTexEnvi(GPTEXENVI fnptr, GLenum  target, GLenum  pname, GLint  param) {
//   (*fnptr)(target, pname, param);
// }
// static void  glowTexEnviv(GPTEXENVIV fnptr, GLenum  target, GLenum  pname, const GLint * params) {
//   (*fnptr)(target, pname, params);
// }
// static void  glowTexGend(GPTEXGEND fnptr, GLenum  coord, GLenum  pname, GLdouble  param) {
//   (*fnptr)(coord, pname, param);
// }
// static void  glowTexGendv(GPTEXGENDV fnptr, GLenum  coord, GLenum  pname, const GLdouble * params) {
//   (*fnptr)(coord, pname, params);
// }
// static void  glowTexGenf(GPTEXGENF fnptr, GLenum  coord, GLenum  pname, GLfloat  param) {
//   (*fnptr)(coord, pname, param);
// }
// static void  glowTexGenfv(GPTEXGENFV fnptr, GLenum  coord, GLenum  pname, const GLfloat * params) {
//   (*fnptr)(coord, pname, params);
// }
// static void  glowTexGeni(GPTEXGENI fnptr, GLenum  coord, GLenum  pname, GLint  param) {
//   (*fnptr)(coord, pname, param);
// }
// static void  glowTexGeniv(GPTEXGENIV fnptr, GLenum  coord, GLenum  pname, const GLint * params) {
//   (*fnptr)(coord, pname, params);
// }
// static void  glowTexImage1D(GPTEXIMAGE1D fnptr, GLenum  target, GLint  level, GLint  internalformat, GLsizei  width, GLint  border, GLenum  format, GLenum  type, const void * pixels) {
//   (*fnptr)(target, level, internalformat, width, border, format, type, pixels);
// }
// static void  glowTexImage2D(GPTEXIMAGE2D fnptr, GLenum  target, GLint  level, GLint  internalformat, GLsizei  width, GLsizei  height, GLint  border, GLenum  format, GLenum  type, const void * pixels) {
//   (*fnptr)(target, level, internalformat, width, height, border, format, type, pixels);
// }
// static void  glowTexImage2DMultisample(GPTEXIMAGE2DMULTISAMPLE fnptr, GLenum  target, GLsizei  samples, GLenum  internalformat, GLsizei  width, GLsizei  height, GLboolean  fixedsamplelocations) {
//   (*fnptr)(target, samples, internalformat, width, height, fixedsamplelocations);
// }
// static void  glowTexImage3D(GPTEXIMAGE3D fnptr, GLenum  target, GLint  level, GLint  internalformat, GLsizei  width, GLsizei  height, GLsizei  depth, GLint  border, GLenum  format, GLenum  type, const void * pixels) {
//   (*fnptr)(target, level, internalformat, width, height, depth, border, format, type, pixels);
// }
// static void  glowTexImage3DMultisample(GPTEXIMAGE3DMULTISAMPLE fnptr, GLenum  target, GLsizei  samples, GLenum  internalformat, GLsizei  width, GLsizei  height, GLsizei  depth, GLboolean  fixedsamplelocations) {
//   (*fnptr)(target, samples, internalformat, width, height, depth, fixedsamplelocations);
// }
// static void  glowTexPageCommitmentARB(GPTEXPAGECOMMITMENTARB fnptr, GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth, GLboolean  resident) {
//   (*fnptr)(target, level, xoffset, yoffset, zoffset, width, height, depth, resident);
// }
// static void  glowTexParameterIiv(GPTEXPARAMETERIIV fnptr, GLenum  target, GLenum  pname, const GLint * params) {
//   (*fnptr)(target, pname, params);
// }
// static void  glowTexParameterIuiv(GPTEXPARAMETERIUIV fnptr, GLenum  target, GLenum  pname, const GLuint * params) {
//   (*fnptr)(target, pname, params);
// }
// static void  glowTexParameterf(GPTEXPARAMETERF fnptr, GLenum  target, GLenum  pname, GLfloat  param) {
//   (*fnptr)(target, pname, param);
// }
// static void  glowTexParameterfv(GPTEXPARAMETERFV fnptr, GLenum  target, GLenum  pname, const GLfloat * params) {
//   (*fnptr)(target, pname, params);
// }
// static void  glowTexParameteri(GPTEXPARAMETERI fnptr, GLenum  target, GLenum  pname, GLint  param) {
//   (*fnptr)(target, pname, param);
// }
// static void  glowTexParameteriv(GPTEXPARAMETERIV fnptr, GLenum  target, GLenum  pname, const GLint * params) {
//   (*fnptr)(target, pname, params);
// }
// static void  glowTexStorage1D(GPTEXSTORAGE1D fnptr, GLenum  target, GLsizei  levels, GLenum  internalformat, GLsizei  width) {
//   (*fnptr)(target, levels, internalformat, width);
// }
// static void  glowTexStorage2D(GPTEXSTORAGE2D fnptr, GLenum  target, GLsizei  levels, GLenum  internalformat, GLsizei  width, GLsizei  height) {
//   (*fnptr)(target, levels, internalformat, width, height);
// }
// static void  glowTexStorage2DMultisample(GPTEXSTORAGE2DMULTISAMPLE fnptr, GLenum  target, GLsizei  samples, GLenum  internalformat, GLsizei  width, GLsizei  height, GLboolean  fixedsamplelocations) {
//   (*fnptr)(target, samples, internalformat, width, height, fixedsamplelocations);
// }
// static void  glowTexStorage3D(GPTEXSTORAGE3D fnptr, GLenum  target, GLsizei  levels, GLenum  internalformat, GLsizei  width, GLsizei  height, GLsizei  depth) {
//   (*fnptr)(target, levels, internalformat, width, height, depth);
// }
// static void  glowTexStorage3DMultisample(GPTEXSTORAGE3DMULTISAMPLE fnptr, GLenum  target, GLsizei  samples, GLenum  internalformat, GLsizei  width, GLsizei  height, GLsizei  depth, GLboolean  fixedsamplelocations) {
//   (*fnptr)(target, samples, internalformat, width, height, depth, fixedsamplelocations);
// }
// static void  glowTexSubImage1D(GPTEXSUBIMAGE1D fnptr, GLenum  target, GLint  level, GLint  xoffset, GLsizei  width, GLenum  format, GLenum  type, const void * pixels) {
//   (*fnptr)(target, level, xoffset, width, format, type, pixels);
// }
// static void  glowTexSubImage2D(GPTEXSUBIMAGE2D fnptr, GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLsizei  width, GLsizei  height, GLenum  format, GLenum  type, const void * pixels) {
//   (*fnptr)(target, level, xoffset, yoffset, width, height, format, type, pixels);
// }
// static void  glowTexSubImage3D(GPTEXSUBIMAGE3D fnptr, GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth, GLenum  format, GLenum  type, const void * pixels) {
//   (*fnptr)(target, level, xoffset, yoffset, zoffset, width, height, depth, format, type, pixels);
// }
// static void  glowTextureBarrier(GPTEXTUREBARRIER fnptr) {
//   (*fnptr)();
// }
// static void  glowTextureBuffer(GPTEXTUREBUFFER fnptr, GLuint  texture, GLenum  internalformat, GLuint  buffer) {
//   (*fnptr)(texture, internalformat, buffer);
// }
// static void  glowTextureBufferRange(GPTEXTUREBUFFERRANGE fnptr, GLuint  texture, GLenum  internalformat, GLuint  buffer, GLintptr  offset, GLsizei  size) {
//   (*fnptr)(texture, internalformat, buffer, offset, size);
// }
// static void  glowTextureParameterIiv(GPTEXTUREPARAMETERIIV fnptr, GLuint  texture, GLenum  pname, const GLint * params) {
//   (*fnptr)(texture, pname, params);
// }
// static void  glowTextureParameterIuiv(GPTEXTUREPARAMETERIUIV fnptr, GLuint  texture, GLenum  pname, const GLuint * params) {
//   (*fnptr)(texture, pname, params);
// }
// static void  glowTextureParameterf(GPTEXTUREPARAMETERF fnptr, GLuint  texture, GLenum  pname, GLfloat  param) {
//   (*fnptr)(texture, pname, param);
// }
// static void  glowTextureParameterfv(GPTEXTUREPARAMETERFV fnptr, GLuint  texture, GLenum  pname, const GLfloat * param) {
//   (*fnptr)(texture, pname, param);
// }
// static void  glowTextureParameteri(GPTEXTUREPARAMETERI fnptr, GLuint  texture, GLenum  pname, GLint  param) {
//   (*fnptr)(texture, pname, param);
// }
// static void  glowTextureParameteriv(GPTEXTUREPARAMETERIV fnptr, GLuint  texture, GLenum  pname, const GLint * param) {
//   (*fnptr)(texture, pname, param);
// }
// static void  glowTextureStorage1D(GPTEXTURESTORAGE1D fnptr, GLuint  texture, GLsizei  levels, GLenum  internalformat, GLsizei  width) {
//   (*fnptr)(texture, levels, internalformat, width);
// }
// static void  glowTextureStorage2D(GPTEXTURESTORAGE2D fnptr, GLuint  texture, GLsizei  levels, GLenum  internalformat, GLsizei  width, GLsizei  height) {
//   (*fnptr)(texture, levels, internalformat, width, height);
// }
// static void  glowTextureStorage2DMultisample(GPTEXTURESTORAGE2DMULTISAMPLE fnptr, GLuint  texture, GLsizei  samples, GLenum  internalformat, GLsizei  width, GLsizei  height, GLboolean  fixedsamplelocations) {
//   (*fnptr)(texture, samples, internalformat, width, height, fixedsamplelocations);
// }
// static void  glowTextureStorage3D(GPTEXTURESTORAGE3D fnptr, GLuint  texture, GLsizei  levels, GLenum  internalformat, GLsizei  width, GLsizei  height, GLsizei  depth) {
//   (*fnptr)(texture, levels, internalformat, width, height, depth);
// }
// static void  glowTextureStorage3DMultisample(GPTEXTURESTORAGE3DMULTISAMPLE fnptr, GLuint  texture, GLsizei  samples, GLenum  internalformat, GLsizei  width, GLsizei  height, GLsizei  depth, GLboolean  fixedsamplelocations) {
//   (*fnptr)(texture, samples, internalformat, width, height, depth, fixedsamplelocations);
// }
// static void  glowTextureSubImage1D(GPTEXTURESUBIMAGE1D fnptr, GLuint  texture, GLint  level, GLint  xoffset, GLsizei  width, GLenum  format, GLenum  type, const void * pixels) {
//   (*fnptr)(texture, level, xoffset, width, format, type, pixels);
// }
// static void  glowTextureSubImage2D(GPTEXTURESUBIMAGE2D fnptr, GLuint  texture, GLint  level, GLint  xoffset, GLint  yoffset, GLsizei  width, GLsizei  height, GLenum  format, GLenum  type, const void * pixels) {
//   (*fnptr)(texture, level, xoffset, yoffset, width, height, format, type, pixels);
// }
// static void  glowTextureSubImage3D(GPTEXTURESUBIMAGE3D fnptr, GLuint  texture, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth, GLenum  format, GLenum  type, const void * pixels) {
//   (*fnptr)(texture, level, xoffset, yoffset, zoffset, width, height, depth, format, type, pixels);
// }
// static void  glowTextureView(GPTEXTUREVIEW fnptr, GLuint  texture, GLenum  target, GLuint  origtexture, GLenum  internalformat, GLuint  minlevel, GLuint  numlevels, GLuint  minlayer, GLuint  numlayers) {
//   (*fnptr)(texture, target, origtexture, internalformat, minlevel, numlevels, minlayer, numlayers);
// }
// static void  glowTransformFeedbackBufferBase(GPTRANSFORMFEEDBACKBUFFERBASE fnptr, GLuint  xfb, GLuint  index, GLuint  buffer) {
//   (*fnptr)(xfb, index, buffer);
// }
// static void  glowTransformFeedbackBufferRange(GPTRANSFORMFEEDBACKBUFFERRANGE fnptr, GLuint  xfb, GLuint  index, GLuint  buffer, GLintptr  offset, GLsizei  size) {
//   (*fnptr)(xfb, index, buffer, offset, size);
// }
// static void  glowTransformFeedbackVaryings(GPTRANSFORMFEEDBACKVARYINGS fnptr, GLuint  program, GLsizei  count, const GLchar *const* varyings, GLenum  bufferMode) {
//   (*fnptr)(program, count, varyings, bufferMode);
// }
// static void  glowTranslated(GPTRANSLATED fnptr, GLdouble  x, GLdouble  y, GLdouble  z) {
//   (*fnptr)(x, y, z);
// }
// static void  glowTranslatef(GPTRANSLATEF fnptr, GLfloat  x, GLfloat  y, GLfloat  z) {
//   (*fnptr)(x, y, z);
// }
// static void  glowUniform1d(GPUNIFORM1D fnptr, GLint  location, GLdouble  x) {
//   (*fnptr)(location, x);
// }
// static void  glowUniform1dv(GPUNIFORM1DV fnptr, GLint  location, GLsizei  count, const GLdouble * value) {
//   (*fnptr)(location, count, value);
// }
// static void  glowUniform1f(GPUNIFORM1F fnptr, GLint  location, GLfloat  v0) {
//   (*fnptr)(location, v0);
// }
// static void  glowUniform1fv(GPUNIFORM1FV fnptr, GLint  location, GLsizei  count, const GLfloat * value) {
//   (*fnptr)(location, count, value);
// }
// static void  glowUniform1i(GPUNIFORM1I fnptr, GLint  location, GLint  v0) {
//   (*fnptr)(location, v0);
// }
// static void  glowUniform1iv(GPUNIFORM1IV fnptr, GLint  location, GLsizei  count, const GLint * value) {
//   (*fnptr)(location, count, value);
// }
// static void  glowUniform1ui(GPUNIFORM1UI fnptr, GLint  location, GLuint  v0) {
//   (*fnptr)(location, v0);
// }
// static void  glowUniform1uiv(GPUNIFORM1UIV fnptr, GLint  location, GLsizei  count, const GLuint * value) {
//   (*fnptr)(location, count, value);
// }
// static void  glowUniform2d(GPUNIFORM2D fnptr, GLint  location, GLdouble  x, GLdouble  y) {
//   (*fnptr)(location, x, y);
// }
// static void  glowUniform2dv(GPUNIFORM2DV fnptr, GLint  location, GLsizei  count, const GLdouble * value) {
//   (*fnptr)(location, count, value);
// }
// static void  glowUniform2f(GPUNIFORM2F fnptr, GLint  location, GLfloat  v0, GLfloat  v1) {
//   (*fnptr)(location, v0, v1);
// }
// static void  glowUniform2fv(GPUNIFORM2FV fnptr, GLint  location, GLsizei  count, const GLfloat * value) {
//   (*fnptr)(location, count, value);
// }
// static void  glowUniform2i(GPUNIFORM2I fnptr, GLint  location, GLint  v0, GLint  v1) {
//   (*fnptr)(location, v0, v1);
// }
// static void  glowUniform2iv(GPUNIFORM2IV fnptr, GLint  location, GLsizei  count, const GLint * value) {
//   (*fnptr)(location, count, value);
// }
// static void  glowUniform2ui(GPUNIFORM2UI fnptr, GLint  location, GLuint  v0, GLuint  v1) {
//   (*fnptr)(location, v0, v1);
// }
// static void  glowUniform2uiv(GPUNIFORM2UIV fnptr, GLint  location, GLsizei  count, const GLuint * value) {
//   (*fnptr)(location, count, value);
// }
// static void  glowUniform3d(GPUNIFORM3D fnptr, GLint  location, GLdouble  x, GLdouble  y, GLdouble  z) {
//   (*fnptr)(location, x, y, z);
// }
// static void  glowUniform3dv(GPUNIFORM3DV fnptr, GLint  location, GLsizei  count, const GLdouble * value) {
//   (*fnptr)(location, count, value);
// }
// static void  glowUniform3f(GPUNIFORM3F fnptr, GLint  location, GLfloat  v0, GLfloat  v1, GLfloat  v2) {
//   (*fnptr)(location, v0, v1, v2);
// }
// static void  glowUniform3fv(GPUNIFORM3FV fnptr, GLint  location, GLsizei  count, const GLfloat * value) {
//   (*fnptr)(location, count, value);
// }
// static void  glowUniform3i(GPUNIFORM3I fnptr, GLint  location, GLint  v0, GLint  v1, GLint  v2) {
//   (*fnptr)(location, v0, v1, v2);
// }
// static void  glowUniform3iv(GPUNIFORM3IV fnptr, GLint  location, GLsizei  count, const GLint * value) {
//   (*fnptr)(location, count, value);
// }
// static void  glowUniform3ui(GPUNIFORM3UI fnptr, GLint  location, GLuint  v0, GLuint  v1, GLuint  v2) {
//   (*fnptr)(location, v0, v1, v2);
// }
// static void  glowUniform3uiv(GPUNIFORM3UIV fnptr, GLint  location, GLsizei  count, const GLuint * value) {
//   (*fnptr)(location, count, value);
// }
// static void  glowUniform4d(GPUNIFORM4D fnptr, GLint  location, GLdouble  x, GLdouble  y, GLdouble  z, GLdouble  w) {
//   (*fnptr)(location, x, y, z, w);
// }
// static void  glowUniform4dv(GPUNIFORM4DV fnptr, GLint  location, GLsizei  count, const GLdouble * value) {
//   (*fnptr)(location, count, value);
// }
// static void  glowUniform4f(GPUNIFORM4F fnptr, GLint  location, GLfloat  v0, GLfloat  v1, GLfloat  v2, GLfloat  v3) {
//   (*fnptr)(location, v0, v1, v2, v3);
// }
// static void  glowUniform4fv(GPUNIFORM4FV fnptr, GLint  location, GLsizei  count, const GLfloat * value) {
//   (*fnptr)(location, count, value);
// }
// static void  glowUniform4i(GPUNIFORM4I fnptr, GLint  location, GLint  v0, GLint  v1, GLint  v2, GLint  v3) {
//   (*fnptr)(location, v0, v1, v2, v3);
// }
// static void  glowUniform4iv(GPUNIFORM4IV fnptr, GLint  location, GLsizei  count, const GLint * value) {
//   (*fnptr)(location, count, value);
// }
// static void  glowUniform4ui(GPUNIFORM4UI fnptr, GLint  location, GLuint  v0, GLuint  v1, GLuint  v2, GLuint  v3) {
//   (*fnptr)(location, v0, v1, v2, v3);
// }
// static void  glowUniform4uiv(GPUNIFORM4UIV fnptr, GLint  location, GLsizei  count, const GLuint * value) {
//   (*fnptr)(location, count, value);
// }
// static void  glowUniformBlockBinding(GPUNIFORMBLOCKBINDING fnptr, GLuint  program, GLuint  uniformBlockIndex, GLuint  uniformBlockBinding) {
//   (*fnptr)(program, uniformBlockIndex, uniformBlockBinding);
// }
// static void  glowUniformHandleui64ARB(GPUNIFORMHANDLEUI64ARB fnptr, GLint  location, GLuint64  value) {
//   (*fnptr)(location, value);
// }
// static void  glowUniformHandleui64vARB(GPUNIFORMHANDLEUI64VARB fnptr, GLint  location, GLsizei  count, const GLuint64 * value) {
//   (*fnptr)(location, count, value);
// }
// static void  glowUniformMatrix2dv(GPUNIFORMMATRIX2DV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value) {
//   (*fnptr)(location, count, transpose, value);
// }
// static void  glowUniformMatrix2fv(GPUNIFORMMATRIX2FV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(location, count, transpose, value);
// }
// static void  glowUniformMatrix2x3dv(GPUNIFORMMATRIX2X3DV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value) {
//   (*fnptr)(location, count, transpose, value);
// }
// static void  glowUniformMatrix2x3fv(GPUNIFORMMATRIX2X3FV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(location, count, transpose, value);
// }
// static void  glowUniformMatrix2x4dv(GPUNIFORMMATRIX2X4DV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value) {
//   (*fnptr)(location, count, transpose, value);
// }
// static void  glowUniformMatrix2x4fv(GPUNIFORMMATRIX2X4FV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(location, count, transpose, value);
// }
// static void  glowUniformMatrix3dv(GPUNIFORMMATRIX3DV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value) {
//   (*fnptr)(location, count, transpose, value);
// }
// static void  glowUniformMatrix3fv(GPUNIFORMMATRIX3FV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(location, count, transpose, value);
// }
// static void  glowUniformMatrix3x2dv(GPUNIFORMMATRIX3X2DV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value) {
//   (*fnptr)(location, count, transpose, value);
// }
// static void  glowUniformMatrix3x2fv(GPUNIFORMMATRIX3X2FV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(location, count, transpose, value);
// }
// static void  glowUniformMatrix3x4dv(GPUNIFORMMATRIX3X4DV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value) {
//   (*fnptr)(location, count, transpose, value);
// }
// static void  glowUniformMatrix3x4fv(GPUNIFORMMATRIX3X4FV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(location, count, transpose, value);
// }
// static void  glowUniformMatrix4dv(GPUNIFORMMATRIX4DV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value) {
//   (*fnptr)(location, count, transpose, value);
// }
// static void  glowUniformMatrix4fv(GPUNIFORMMATRIX4FV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(location, count, transpose, value);
// }
// static void  glowUniformMatrix4x2dv(GPUNIFORMMATRIX4X2DV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value) {
//   (*fnptr)(location, count, transpose, value);
// }
// static void  glowUniformMatrix4x2fv(GPUNIFORMMATRIX4X2FV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(location, count, transpose, value);
// }
// static void  glowUniformMatrix4x3dv(GPUNIFORMMATRIX4X3DV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value) {
//   (*fnptr)(location, count, transpose, value);
// }
// static void  glowUniformMatrix4x3fv(GPUNIFORMMATRIX4X3FV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(location, count, transpose, value);
// }
// static void  glowUniformSubroutinesuiv(GPUNIFORMSUBROUTINESUIV fnptr, GLenum  shadertype, GLsizei  count, const GLuint * indices) {
//   (*fnptr)(shadertype, count, indices);
// }
// static GLboolean  glowUnmapBuffer(GPUNMAPBUFFER fnptr, GLenum  target) {
//   return (*fnptr)(target);
// }
// static GLboolean  glowUnmapNamedBuffer(GPUNMAPNAMEDBUFFER fnptr, GLuint  buffer) {
//   return (*fnptr)(buffer);
// }
// static void  glowUseProgram(GPUSEPROGRAM fnptr, GLuint  program) {
//   (*fnptr)(program);
// }
// static void  glowUseProgramStages(GPUSEPROGRAMSTAGES fnptr, GLuint  pipeline, GLbitfield  stages, GLuint  program) {
//   (*fnptr)(pipeline, stages, program);
// }
// static void  glowValidateProgram(GPVALIDATEPROGRAM fnptr, GLuint  program) {
//   (*fnptr)(program);
// }
// static void  glowValidateProgramPipeline(GPVALIDATEPROGRAMPIPELINE fnptr, GLuint  pipeline) {
//   (*fnptr)(pipeline);
// }
// static void  glowVertex2d(GPVERTEX2D fnptr, GLdouble  x, GLdouble  y) {
//   (*fnptr)(x, y);
// }
// static void  glowVertex2dv(GPVERTEX2DV fnptr, const GLdouble * v) {
//   (*fnptr)(v);
// }
// static void  glowVertex2f(GPVERTEX2F fnptr, GLfloat  x, GLfloat  y) {
//   (*fnptr)(x, y);
// }
// static void  glowVertex2fv(GPVERTEX2FV fnptr, const GLfloat * v) {
//   (*fnptr)(v);
// }
// static void  glowVertex2i(GPVERTEX2I fnptr, GLint  x, GLint  y) {
//   (*fnptr)(x, y);
// }
// static void  glowVertex2iv(GPVERTEX2IV fnptr, const GLint * v) {
//   (*fnptr)(v);
// }
// static void  glowVertex2s(GPVERTEX2S fnptr, GLshort  x, GLshort  y) {
//   (*fnptr)(x, y);
// }
// static void  glowVertex2sv(GPVERTEX2SV fnptr, const GLshort * v) {
//   (*fnptr)(v);
// }
// static void  glowVertex3d(GPVERTEX3D fnptr, GLdouble  x, GLdouble  y, GLdouble  z) {
//   (*fnptr)(x, y, z);
// }
// static void  glowVertex3dv(GPVERTEX3DV fnptr, const GLdouble * v) {
//   (*fnptr)(v);
// }
// static void  glowVertex3f(GPVERTEX3F fnptr, GLfloat  x, GLfloat  y, GLfloat  z) {
//   (*fnptr)(x, y, z);
// }
// static void  glowVertex3fv(GPVERTEX3FV fnptr, const GLfloat * v) {
//   (*fnptr)(v);
// }
// static void  glowVertex3i(GPVERTEX3I fnptr, GLint  x, GLint  y, GLint  z) {
//   (*fnptr)(x, y, z);
// }
// static void  glowVertex3iv(GPVERTEX3IV fnptr, const GLint * v) {
//   (*fnptr)(v);
// }
// static void  glowVertex3s(GPVERTEX3S fnptr, GLshort  x, GLshort  y, GLshort  z) {
//   (*fnptr)(x, y, z);
// }
// static void  glowVertex3sv(GPVERTEX3SV fnptr, const GLshort * v) {
//   (*fnptr)(v);
// }
// static void  glowVertex4d(GPVERTEX4D fnptr, GLdouble  x, GLdouble  y, GLdouble  z, GLdouble  w) {
//   (*fnptr)(x, y, z, w);
// }
// static void  glowVertex4dv(GPVERTEX4DV fnptr, const GLdouble * v) {
//   (*fnptr)(v);
// }
// static void  glowVertex4f(GPVERTEX4F fnptr, GLfloat  x, GLfloat  y, GLfloat  z, GLfloat  w) {
//   (*fnptr)(x, y, z, w);
// }
// static void  glowVertex4fv(GPVERTEX4FV fnptr, const GLfloat * v) {
//   (*fnptr)(v);
// }
// static void  glowVertex4i(GPVERTEX4I fnptr, GLint  x, GLint  y, GLint  z, GLint  w) {
//   (*fnptr)(x, y, z, w);
// }
// static void  glowVertex4iv(GPVERTEX4IV fnptr, const GLint * v) {
//   (*fnptr)(v);
// }
// static void  glowVertex4s(GPVERTEX4S fnptr, GLshort  x, GLshort  y, GLshort  z, GLshort  w) {
//   (*fnptr)(x, y, z, w);
// }
// static void  glowVertex4sv(GPVERTEX4SV fnptr, const GLshort * v) {
//   (*fnptr)(v);
// }
// static void  glowVertexArrayAttribBinding(GPVERTEXARRAYATTRIBBINDING fnptr, GLuint  vaobj, GLuint  attribindex, GLuint  bindingindex) {
//   (*fnptr)(vaobj, attribindex, bindingindex);
// }
// static void  glowVertexArrayAttribFormat(GPVERTEXARRAYATTRIBFORMAT fnptr, GLuint  vaobj, GLuint  attribindex, GLint  size, GLenum  type, GLboolean  normalized, GLuint  relativeoffset) {
//   (*fnptr)(vaobj, attribindex, size, type, normalized, relativeoffset);
// }
// static void  glowVertexArrayAttribIFormat(GPVERTEXARRAYATTRIBIFORMAT fnptr, GLuint  vaobj, GLuint  attribindex, GLint  size, GLenum  type, GLuint  relativeoffset) {
//   (*fnptr)(vaobj, attribindex, size, type, relativeoffset);
// }
// static void  glowVertexArrayAttribLFormat(GPVERTEXARRAYATTRIBLFORMAT fnptr, GLuint  vaobj, GLuint  attribindex, GLint  size, GLenum  type, GLuint  relativeoffset) {
//   (*fnptr)(vaobj, attribindex, size, type, relativeoffset);
// }
// static void  glowVertexArrayBindingDivisor(GPVERTEXARRAYBINDINGDIVISOR fnptr, GLuint  vaobj, GLuint  bindingindex, GLuint  divisor) {
//   (*fnptr)(vaobj, bindingindex, divisor);
// }
// static void  glowVertexArrayElementBuffer(GPVERTEXARRAYELEMENTBUFFER fnptr, GLuint  vaobj, GLuint  buffer) {
//   (*fnptr)(vaobj, buffer);
// }
// static void  glowVertexArrayVertexBuffer(GPVERTEXARRAYVERTEXBUFFER fnptr, GLuint  vaobj, GLuint  bindingindex, GLuint  buffer, GLintptr  offset, GLsizei  stride) {
//   (*fnptr)(vaobj, bindingindex, buffer, offset, stride);
// }
// static void  glowVertexArrayVertexBuffers(GPVERTEXARRAYVERTEXBUFFERS fnptr, GLuint  vaobj, GLuint  first, GLsizei  count, const GLuint * buffers, const GLintptr * offsets, const GLsizei * strides) {
//   (*fnptr)(vaobj, first, count, buffers, offsets, strides);
// }
// static void  glowVertexAttrib1d(GPVERTEXATTRIB1D fnptr, GLuint  index, GLdouble  x) {
//   (*fnptr)(index, x);
// }
// static void  glowVertexAttrib1dv(GPVERTEXATTRIB1DV fnptr, GLuint  index, const GLdouble * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttrib1f(GPVERTEXATTRIB1F fnptr, GLuint  index, GLfloat  x) {
//   (*fnptr)(index, x);
// }
// static void  glowVertexAttrib1fv(GPVERTEXATTRIB1FV fnptr, GLuint  index, const GLfloat * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttrib1s(GPVERTEXATTRIB1S fnptr, GLuint  index, GLshort  x) {
//   (*fnptr)(index, x);
// }
// static void  glowVertexAttrib1sv(GPVERTEXATTRIB1SV fnptr, GLuint  index, const GLshort * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttrib2d(GPVERTEXATTRIB2D fnptr, GLuint  index, GLdouble  x, GLdouble  y) {
//   (*fnptr)(index, x, y);
// }
// static void  glowVertexAttrib2dv(GPVERTEXATTRIB2DV fnptr, GLuint  index, const GLdouble * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttrib2f(GPVERTEXATTRIB2F fnptr, GLuint  index, GLfloat  x, GLfloat  y) {
//   (*fnptr)(index, x, y);
// }
// static void  glowVertexAttrib2fv(GPVERTEXATTRIB2FV fnptr, GLuint  index, const GLfloat * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttrib2s(GPVERTEXATTRIB2S fnptr, GLuint  index, GLshort  x, GLshort  y) {
//   (*fnptr)(index, x, y);
// }
// static void  glowVertexAttrib2sv(GPVERTEXATTRIB2SV fnptr, GLuint  index, const GLshort * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttrib3d(GPVERTEXATTRIB3D fnptr, GLuint  index, GLdouble  x, GLdouble  y, GLdouble  z) {
//   (*fnptr)(index, x, y, z);
// }
// static void  glowVertexAttrib3dv(GPVERTEXATTRIB3DV fnptr, GLuint  index, const GLdouble * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttrib3f(GPVERTEXATTRIB3F fnptr, GLuint  index, GLfloat  x, GLfloat  y, GLfloat  z) {
//   (*fnptr)(index, x, y, z);
// }
// static void  glowVertexAttrib3fv(GPVERTEXATTRIB3FV fnptr, GLuint  index, const GLfloat * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttrib3s(GPVERTEXATTRIB3S fnptr, GLuint  index, GLshort  x, GLshort  y, GLshort  z) {
//   (*fnptr)(index, x, y, z);
// }
// static void  glowVertexAttrib3sv(GPVERTEXATTRIB3SV fnptr, GLuint  index, const GLshort * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttrib4Nbv(GPVERTEXATTRIB4NBV fnptr, GLuint  index, const GLbyte * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttrib4Niv(GPVERTEXATTRIB4NIV fnptr, GLuint  index, const GLint * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttrib4Nsv(GPVERTEXATTRIB4NSV fnptr, GLuint  index, const GLshort * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttrib4Nub(GPVERTEXATTRIB4NUB fnptr, GLuint  index, GLubyte  x, GLubyte  y, GLubyte  z, GLubyte  w) {
//   (*fnptr)(index, x, y, z, w);
// }
// static void  glowVertexAttrib4Nubv(GPVERTEXATTRIB4NUBV fnptr, GLuint  index, const GLubyte * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttrib4Nuiv(GPVERTEXATTRIB4NUIV fnptr, GLuint  index, const GLuint * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttrib4Nusv(GPVERTEXATTRIB4NUSV fnptr, GLuint  index, const GLushort * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttrib4bv(GPVERTEXATTRIB4BV fnptr, GLuint  index, const GLbyte * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttrib4d(GPVERTEXATTRIB4D fnptr, GLuint  index, GLdouble  x, GLdouble  y, GLdouble  z, GLdouble  w) {
//   (*fnptr)(index, x, y, z, w);
// }
// static void  glowVertexAttrib4dv(GPVERTEXATTRIB4DV fnptr, GLuint  index, const GLdouble * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttrib4f(GPVERTEXATTRIB4F fnptr, GLuint  index, GLfloat  x, GLfloat  y, GLfloat  z, GLfloat  w) {
//   (*fnptr)(index, x, y, z, w);
// }
// static void  glowVertexAttrib4fv(GPVERTEXATTRIB4FV fnptr, GLuint  index, const GLfloat * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttrib4iv(GPVERTEXATTRIB4IV fnptr, GLuint  index, const GLint * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttrib4s(GPVERTEXATTRIB4S fnptr, GLuint  index, GLshort  x, GLshort  y, GLshort  z, GLshort  w) {
//   (*fnptr)(index, x, y, z, w);
// }
// static void  glowVertexAttrib4sv(GPVERTEXATTRIB4SV fnptr, GLuint  index, const GLshort * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttrib4ubv(GPVERTEXATTRIB4UBV fnptr, GLuint  index, const GLubyte * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttrib4uiv(GPVERTEXATTRIB4UIV fnptr, GLuint  index, const GLuint * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttrib4usv(GPVERTEXATTRIB4USV fnptr, GLuint  index, const GLushort * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttribBinding(GPVERTEXATTRIBBINDING fnptr, GLuint  attribindex, GLuint  bindingindex) {
//   (*fnptr)(attribindex, bindingindex);
// }
// static void  glowVertexAttribDivisor(GPVERTEXATTRIBDIVISOR fnptr, GLuint  index, GLuint  divisor) {
//   (*fnptr)(index, divisor);
// }
// static void  glowVertexAttribFormat(GPVERTEXATTRIBFORMAT fnptr, GLuint  attribindex, GLint  size, GLenum  type, GLboolean  normalized, GLuint  relativeoffset) {
//   (*fnptr)(attribindex, size, type, normalized, relativeoffset);
// }
// static void  glowVertexAttribI1i(GPVERTEXATTRIBI1I fnptr, GLuint  index, GLint  x) {
//   (*fnptr)(index, x);
// }
// static void  glowVertexAttribI1iv(GPVERTEXATTRIBI1IV fnptr, GLuint  index, const GLint * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttribI1ui(GPVERTEXATTRIBI1UI fnptr, GLuint  index, GLuint  x) {
//   (*fnptr)(index, x);
// }
// static void  glowVertexAttribI1uiv(GPVERTEXATTRIBI1UIV fnptr, GLuint  index, const GLuint * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttribI2i(GPVERTEXATTRIBI2I fnptr, GLuint  index, GLint  x, GLint  y) {
//   (*fnptr)(index, x, y);
// }
// static void  glowVertexAttribI2iv(GPVERTEXATTRIBI2IV fnptr, GLuint  index, const GLint * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttribI2ui(GPVERTEXATTRIBI2UI fnptr, GLuint  index, GLuint  x, GLuint  y) {
//   (*fnptr)(index, x, y);
// }
// static void  glowVertexAttribI2uiv(GPVERTEXATTRIBI2UIV fnptr, GLuint  index, const GLuint * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttribI3i(GPVERTEXATTRIBI3I fnptr, GLuint  index, GLint  x, GLint  y, GLint  z) {
//   (*fnptr)(index, x, y, z);
// }
// static void  glowVertexAttribI3iv(GPVERTEXATTRIBI3IV fnptr, GLuint  index, const GLint * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttribI3ui(GPVERTEXATTRIBI3UI fnptr, GLuint  index, GLuint  x, GLuint  y, GLuint  z) {
//   (*fnptr)(index, x, y, z);
// }
// static void  glowVertexAttribI3uiv(GPVERTEXATTRIBI3UIV fnptr, GLuint  index, const GLuint * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttribI4bv(GPVERTEXATTRIBI4BV fnptr, GLuint  index, const GLbyte * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttribI4i(GPVERTEXATTRIBI4I fnptr, GLuint  index, GLint  x, GLint  y, GLint  z, GLint  w) {
//   (*fnptr)(index, x, y, z, w);
// }
// static void  glowVertexAttribI4iv(GPVERTEXATTRIBI4IV fnptr, GLuint  index, const GLint * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttribI4sv(GPVERTEXATTRIBI4SV fnptr, GLuint  index, const GLshort * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttribI4ubv(GPVERTEXATTRIBI4UBV fnptr, GLuint  index, const GLubyte * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttribI4ui(GPVERTEXATTRIBI4UI fnptr, GLuint  index, GLuint  x, GLuint  y, GLuint  z, GLuint  w) {
//   (*fnptr)(index, x, y, z, w);
// }
// static void  glowVertexAttribI4uiv(GPVERTEXATTRIBI4UIV fnptr, GLuint  index, const GLuint * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttribI4usv(GPVERTEXATTRIBI4USV fnptr, GLuint  index, const GLushort * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttribIFormat(GPVERTEXATTRIBIFORMAT fnptr, GLuint  attribindex, GLint  size, GLenum  type, GLuint  relativeoffset) {
//   (*fnptr)(attribindex, size, type, relativeoffset);
// }
// static void  glowVertexAttribIPointer(GPVERTEXATTRIBIPOINTER fnptr, GLuint  index, GLint  size, GLenum  type, GLsizei  stride, const void * pointer) {
//   (*fnptr)(index, size, type, stride, pointer);
// }
// static void  glowVertexAttribL1d(GPVERTEXATTRIBL1D fnptr, GLuint  index, GLdouble  x) {
//   (*fnptr)(index, x);
// }
// static void  glowVertexAttribL1dv(GPVERTEXATTRIBL1DV fnptr, GLuint  index, const GLdouble * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttribL1ui64ARB(GPVERTEXATTRIBL1UI64ARB fnptr, GLuint  index, GLuint64EXT  x) {
//   (*fnptr)(index, x);
// }
// static void  glowVertexAttribL1ui64vARB(GPVERTEXATTRIBL1UI64VARB fnptr, GLuint  index, const GLuint64EXT * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttribL2d(GPVERTEXATTRIBL2D fnptr, GLuint  index, GLdouble  x, GLdouble  y) {
//   (*fnptr)(index, x, y);
// }
// static void  glowVertexAttribL2dv(GPVERTEXATTRIBL2DV fnptr, GLuint  index, const GLdouble * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttribL3d(GPVERTEXATTRIBL3D fnptr, GLuint  index, GLdouble  x, GLdouble  y, GLdouble  z) {
//   (*fnptr)(index, x, y, z);
// }
// static void  glowVertexAttribL3dv(GPVERTEXATTRIBL3DV fnptr, GLuint  index, const GLdouble * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttribL4d(GPVERTEXATTRIBL4D fnptr, GLuint  index, GLdouble  x, GLdouble  y, GLdouble  z, GLdouble  w) {
//   (*fnptr)(index, x, y, z, w);
// }
// static void  glowVertexAttribL4dv(GPVERTEXATTRIBL4DV fnptr, GLuint  index, const GLdouble * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttribLFormat(GPVERTEXATTRIBLFORMAT fnptr, GLuint  attribindex, GLint  size, GLenum  type, GLuint  relativeoffset) {
//   (*fnptr)(attribindex, size, type, relativeoffset);
// }
// static void  glowVertexAttribLPointer(GPVERTEXATTRIBLPOINTER fnptr, GLuint  index, GLint  size, GLenum  type, GLsizei  stride, const void * pointer) {
//   (*fnptr)(index, size, type, stride, pointer);
// }
// static void  glowVertexAttribP1ui(GPVERTEXATTRIBP1UI fnptr, GLuint  index, GLenum  type, GLboolean  normalized, GLuint  value) {
//   (*fnptr)(index, type, normalized, value);
// }
// static void  glowVertexAttribP1uiv(GPVERTEXATTRIBP1UIV fnptr, GLuint  index, GLenum  type, GLboolean  normalized, const GLuint * value) {
//   (*fnptr)(index, type, normalized, value);
// }
// static void  glowVertexAttribP2ui(GPVERTEXATTRIBP2UI fnptr, GLuint  index, GLenum  type, GLboolean  normalized, GLuint  value) {
//   (*fnptr)(index, type, normalized, value);
// }
// static void  glowVertexAttribP2uiv(GPVERTEXATTRIBP2UIV fnptr, GLuint  index, GLenum  type, GLboolean  normalized, const GLuint * value) {
//   (*fnptr)(index, type, normalized, value);
// }
// static void  glowVertexAttribP3ui(GPVERTEXATTRIBP3UI fnptr, GLuint  index, GLenum  type, GLboolean  normalized, GLuint  value) {
//   (*fnptr)(index, type, normalized, value);
// }
// static void  glowVertexAttribP3uiv(GPVERTEXATTRIBP3UIV fnptr, GLuint  index, GLenum  type, GLboolean  normalized, const GLuint * value) {
//   (*fnptr)(index, type, normalized, value);
// }
// static void  glowVertexAttribP4ui(GPVERTEXATTRIBP4UI fnptr, GLuint  index, GLenum  type, GLboolean  normalized, GLuint  value) {
//   (*fnptr)(index, type, normalized, value);
// }
// static void  glowVertexAttribP4uiv(GPVERTEXATTRIBP4UIV fnptr, GLuint  index, GLenum  type, GLboolean  normalized, const GLuint * value) {
//   (*fnptr)(index, type, normalized, value);
// }
// static void  glowVertexAttribPointer(GPVERTEXATTRIBPOINTER fnptr, GLuint  index, GLint  size, GLenum  type, GLboolean  normalized, GLsizei  stride, const void * pointer) {
//   (*fnptr)(index, size, type, normalized, stride, pointer);
// }
// static void  glowVertexBindingDivisor(GPVERTEXBINDINGDIVISOR fnptr, GLuint  bindingindex, GLuint  divisor) {
//   (*fnptr)(bindingindex, divisor);
// }
// static void  glowVertexPointer(GPVERTEXPOINTER fnptr, GLint  size, GLenum  type, GLsizei  stride, const void * pointer) {
//   (*fnptr)(size, type, stride, pointer);
// }
// static void  glowViewport(GPVIEWPORT fnptr, GLint  x, GLint  y, GLsizei  width, GLsizei  height) {
//   (*fnptr)(x, y, width, height);
// }
// static void  glowViewportArrayv(GPVIEWPORTARRAYV fnptr, GLuint  first, GLsizei  count, const GLfloat * v) {
//   (*fnptr)(first, count, v);
// }
// static void  glowViewportIndexedf(GPVIEWPORTINDEXEDF fnptr, GLuint  index, GLfloat  x, GLfloat  y, GLfloat  w, GLfloat  h) {
//   (*fnptr)(index, x, y, w, h);
// }
// static void  glowViewportIndexedfv(GPVIEWPORTINDEXEDFV fnptr, GLuint  index, const GLfloat * v) {
//   (*fnptr)(index, v);
// }
// static void  glowWaitSync(GPWAITSYNC fnptr, GLsync  sync, GLbitfield  flags, GLuint64  timeout) {
//   (*fnptr)(sync, flags, timeout);
// }
// static void  glowWindowPos2d(GPWINDOWPOS2D fnptr, GLdouble  x, GLdouble  y) {
//   (*fnptr)(x, y);
// }
// static void  glowWindowPos2dv(GPWINDOWPOS2DV fnptr, const GLdouble * v) {
//   (*fnptr)(v);
// }
// static void  glowWindowPos2f(GPWINDOWPOS2F fnptr, GLfloat  x, GLfloat  y) {
//   (*fnptr)(x, y);
// }
// static void  glowWindowPos2fv(GPWINDOWPOS2FV fnptr, const GLfloat * v) {
//   (*fnptr)(v);
// }
// static void  glowWindowPos2i(GPWINDOWPOS2I fnptr, GLint  x, GLint  y) {
//   (*fnptr)(x, y);
// }
// static void  glowWindowPos2iv(GPWINDOWPOS2IV fnptr, const GLint * v) {
//   (*fnptr)(v);
// }
// static void  glowWindowPos2s(GPWINDOWPOS2S fnptr, GLshort  x, GLshort  y) {
//   (*fnptr)(x, y);
// }
// static void  glowWindowPos2sv(GPWINDOWPOS2SV fnptr, const GLshort * v) {
//   (*fnptr)(v);
// }
// static void  glowWindowPos3d(GPWINDOWPOS3D fnptr, GLdouble  x, GLdouble  y, GLdouble  z) {
//   (*fnptr)(x, y, z);
// }
// static void  glowWindowPos3dv(GPWINDOWPOS3DV fnptr, const GLdouble * v) {
//   (*fnptr)(v);
// }
// static void  glowWindowPos3f(GPWINDOWPOS3F fnptr, GLfloat  x, GLfloat  y, GLfloat  z) {
//   (*fnptr)(x, y, z);
// }
// static void  glowWindowPos3fv(GPWINDOWPOS3FV fnptr, const GLfloat * v) {
//   (*fnptr)(v);
// }
// static void  glowWindowPos3i(GPWINDOWPOS3I fnptr, GLint  x, GLint  y, GLint  z) {
//   (*fnptr)(x, y, z);
// }
// static void  glowWindowPos3iv(GPWINDOWPOS3IV fnptr, const GLint * v) {
//   (*fnptr)(v);
// }
// static void  glowWindowPos3s(GPWINDOWPOS3S fnptr, GLshort  x, GLshort  y, GLshort  z) {
//   (*fnptr)(x, y, z);
// }
// static void  glowWindowPos3sv(GPWINDOWPOS3SV fnptr, const GLshort * v) {
//   (*fnptr)(v);
// }
import "C"
import (
	"unsafe"
)

const (
	GL_2D                                                      = 0x0600
	GL_2_BYTES                                                 = 0x1407
	GL_3D                                                      = 0x0601
	GL_3D_COLOR                                                = 0x0602
	GL_3D_COLOR_TEXTURE                                        = 0x0603
	GL_3_BYTES                                                 = 0x1408
	GL_4D_COLOR_TEXTURE                                        = 0x0604
	GL_4_BYTES                                                 = 0x1409
	ACCUM                                                      = 0x0100
	ACCUM_ALPHA_BITS                                           = 0x0D5B
	ACCUM_BLUE_BITS                                            = 0x0D5A
	ACCUM_BUFFER_BIT                                           = 0x00000200
	ACCUM_CLEAR_VALUE                                          = 0x0B80
	ACCUM_GREEN_BITS                                           = 0x0D59
	ACCUM_RED_BITS                                             = 0x0D58
	ACTIVE_ATOMIC_COUNTER_BUFFERS                              = 0x92D9
	ACTIVE_ATTRIBUTES                                          = 0x8B89
	ACTIVE_ATTRIBUTE_MAX_LENGTH                                = 0x8B8A
	ACTIVE_PROGRAM                                             = 0x8259
	ACTIVE_RESOURCES                                           = 0x92F5
	ACTIVE_SUBROUTINES                                         = 0x8DE5
	ACTIVE_SUBROUTINE_MAX_LENGTH                               = 0x8E48
	ACTIVE_SUBROUTINE_UNIFORMS                                 = 0x8DE6
	ACTIVE_SUBROUTINE_UNIFORM_LOCATIONS                        = 0x8E47
	ACTIVE_SUBROUTINE_UNIFORM_MAX_LENGTH                       = 0x8E49
	ACTIVE_TEXTURE                                             = 0x84E0
	ACTIVE_UNIFORMS                                            = 0x8B86
	ACTIVE_UNIFORM_BLOCKS                                      = 0x8A36
	ACTIVE_UNIFORM_BLOCK_MAX_NAME_LENGTH                       = 0x8A35
	ACTIVE_UNIFORM_MAX_LENGTH                                  = 0x8B87
	ACTIVE_VARIABLES                                           = 0x9305
	ADD                                                        = 0x0104
	ADD_SIGNED                                                 = 0x8574
	ALIASED_LINE_WIDTH_RANGE                                   = 0x846E
	ALIASED_POINT_SIZE_RANGE                                   = 0x846D
	ALL_ATTRIB_BITS                                            = 0xFFFFFFFF
	ALL_BARRIER_BITS                                           = 0xFFFFFFFF
	ALL_SHADER_BITS                                            = 0xFFFFFFFF
	ALPHA                                                      = 0x1906
	ALPHA12                                                    = 0x803D
	ALPHA16                                                    = 0x803E
	ALPHA4                                                     = 0x803B
	ALPHA8                                                     = 0x803C
	ALPHA_BIAS                                                 = 0x0D1D
	ALPHA_BITS                                                 = 0x0D55
	ALPHA_INTEGER                                              = 0x8D97
	ALPHA_SCALE                                                = 0x0D1C
	ALPHA_TEST                                                 = 0x0BC0
	ALPHA_TEST_FUNC                                            = 0x0BC1
	ALPHA_TEST_REF                                             = 0x0BC2
	ALREADY_SIGNALED                                           = 0x911A
	ALWAYS                                                     = 0x0207
	AMBIENT                                                    = 0x1200
	AMBIENT_AND_DIFFUSE                                        = 0x1602
	AND                                                        = 0x1501
	AND_INVERTED                                               = 0x1504
	AND_REVERSE                                                = 0x1502
	ANY_SAMPLES_PASSED                                         = 0x8C2F
	ANY_SAMPLES_PASSED_CONSERVATIVE                            = 0x8D6A
	ARRAY_BUFFER                                               = 0x8892
	ARRAY_BUFFER_BINDING                                       = 0x8894
	ARRAY_SIZE                                                 = 0x92FB
	ARRAY_STRIDE                                               = 0x92FE
	ATOMIC_COUNTER_BARRIER_BIT                                 = 0x00001000
	ATOMIC_COUNTER_BUFFER                                      = 0x92C0
	ATOMIC_COUNTER_BUFFER_ACTIVE_ATOMIC_COUNTERS               = 0x92C5
	ATOMIC_COUNTER_BUFFER_ACTIVE_ATOMIC_COUNTER_INDICES        = 0x92C6
	ATOMIC_COUNTER_BUFFER_BINDING                              = 0x92C1
	ATOMIC_COUNTER_BUFFER_DATA_SIZE                            = 0x92C4
	ATOMIC_COUNTER_BUFFER_INDEX                                = 0x9301
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_COMPUTE_SHADER         = 0x90ED
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_FRAGMENT_SHADER        = 0x92CB
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_GEOMETRY_SHADER        = 0x92CA
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_TESS_CONTROL_SHADER    = 0x92C8
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_TESS_EVALUATION_SHADER = 0x92C9
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_VERTEX_SHADER          = 0x92C7
	ATOMIC_COUNTER_BUFFER_SIZE                                 = 0x92C3
	ATOMIC_COUNTER_BUFFER_START                                = 0x92C2
	ATTACHED_SHADERS                                           = 0x8B85
	ATTRIB_STACK_DEPTH                                         = 0x0BB0
	AUTO_GENERATE_MIPMAP                                       = 0x8295
	AUTO_NORMAL                                                = 0x0D80
	AUX0                                                       = 0x0409
	AUX1                                                       = 0x040A
	AUX2                                                       = 0x040B
	AUX3                                                       = 0x040C
	AUX_BUFFERS                                                = 0x0C00
	BACK                                                       = 0x0405
	BACK_LEFT                                                  = 0x0402
	BACK_RIGHT                                                 = 0x0403
	BGR                                                        = 0x80E0
	BGRA                                                       = 0x80E1
	BGRA_INTEGER                                               = 0x8D9B
	BGR_INTEGER                                                = 0x8D9A
	BITMAP                                                     = 0x1A00
	BITMAP_TOKEN                                               = 0x0704
	BLEND                                                      = 0x0BE2
	BLEND_COLOR                                                = 0x8005
	BLEND_DST                                                  = 0x0BE0
	BLEND_DST_ALPHA                                            = 0x80CA
	BLEND_DST_RGB                                              = 0x80C8
	BLEND_EQUATION                                             = 0x8009
	BLEND_EQUATION_ALPHA                                       = 0x883D
	BLEND_EQUATION_RGB                                         = 0x8009
	BLEND_SRC                                                  = 0x0BE1
	BLEND_SRC_ALPHA                                            = 0x80CB
	BLEND_SRC_RGB                                              = 0x80C9
	BLOCK_INDEX                                                = 0x92FD
	BLUE                                                       = 0x1905
	BLUE_BIAS                                                  = 0x0D1B
	BLUE_BITS                                                  = 0x0D54
	BLUE_INTEGER                                               = 0x8D96
	BLUE_SCALE                                                 = 0x0D1A
	BOOL                                                       = 0x8B56
	BOOL_VEC2                                                  = 0x8B57
	BOOL_VEC3                                                  = 0x8B58
	BOOL_VEC4                                                  = 0x8B59
	BUFFER                                                     = 0x82E0
	BUFFER_ACCESS                                              = 0x88BB
	BUFFER_ACCESS_FLAGS                                        = 0x911F
	BUFFER_BINDING                                             = 0x9302
	BUFFER_DATA_SIZE                                           = 0x9303
	BUFFER_IMMUTABLE_STORAGE                                   = 0x821F
	BUFFER_KHR                                                 = 0x82E0
	BUFFER_MAPPED                                              = 0x88BC
	BUFFER_MAP_LENGTH                                          = 0x9120
	BUFFER_MAP_OFFSET                                          = 0x9121
	BUFFER_MAP_POINTER                                         = 0x88BD
	BUFFER_SIZE                                                = 0x8764
	BUFFER_STORAGE_FLAGS                                       = 0x8220
	BUFFER_UPDATE_BARRIER_BIT                                  = 0x00000200
	BUFFER_USAGE                                               = 0x8765
	BUFFER_VARIABLE                                            = 0x92E5
	BYTE                                                       = 0x1400
	C3F_V3F                                                    = 0x2A24
	C4F_N3F_V3F                                                = 0x2A26
	C4UB_V2F                                                   = 0x2A22
	C4UB_V3F                                                   = 0x2A23
	CAVEAT_SUPPORT                                             = 0x82B8
	CCW                                                        = 0x0901
	CLAMP                                                      = 0x2900
	CLAMP_FRAGMENT_COLOR                                       = 0x891B
	CLAMP_READ_COLOR                                           = 0x891C
	CLAMP_TO_BORDER                                            = 0x812D
	CLAMP_TO_EDGE                                              = 0x812F
	CLAMP_VERTEX_COLOR                                         = 0x891A
	CLEAR                                                      = 0x1500
	CLEAR_BUFFER                                               = 0x82B4
	CLEAR_TEXTURE                                              = 0x9365
	CLIENT_ACTIVE_TEXTURE                                      = 0x84E1
	CLIENT_ALL_ATTRIB_BITS                                     = 0xFFFFFFFF
	CLIENT_ATTRIB_STACK_DEPTH                                  = 0x0BB1
	CLIENT_MAPPED_BUFFER_BARRIER_BIT                           = 0x00004000
	CLIENT_PIXEL_STORE_BIT                                     = 0x00000001
	CLIENT_STORAGE_BIT                                         = 0x0200
	CLIENT_VERTEX_ARRAY_BIT                                    = 0x00000002
	CLIPPING_INPUT_PRIMITIVES_ARB                              = 0x82F6
	CLIPPING_OUTPUT_PRIMITIVES_ARB                             = 0x82F7
	CLIP_DEPTH_MODE                                            = 0x935D
	CLIP_DISTANCE0                                             = 0x3000
	CLIP_DISTANCE1                                             = 0x3001
	CLIP_DISTANCE2                                             = 0x3002
	CLIP_DISTANCE3                                             = 0x3003
	CLIP_DISTANCE4                                             = 0x3004
	CLIP_DISTANCE5                                             = 0x3005
	CLIP_DISTANCE6                                             = 0x3006
	CLIP_DISTANCE7                                             = 0x3007
	CLIP_ORIGIN                                                = 0x935C
	CLIP_PLANE0                                                = 0x3000
	CLIP_PLANE1                                                = 0x3001
	CLIP_PLANE2                                                = 0x3002
	CLIP_PLANE3                                                = 0x3003
	CLIP_PLANE4                                                = 0x3004
	CLIP_PLANE5                                                = 0x3005
	COEFF                                                      = 0x0A00
	COLOR                                                      = 0x1800
	COLOR_ARRAY                                                = 0x8076
	COLOR_ARRAY_BUFFER_BINDING                                 = 0x8898
	COLOR_ARRAY_POINTER                                        = 0x8090
	COLOR_ARRAY_SIZE                                           = 0x8081
	COLOR_ARRAY_STRIDE                                         = 0x8083
	COLOR_ARRAY_TYPE                                           = 0x8082
	COLOR_ATTACHMENT0                                          = 0x8CE0
	COLOR_ATTACHMENT1                                          = 0x8CE1
	COLOR_ATTACHMENT10                                         = 0x8CEA
	COLOR_ATTACHMENT11                                         = 0x8CEB
	COLOR_ATTACHMENT12                                         = 0x8CEC
	COLOR_ATTACHMENT13                                         = 0x8CED
	COLOR_ATTACHMENT14                                         = 0x8CEE
	COLOR_ATTACHMENT15                                         = 0x8CEF
	COLOR_ATTACHMENT2                                          = 0x8CE2
	COLOR_ATTACHMENT3                                          = 0x8CE3
	COLOR_ATTACHMENT4                                          = 0x8CE4
	COLOR_ATTACHMENT5                                          = 0x8CE5
	COLOR_ATTACHMENT6                                          = 0x8CE6
	COLOR_ATTACHMENT7                                          = 0x8CE7
	COLOR_ATTACHMENT8                                          = 0x8CE8
	COLOR_ATTACHMENT9                                          = 0x8CE9
	COLOR_BUFFER_BIT                                           = 0x00004000
	COLOR_CLEAR_VALUE                                          = 0x0C22
	COLOR_COMPONENTS                                           = 0x8283
	COLOR_ENCODING                                             = 0x8296
	COLOR_INDEX                                                = 0x1900
	COLOR_INDEXES                                              = 0x1603
	COLOR_LOGIC_OP                                             = 0x0BF2
	COLOR_MATERIAL                                             = 0x0B57
	COLOR_MATERIAL_FACE                                        = 0x0B55
	COLOR_MATERIAL_PARAMETER                                   = 0x0B56
	COLOR_RENDERABLE                                           = 0x8286
	COLOR_SUM                                                  = 0x8458
	COLOR_WRITEMASK                                            = 0x0C23
	COMBINE                                                    = 0x8570
	COMBINE_ALPHA                                              = 0x8572
	COMBINE_RGB                                                = 0x8571
	COMMAND_BARRIER_BIT                                        = 0x00000040
	COMPARE_REF_TO_TEXTURE                                     = 0x884E
	COMPARE_R_TO_TEXTURE                                       = 0x884E
	COMPATIBLE_SUBROUTINES                                     = 0x8E4B
	COMPILE                                                    = 0x1300
	COMPILE_AND_EXECUTE                                        = 0x1301
	COMPILE_STATUS                                             = 0x8B81
	COMPRESSED_ALPHA                                           = 0x84E9
	COMPRESSED_INTENSITY                                       = 0x84EC
	COMPRESSED_LUMINANCE                                       = 0x84EA
	COMPRESSED_LUMINANCE_ALPHA                                 = 0x84EB
	COMPRESSED_R11_EAC                                         = 0x9270
	COMPRESSED_RED                                             = 0x8225
	COMPRESSED_RED_RGTC1                                       = 0x8DBB
	COMPRESSED_RG                                              = 0x8226
	COMPRESSED_RG11_EAC                                        = 0x9272
	COMPRESSED_RGB                                             = 0x84ED
	COMPRESSED_RGB8_ETC2                                       = 0x9274
	COMPRESSED_RGB8_PUNCHTHROUGH_ALPHA1_ETC2                   = 0x9276
	COMPRESSED_RGBA                                            = 0x84EE
	COMPRESSED_RGBA8_ETC2_EAC                                  = 0x9278
	COMPRESSED_RGBA_ASTC_10x10_KHR                             = 0x93BB
	COMPRESSED_RGBA_ASTC_10x5_KHR                              = 0x93B8
	COMPRESSED_RGBA_ASTC_10x6_KHR                              = 0x93B9
	COMPRESSED_RGBA_ASTC_10x8_KHR                              = 0x93BA
	COMPRESSED_RGBA_ASTC_12x10_KHR                             = 0x93BC
	COMPRESSED_RGBA_ASTC_12x12_KHR                             = 0x93BD
	COMPRESSED_RGBA_ASTC_4x4_KHR                               = 0x93B0
	COMPRESSED_RGBA_ASTC_5x4_KHR                               = 0x93B1
	COMPRESSED_RGBA_ASTC_5x5_KHR                               = 0x93B2
	COMPRESSED_RGBA_ASTC_6x5_KHR                               = 0x93B3
	COMPRESSED_RGBA_ASTC_6x6_KHR                               = 0x93B4
	COMPRESSED_RGBA_ASTC_8x5_KHR                               = 0x93B5
	COMPRESSED_RGBA_ASTC_8x6_KHR                               = 0x93B6
	COMPRESSED_RGBA_ASTC_8x8_KHR                               = 0x93B7
	COMPRESSED_RGBA_BPTC_UNORM                                 = 0x8E8C
	COMPRESSED_RGBA_BPTC_UNORM_ARB                             = 0x8E8C
	COMPRESSED_RGB_BPTC_SIGNED_FLOAT                           = 0x8E8E
	COMPRESSED_RGB_BPTC_SIGNED_FLOAT_ARB                       = 0x8E8E
	COMPRESSED_RGB_BPTC_UNSIGNED_FLOAT                         = 0x8E8F
	COMPRESSED_RGB_BPTC_UNSIGNED_FLOAT_ARB                     = 0x8E8F
	COMPRESSED_RG_RGTC2                                        = 0x8DBD
	COMPRESSED_SIGNED_R11_EAC                                  = 0x9271
	COMPRESSED_SIGNED_RED_RGTC1                                = 0x8DBC
	COMPRESSED_SIGNED_RG11_EAC                                 = 0x9273
	COMPRESSED_SIGNED_RG_RGTC2                                 = 0x8DBE
	COMPRESSED_SLUMINANCE                                      = 0x8C4A
	COMPRESSED_SLUMINANCE_ALPHA                                = 0x8C4B
	COMPRESSED_SRGB                                            = 0x8C48
	COMPRESSED_SRGB8_ALPHA8_ASTC_10x10_KHR                     = 0x93DB
	COMPRESSED_SRGB8_ALPHA8_ASTC_10x5_KHR                      = 0x93D8
	COMPRESSED_SRGB8_ALPHA8_ASTC_10x6_KHR                      = 0x93D9
	COMPRESSED_SRGB8_ALPHA8_ASTC_10x8_KHR                      = 0x93DA
	COMPRESSED_SRGB8_ALPHA8_ASTC_12x10_KHR                     = 0x93DC
	COMPRESSED_SRGB8_ALPHA8_ASTC_12x12_KHR                     = 0x93DD
	COMPRESSED_SRGB8_ALPHA8_ASTC_4x4_KHR                       = 0x93D0
	COMPRESSED_SRGB8_ALPHA8_ASTC_5x4_KHR                       = 0x93D1
	COMPRESSED_SRGB8_ALPHA8_ASTC_5x5_KHR                       = 0x93D2
	COMPRESSED_SRGB8_ALPHA8_ASTC_6x5_KHR                       = 0x93D3
	COMPRESSED_SRGB8_ALPHA8_ASTC_6x6_KHR                       = 0x93D4
	COMPRESSED_SRGB8_ALPHA8_ASTC_8x5_KHR                       = 0x93D5
	COMPRESSED_SRGB8_ALPHA8_ASTC_8x6_KHR                       = 0x93D6
	COMPRESSED_SRGB8_ALPHA8_ASTC_8x8_KHR                       = 0x93D7
	COMPRESSED_SRGB8_ALPHA8_ETC2_EAC                           = 0x9279
	COMPRESSED_SRGB8_ETC2                                      = 0x9275
	COMPRESSED_SRGB8_PUNCHTHROUGH_ALPHA1_ETC2                  = 0x9277
	COMPRESSED_SRGB_ALPHA                                      = 0x8C49
	COMPRESSED_SRGB_ALPHA_BPTC_UNORM                           = 0x8E8D
	COMPRESSED_SRGB_ALPHA_BPTC_UNORM_ARB                       = 0x8E8D
	COMPRESSED_TEXTURE_FORMATS                                 = 0x86A3
	COMPUTE_SHADER                                             = 0x91B9
	COMPUTE_SHADER_BIT                                         = 0x00000020
	COMPUTE_SHADER_INVOCATIONS_ARB                             = 0x82F5
	COMPUTE_SUBROUTINE                                         = 0x92ED
	COMPUTE_SUBROUTINE_UNIFORM                                 = 0x92F3
	COMPUTE_TEXTURE                                            = 0x82A0
	COMPUTE_WORK_GROUP_SIZE                                    = 0x8267
	CONDITION_SATISFIED                                        = 0x911C
	CONSTANT                                                   = 0x8576
	CONSTANT_ALPHA                                             = 0x8003
	CONSTANT_ATTENUATION                                       = 0x1207
	CONSTANT_COLOR                                             = 0x8001
	CONTEXT_COMPATIBILITY_PROFILE_BIT                          = 0x00000002
	CONTEXT_CORE_PROFILE_BIT                                   = 0x00000001
	CONTEXT_FLAGS                                              = 0x821E
	CONTEXT_FLAG_DEBUG_BIT                                     = 0x00000002
	CONTEXT_FLAG_DEBUG_BIT_KHR                                 = 0x00000002
	CONTEXT_FLAG_FORWARD_COMPATIBLE_BIT                        = 0x00000001
	CONTEXT_FLAG_ROBUST_ACCESS_BIT                             = 0x00000004
	CONTEXT_FLAG_ROBUST_ACCESS_BIT_ARB                         = 0x00000004
	CONTEXT_LOST                                               = 0x0507
	CONTEXT_LOST_KHR                                           = 0x0507
	CONTEXT_PROFILE_MASK                                       = 0x9126
	CONTEXT_RELEASE_BEHAVIOR                                   = 0x82FB
	CONTEXT_RELEASE_BEHAVIOR_FLUSH                             = 0x82FC
	CONTEXT_RELEASE_BEHAVIOR_FLUSH_KHR                         = 0x82FC
	CONTEXT_RELEASE_BEHAVIOR_KHR                               = 0x82FB
	CONTEXT_ROBUST_ACCESS                                      = 0x90F3
	CONTEXT_ROBUST_ACCESS_KHR                                  = 0x90F3
	COORD_REPLACE                                              = 0x8862
	COPY                                                       = 0x1503
	COPY_INVERTED                                              = 0x150C
	COPY_PIXEL_TOKEN                                           = 0x0706
	COPY_READ_BUFFER                                           = 0x8F36
	COPY_READ_BUFFER_BINDING                                   = 0x8F36
	COPY_WRITE_BUFFER                                          = 0x8F37
	COPY_WRITE_BUFFER_BINDING                                  = 0x8F37
	CULL_FACE                                                  = 0x0B44
	CULL_FACE_MODE                                             = 0x0B45
	CURRENT_BIT                                                = 0x00000001
	CURRENT_COLOR                                              = 0x0B00
	CURRENT_FOG_COORD                                          = 0x8453
	CURRENT_FOG_COORDINATE                                     = 0x8453
	CURRENT_INDEX                                              = 0x0B01
	CURRENT_NORMAL                                             = 0x0B02
	CURRENT_PROGRAM                                            = 0x8B8D
	CURRENT_QUERY                                              = 0x8865
	CURRENT_RASTER_COLOR                                       = 0x0B04
	CURRENT_RASTER_DISTANCE                                    = 0x0B09
	CURRENT_RASTER_INDEX                                       = 0x0B05
	CURRENT_RASTER_POSITION                                    = 0x0B07
	CURRENT_RASTER_POSITION_VALID                              = 0x0B08
	CURRENT_RASTER_SECONDARY_COLOR                             = 0x845F
	CURRENT_RASTER_TEXTURE_COORDS                              = 0x0B06
	CURRENT_SECONDARY_COLOR                                    = 0x8459
	CURRENT_TEXTURE_COORDS                                     = 0x0B03
	CURRENT_VERTEX_ATTRIB                                      = 0x8626
	CW                                                         = 0x0900
	DEBUG_CALLBACK_FUNCTION                                    = 0x8244
	DEBUG_CALLBACK_FUNCTION_ARB                                = 0x8244
	DEBUG_CALLBACK_FUNCTION_KHR                                = 0x8244
	DEBUG_CALLBACK_USER_PARAM                                  = 0x8245
	DEBUG_CALLBACK_USER_PARAM_ARB                              = 0x8245
	DEBUG_CALLBACK_USER_PARAM_KHR                              = 0x8245
	DEBUG_GROUP_STACK_DEPTH                                    = 0x826D
	DEBUG_GROUP_STACK_DEPTH_KHR                                = 0x826D
	DEBUG_LOGGED_MESSAGES                                      = 0x9145
	DEBUG_LOGGED_MESSAGES_ARB                                  = 0x9145
	DEBUG_LOGGED_MESSAGES_KHR                                  = 0x9145
	DEBUG_NEXT_LOGGED_MESSAGE_LENGTH                           = 0x8243
	DEBUG_NEXT_LOGGED_MESSAGE_LENGTH_ARB                       = 0x8243
	DEBUG_NEXT_LOGGED_MESSAGE_LENGTH_KHR                       = 0x8243
	DEBUG_OUTPUT                                               = 0x92E0
	DEBUG_OUTPUT_KHR                                           = 0x92E0
	DEBUG_OUTPUT_SYNCHRONOUS                                   = 0x8242
	DEBUG_OUTPUT_SYNCHRONOUS_ARB                               = 0x8242
	DEBUG_OUTPUT_SYNCHRONOUS_KHR                               = 0x8242
	DEBUG_SEVERITY_HIGH                                        = 0x9146
	DEBUG_SEVERITY_HIGH_ARB                                    = 0x9146
	DEBUG_SEVERITY_HIGH_KHR                                    = 0x9146
	DEBUG_SEVERITY_LOW                                         = 0x9148
	DEBUG_SEVERITY_LOW_ARB                                     = 0x9148
	DEBUG_SEVERITY_LOW_KHR                                     = 0x9148
	DEBUG_SEVERITY_MEDIUM                                      = 0x9147
	DEBUG_SEVERITY_MEDIUM_ARB                                  = 0x9147
	DEBUG_SEVERITY_MEDIUM_KHR                                  = 0x9147
	DEBUG_SEVERITY_NOTIFICATION                                = 0x826B
	DEBUG_SEVERITY_NOTIFICATION_KHR                            = 0x826B
	DEBUG_SOURCE_API                                           = 0x8246
	DEBUG_SOURCE_API_ARB                                       = 0x8246
	DEBUG_SOURCE_API_KHR                                       = 0x8246
	DEBUG_SOURCE_APPLICATION                                   = 0x824A
	DEBUG_SOURCE_APPLICATION_ARB                               = 0x824A
	DEBUG_SOURCE_APPLICATION_KHR                               = 0x824A
	DEBUG_SOURCE_OTHER                                         = 0x824B
	DEBUG_SOURCE_OTHER_ARB                                     = 0x824B
	DEBUG_SOURCE_OTHER_KHR                                     = 0x824B
	DEBUG_SOURCE_SHADER_COMPILER                               = 0x8248
	DEBUG_SOURCE_SHADER_COMPILER_ARB                           = 0x8248
	DEBUG_SOURCE_SHADER_COMPILER_KHR                           = 0x8248
	DEBUG_SOURCE_THIRD_PARTY                                   = 0x8249
	DEBUG_SOURCE_THIRD_PARTY_ARB                               = 0x8249
	DEBUG_SOURCE_THIRD_PARTY_KHR                               = 0x8249
	DEBUG_SOURCE_WINDOW_SYSTEM                                 = 0x8247
	DEBUG_SOURCE_WINDOW_SYSTEM_ARB                             = 0x8247
	DEBUG_SOURCE_WINDOW_SYSTEM_KHR                             = 0x8247
	DEBUG_TYPE_DEPRECATED_BEHAVIOR                             = 0x824D
	DEBUG_TYPE_DEPRECATED_BEHAVIOR_ARB                         = 0x824D
	DEBUG_TYPE_DEPRECATED_BEHAVIOR_KHR                         = 0x824D
	DEBUG_TYPE_ERROR                                           = 0x824C
	DEBUG_TYPE_ERROR_ARB                                       = 0x824C
	DEBUG_TYPE_ERROR_KHR                                       = 0x824C
	DEBUG_TYPE_MARKER                                          = 0x8268
	DEBUG_TYPE_MARKER_KHR                                      = 0x8268
	DEBUG_TYPE_OTHER                                           = 0x8251
	DEBUG_TYPE_OTHER_ARB                                       = 0x8251
	DEBUG_TYPE_OTHER_KHR                                       = 0x8251
	DEBUG_TYPE_PERFORMANCE                                     = 0x8250
	DEBUG_TYPE_PERFORMANCE_ARB                                 = 0x8250
	DEBUG_TYPE_PERFORMANCE_KHR                                 = 0x8250
	DEBUG_TYPE_POP_GROUP                                       = 0x826A
	DEBUG_TYPE_POP_GROUP_KHR                                   = 0x826A
	DEBUG_TYPE_PORTABILITY                                     = 0x824F
	DEBUG_TYPE_PORTABILITY_ARB                                 = 0x824F
	DEBUG_TYPE_PORTABILITY_KHR                                 = 0x824F
	DEBUG_TYPE_PUSH_GROUP                                      = 0x8269
	DEBUG_TYPE_PUSH_GROUP_KHR                                  = 0x8269
	DEBUG_TYPE_UNDEFINED_BEHAVIOR                              = 0x824E
	DEBUG_TYPE_UNDEFINED_BEHAVIOR_ARB                          = 0x824E
	DEBUG_TYPE_UNDEFINED_BEHAVIOR_KHR                          = 0x824E
	DECAL                                                      = 0x2101
	DECR                                                       = 0x1E03
	DECR_WRAP                                                  = 0x8508
	DELETE_STATUS                                              = 0x8B80
	DEPTH                                                      = 0x1801
	DEPTH24_STENCIL8                                           = 0x88F0
	DEPTH32F_STENCIL8                                          = 0x8CAD
	DEPTH_ATTACHMENT                                           = 0x8D00
	DEPTH_BIAS                                                 = 0x0D1F
	DEPTH_BITS                                                 = 0x0D56
	DEPTH_BUFFER_BIT                                           = 0x00000100
	DEPTH_CLAMP                                                = 0x864F
	DEPTH_CLEAR_VALUE                                          = 0x0B73
	DEPTH_COMPONENT                                            = 0x1902
	DEPTH_COMPONENT16                                          = 0x81A5
	DEPTH_COMPONENT24                                          = 0x81A6
	DEPTH_COMPONENT32                                          = 0x81A7
	DEPTH_COMPONENT32F                                         = 0x8CAC
	DEPTH_COMPONENTS                                           = 0x8284
	DEPTH_FUNC                                                 = 0x0B74
	DEPTH_RANGE                                                = 0x0B70
	DEPTH_RENDERABLE                                           = 0x8287
	DEPTH_SCALE                                                = 0x0D1E
	DEPTH_STENCIL                                              = 0x84F9
	DEPTH_STENCIL_ATTACHMENT                                   = 0x821A
	DEPTH_STENCIL_TEXTURE_MODE                                 = 0x90EA
	DEPTH_TEST                                                 = 0x0B71
	DEPTH_TEXTURE_MODE                                         = 0x884B
	DEPTH_WRITEMASK                                            = 0x0B72
	DIFFUSE                                                    = 0x1201
	DISPATCH_INDIRECT_BUFFER                                   = 0x90EE
	DISPATCH_INDIRECT_BUFFER_BINDING                           = 0x90EF
	DITHER                                                     = 0x0BD0
	DOMAIN                                                     = 0x0A02
	DONT_CARE                                                  = 0x1100
	DOT3_RGB                                                   = 0x86AE
	DOT3_RGBA                                                  = 0x86AF
	DOUBLE                                                     = 0x140A
	DOUBLEBUFFER                                               = 0x0C32
	DOUBLE_MAT2                                                = 0x8F46
	DOUBLE_MAT2x3                                              = 0x8F49
	DOUBLE_MAT2x4                                              = 0x8F4A
	DOUBLE_MAT3                                                = 0x8F47
	DOUBLE_MAT3x2                                              = 0x8F4B
	DOUBLE_MAT3x4                                              = 0x8F4C
	DOUBLE_MAT4                                                = 0x8F48
	DOUBLE_MAT4x2                                              = 0x8F4D
	DOUBLE_MAT4x3                                              = 0x8F4E
	DOUBLE_VEC2                                                = 0x8FFC
	DOUBLE_VEC3                                                = 0x8FFD
	DOUBLE_VEC4                                                = 0x8FFE
	DRAW_BUFFER                                                = 0x0C01
	DRAW_BUFFER0                                               = 0x8825
	DRAW_BUFFER1                                               = 0x8826
	DRAW_BUFFER10                                              = 0x882F
	DRAW_BUFFER11                                              = 0x8830
	DRAW_BUFFER12                                              = 0x8831
	DRAW_BUFFER13                                              = 0x8832
	DRAW_BUFFER14                                              = 0x8833
	DRAW_BUFFER15                                              = 0x8834
	DRAW_BUFFER2                                               = 0x8827
	DRAW_BUFFER3                                               = 0x8828
	DRAW_BUFFER4                                               = 0x8829
	DRAW_BUFFER5                                               = 0x882A
	DRAW_BUFFER6                                               = 0x882B
	DRAW_BUFFER7                                               = 0x882C
	DRAW_BUFFER8                                               = 0x882D
	DRAW_BUFFER9                                               = 0x882E
	DRAW_FRAMEBUFFER                                           = 0x8CA9
	DRAW_FRAMEBUFFER_BINDING                                   = 0x8CA6
	DRAW_INDIRECT_BUFFER                                       = 0x8F3F
	DRAW_INDIRECT_BUFFER_BINDING                               = 0x8F43
	DRAW_PIXEL_TOKEN                                           = 0x0705
	DST_ALPHA                                                  = 0x0304
	DST_COLOR                                                  = 0x0306
	DYNAMIC_COPY                                               = 0x88EA
	DYNAMIC_DRAW                                               = 0x88E8
	DYNAMIC_READ                                               = 0x88E9
	DYNAMIC_STORAGE_BIT                                        = 0x0100
	EDGE_FLAG                                                  = 0x0B43
	EDGE_FLAG_ARRAY                                            = 0x8079
	EDGE_FLAG_ARRAY_BUFFER_BINDING                             = 0x889B
	EDGE_FLAG_ARRAY_POINTER                                    = 0x8093
	EDGE_FLAG_ARRAY_STRIDE                                     = 0x808C
	ELEMENT_ARRAY_BARRIER_BIT                                  = 0x00000002
	ELEMENT_ARRAY_BUFFER                                       = 0x8893
	ELEMENT_ARRAY_BUFFER_BINDING                               = 0x8895
	EMISSION                                                   = 0x1600
	ENABLE_BIT                                                 = 0x00002000
	EQUAL                                                      = 0x0202
	EQUIV                                                      = 0x1509
	EVAL_BIT                                                   = 0x00010000
	EXP                                                        = 0x0800
	EXP2                                                       = 0x0801
	EXTENSIONS                                                 = 0x1F03
	EYE_LINEAR                                                 = 0x2400
	EYE_PLANE                                                  = 0x2502
	FALSE                                                      = 0
	FASTEST                                                    = 0x1101
	FEEDBACK                                                   = 0x1C01
	FEEDBACK_BUFFER_POINTER                                    = 0x0DF0
	FEEDBACK_BUFFER_SIZE                                       = 0x0DF1
	FEEDBACK_BUFFER_TYPE                                       = 0x0DF2
	FILL                                                       = 0x1B02
	FILTER                                                     = 0x829A
	FIRST_VERTEX_CONVENTION                                    = 0x8E4D
	FIXED                                                      = 0x140C
	FIXED_ONLY                                                 = 0x891D
	FLAT                                                       = 0x1D00
	FLOAT                                                      = 0x1406
	FLOAT_32_UNSIGNED_INT_24_8_REV                             = 0x8DAD
	FLOAT_MAT2                                                 = 0x8B5A
	FLOAT_MAT2x3                                               = 0x8B65
	FLOAT_MAT2x4                                               = 0x8B66
	FLOAT_MAT3                                                 = 0x8B5B
	FLOAT_MAT3x2                                               = 0x8B67
	FLOAT_MAT3x4                                               = 0x8B68
	FLOAT_MAT4                                                 = 0x8B5C
	FLOAT_MAT4x2                                               = 0x8B69
	FLOAT_MAT4x3                                               = 0x8B6A
	FLOAT_VEC2                                                 = 0x8B50
	FLOAT_VEC3                                                 = 0x8B51
	FLOAT_VEC4                                                 = 0x8B52
	FOG                                                        = 0x0B60
	FOG_BIT                                                    = 0x00000080
	FOG_COLOR                                                  = 0x0B66
	FOG_COORD                                                  = 0x8451
	FOG_COORDINATE                                             = 0x8451
	FOG_COORDINATE_ARRAY                                       = 0x8457
	FOG_COORDINATE_ARRAY_BUFFER_BINDING                        = 0x889D
	FOG_COORDINATE_ARRAY_POINTER                               = 0x8456
	FOG_COORDINATE_ARRAY_STRIDE                                = 0x8455
	FOG_COORDINATE_ARRAY_TYPE                                  = 0x8454
	FOG_COORDINATE_SOURCE                                      = 0x8450
	FOG_COORD_ARRAY                                            = 0x8457
	FOG_COORD_ARRAY_BUFFER_BINDING                             = 0x889D
	FOG_COORD_ARRAY_POINTER                                    = 0x8456
	FOG_COORD_ARRAY_STRIDE                                     = 0x8455
	FOG_COORD_ARRAY_TYPE                                       = 0x8454
	FOG_COORD_SRC                                              = 0x8450
	FOG_DENSITY                                                = 0x0B62
	FOG_END                                                    = 0x0B64
	FOG_HINT                                                   = 0x0C54
	FOG_INDEX                                                  = 0x0B61
	FOG_MODE                                                   = 0x0B65
	FOG_START                                                  = 0x0B63
	FRACTIONAL_EVEN                                            = 0x8E7C
	FRACTIONAL_ODD                                             = 0x8E7B
	FRAGMENT_DEPTH                                             = 0x8452
	FRAGMENT_INTERPOLATION_OFFSET_BITS                         = 0x8E5D
	FRAGMENT_SHADER                                            = 0x8B30
	FRAGMENT_SHADER_BIT                                        = 0x00000002
	FRAGMENT_SHADER_DERIVATIVE_HINT                            = 0x8B8B
	FRAGMENT_SHADER_INVOCATIONS_ARB                            = 0x82F4
	FRAGMENT_SUBROUTINE                                        = 0x92EC
	FRAGMENT_SUBROUTINE_UNIFORM                                = 0x92F2
	FRAGMENT_TEXTURE                                           = 0x829F
	FRAMEBUFFER                                                = 0x8D40
	FRAMEBUFFER_ATTACHMENT_ALPHA_SIZE                          = 0x8215
	FRAMEBUFFER_ATTACHMENT_BLUE_SIZE                           = 0x8214
	FRAMEBUFFER_ATTACHMENT_COLOR_ENCODING                      = 0x8210
	FRAMEBUFFER_ATTACHMENT_COMPONENT_TYPE                      = 0x8211
	FRAMEBUFFER_ATTACHMENT_DEPTH_SIZE                          = 0x8216
	FRAMEBUFFER_ATTACHMENT_GREEN_SIZE                          = 0x8213
	FRAMEBUFFER_ATTACHMENT_LAYERED                             = 0x8DA7
	FRAMEBUFFER_ATTACHMENT_OBJECT_NAME                         = 0x8CD1
	FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE                         = 0x8CD0
	FRAMEBUFFER_ATTACHMENT_RED_SIZE                            = 0x8212
	FRAMEBUFFER_ATTACHMENT_STENCIL_SIZE                        = 0x8217
	FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE               = 0x8CD3
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LAYER                       = 0x8CD4
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL                       = 0x8CD2
	FRAMEBUFFER_BARRIER_BIT                                    = 0x00000400
	FRAMEBUFFER_BINDING                                        = 0x8CA6
	FRAMEBUFFER_BLEND                                          = 0x828B
	FRAMEBUFFER_COMPLETE                                       = 0x8CD5
	FRAMEBUFFER_DEFAULT                                        = 0x8218
	FRAMEBUFFER_DEFAULT_FIXED_SAMPLE_LOCATIONS                 = 0x9314
	FRAMEBUFFER_DEFAULT_HEIGHT                                 = 0x9311
	FRAMEBUFFER_DEFAULT_LAYERS                                 = 0x9312
	FRAMEBUFFER_DEFAULT_SAMPLES                                = 0x9313
	FRAMEBUFFER_DEFAULT_WIDTH                                  = 0x9310
	FRAMEBUFFER_INCOMPLETE_ATTACHMENT                          = 0x8CD6
	FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER                         = 0x8CDB
	FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS                       = 0x8DA8
	FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT                  = 0x8CD7
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE                         = 0x8D56
	FRAMEBUFFER_INCOMPLETE_READ_BUFFER                         = 0x8CDC
	FRAMEBUFFER_RENDERABLE                                     = 0x8289
	FRAMEBUFFER_RENDERABLE_LAYERED                             = 0x828A
	FRAMEBUFFER_SRGB                                           = 0x8DB9
	FRAMEBUFFER_UNDEFINED                                      = 0x8219
	FRAMEBUFFER_UNSUPPORTED                                    = 0x8CDD
	FRONT                                                      = 0x0404
	FRONT_AND_BACK                                             = 0x0408
	FRONT_FACE                                                 = 0x0B46
	FRONT_LEFT                                                 = 0x0400
	FRONT_RIGHT                                                = 0x0401
	FULL_SUPPORT                                               = 0x82B7
	FUNC_ADD                                                   = 0x8006
	FUNC_REVERSE_SUBTRACT                                      = 0x800B
	FUNC_SUBTRACT                                              = 0x800A
	GENERATE_MIPMAP                                            = 0x8191
	GENERATE_MIPMAP_HINT                                       = 0x8192
	GEOMETRY_INPUT_TYPE                                        = 0x8917
	GEOMETRY_OUTPUT_TYPE                                       = 0x8918
	GEOMETRY_SHADER                                            = 0x8DD9
	GEOMETRY_SHADER_BIT                                        = 0x00000004
	GEOMETRY_SHADER_INVOCATIONS                                = 0x887F
	GEOMETRY_SHADER_PRIMITIVES_EMITTED_ARB                     = 0x82F3
	GEOMETRY_SUBROUTINE                                        = 0x92EB
	GEOMETRY_SUBROUTINE_UNIFORM                                = 0x92F1
	GEOMETRY_TEXTURE                                           = 0x829E
	GEOMETRY_VERTICES_OUT                                      = 0x8916
	GEQUAL                                                     = 0x0206
	GET_TEXTURE_IMAGE_FORMAT                                   = 0x8291
	GET_TEXTURE_IMAGE_TYPE                                     = 0x8292
	GREATER                                                    = 0x0204
	GREEN                                                      = 0x1904
	GREEN_BIAS                                                 = 0x0D19
	GREEN_BITS                                                 = 0x0D53
	GREEN_INTEGER                                              = 0x8D95
	GREEN_SCALE                                                = 0x0D18
	GUILTY_CONTEXT_RESET                                       = 0x8253
	GUILTY_CONTEXT_RESET_ARB                                   = 0x8253
	GUILTY_CONTEXT_RESET_KHR                                   = 0x8253
	HALF_FLOAT                                                 = 0x140B
	HIGH_FLOAT                                                 = 0x8DF2
	HIGH_INT                                                   = 0x8DF5
	HINT_BIT                                                   = 0x00008000
	IMAGE_1D                                                   = 0x904C
	IMAGE_1D_ARRAY                                             = 0x9052
	IMAGE_2D                                                   = 0x904D
	IMAGE_2D_ARRAY                                             = 0x9053
	IMAGE_2D_MULTISAMPLE                                       = 0x9055
	IMAGE_2D_MULTISAMPLE_ARRAY                                 = 0x9056
	IMAGE_2D_RECT                                              = 0x904F
	IMAGE_3D                                                   = 0x904E
	IMAGE_BINDING_ACCESS                                       = 0x8F3E
	IMAGE_BINDING_FORMAT                                       = 0x906E
	IMAGE_BINDING_LAYER                                        = 0x8F3D
	IMAGE_BINDING_LAYERED                                      = 0x8F3C
	IMAGE_BINDING_LEVEL                                        = 0x8F3B
	IMAGE_BINDING_NAME                                         = 0x8F3A
	IMAGE_BUFFER                                               = 0x9051
	IMAGE_CLASS_10_10_10_2                                     = 0x82C3
	IMAGE_CLASS_11_11_10                                       = 0x82C2
	IMAGE_CLASS_1_X_16                                         = 0x82BE
	IMAGE_CLASS_1_X_32                                         = 0x82BB
	IMAGE_CLASS_1_X_8                                          = 0x82C1
	IMAGE_CLASS_2_X_16                                         = 0x82BD
	IMAGE_CLASS_2_X_32                                         = 0x82BA
	IMAGE_CLASS_2_X_8                                          = 0x82C0
	IMAGE_CLASS_4_X_16                                         = 0x82BC
	IMAGE_CLASS_4_X_32                                         = 0x82B9
	IMAGE_CLASS_4_X_8                                          = 0x82BF
	IMAGE_COMPATIBILITY_CLASS                                  = 0x82A8
	IMAGE_CUBE                                                 = 0x9050
	IMAGE_CUBE_MAP_ARRAY                                       = 0x9054
	IMAGE_FORMAT_COMPATIBILITY_BY_CLASS                        = 0x90C9
	IMAGE_FORMAT_COMPATIBILITY_BY_SIZE                         = 0x90C8
	IMAGE_FORMAT_COMPATIBILITY_TYPE                            = 0x90C7
	IMAGE_PIXEL_FORMAT                                         = 0x82A9
	IMAGE_PIXEL_TYPE                                           = 0x82AA
	IMAGE_TEXEL_SIZE                                           = 0x82A7
	IMPLEMENTATION_COLOR_READ_FORMAT                           = 0x8B9B
	IMPLEMENTATION_COLOR_READ_TYPE                             = 0x8B9A
	INCR                                                       = 0x1E02
	INCR_WRAP                                                  = 0x8507
	INDEX_ARRAY                                                = 0x8077
	INDEX_ARRAY_BUFFER_BINDING                                 = 0x8899
	INDEX_ARRAY_POINTER                                        = 0x8091
	INDEX_ARRAY_STRIDE                                         = 0x8086
	INDEX_ARRAY_TYPE                                           = 0x8085
	INDEX_BITS                                                 = 0x0D51
	INDEX_CLEAR_VALUE                                          = 0x0C20
	INDEX_LOGIC_OP                                             = 0x0BF1
	INDEX_MODE                                                 = 0x0C30
	INDEX_OFFSET                                               = 0x0D13
	INDEX_SHIFT                                                = 0x0D12
	INDEX_WRITEMASK                                            = 0x0C21
	INFO_LOG_LENGTH                                            = 0x8B84
	INNOCENT_CONTEXT_RESET                                     = 0x8254
	INNOCENT_CONTEXT_RESET_ARB                                 = 0x8254
	INNOCENT_CONTEXT_RESET_KHR                                 = 0x8254
	INT                                                        = 0x1404
	INTENSITY                                                  = 0x8049
	INTENSITY12                                                = 0x804C
	INTENSITY16                                                = 0x804D
	INTENSITY4                                                 = 0x804A
	INTENSITY8                                                 = 0x804B
	INTERLEAVED_ATTRIBS                                        = 0x8C8C
	INTERNALFORMAT_ALPHA_SIZE                                  = 0x8274
	INTERNALFORMAT_ALPHA_TYPE                                  = 0x827B
	INTERNALFORMAT_BLUE_SIZE                                   = 0x8273
	INTERNALFORMAT_BLUE_TYPE                                   = 0x827A
	INTERNALFORMAT_DEPTH_SIZE                                  = 0x8275
	INTERNALFORMAT_DEPTH_TYPE                                  = 0x827C
	INTERNALFORMAT_GREEN_SIZE                                  = 0x8272
	INTERNALFORMAT_GREEN_TYPE                                  = 0x8279
	INTERNALFORMAT_PREFERRED                                   = 0x8270
	INTERNALFORMAT_RED_SIZE                                    = 0x8271
	INTERNALFORMAT_RED_TYPE                                    = 0x8278
	INTERNALFORMAT_SHARED_SIZE                                 = 0x8277
	INTERNALFORMAT_STENCIL_SIZE                                = 0x8276
	INTERNALFORMAT_STENCIL_TYPE                                = 0x827D
	INTERNALFORMAT_SUPPORTED                                   = 0x826F
	INTERPOLATE                                                = 0x8575
	INT_2_10_10_10_REV                                         = 0x8D9F
	INT_IMAGE_1D                                               = 0x9057
	INT_IMAGE_1D_ARRAY                                         = 0x905D
	INT_IMAGE_2D                                               = 0x9058
	INT_IMAGE_2D_ARRAY                                         = 0x905E
	INT_IMAGE_2D_MULTISAMPLE                                   = 0x9060
	INT_IMAGE_2D_MULTISAMPLE_ARRAY                             = 0x9061
	INT_IMAGE_2D_RECT                                          = 0x905A
	INT_IMAGE_3D                                               = 0x9059
	INT_IMAGE_BUFFER                                           = 0x905C
	INT_IMAGE_CUBE                                             = 0x905B
	INT_IMAGE_CUBE_MAP_ARRAY                                   = 0x905F
	INT_SAMPLER_1D                                             = 0x8DC9
	INT_SAMPLER_1D_ARRAY                                       = 0x8DCE
	INT_SAMPLER_2D                                             = 0x8DCA
	INT_SAMPLER_2D_ARRAY                                       = 0x8DCF
	INT_SAMPLER_2D_MULTISAMPLE                                 = 0x9109
	INT_SAMPLER_2D_MULTISAMPLE_ARRAY                           = 0x910C
	INT_SAMPLER_2D_RECT                                        = 0x8DCD
	INT_SAMPLER_3D                                             = 0x8DCB
	INT_SAMPLER_BUFFER                                         = 0x8DD0
	INT_SAMPLER_CUBE                                           = 0x8DCC
	INT_SAMPLER_CUBE_MAP_ARRAY                                 = 0x900E
	INT_SAMPLER_CUBE_MAP_ARRAY_ARB                             = 0x900E
	INT_VEC2                                                   = 0x8B53
	INT_VEC3                                                   = 0x8B54
	INT_VEC4                                                   = 0x8B55
	INVALID_ENUM                                               = 0x0500
	INVALID_FRAMEBUFFER_OPERATION                              = 0x0506
	INVALID_INDEX                                              = 0xFFFFFFFF
	INVALID_OPERATION                                          = 0x0502
	INVALID_VALUE                                              = 0x0501
	INVERT                                                     = 0x150A
	ISOLINES                                                   = 0x8E7A
	IS_PER_PATCH                                               = 0x92E7
	IS_ROW_MAJOR                                               = 0x9300
	KEEP                                                       = 0x1E00
	LAST_VERTEX_CONVENTION                                     = 0x8E4E
	LAYER_PROVOKING_VERTEX                                     = 0x825E
	LEFT                                                       = 0x0406
	LEQUAL                                                     = 0x0203
	LESS                                                       = 0x0201
	LIGHT0                                                     = 0x4000
	LIGHT1                                                     = 0x4001
	LIGHT2                                                     = 0x4002
	LIGHT3                                                     = 0x4003
	LIGHT4                                                     = 0x4004
	LIGHT5                                                     = 0x4005
	LIGHT6                                                     = 0x4006
	LIGHT7                                                     = 0x4007
	LIGHTING                                                   = 0x0B50
	LIGHTING_BIT                                               = 0x00000040
	LIGHT_MODEL_AMBIENT                                        = 0x0B53
	LIGHT_MODEL_COLOR_CONTROL                                  = 0x81F8
	LIGHT_MODEL_LOCAL_VIEWER                                   = 0x0B51
	LIGHT_MODEL_TWO_SIDE                                       = 0x0B52
	LINE                                                       = 0x1B01
	LINEAR                                                     = 0x2601
	LINEAR_ATTENUATION                                         = 0x1208
	LINEAR_MIPMAP_LINEAR                                       = 0x2703
	LINEAR_MIPMAP_NEAREST                                      = 0x2701
	LINES                                                      = 0x0001
	LINES_ADJACENCY                                            = 0x000A
	LINE_BIT                                                   = 0x00000004
	LINE_LOOP                                                  = 0x0002
	LINE_RESET_TOKEN                                           = 0x0707
	LINE_SMOOTH                                                = 0x0B20
	LINE_SMOOTH_HINT                                           = 0x0C52
	LINE_STIPPLE                                               = 0x0B24
	LINE_STIPPLE_PATTERN                                       = 0x0B25
	LINE_STIPPLE_REPEAT                                        = 0x0B26
	LINE_STRIP                                                 = 0x0003
	LINE_STRIP_ADJACENCY                                       = 0x000B
	LINE_TOKEN                                                 = 0x0702
	LINE_WIDTH                                                 = 0x0B21
	LINE_WIDTH_GRANULARITY                                     = 0x0B23
	LINE_WIDTH_RANGE                                           = 0x0B22
	LINK_STATUS                                                = 0x8B82
	LIST_BASE                                                  = 0x0B32
	LIST_BIT                                                   = 0x00020000
	LIST_INDEX                                                 = 0x0B33
	LIST_MODE                                                  = 0x0B30
	LOAD                                                       = 0x0101
	LOCATION                                                   = 0x930E
	LOCATION_COMPONENT                                         = 0x934A
	LOCATION_INDEX                                             = 0x930F
	LOGIC_OP                                                   = 0x0BF1
	LOGIC_OP_MODE                                              = 0x0BF0
	LOSE_CONTEXT_ON_RESET                                      = 0x8252
	LOSE_CONTEXT_ON_RESET_ARB                                  = 0x8252
	LOSE_CONTEXT_ON_RESET_KHR                                  = 0x8252
	LOWER_LEFT                                                 = 0x8CA1
	LOW_FLOAT                                                  = 0x8DF0
	LOW_INT                                                    = 0x8DF3
	LUMINANCE                                                  = 0x1909
	LUMINANCE12                                                = 0x8041
	LUMINANCE12_ALPHA12                                        = 0x8047
	LUMINANCE12_ALPHA4                                         = 0x8046
	LUMINANCE16                                                = 0x8042
	LUMINANCE16_ALPHA16                                        = 0x8048
	LUMINANCE4                                                 = 0x803F
	LUMINANCE4_ALPHA4                                          = 0x8043
	LUMINANCE6_ALPHA2                                          = 0x8044
	LUMINANCE8                                                 = 0x8040
	LUMINANCE8_ALPHA8                                          = 0x8045
	LUMINANCE_ALPHA                                            = 0x190A
	MAJOR_VERSION                                              = 0x821B
	MANUAL_GENERATE_MIPMAP                                     = 0x8294
	MAP1_COLOR_4                                               = 0x0D90
	MAP1_GRID_DOMAIN                                           = 0x0DD0
	MAP1_GRID_SEGMENTS                                         = 0x0DD1
	MAP1_INDEX                                                 = 0x0D91
	MAP1_NORMAL                                                = 0x0D92
	MAP1_TEXTURE_COORD_1                                       = 0x0D93
	MAP1_TEXTURE_COORD_2                                       = 0x0D94
	MAP1_TEXTURE_COORD_3                                       = 0x0D95
	MAP1_TEXTURE_COORD_4                                       = 0x0D96
	MAP1_VERTEX_3                                              = 0x0D97
	MAP1_VERTEX_4                                              = 0x0D98
	MAP2_COLOR_4                                               = 0x0DB0
	MAP2_GRID_DOMAIN                                           = 0x0DD2
	MAP2_GRID_SEGMENTS                                         = 0x0DD3
	MAP2_INDEX                                                 = 0x0DB1
	MAP2_NORMAL                                                = 0x0DB2
	MAP2_TEXTURE_COORD_1                                       = 0x0DB3
	MAP2_TEXTURE_COORD_2                                       = 0x0DB4
	MAP2_TEXTURE_COORD_3                                       = 0x0DB5
	MAP2_TEXTURE_COORD_4                                       = 0x0DB6
	MAP2_VERTEX_3                                              = 0x0DB7
	MAP2_VERTEX_4                                              = 0x0DB8
	MAP_COHERENT_BIT                                           = 0x0080
	MAP_COLOR                                                  = 0x0D10
	MAP_FLUSH_EXPLICIT_BIT                                     = 0x0010
	MAP_INVALIDATE_BUFFER_BIT                                  = 0x0008
	MAP_INVALIDATE_RANGE_BIT                                   = 0x0004
	MAP_PERSISTENT_BIT                                         = 0x0040
	MAP_READ_BIT                                               = 0x0001
	MAP_STENCIL                                                = 0x0D11
	MAP_UNSYNCHRONIZED_BIT                                     = 0x0020
	MAP_WRITE_BIT                                              = 0x0002
	MATRIX_MODE                                                = 0x0BA0
	MATRIX_STRIDE                                              = 0x92FF
	MAX                                                        = 0x8008
	MAX_3D_TEXTURE_SIZE                                        = 0x8073
	MAX_ARRAY_TEXTURE_LAYERS                                   = 0x88FF
	MAX_ATOMIC_COUNTER_BUFFER_BINDINGS                         = 0x92DC
	MAX_ATOMIC_COUNTER_BUFFER_SIZE                             = 0x92D8
	MAX_ATTRIB_STACK_DEPTH                                     = 0x0D35
	MAX_CLIENT_ATTRIB_STACK_DEPTH                              = 0x0D3B
	MAX_CLIP_DISTANCES                                         = 0x0D32
	MAX_CLIP_PLANES                                            = 0x0D32
	MAX_COLOR_ATTACHMENTS                                      = 0x8CDF
	MAX_COLOR_TEXTURE_SAMPLES                                  = 0x910E
	MAX_COMBINED_ATOMIC_COUNTERS                               = 0x92D7
	MAX_COMBINED_ATOMIC_COUNTER_BUFFERS                        = 0x92D1
	MAX_COMBINED_CLIP_AND_CULL_DISTANCES                       = 0x82FA
	MAX_COMBINED_COMPUTE_UNIFORM_COMPONENTS                    = 0x8266
	MAX_COMBINED_DIMENSIONS                                    = 0x8282
	MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS                   = 0x8A33
	MAX_COMBINED_GEOMETRY_UNIFORM_COMPONENTS                   = 0x8A32
	MAX_COMBINED_IMAGE_UNIFORMS                                = 0x90CF
	MAX_COMBINED_IMAGE_UNITS_AND_FRAGMENT_OUTPUTS              = 0x8F39
	MAX_COMBINED_SHADER_OUTPUT_RESOURCES                       = 0x8F39
	MAX_COMBINED_SHADER_STORAGE_BLOCKS                         = 0x90DC
	MAX_COMBINED_TESS_CONTROL_UNIFORM_COMPONENTS               = 0x8E1E
	MAX_COMBINED_TESS_EVALUATION_UNIFORM_COMPONENTS            = 0x8E1F
	MAX_COMBINED_TEXTURE_IMAGE_UNITS                           = 0x8B4D
	MAX_COMBINED_UNIFORM_BLOCKS                                = 0x8A2E
	MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS                     = 0x8A31
	MAX_COMPUTE_ATOMIC_COUNTERS                                = 0x8265
	MAX_COMPUTE_ATOMIC_COUNTER_BUFFERS                         = 0x8264
	MAX_COMPUTE_FIXED_GROUP_INVOCATIONS_ARB                    = 0x90EB
	MAX_COMPUTE_FIXED_GROUP_SIZE_ARB                           = 0x91BF
	MAX_COMPUTE_IMAGE_UNIFORMS                                 = 0x91BD
	MAX_COMPUTE_SHADER_STORAGE_BLOCKS                          = 0x90DB
	MAX_COMPUTE_SHARED_MEMORY_SIZE                             = 0x8262
	MAX_COMPUTE_TEXTURE_IMAGE_UNITS                            = 0x91BC
	MAX_COMPUTE_UNIFORM_BLOCKS                                 = 0x91BB
	MAX_COMPUTE_UNIFORM_COMPONENTS                             = 0x8263
	MAX_COMPUTE_VARIABLE_GROUP_INVOCATIONS_ARB                 = 0x9344
	MAX_COMPUTE_VARIABLE_GROUP_SIZE_ARB                        = 0x9345
	MAX_COMPUTE_WORK_GROUP_COUNT                               = 0x91BE
	MAX_COMPUTE_WORK_GROUP_INVOCATIONS                         = 0x90EB
	MAX_COMPUTE_WORK_GROUP_SIZE                                = 0x91BF
	MAX_CUBE_MAP_TEXTURE_SIZE                                  = 0x851C
	MAX_CULL_DISTANCES                                         = 0x82F9
	MAX_DEBUG_GROUP_STACK_DEPTH                                = 0x826C
	MAX_DEBUG_GROUP_STACK_DEPTH_KHR                            = 0x826C
	MAX_DEBUG_LOGGED_MESSAGES                                  = 0x9144
	MAX_DEBUG_LOGGED_MESSAGES_ARB                              = 0x9144
	MAX_DEBUG_LOGGED_MESSAGES_KHR                              = 0x9144
	MAX_DEBUG_MESSAGE_LENGTH                                   = 0x9143
	MAX_DEBUG_MESSAGE_LENGTH_ARB                               = 0x9143
	MAX_DEBUG_MESSAGE_LENGTH_KHR                               = 0x9143
	MAX_DEPTH                                                  = 0x8280
	MAX_DEPTH_TEXTURE_SAMPLES                                  = 0x910F
	MAX_DRAW_BUFFERS                                           = 0x8824
	MAX_DUAL_SOURCE_DRAW_BUFFERS                               = 0x88FC
	MAX_ELEMENTS_INDICES                                       = 0x80E9
	MAX_ELEMENTS_VERTICES                                      = 0x80E8
	MAX_ELEMENT_INDEX                                          = 0x8D6B
	MAX_EVAL_ORDER                                             = 0x0D30
	MAX_FRAGMENT_ATOMIC_COUNTERS                               = 0x92D6
	MAX_FRAGMENT_ATOMIC_COUNTER_BUFFERS                        = 0x92D0
	MAX_FRAGMENT_IMAGE_UNIFORMS                                = 0x90CE
	MAX_FRAGMENT_INPUT_COMPONENTS                              = 0x9125
	MAX_FRAGMENT_INTERPOLATION_OFFSET                          = 0x8E5C
	MAX_FRAGMENT_SHADER_STORAGE_BLOCKS                         = 0x90DA
	MAX_FRAGMENT_UNIFORM_BLOCKS                                = 0x8A2D
	MAX_FRAGMENT_UNIFORM_COMPONENTS                            = 0x8B49
	MAX_FRAGMENT_UNIFORM_VECTORS                               = 0x8DFD
	MAX_FRAMEBUFFER_HEIGHT                                     = 0x9316
	MAX_FRAMEBUFFER_LAYERS                                     = 0x9317
	MAX_FRAMEBUFFER_SAMPLES                                    = 0x9318
	MAX_FRAMEBUFFER_WIDTH                                      = 0x9315
	MAX_GEOMETRY_ATOMIC_COUNTERS                               = 0x92D5
	MAX_GEOMETRY_ATOMIC_COUNTER_BUFFERS                        = 0x92CF
	MAX_GEOMETRY_IMAGE_UNIFORMS                                = 0x90CD
	MAX_GEOMETRY_INPUT_COMPONENTS                              = 0x9123
	MAX_GEOMETRY_OUTPUT_COMPONENTS                             = 0x9124
	MAX_GEOMETRY_OUTPUT_VERTICES                               = 0x8DE0
	MAX_GEOMETRY_SHADER_INVOCATIONS                            = 0x8E5A
	MAX_GEOMETRY_SHADER_STORAGE_BLOCKS                         = 0x90D7
	MAX_GEOMETRY_TEXTURE_IMAGE_UNITS                           = 0x8C29
	MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS                       = 0x8DE1
	MAX_GEOMETRY_UNIFORM_BLOCKS                                = 0x8A2C
	MAX_GEOMETRY_UNIFORM_COMPONENTS                            = 0x8DDF
	MAX_HEIGHT                                                 = 0x827F
	MAX_IMAGE_SAMPLES                                          = 0x906D
	MAX_IMAGE_UNITS                                            = 0x8F38
	MAX_INTEGER_SAMPLES                                        = 0x9110
	MAX_LABEL_LENGTH                                           = 0x82E8
	MAX_LABEL_LENGTH_KHR                                       = 0x82E8
	MAX_LAYERS                                                 = 0x8281
	MAX_LIGHTS                                                 = 0x0D31
	MAX_LIST_NESTING                                           = 0x0B31
	MAX_MODELVIEW_STACK_DEPTH                                  = 0x0D36
	MAX_NAME_LENGTH                                            = 0x92F6
	MAX_NAME_STACK_DEPTH                                       = 0x0D37
	MAX_NUM_ACTIVE_VARIABLES                                   = 0x92F7
	MAX_NUM_COMPATIBLE_SUBROUTINES                             = 0x92F8
	MAX_PATCH_VERTICES                                         = 0x8E7D
	MAX_PIXEL_MAP_TABLE                                        = 0x0D34
	MAX_PROGRAM_TEXEL_OFFSET                                   = 0x8905
	MAX_PROGRAM_TEXTURE_GATHER_COMPONENTS_ARB                  = 0x8F9F
	MAX_PROGRAM_TEXTURE_GATHER_OFFSET                          = 0x8E5F
	MAX_PROGRAM_TEXTURE_GATHER_OFFSET_ARB                      = 0x8E5F
	MAX_PROJECTION_STACK_DEPTH                                 = 0x0D38
	MAX_RECTANGLE_TEXTURE_SIZE                                 = 0x84F8
	MAX_RENDERBUFFER_SIZE                                      = 0x84E8
	MAX_SAMPLES                                                = 0x8D57
	MAX_SAMPLE_MASK_WORDS                                      = 0x8E59
	MAX_SERVER_WAIT_TIMEOUT                                    = 0x9111
	MAX_SHADER_STORAGE_BLOCK_SIZE                              = 0x90DE
	MAX_SHADER_STORAGE_BUFFER_BINDINGS                         = 0x90DD
	MAX_SPARSE_3D_TEXTURE_SIZE_ARB                             = 0x9199
	MAX_SPARSE_ARRAY_TEXTURE_LAYERS_ARB                        = 0x919A
	MAX_SPARSE_TEXTURE_SIZE_ARB                                = 0x9198
	MAX_SUBROUTINES                                            = 0x8DE7
	MAX_SUBROUTINE_UNIFORM_LOCATIONS                           = 0x8DE8
	MAX_TESS_CONTROL_ATOMIC_COUNTERS                           = 0x92D3
	MAX_TESS_CONTROL_ATOMIC_COUNTER_BUFFERS                    = 0x92CD
	MAX_TESS_CONTROL_IMAGE_UNIFORMS                            = 0x90CB
	MAX_TESS_CONTROL_INPUT_COMPONENTS                          = 0x886C
	MAX_TESS_CONTROL_OUTPUT_COMPONENTS                         = 0x8E83
	MAX_TESS_CONTROL_SHADER_STORAGE_BLOCKS                     = 0x90D8
	MAX_TESS_CONTROL_TEXTURE_IMAGE_UNITS                       = 0x8E81
	MAX_TESS_CONTROL_TOTAL_OUTPUT_COMPONENTS                   = 0x8E85
	MAX_TESS_CONTROL_UNIFORM_BLOCKS                            = 0x8E89
	MAX_TESS_CONTROL_UNIFORM_COMPONENTS                        = 0x8E7F
	MAX_TESS_EVALUATION_ATOMIC_COUNTERS                        = 0x92D4
	MAX_TESS_EVALUATION_ATOMIC_COUNTER_BUFFERS                 = 0x92CE
	MAX_TESS_EVALUATION_IMAGE_UNIFORMS                         = 0x90CC
	MAX_TESS_EVALUATION_INPUT_COMPONENTS                       = 0x886D
	MAX_TESS_EVALUATION_OUTPUT_COMPONENTS                      = 0x8E86
	MAX_TESS_EVALUATION_SHADER_STORAGE_BLOCKS                  = 0x90D9
	MAX_TESS_EVALUATION_TEXTURE_IMAGE_UNITS                    = 0x8E82
	MAX_TESS_EVALUATION_UNIFORM_BLOCKS                         = 0x8E8A
	MAX_TESS_EVALUATION_UNIFORM_COMPONENTS                     = 0x8E80
	MAX_TESS_GEN_LEVEL                                         = 0x8E7E
	MAX_TESS_PATCH_COMPONENTS                                  = 0x8E84
	MAX_TEXTURE_BUFFER_SIZE                                    = 0x8C2B
	MAX_TEXTURE_COORDS                                         = 0x8871
	MAX_TEXTURE_IMAGE_UNITS                                    = 0x8872
	MAX_TEXTURE_LOD_BIAS                                       = 0x84FD
	MAX_TEXTURE_SIZE                                           = 0x0D33
	MAX_TEXTURE_STACK_DEPTH                                    = 0x0D39
	MAX_TEXTURE_UNITS                                          = 0x84E2
	MAX_TRANSFORM_FEEDBACK_BUFFERS                             = 0x8E70
	MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS              = 0x8C8A
	MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS                    = 0x8C8B
	MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS                 = 0x8C80
	MAX_UNIFORM_BLOCK_SIZE                                     = 0x8A30
	MAX_UNIFORM_BUFFER_BINDINGS                                = 0x8A2F
	MAX_UNIFORM_LOCATIONS                                      = 0x826E
	MAX_VARYING_COMPONENTS                                     = 0x8B4B
	MAX_VARYING_FLOATS                                         = 0x8B4B
	MAX_VARYING_VECTORS                                        = 0x8DFC
	MAX_VERTEX_ATOMIC_COUNTERS                                 = 0x92D2
	MAX_VERTEX_ATOMIC_COUNTER_BUFFERS                          = 0x92CC
	MAX_VERTEX_ATTRIBS                                         = 0x8869
	MAX_VERTEX_ATTRIB_BINDINGS                                 = 0x82DA
	MAX_VERTEX_ATTRIB_RELATIVE_OFFSET                          = 0x82D9
	MAX_VERTEX_ATTRIB_STRIDE                                   = 0x82E5
	MAX_VERTEX_IMAGE_UNIFORMS                                  = 0x90CA
	MAX_VERTEX_OUTPUT_COMPONENTS                               = 0x9122
	MAX_VERTEX_SHADER_STORAGE_BLOCKS                           = 0x90D6
	MAX_VERTEX_STREAMS                                         = 0x8E71
	MAX_VERTEX_TEXTURE_IMAGE_UNITS                             = 0x8B4C
	MAX_VERTEX_UNIFORM_BLOCKS                                  = 0x8A2B
	MAX_VERTEX_UNIFORM_COMPONENTS                              = 0x8B4A
	MAX_VERTEX_UNIFORM_VECTORS                                 = 0x8DFB
	MAX_VIEWPORTS                                              = 0x825B
	MAX_VIEWPORT_DIMS                                          = 0x0D3A
	MAX_WIDTH                                                  = 0x827E
	MEDIUM_FLOAT                                               = 0x8DF1
	MEDIUM_INT                                                 = 0x8DF4
	MIN                                                        = 0x8007
	MINOR_VERSION                                              = 0x821C
	MIN_FRAGMENT_INTERPOLATION_OFFSET                          = 0x8E5B
	MIN_MAP_BUFFER_ALIGNMENT                                   = 0x90BC
	MIN_PROGRAM_TEXEL_OFFSET                                   = 0x8904
	MIN_PROGRAM_TEXTURE_GATHER_OFFSET                          = 0x8E5E
	MIN_PROGRAM_TEXTURE_GATHER_OFFSET_ARB                      = 0x8E5E
	MIN_SAMPLE_SHADING_VALUE                                   = 0x8C37
	MIN_SAMPLE_SHADING_VALUE_ARB                               = 0x8C37
	MIN_SPARSE_LEVEL_ARB                                       = 0x919B
	MIPMAP                                                     = 0x8293
	MIRRORED_REPEAT                                            = 0x8370
	MIRROR_CLAMP_TO_EDGE                                       = 0x8743
	MODELVIEW                                                  = 0x1700
	MODELVIEW_MATRIX                                           = 0x0BA6
	MODELVIEW_STACK_DEPTH                                      = 0x0BA3
	MODULATE                                                   = 0x2100
	MULT                                                       = 0x0103
	MULTISAMPLE                                                = 0x809D
	MULTISAMPLE_BIT                                            = 0x20000000
	N3F_V3F                                                    = 0x2A25
	NAMED_STRING_LENGTH_ARB                                    = 0x8DE9
	NAMED_STRING_TYPE_ARB                                      = 0x8DEA
	NAME_LENGTH                                                = 0x92F9
	NAME_STACK_DEPTH                                           = 0x0D70
	NAND                                                       = 0x150E
	NEAREST                                                    = 0x2600
	NEAREST_MIPMAP_LINEAR                                      = 0x2702
	NEAREST_MIPMAP_NEAREST                                     = 0x2700
	NEGATIVE_ONE_TO_ONE                                        = 0x935E
	NEVER                                                      = 0x0200
	NICEST                                                     = 0x1102
	NONE                                                       = 0
	NOOP                                                       = 0x1505
	NOR                                                        = 0x1508
	NORMALIZE                                                  = 0x0BA1
	NORMAL_ARRAY                                               = 0x8075
	NORMAL_ARRAY_BUFFER_BINDING                                = 0x8897
	NORMAL_ARRAY_POINTER                                       = 0x808F
	NORMAL_ARRAY_STRIDE                                        = 0x807F
	NORMAL_ARRAY_TYPE                                          = 0x807E
	NORMAL_MAP                                                 = 0x8511
	NOTEQUAL                                                   = 0x0205
	NO_ERROR                                                   = 0
	NO_RESET_NOTIFICATION                                      = 0x8261
	NO_RESET_NOTIFICATION_ARB                                  = 0x8261
	NO_RESET_NOTIFICATION_KHR                                  = 0x8261
	NUM_ACTIVE_VARIABLES                                       = 0x9304
	NUM_COMPATIBLE_SUBROUTINES                                 = 0x8E4A
	NUM_COMPRESSED_TEXTURE_FORMATS                             = 0x86A2
	NUM_EXTENSIONS                                             = 0x821D
	NUM_PROGRAM_BINARY_FORMATS                                 = 0x87FE
	NUM_SAMPLE_COUNTS                                          = 0x9380
	NUM_SHADER_BINARY_FORMATS                                  = 0x8DF9
	NUM_SHADING_LANGUAGE_VERSIONS                              = 0x82E9
	NUM_VIRTUAL_PAGE_SIZES_ARB                                 = 0x91A8
	OBJECT_LINEAR                                              = 0x2401
	OBJECT_PLANE                                               = 0x2501
	OBJECT_TYPE                                                = 0x9112
	OFFSET                                                     = 0x92FC
	ONE                                                        = 1
	ONE_MINUS_CONSTANT_ALPHA                                   = 0x8004
	ONE_MINUS_CONSTANT_COLOR                                   = 0x8002
	ONE_MINUS_DST_ALPHA                                        = 0x0305
	ONE_MINUS_DST_COLOR                                        = 0x0307
	ONE_MINUS_SRC1_ALPHA                                       = 0x88FB
	ONE_MINUS_SRC1_COLOR                                       = 0x88FA
	ONE_MINUS_SRC_ALPHA                                        = 0x0303
	ONE_MINUS_SRC_COLOR                                        = 0x0301
	OPERAND0_ALPHA                                             = 0x8598
	OPERAND0_RGB                                               = 0x8590
	OPERAND1_ALPHA                                             = 0x8599
	OPERAND1_RGB                                               = 0x8591
	OPERAND2_ALPHA                                             = 0x859A
	OPERAND2_RGB                                               = 0x8592
	OR                                                         = 0x1507
	ORDER                                                      = 0x0A01
	OR_INVERTED                                                = 0x150D
	OR_REVERSE                                                 = 0x150B
	OUT_OF_MEMORY                                              = 0x0505
	PACK_ALIGNMENT                                             = 0x0D05
	PACK_COMPRESSED_BLOCK_DEPTH                                = 0x912D
	PACK_COMPRESSED_BLOCK_HEIGHT                               = 0x912C
	PACK_COMPRESSED_BLOCK_SIZE                                 = 0x912E
	PACK_COMPRESSED_BLOCK_WIDTH                                = 0x912B
	PACK_IMAGE_HEIGHT                                          = 0x806C
	PACK_LSB_FIRST                                             = 0x0D01
	PACK_ROW_LENGTH                                            = 0x0D02
	PACK_SKIP_IMAGES                                           = 0x806B
	PACK_SKIP_PIXELS                                           = 0x0D04
	PACK_SKIP_ROWS                                             = 0x0D03
	PACK_SWAP_BYTES                                            = 0x0D00
	PARAMETER_BUFFER_ARB                                       = 0x80EE
	PARAMETER_BUFFER_BINDING_ARB                               = 0x80EF
	PASS_THROUGH_TOKEN                                         = 0x0700
	PATCHES                                                    = 0x000E
	PATCH_DEFAULT_INNER_LEVEL                                  = 0x8E73
	PATCH_DEFAULT_OUTER_LEVEL                                  = 0x8E74
	PATCH_VERTICES                                             = 0x8E72
	PERSPECTIVE_CORRECTION_HINT                                = 0x0C50
	PIXEL_BUFFER_BARRIER_BIT                                   = 0x00000080
	PIXEL_MAP_A_TO_A                                           = 0x0C79
	PIXEL_MAP_A_TO_A_SIZE                                      = 0x0CB9
	PIXEL_MAP_B_TO_B                                           = 0x0C78
	PIXEL_MAP_B_TO_B_SIZE                                      = 0x0CB8
	PIXEL_MAP_G_TO_G                                           = 0x0C77
	PIXEL_MAP_G_TO_G_SIZE                                      = 0x0CB7
	PIXEL_MAP_I_TO_A                                           = 0x0C75
	PIXEL_MAP_I_TO_A_SIZE                                      = 0x0CB5
	PIXEL_MAP_I_TO_B                                           = 0x0C74
	PIXEL_MAP_I_TO_B_SIZE                                      = 0x0CB4
	PIXEL_MAP_I_TO_G                                           = 0x0C73
	PIXEL_MAP_I_TO_G_SIZE                                      = 0x0CB3
	PIXEL_MAP_I_TO_I                                           = 0x0C70
	PIXEL_MAP_I_TO_I_SIZE                                      = 0x0CB0
	PIXEL_MAP_I_TO_R                                           = 0x0C72
	PIXEL_MAP_I_TO_R_SIZE                                      = 0x0CB2
	PIXEL_MAP_R_TO_R                                           = 0x0C76
	PIXEL_MAP_R_TO_R_SIZE                                      = 0x0CB6
	PIXEL_MAP_S_TO_S                                           = 0x0C71
	PIXEL_MAP_S_TO_S_SIZE                                      = 0x0CB1
	PIXEL_MODE_BIT                                             = 0x00000020
	PIXEL_PACK_BUFFER                                          = 0x88EB
	PIXEL_PACK_BUFFER_BINDING                                  = 0x88ED
	PIXEL_UNPACK_BUFFER                                        = 0x88EC
	PIXEL_UNPACK_BUFFER_BINDING                                = 0x88EF
	POINT                                                      = 0x1B00
	POINTS                                                     = 0x0000
	POINT_BIT                                                  = 0x00000002
	POINT_DISTANCE_ATTENUATION                                 = 0x8129
	POINT_FADE_THRESHOLD_SIZE                                  = 0x8128
	POINT_SIZE                                                 = 0x0B11
	POINT_SIZE_GRANULARITY                                     = 0x0B13
	POINT_SIZE_MAX                                             = 0x8127
	POINT_SIZE_MIN                                             = 0x8126
	POINT_SIZE_RANGE                                           = 0x0B12
	POINT_SMOOTH                                               = 0x0B10
	POINT_SMOOTH_HINT                                          = 0x0C51
	POINT_SPRITE                                               = 0x8861
	POINT_SPRITE_COORD_ORIGIN                                  = 0x8CA0
	POINT_TOKEN                                                = 0x0701
	POLYGON                                                    = 0x0009
	POLYGON_BIT                                                = 0x00000008
	POLYGON_MODE                                               = 0x0B40
	POLYGON_OFFSET_FACTOR                                      = 0x8038
	POLYGON_OFFSET_FILL                                        = 0x8037
	POLYGON_OFFSET_LINE                                        = 0x2A02
	POLYGON_OFFSET_POINT                                       = 0x2A01
	POLYGON_OFFSET_UNITS                                       = 0x2A00
	POLYGON_SMOOTH                                             = 0x0B41
	POLYGON_SMOOTH_HINT                                        = 0x0C53
	POLYGON_STIPPLE                                            = 0x0B42
	POLYGON_STIPPLE_BIT                                        = 0x00000010
	POLYGON_TOKEN                                              = 0x0703
	POSITION                                                   = 0x1203
	PREVIOUS                                                   = 0x8578
	PRIMARY_COLOR                                              = 0x8577
	PRIMITIVES_GENERATED                                       = 0x8C87
	PRIMITIVES_SUBMITTED_ARB                                   = 0x82EF
	PRIMITIVE_RESTART                                          = 0x8F9D
	PRIMITIVE_RESTART_FIXED_INDEX                              = 0x8D69
	PRIMITIVE_RESTART_FOR_PATCHES_SUPPORTED                    = 0x8221
	PRIMITIVE_RESTART_INDEX                                    = 0x8F9E
	PROGRAM                                                    = 0x82E2
	PROGRAM_BINARY_FORMATS                                     = 0x87FF
	PROGRAM_BINARY_LENGTH                                      = 0x8741
	PROGRAM_BINARY_RETRIEVABLE_HINT                            = 0x8257
	PROGRAM_INPUT                                              = 0x92E3
	PROGRAM_KHR                                                = 0x82E2
	PROGRAM_OUTPUT                                             = 0x92E4
	PROGRAM_PIPELINE                                           = 0x82E4
	PROGRAM_PIPELINE_BINDING                                   = 0x825A
	PROGRAM_POINT_SIZE                                         = 0x8642
	PROGRAM_SEPARABLE                                          = 0x8258
	PROJECTION                                                 = 0x1701
	PROJECTION_MATRIX                                          = 0x0BA7
	PROJECTION_STACK_DEPTH                                     = 0x0BA4
	PROVOKING_VERTEX                                           = 0x8E4F
	PROXY_TEXTURE_1D                                           = 0x8063
	PROXY_TEXTURE_1D_ARRAY                                     = 0x8C19
	PROXY_TEXTURE_2D                                           = 0x8064
	PROXY_TEXTURE_2D_ARRAY                                     = 0x8C1B
	PROXY_TEXTURE_2D_MULTISAMPLE                               = 0x9101
	PROXY_TEXTURE_2D_MULTISAMPLE_ARRAY                         = 0x9103
	PROXY_TEXTURE_3D                                           = 0x8070
	PROXY_TEXTURE_CUBE_MAP                                     = 0x851B
	PROXY_TEXTURE_CUBE_MAP_ARRAY                               = 0x900B
	PROXY_TEXTURE_CUBE_MAP_ARRAY_ARB                           = 0x900B
	PROXY_TEXTURE_RECTANGLE                                    = 0x84F7
	Q                                                          = 0x2003
	QUADRATIC_ATTENUATION                                      = 0x1209
	QUADS                                                      = 0x0007
	QUADS_FOLLOW_PROVOKING_VERTEX_CONVENTION                   = 0x8E4C
	QUAD_STRIP                                                 = 0x0008
	QUERY                                                      = 0x82E3
	QUERY_BUFFER                                               = 0x9192
	QUERY_BUFFER_BARRIER_BIT                                   = 0x00008000
	QUERY_BUFFER_BINDING                                       = 0x9193
	QUERY_BY_REGION_NO_WAIT                                    = 0x8E16
	QUERY_BY_REGION_NO_WAIT_INVERTED                           = 0x8E1A
	QUERY_BY_REGION_WAIT                                       = 0x8E15
	QUERY_BY_REGION_WAIT_INVERTED                              = 0x8E19
	QUERY_COUNTER_BITS                                         = 0x8864
	QUERY_KHR                                                  = 0x82E3
	QUERY_NO_WAIT                                              = 0x8E14
	QUERY_NO_WAIT_INVERTED                                     = 0x8E18
	QUERY_RESULT                                               = 0x8866
	QUERY_RESULT_AVAILABLE                                     = 0x8867
	QUERY_RESULT_NO_WAIT                                       = 0x9194
	QUERY_TARGET                                               = 0x82EA
	QUERY_WAIT                                                 = 0x8E13
	QUERY_WAIT_INVERTED                                        = 0x8E17
	R                                                          = 0x2002
	R11F_G11F_B10F                                             = 0x8C3A
	R16                                                        = 0x822A
	R16F                                                       = 0x822D
	R16I                                                       = 0x8233
	R16UI                                                      = 0x8234
	R16_SNORM                                                  = 0x8F98
	R32F                                                       = 0x822E
	R32I                                                       = 0x8235
	R32UI                                                      = 0x8236
	R3_G3_B2                                                   = 0x2A10
	R8                                                         = 0x8229
	R8I                                                        = 0x8231
	R8UI                                                       = 0x8232
	R8_SNORM                                                   = 0x8F94
	RASTERIZER_DISCARD                                         = 0x8C89
	READ_BUFFER                                                = 0x0C02
	READ_FRAMEBUFFER                                           = 0x8CA8
	READ_FRAMEBUFFER_BINDING                                   = 0x8CAA
	READ_ONLY                                                  = 0x88B8
	READ_PIXELS                                                = 0x828C
	READ_PIXELS_FORMAT                                         = 0x828D
	READ_PIXELS_TYPE                                           = 0x828E
	READ_WRITE                                                 = 0x88BA
	RED                                                        = 0x1903
	RED_BIAS                                                   = 0x0D15
	RED_BITS                                                   = 0x0D52
	RED_INTEGER                                                = 0x8D94
	RED_SCALE                                                  = 0x0D14
	REFERENCED_BY_COMPUTE_SHADER                               = 0x930B
	REFERENCED_BY_FRAGMENT_SHADER                              = 0x930A
	REFERENCED_BY_GEOMETRY_SHADER                              = 0x9309
	REFERENCED_BY_TESS_CONTROL_SHADER                          = 0x9307
	REFERENCED_BY_TESS_EVALUATION_SHADER                       = 0x9308
	REFERENCED_BY_VERTEX_SHADER                                = 0x9306
	REFLECTION_MAP                                             = 0x8512
	RENDER                                                     = 0x1C00
	RENDERBUFFER                                               = 0x8D41
	RENDERBUFFER_ALPHA_SIZE                                    = 0x8D53
	RENDERBUFFER_BINDING                                       = 0x8CA7
	RENDERBUFFER_BLUE_SIZE                                     = 0x8D52
	RENDERBUFFER_DEPTH_SIZE                                    = 0x8D54
	RENDERBUFFER_GREEN_SIZE                                    = 0x8D51
	RENDERBUFFER_HEIGHT                                        = 0x8D43
	RENDERBUFFER_INTERNAL_FORMAT                               = 0x8D44
	RENDERBUFFER_RED_SIZE                                      = 0x8D50
	RENDERBUFFER_SAMPLES                                       = 0x8CAB
	RENDERBUFFER_STENCIL_SIZE                                  = 0x8D55
	RENDERBUFFER_WIDTH                                         = 0x8D42
	RENDERER                                                   = 0x1F01
	RENDER_MODE                                                = 0x0C40
	REPEAT                                                     = 0x2901
	REPLACE                                                    = 0x1E01
	RESCALE_NORMAL                                             = 0x803A
	RESET_NOTIFICATION_STRATEGY                                = 0x8256
	RESET_NOTIFICATION_STRATEGY_ARB                            = 0x8256
	RESET_NOTIFICATION_STRATEGY_KHR                            = 0x8256
	RETURN                                                     = 0x0102
	RG                                                         = 0x8227
	RG16                                                       = 0x822C
	RG16F                                                      = 0x822F
	RG16I                                                      = 0x8239
	RG16UI                                                     = 0x823A
	RG16_SNORM                                                 = 0x8F99
	RG32F                                                      = 0x8230
	RG32I                                                      = 0x823B
	RG32UI                                                     = 0x823C
	RG8                                                        = 0x822B
	RG8I                                                       = 0x8237
	RG8UI                                                      = 0x8238
	RG8_SNORM                                                  = 0x8F95
	RGB                                                        = 0x1907
	RGB10                                                      = 0x8052
	RGB10_A2                                                   = 0x8059
	RGB10_A2UI                                                 = 0x906F
	RGB12                                                      = 0x8053
	RGB16                                                      = 0x8054
	RGB16F                                                     = 0x881B
	RGB16I                                                     = 0x8D89
	RGB16UI                                                    = 0x8D77
	RGB16_SNORM                                                = 0x8F9A
	RGB32F                                                     = 0x8815
	RGB32I                                                     = 0x8D83
	RGB32UI                                                    = 0x8D71
	RGB4                                                       = 0x804F
	RGB5                                                       = 0x8050
	RGB565                                                     = 0x8D62
	RGB5_A1                                                    = 0x8057
	RGB8                                                       = 0x8051
	RGB8I                                                      = 0x8D8F
	RGB8UI                                                     = 0x8D7D
	RGB8_SNORM                                                 = 0x8F96
	RGB9_E5                                                    = 0x8C3D
	RGBA                                                       = 0x1908
	RGBA12                                                     = 0x805A
	RGBA16                                                     = 0x805B
	RGBA16F                                                    = 0x881A
	RGBA16I                                                    = 0x8D88
	RGBA16UI                                                   = 0x8D76
	RGBA16_SNORM                                               = 0x8F9B
	RGBA2                                                      = 0x8055
	RGBA32F                                                    = 0x8814
	RGBA32I                                                    = 0x8D82
	RGBA32UI                                                   = 0x8D70
	RGBA4                                                      = 0x8056
	RGBA8                                                      = 0x8058
	RGBA8I                                                     = 0x8D8E
	RGBA8UI                                                    = 0x8D7C
	RGBA8_SNORM                                                = 0x8F97
	RGBA_INTEGER                                               = 0x8D99
	RGBA_MODE                                                  = 0x0C31
	RGB_INTEGER                                                = 0x8D98
	RGB_SCALE                                                  = 0x8573
	RG_INTEGER                                                 = 0x8228
	RIGHT                                                      = 0x0407
	S                                                          = 0x2000
	SAMPLER                                                    = 0x82E6
	SAMPLER_1D                                                 = 0x8B5D
	SAMPLER_1D_ARRAY                                           = 0x8DC0
	SAMPLER_1D_ARRAY_SHADOW                                    = 0x8DC3
	SAMPLER_1D_SHADOW                                          = 0x8B61
	SAMPLER_2D                                                 = 0x8B5E
	SAMPLER_2D_ARRAY                                           = 0x8DC1
	SAMPLER_2D_ARRAY_SHADOW                                    = 0x8DC4
	SAMPLER_2D_MULTISAMPLE                                     = 0x9108
	SAMPLER_2D_MULTISAMPLE_ARRAY                               = 0x910B
	SAMPLER_2D_RECT                                            = 0x8B63
	SAMPLER_2D_RECT_SHADOW                                     = 0x8B64
	SAMPLER_2D_SHADOW                                          = 0x8B62
	SAMPLER_3D                                                 = 0x8B5F
	SAMPLER_BINDING                                            = 0x8919
	SAMPLER_BUFFER                                             = 0x8DC2
	SAMPLER_CUBE                                               = 0x8B60
	SAMPLER_CUBE_MAP_ARRAY                                     = 0x900C
	SAMPLER_CUBE_MAP_ARRAY_ARB                                 = 0x900C
	SAMPLER_CUBE_MAP_ARRAY_SHADOW                              = 0x900D
	SAMPLER_CUBE_MAP_ARRAY_SHADOW_ARB                          = 0x900D
	SAMPLER_CUBE_SHADOW                                        = 0x8DC5
	SAMPLER_KHR                                                = 0x82E6
	SAMPLES                                                    = 0x80A9
	SAMPLES_PASSED                                             = 0x8914
	SAMPLE_ALPHA_TO_COVERAGE                                   = 0x809E
	SAMPLE_ALPHA_TO_ONE                                        = 0x809F
	SAMPLE_BUFFERS                                             = 0x80A8
	SAMPLE_COVERAGE                                            = 0x80A0
	SAMPLE_COVERAGE_INVERT                                     = 0x80AB
	SAMPLE_COVERAGE_VALUE                                      = 0x80AA
	SAMPLE_MASK                                                = 0x8E51
	SAMPLE_MASK_VALUE                                          = 0x8E52
	SAMPLE_POSITION                                            = 0x8E50
	SAMPLE_SHADING                                             = 0x8C36
	SAMPLE_SHADING_ARB                                         = 0x8C36
	SCISSOR_BIT                                                = 0x00080000
	SCISSOR_BOX                                                = 0x0C10
	SCISSOR_TEST                                               = 0x0C11
	SECONDARY_COLOR_ARRAY                                      = 0x845E
	SECONDARY_COLOR_ARRAY_BUFFER_BINDING                       = 0x889C
	SECONDARY_COLOR_ARRAY_POINTER                              = 0x845D
	SECONDARY_COLOR_ARRAY_SIZE                                 = 0x845A
	SECONDARY_COLOR_ARRAY_STRIDE                               = 0x845C
	SECONDARY_COLOR_ARRAY_TYPE                                 = 0x845B
	SELECT                                                     = 0x1C02
	SELECTION_BUFFER_POINTER                                   = 0x0DF3
	SELECTION_BUFFER_SIZE                                      = 0x0DF4
	SEPARATE_ATTRIBS                                           = 0x8C8D
	SEPARATE_SPECULAR_COLOR                                    = 0x81FA
	SET                                                        = 0x150F
	SHADER                                                     = 0x82E1
	SHADER_BINARY_FORMATS                                      = 0x8DF8
	SHADER_COMPILER                                            = 0x8DFA
	SHADER_IMAGE_ACCESS_BARRIER_BIT                            = 0x00000020
	SHADER_IMAGE_ATOMIC                                        = 0x82A6
	SHADER_IMAGE_LOAD                                          = 0x82A4
	SHADER_IMAGE_STORE                                         = 0x82A5
	SHADER_INCLUDE_ARB                                         = 0x8DAE
	SHADER_KHR                                                 = 0x82E1
	SHADER_SOURCE_LENGTH                                       = 0x8B88
	SHADER_STORAGE_BARRIER_BIT                                 = 0x00002000
	SHADER_STORAGE_BLOCK                                       = 0x92E6
	SHADER_STORAGE_BUFFER                                      = 0x90D2
	SHADER_STORAGE_BUFFER_BINDING                              = 0x90D3
	SHADER_STORAGE_BUFFER_OFFSET_ALIGNMENT                     = 0x90DF
	SHADER_STORAGE_BUFFER_SIZE                                 = 0x90D5
	SHADER_STORAGE_BUFFER_START                                = 0x90D4
	SHADER_TYPE                                                = 0x8B4F
	SHADE_MODEL                                                = 0x0B54
	SHADING_LANGUAGE_VERSION                                   = 0x8B8C
	SHININESS                                                  = 0x1601
	SHORT                                                      = 0x1402
	SIGNALED                                                   = 0x9119
	SIGNED_NORMALIZED                                          = 0x8F9C
	SIMULTANEOUS_TEXTURE_AND_DEPTH_TEST                        = 0x82AC
	SIMULTANEOUS_TEXTURE_AND_DEPTH_WRITE                       = 0x82AE
	SIMULTANEOUS_TEXTURE_AND_STENCIL_TEST                      = 0x82AD
	SIMULTANEOUS_TEXTURE_AND_STENCIL_WRITE                     = 0x82AF
	SINGLE_COLOR                                               = 0x81F9
	SLUMINANCE                                                 = 0x8C46
	SLUMINANCE8                                                = 0x8C47
	SLUMINANCE8_ALPHA8                                         = 0x8C45
	SLUMINANCE_ALPHA                                           = 0x8C44
	SMOOTH                                                     = 0x1D01
	SMOOTH_LINE_WIDTH_GRANULARITY                              = 0x0B23
	SMOOTH_LINE_WIDTH_RANGE                                    = 0x0B22
	SMOOTH_POINT_SIZE_GRANULARITY                              = 0x0B13
	SMOOTH_POINT_SIZE_RANGE                                    = 0x0B12
	SOURCE0_ALPHA                                              = 0x8588
	SOURCE0_RGB                                                = 0x8580
	SOURCE1_ALPHA                                              = 0x8589
	SOURCE1_RGB                                                = 0x8581
	SOURCE2_ALPHA                                              = 0x858A
	SOURCE2_RGB                                                = 0x8582
	SPARSE_BUFFER_PAGE_SIZE_ARB                                = 0x82F8
	SPARSE_STORAGE_BIT_ARB                                     = 0x0400
	SPARSE_TEXTURE_FULL_ARRAY_CUBE_MIPMAPS_ARB                 = 0x91A9
	SPECULAR                                                   = 0x1202
	SPHERE_MAP                                                 = 0x2402
	SPOT_CUTOFF                                                = 0x1206
	SPOT_DIRECTION                                             = 0x1204
	SPOT_EXPONENT                                              = 0x1205
	SRC0_ALPHA                                                 = 0x8588
	SRC0_RGB                                                   = 0x8580
	SRC1_ALPHA                                                 = 0x8589
	SRC1_COLOR                                                 = 0x88F9
	SRC1_RGB                                                   = 0x8581
	SRC2_ALPHA                                                 = 0x858A
	SRC2_RGB                                                   = 0x8582
	SRC_ALPHA                                                  = 0x0302
	SRC_ALPHA_SATURATE                                         = 0x0308
	SRC_COLOR                                                  = 0x0300
	SRGB                                                       = 0x8C40
	SRGB8                                                      = 0x8C41
	SRGB8_ALPHA8                                               = 0x8C43
	SRGB_ALPHA                                                 = 0x8C42
	SRGB_DECODE_ARB                                            = 0x8299
	SRGB_READ                                                  = 0x8297
	SRGB_WRITE                                                 = 0x8298
	STACK_OVERFLOW                                             = 0x0503
	STACK_OVERFLOW_KHR                                         = 0x0503
	STACK_UNDERFLOW                                            = 0x0504
	STACK_UNDERFLOW_KHR                                        = 0x0504
	STATIC_COPY                                                = 0x88E6
	STATIC_DRAW                                                = 0x88E4
	STATIC_READ                                                = 0x88E5
	STENCIL                                                    = 0x1802
	STENCIL_ATTACHMENT                                         = 0x8D20
	STENCIL_BACK_FAIL                                          = 0x8801
	STENCIL_BACK_FUNC                                          = 0x8800
	STENCIL_BACK_PASS_DEPTH_FAIL                               = 0x8802
	STENCIL_BACK_PASS_DEPTH_PASS                               = 0x8803
	STENCIL_BACK_REF                                           = 0x8CA3
	STENCIL_BACK_VALUE_MASK                                    = 0x8CA4
	STENCIL_BACK_WRITEMASK                                     = 0x8CA5
	STENCIL_BITS                                               = 0x0D57
	STENCIL_BUFFER_BIT                                         = 0x00000400
	STENCIL_CLEAR_VALUE                                        = 0x0B91
	STENCIL_COMPONENTS                                         = 0x8285
	STENCIL_FAIL                                               = 0x0B94
	STENCIL_FUNC                                               = 0x0B92
	STENCIL_INDEX                                              = 0x1901
	STENCIL_INDEX1                                             = 0x8D46
	STENCIL_INDEX16                                            = 0x8D49
	STENCIL_INDEX4                                             = 0x8D47
	STENCIL_INDEX8                                             = 0x8D48
	STENCIL_PASS_DEPTH_FAIL                                    = 0x0B95
	STENCIL_PASS_DEPTH_PASS                                    = 0x0B96
	STENCIL_REF                                                = 0x0B97
	STENCIL_RENDERABLE                                         = 0x8288
	STENCIL_TEST                                               = 0x0B90
	STENCIL_VALUE_MASK                                         = 0x0B93
	STENCIL_WRITEMASK                                          = 0x0B98
	STEREO                                                     = 0x0C33
	STREAM_COPY                                                = 0x88E2
	STREAM_DRAW                                                = 0x88E0
	STREAM_READ                                                = 0x88E1
	SUBPIXEL_BITS                                              = 0x0D50
	SUBTRACT                                                   = 0x84E7
	SYNC_CL_EVENT_ARB                                          = 0x8240
	SYNC_CL_EVENT_COMPLETE_ARB                                 = 0x8241
	SYNC_CONDITION                                             = 0x9113
	SYNC_FENCE                                                 = 0x9116
	SYNC_FLAGS                                                 = 0x9115
	SYNC_FLUSH_COMMANDS_BIT                                    = 0x00000001
	SYNC_GPU_COMMANDS_COMPLETE                                 = 0x9117
	SYNC_STATUS                                                = 0x9114
	T                                                          = 0x2001
	T2F_C3F_V3F                                                = 0x2A2A
	T2F_C4F_N3F_V3F                                            = 0x2A2C
	T2F_C4UB_V3F                                               = 0x2A29
	T2F_N3F_V3F                                                = 0x2A2B
	T2F_V3F                                                    = 0x2A27
	T4F_C4F_N3F_V4F                                            = 0x2A2D
	T4F_V4F                                                    = 0x2A28
	TESS_CONTROL_OUTPUT_VERTICES                               = 0x8E75
	TESS_CONTROL_SHADER                                        = 0x8E88
	TESS_CONTROL_SHADER_BIT                                    = 0x00000008
	TESS_CONTROL_SHADER_PATCHES_ARB                            = 0x82F1
	TESS_CONTROL_SUBROUTINE                                    = 0x92E9
	TESS_CONTROL_SUBROUTINE_UNIFORM                            = 0x92EF
	TESS_CONTROL_TEXTURE                                       = 0x829C
	TESS_EVALUATION_SHADER                                     = 0x8E87
	TESS_EVALUATION_SHADER_BIT                                 = 0x00000010
	TESS_EVALUATION_SHADER_INVOCATIONS_ARB                     = 0x82F2
	TESS_EVALUATION_SUBROUTINE                                 = 0x92EA
	TESS_EVALUATION_SUBROUTINE_UNIFORM                         = 0x92F0
	TESS_EVALUATION_TEXTURE                                    = 0x829D
	TESS_GEN_MODE                                              = 0x8E76
	TESS_GEN_POINT_MODE                                        = 0x8E79
	TESS_GEN_SPACING                                           = 0x8E77
	TESS_GEN_VERTEX_ORDER                                      = 0x8E78
	TEXTURE                                                    = 0x1702
	TEXTURE0                                                   = 0x84C0
	TEXTURE1                                                   = 0x84C1
	TEXTURE10                                                  = 0x84CA
	TEXTURE11                                                  = 0x84CB
	TEXTURE12                                                  = 0x84CC
	TEXTURE13                                                  = 0x84CD
	TEXTURE14                                                  = 0x84CE
	TEXTURE15                                                  = 0x84CF
	TEXTURE16                                                  = 0x84D0
	TEXTURE17                                                  = 0x84D1
	TEXTURE18                                                  = 0x84D2
	TEXTURE19                                                  = 0x84D3
	TEXTURE2                                                   = 0x84C2
	TEXTURE20                                                  = 0x84D4
	TEXTURE21                                                  = 0x84D5
	TEXTURE22                                                  = 0x84D6
	TEXTURE23                                                  = 0x84D7
	TEXTURE24                                                  = 0x84D8
	TEXTURE25                                                  = 0x84D9
	TEXTURE26                                                  = 0x84DA
	TEXTURE27                                                  = 0x84DB
	TEXTURE28                                                  = 0x84DC
	TEXTURE29                                                  = 0x84DD
	TEXTURE3                                                   = 0x84C3
	TEXTURE30                                                  = 0x84DE
	TEXTURE31                                                  = 0x84DF
	TEXTURE4                                                   = 0x84C4
	TEXTURE5                                                   = 0x84C5
	TEXTURE6                                                   = 0x84C6
	TEXTURE7                                                   = 0x84C7
	TEXTURE8                                                   = 0x84C8
	TEXTURE9                                                   = 0x84C9
	TEXTURE_1D                                                 = 0x0DE0
	TEXTURE_1D_ARRAY                                           = 0x8C18
	TEXTURE_2D                                                 = 0x0DE1
	TEXTURE_2D_ARRAY                                           = 0x8C1A
	TEXTURE_2D_MULTISAMPLE                                     = 0x9100
	TEXTURE_2D_MULTISAMPLE_ARRAY                               = 0x9102
	TEXTURE_3D                                                 = 0x806F
	TEXTURE_ALPHA_SIZE                                         = 0x805F
	TEXTURE_ALPHA_TYPE                                         = 0x8C13
	TEXTURE_BASE_LEVEL                                         = 0x813C
	TEXTURE_BINDING                                            = 0x82EB
	TEXTURE_BINDING_1D                                         = 0x8068
	TEXTURE_BINDING_1D_ARRAY                                   = 0x8C1C
	TEXTURE_BINDING_2D                                         = 0x8069
	TEXTURE_BINDING_2D_ARRAY                                   = 0x8C1D
	TEXTURE_BINDING_2D_MULTISAMPLE                             = 0x9104
	TEXTURE_BINDING_2D_MULTISAMPLE_ARRAY                       = 0x9105
	TEXTURE_BINDING_3D                                         = 0x806A
	TEXTURE_BINDING_BUFFER                                     = 0x8C2C
	TEXTURE_BINDING_CUBE_MAP                                   = 0x8514
	TEXTURE_BINDING_CUBE_MAP_ARRAY                             = 0x900A
	TEXTURE_BINDING_CUBE_MAP_ARRAY_ARB                         = 0x900A
	TEXTURE_BINDING_RECTANGLE                                  = 0x84F6
	TEXTURE_BIT                                                = 0x00040000
	TEXTURE_BLUE_SIZE                                          = 0x805E
	TEXTURE_BLUE_TYPE                                          = 0x8C12
	TEXTURE_BORDER                                             = 0x1005
	TEXTURE_BORDER_COLOR                                       = 0x1004
	TEXTURE_BUFFER                                             = 0x8C2A
	TEXTURE_BUFFER_BINDING                                     = 0x8C2A
	TEXTURE_BUFFER_DATA_STORE_BINDING                          = 0x8C2D
	TEXTURE_BUFFER_OFFSET                                      = 0x919D
	TEXTURE_BUFFER_OFFSET_ALIGNMENT                            = 0x919F
	TEXTURE_BUFFER_SIZE                                        = 0x919E
	TEXTURE_COMPARE_FUNC                                       = 0x884D
	TEXTURE_COMPARE_MODE                                       = 0x884C
	TEXTURE_COMPONENTS                                         = 0x1003
	TEXTURE_COMPRESSED                                         = 0x86A1
	TEXTURE_COMPRESSED_BLOCK_HEIGHT                            = 0x82B2
	TEXTURE_COMPRESSED_BLOCK_SIZE                              = 0x82B3
	TEXTURE_COMPRESSED_BLOCK_WIDTH                             = 0x82B1
	TEXTURE_COMPRESSED_IMAGE_SIZE                              = 0x86A0
	TEXTURE_COMPRESSION_HINT                                   = 0x84EF
	TEXTURE_COORD_ARRAY                                        = 0x8078
	TEXTURE_COORD_ARRAY_BUFFER_BINDING                         = 0x889A
	TEXTURE_COORD_ARRAY_POINTER                                = 0x8092
	TEXTURE_COORD_ARRAY_SIZE                                   = 0x8088
	TEXTURE_COORD_ARRAY_STRIDE                                 = 0x808A
	TEXTURE_COORD_ARRAY_TYPE                                   = 0x8089
	TEXTURE_CUBE_MAP                                           = 0x8513
	TEXTURE_CUBE_MAP_ARRAY                                     = 0x9009
	TEXTURE_CUBE_MAP_ARRAY_ARB                                 = 0x9009
	TEXTURE_CUBE_MAP_NEGATIVE_X                                = 0x8516
	TEXTURE_CUBE_MAP_NEGATIVE_Y                                = 0x8518
	TEXTURE_CUBE_MAP_NEGATIVE_Z                                = 0x851A
	TEXTURE_CUBE_MAP_POSITIVE_X                                = 0x8515
	TEXTURE_CUBE_MAP_POSITIVE_Y                                = 0x8517
	TEXTURE_CUBE_MAP_POSITIVE_Z                                = 0x8519
	TEXTURE_CUBE_MAP_SEAMLESS                                  = 0x884F
	TEXTURE_DEPTH                                              = 0x8071
	TEXTURE_DEPTH_SIZE                                         = 0x884A
	TEXTURE_DEPTH_TYPE                                         = 0x8C16
	TEXTURE_ENV                                                = 0x2300
	TEXTURE_ENV_COLOR                                          = 0x2201
	TEXTURE_ENV_MODE                                           = 0x2200
	TEXTURE_FETCH_BARRIER_BIT                                  = 0x00000008
	TEXTURE_FILTER_CONTROL                                     = 0x8500
	TEXTURE_FIXED_SAMPLE_LOCATIONS                             = 0x9107
	TEXTURE_GATHER                                             = 0x82A2
	TEXTURE_GATHER_SHADOW                                      = 0x82A3
	TEXTURE_GEN_MODE                                           = 0x2500
	TEXTURE_GEN_Q                                              = 0x0C63
	TEXTURE_GEN_R                                              = 0x0C62
	TEXTURE_GEN_S                                              = 0x0C60
	TEXTURE_GEN_T                                              = 0x0C61
	TEXTURE_GREEN_SIZE                                         = 0x805D
	TEXTURE_GREEN_TYPE                                         = 0x8C11
	TEXTURE_HEIGHT                                             = 0x1001
	TEXTURE_IMAGE_FORMAT                                       = 0x828F
	TEXTURE_IMAGE_TYPE                                         = 0x8290
	TEXTURE_IMMUTABLE_FORMAT                                   = 0x912F
	TEXTURE_IMMUTABLE_LEVELS                                   = 0x82DF
	TEXTURE_INTENSITY_SIZE                                     = 0x8061
	TEXTURE_INTERNAL_FORMAT                                    = 0x1003
	TEXTURE_LOD_BIAS                                           = 0x8501
	TEXTURE_LUMINANCE_SIZE                                     = 0x8060
	TEXTURE_MAG_FILTER                                         = 0x2800
	TEXTURE_MATRIX                                             = 0x0BA8
	TEXTURE_MAX_LEVEL                                          = 0x813D
	TEXTURE_MAX_LOD                                            = 0x813B
	TEXTURE_MIN_FILTER                                         = 0x2801
	TEXTURE_MIN_LOD                                            = 0x813A
	TEXTURE_PRIORITY                                           = 0x8066
	TEXTURE_RECTANGLE                                          = 0x84F5
	TEXTURE_RED_SIZE                                           = 0x805C
	TEXTURE_RED_TYPE                                           = 0x8C10
	TEXTURE_RESIDENT                                           = 0x8067
	TEXTURE_SAMPLES                                            = 0x9106
	TEXTURE_SHADOW                                             = 0x82A1
	TEXTURE_SHARED_SIZE                                        = 0x8C3F
	TEXTURE_SPARSE_ARB                                         = 0x91A6
	TEXTURE_STACK_DEPTH                                        = 0x0BA5
	TEXTURE_STENCIL_SIZE                                       = 0x88F1
	TEXTURE_SWIZZLE_A                                          = 0x8E45
	TEXTURE_SWIZZLE_B                                          = 0x8E44
	TEXTURE_SWIZZLE_G                                          = 0x8E43
	TEXTURE_SWIZZLE_R                                          = 0x8E42
	TEXTURE_SWIZZLE_RGBA                                       = 0x8E46
	TEXTURE_TARGET                                             = 0x1006
	TEXTURE_UPDATE_BARRIER_BIT                                 = 0x00000100
	TEXTURE_VIEW                                               = 0x82B5
	TEXTURE_VIEW_MIN_LAYER                                     = 0x82DD
	TEXTURE_VIEW_MIN_LEVEL                                     = 0x82DB
	TEXTURE_VIEW_NUM_LAYERS                                    = 0x82DE
	TEXTURE_VIEW_NUM_LEVELS                                    = 0x82DC
	TEXTURE_WIDTH                                              = 0x1000
	TEXTURE_WRAP_R                                             = 0x8072
	TEXTURE_WRAP_S                                             = 0x2802
	TEXTURE_WRAP_T                                             = 0x2803
	TIMEOUT_EXPIRED                                            = 0x911B
	TIMEOUT_IGNORED                                            = 0xFFFFFFFFFFFFFFFF
	TIMESTAMP                                                  = 0x8E28
	TIME_ELAPSED                                               = 0x88BF
	TOP_LEVEL_ARRAY_SIZE                                       = 0x930C
	TOP_LEVEL_ARRAY_STRIDE                                     = 0x930D
	TRANSFORM_BIT                                              = 0x00001000
	TRANSFORM_FEEDBACK                                         = 0x8E22
	TRANSFORM_FEEDBACK_ACTIVE                                  = 0x8E24
	TRANSFORM_FEEDBACK_BARRIER_BIT                             = 0x00000800
	TRANSFORM_FEEDBACK_BINDING                                 = 0x8E25
	TRANSFORM_FEEDBACK_BUFFER                                  = 0x8C8E
	TRANSFORM_FEEDBACK_BUFFER_ACTIVE                           = 0x8E24
	TRANSFORM_FEEDBACK_BUFFER_BINDING                          = 0x8C8F
	TRANSFORM_FEEDBACK_BUFFER_INDEX                            = 0x934B
	TRANSFORM_FEEDBACK_BUFFER_MODE                             = 0x8C7F
	TRANSFORM_FEEDBACK_BUFFER_PAUSED                           = 0x8E23
	TRANSFORM_FEEDBACK_BUFFER_SIZE                             = 0x8C85
	TRANSFORM_FEEDBACK_BUFFER_START                            = 0x8C84
	TRANSFORM_FEEDBACK_BUFFER_STRIDE                           = 0x934C
	TRANSFORM_FEEDBACK_OVERFLOW_ARB                            = 0x82EC
	TRANSFORM_FEEDBACK_PAUSED                                  = 0x8E23
	TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN                      = 0x8C88
	TRANSFORM_FEEDBACK_STREAM_OVERFLOW_ARB                     = 0x82ED
	TRANSFORM_FEEDBACK_VARYING                                 = 0x92F4
	TRANSFORM_FEEDBACK_VARYINGS                                = 0x8C83
	TRANSFORM_FEEDBACK_VARYING_MAX_LENGTH                      = 0x8C76
	TRANSPOSE_COLOR_MATRIX                                     = 0x84E6
	TRANSPOSE_MODELVIEW_MATRIX                                 = 0x84E3
	TRANSPOSE_PROJECTION_MATRIX                                = 0x84E4
	TRANSPOSE_TEXTURE_MATRIX                                   = 0x84E5
	TRIANGLES                                                  = 0x0004
	TRIANGLES_ADJACENCY                                        = 0x000C
	TRIANGLE_FAN                                               = 0x0006
	TRIANGLE_STRIP                                             = 0x0005
	TRIANGLE_STRIP_ADJACENCY                                   = 0x000D
	TRUE                                                       = 1
	TYPE                                                       = 0x92FA
	UNDEFINED_VERTEX                                           = 0x8260
	UNIFORM                                                    = 0x92E1
	UNIFORM_ARRAY_STRIDE                                       = 0x8A3C
	UNIFORM_ATOMIC_COUNTER_BUFFER_INDEX                        = 0x92DA
	UNIFORM_BARRIER_BIT                                        = 0x00000004
	UNIFORM_BLOCK                                              = 0x92E2
	UNIFORM_BLOCK_ACTIVE_UNIFORMS                              = 0x8A42
	UNIFORM_BLOCK_ACTIVE_UNIFORM_INDICES                       = 0x8A43
	UNIFORM_BLOCK_BINDING                                      = 0x8A3F
	UNIFORM_BLOCK_DATA_SIZE                                    = 0x8A40
	UNIFORM_BLOCK_INDEX                                        = 0x8A3A
	UNIFORM_BLOCK_NAME_LENGTH                                  = 0x8A41
	UNIFORM_BLOCK_REFERENCED_BY_COMPUTE_SHADER                 = 0x90EC
	UNIFORM_BLOCK_REFERENCED_BY_FRAGMENT_SHADER                = 0x8A46
	UNIFORM_BLOCK_REFERENCED_BY_GEOMETRY_SHADER                = 0x8A45
	UNIFORM_BLOCK_REFERENCED_BY_TESS_CONTROL_SHADER            = 0x84F0
	UNIFORM_BLOCK_REFERENCED_BY_TESS_EVALUATION_SHADER         = 0x84F1
	UNIFORM_BLOCK_REFERENCED_BY_VERTEX_SHADER                  = 0x8A44
	UNIFORM_BUFFER                                             = 0x8A11
	UNIFORM_BUFFER_BINDING                                     = 0x8A28
	UNIFORM_BUFFER_OFFSET_ALIGNMENT                            = 0x8A34
	UNIFORM_BUFFER_SIZE                                        = 0x8A2A
	UNIFORM_BUFFER_START                                       = 0x8A29
	UNIFORM_IS_ROW_MAJOR                                       = 0x8A3E
	UNIFORM_MATRIX_STRIDE                                      = 0x8A3D
	UNIFORM_NAME_LENGTH                                        = 0x8A39
	UNIFORM_OFFSET                                             = 0x8A3B
	UNIFORM_SIZE                                               = 0x8A38
	UNIFORM_TYPE                                               = 0x8A37
	UNKNOWN_CONTEXT_RESET                                      = 0x8255
	UNKNOWN_CONTEXT_RESET_ARB                                  = 0x8255
	UNKNOWN_CONTEXT_RESET_KHR                                  = 0x8255
	UNPACK_ALIGNMENT                                           = 0x0CF5
	UNPACK_COMPRESSED_BLOCK_DEPTH                              = 0x9129
	UNPACK_COMPRESSED_BLOCK_HEIGHT                             = 0x9128
	UNPACK_COMPRESSED_BLOCK_SIZE                               = 0x912A
	UNPACK_COMPRESSED_BLOCK_WIDTH                              = 0x9127
	UNPACK_IMAGE_HEIGHT                                        = 0x806E
	UNPACK_LSB_FIRST                                           = 0x0CF1
	UNPACK_ROW_LENGTH                                          = 0x0CF2
	UNPACK_SKIP_IMAGES                                         = 0x806D
	UNPACK_SKIP_PIXELS                                         = 0x0CF4
	UNPACK_SKIP_ROWS                                           = 0x0CF3
	UNPACK_SWAP_BYTES                                          = 0x0CF0
	UNSIGNALED                                                 = 0x9118
	UNSIGNED_BYTE                                              = 0x1401
	UNSIGNED_BYTE_2_3_3_REV                                    = 0x8362
	UNSIGNED_BYTE_3_3_2                                        = 0x8032
	UNSIGNED_INT                                               = 0x1405
	UNSIGNED_INT64_ARB                                         = 0x140F
	UNSIGNED_INT_10F_11F_11F_REV                               = 0x8C3B
	UNSIGNED_INT_10_10_10_2                                    = 0x8036
	UNSIGNED_INT_24_8                                          = 0x84FA
	UNSIGNED_INT_2_10_10_10_REV                                = 0x8368
	UNSIGNED_INT_5_9_9_9_REV                                   = 0x8C3E
	UNSIGNED_INT_8_8_8_8                                       = 0x8035
	UNSIGNED_INT_8_8_8_8_REV                                   = 0x8367
	UNSIGNED_INT_ATOMIC_COUNTER                                = 0x92DB
	UNSIGNED_INT_IMAGE_1D                                      = 0x9062
	UNSIGNED_INT_IMAGE_1D_ARRAY                                = 0x9068
	UNSIGNED_INT_IMAGE_2D                                      = 0x9063
	UNSIGNED_INT_IMAGE_2D_ARRAY                                = 0x9069
	UNSIGNED_INT_IMAGE_2D_MULTISAMPLE                          = 0x906B
	UNSIGNED_INT_IMAGE_2D_MULTISAMPLE_ARRAY                    = 0x906C
	UNSIGNED_INT_IMAGE_2D_RECT                                 = 0x9065
	UNSIGNED_INT_IMAGE_3D                                      = 0x9064
	UNSIGNED_INT_IMAGE_BUFFER                                  = 0x9067
	UNSIGNED_INT_IMAGE_CUBE                                    = 0x9066
	UNSIGNED_INT_IMAGE_CUBE_MAP_ARRAY                          = 0x906A
	UNSIGNED_INT_SAMPLER_1D                                    = 0x8DD1
	UNSIGNED_INT_SAMPLER_1D_ARRAY                              = 0x8DD6
	UNSIGNED_INT_SAMPLER_2D                                    = 0x8DD2
	UNSIGNED_INT_SAMPLER_2D_ARRAY                              = 0x8DD7
	UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE                        = 0x910A
	UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE_ARRAY                  = 0x910D
	UNSIGNED_INT_SAMPLER_2D_RECT                               = 0x8DD5
	UNSIGNED_INT_SAMPLER_3D                                    = 0x8DD3
	UNSIGNED_INT_SAMPLER_BUFFER                                = 0x8DD8
	UNSIGNED_INT_SAMPLER_CUBE                                  = 0x8DD4
	UNSIGNED_INT_SAMPLER_CUBE_MAP_ARRAY                        = 0x900F
	UNSIGNED_INT_SAMPLER_CUBE_MAP_ARRAY_ARB                    = 0x900F
	UNSIGNED_INT_VEC2                                          = 0x8DC6
	UNSIGNED_INT_VEC3                                          = 0x8DC7
	UNSIGNED_INT_VEC4                                          = 0x8DC8
	UNSIGNED_NORMALIZED                                        = 0x8C17
	UNSIGNED_SHORT                                             = 0x1403
	UNSIGNED_SHORT_1_5_5_5_REV                                 = 0x8366
	UNSIGNED_SHORT_4_4_4_4                                     = 0x8033
	UNSIGNED_SHORT_4_4_4_4_REV                                 = 0x8365
	UNSIGNED_SHORT_5_5_5_1                                     = 0x8034
	UNSIGNED_SHORT_5_6_5                                       = 0x8363
	UNSIGNED_SHORT_5_6_5_REV                                   = 0x8364
	UPPER_LEFT                                                 = 0x8CA2
	V2F                                                        = 0x2A20
	V3F                                                        = 0x2A21
	VALIDATE_STATUS                                            = 0x8B83
	VENDOR                                                     = 0x1F00
	VERSION                                                    = 0x1F02
	VERTEX_ARRAY                                               = 0x8074
	VERTEX_ARRAY_BINDING                                       = 0x85B5
	VERTEX_ARRAY_BUFFER_BINDING                                = 0x8896
	VERTEX_ARRAY_KHR                                           = 0x8074
	VERTEX_ARRAY_POINTER                                       = 0x808E
	VERTEX_ARRAY_SIZE                                          = 0x807A
	VERTEX_ARRAY_STRIDE                                        = 0x807C
	VERTEX_ARRAY_TYPE                                          = 0x807B
	VERTEX_ATTRIB_ARRAY_BARRIER_BIT                            = 0x00000001
	VERTEX_ATTRIB_ARRAY_BUFFER_BINDING                         = 0x889F
	VERTEX_ATTRIB_ARRAY_DIVISOR                                = 0x88FE
	VERTEX_ATTRIB_ARRAY_ENABLED                                = 0x8622
	VERTEX_ATTRIB_ARRAY_INTEGER                                = 0x88FD
	VERTEX_ATTRIB_ARRAY_LONG                                   = 0x874E
	VERTEX_ATTRIB_ARRAY_NORMALIZED                             = 0x886A
	VERTEX_ATTRIB_ARRAY_POINTER                                = 0x8645
	VERTEX_ATTRIB_ARRAY_SIZE                                   = 0x8623
	VERTEX_ATTRIB_ARRAY_STRIDE                                 = 0x8624
	VERTEX_ATTRIB_ARRAY_TYPE                                   = 0x8625
	VERTEX_ATTRIB_BINDING                                      = 0x82D4
	VERTEX_ATTRIB_RELATIVE_OFFSET                              = 0x82D5
	VERTEX_BINDING_BUFFER                                      = 0x8F4F
	VERTEX_BINDING_DIVISOR                                     = 0x82D6
	VERTEX_BINDING_OFFSET                                      = 0x82D7
	VERTEX_BINDING_STRIDE                                      = 0x82D8
	VERTEX_PROGRAM_POINT_SIZE                                  = 0x8642
	VERTEX_PROGRAM_TWO_SIDE                                    = 0x8643
	VERTEX_SHADER                                              = 0x8B31
	VERTEX_SHADER_BIT                                          = 0x00000001
	VERTEX_SHADER_INVOCATIONS_ARB                              = 0x82F0
	VERTEX_SUBROUTINE                                          = 0x92E8
	VERTEX_SUBROUTINE_UNIFORM                                  = 0x92EE
	VERTEX_TEXTURE                                             = 0x829B
	VERTICES_SUBMITTED_ARB                                     = 0x82EE
	VIEWPORT                                                   = 0x0BA2
	VIEWPORT_BIT                                               = 0x00000800
	VIEWPORT_BOUNDS_RANGE                                      = 0x825D
	VIEWPORT_INDEX_PROVOKING_VERTEX                            = 0x825F
	VIEWPORT_SUBPIXEL_BITS                                     = 0x825C
	VIEW_CLASS_128_BITS                                        = 0x82C4
	VIEW_CLASS_16_BITS                                         = 0x82CA
	VIEW_CLASS_24_BITS                                         = 0x82C9
	VIEW_CLASS_32_BITS                                         = 0x82C8
	VIEW_CLASS_48_BITS                                         = 0x82C7
	VIEW_CLASS_64_BITS                                         = 0x82C6
	VIEW_CLASS_8_BITS                                          = 0x82CB
	VIEW_CLASS_96_BITS                                         = 0x82C5
	VIEW_CLASS_BPTC_FLOAT                                      = 0x82D3
	VIEW_CLASS_BPTC_UNORM                                      = 0x82D2
	VIEW_CLASS_RGTC1_RED                                       = 0x82D0
	VIEW_CLASS_RGTC2_RG                                        = 0x82D1
	VIEW_CLASS_S3TC_DXT1_RGB                                   = 0x82CC
	VIEW_CLASS_S3TC_DXT1_RGBA                                  = 0x82CD
	VIEW_CLASS_S3TC_DXT3_RGBA                                  = 0x82CE
	VIEW_CLASS_S3TC_DXT5_RGBA                                  = 0x82CF
	VIEW_COMPATIBILITY_CLASS                                   = 0x82B6
	VIRTUAL_PAGE_SIZE_INDEX_ARB                                = 0x91A7
	VIRTUAL_PAGE_SIZE_X_ARB                                    = 0x9195
	VIRTUAL_PAGE_SIZE_Y_ARB                                    = 0x9196
	VIRTUAL_PAGE_SIZE_Z_ARB                                    = 0x9197
	WAIT_FAILED                                                = 0x911D
	WEIGHT_ARRAY_BUFFER_BINDING                                = 0x889E
	WRITE_ONLY                                                 = 0x88B9
	XOR                                                        = 0x1506
	ZERO                                                       = 0
	ZERO_TO_ONE                                                = 0x935F
	ZOOM_X                                                     = 0x0D16
	ZOOM_Y                                                     = 0x0D17
)

var (
	gpAccum                                       C.GPACCUM
	gpActiveShaderProgram                         C.GPACTIVESHADERPROGRAM
	gpActiveTexture                               C.GPACTIVETEXTURE
	gpAlphaFunc                                   C.GPALPHAFUNC
	gpAreTexturesResident                         C.GPARETEXTURESRESIDENT
	gpArrayElement                                C.GPARRAYELEMENT
	gpAttachShader                                C.GPATTACHSHADER
	gpBegin                                       C.GPBEGIN
	gpBeginConditionalRender                      C.GPBEGINCONDITIONALRENDER
	gpBeginQuery                                  C.GPBEGINQUERY
	gpBeginQueryIndexed                           C.GPBEGINQUERYINDEXED
	gpBeginTransformFeedback                      C.GPBEGINTRANSFORMFEEDBACK
	gpBindAttribLocation                          C.GPBINDATTRIBLOCATION
	gpBindBuffer                                  C.GPBINDBUFFER
	gpBindBufferBase                              C.GPBINDBUFFERBASE
	gpBindBufferRange                             C.GPBINDBUFFERRANGE
	gpBindBuffersBase                             C.GPBINDBUFFERSBASE
	gpBindBuffersRange                            C.GPBINDBUFFERSRANGE
	gpBindFragDataLocation                        C.GPBINDFRAGDATALOCATION
	gpBindFragDataLocationIndexed                 C.GPBINDFRAGDATALOCATIONINDEXED
	gpBindFramebuffer                             C.GPBINDFRAMEBUFFER
	gpBindImageTexture                            C.GPBINDIMAGETEXTURE
	gpBindImageTextures                           C.GPBINDIMAGETEXTURES
	gpBindProgramPipeline                         C.GPBINDPROGRAMPIPELINE
	gpBindRenderbuffer                            C.GPBINDRENDERBUFFER
	gpBindSampler                                 C.GPBINDSAMPLER
	gpBindSamplers                                C.GPBINDSAMPLERS
	gpBindTexture                                 C.GPBINDTEXTURE
	gpBindTextureUnit                             C.GPBINDTEXTUREUNIT
	gpBindTextures                                C.GPBINDTEXTURES
	gpBindTransformFeedback                       C.GPBINDTRANSFORMFEEDBACK
	gpBindVertexArray                             C.GPBINDVERTEXARRAY
	gpBindVertexBuffer                            C.GPBINDVERTEXBUFFER
	gpBindVertexBuffers                           C.GPBINDVERTEXBUFFERS
	gpBitmap                                      C.GPBITMAP
	gpBlendColor                                  C.GPBLENDCOLOR
	gpBlendEquation                               C.GPBLENDEQUATION
	gpBlendEquationSeparate                       C.GPBLENDEQUATIONSEPARATE
	gpBlendEquationSeparatei                      C.GPBLENDEQUATIONSEPARATEI
	gpBlendEquationSeparateiARB                   C.GPBLENDEQUATIONSEPARATEIARB
	gpBlendEquationi                              C.GPBLENDEQUATIONI
	gpBlendEquationiARB                           C.GPBLENDEQUATIONIARB
	gpBlendFunc                                   C.GPBLENDFUNC
	gpBlendFuncSeparate                           C.GPBLENDFUNCSEPARATE
	gpBlendFuncSeparatei                          C.GPBLENDFUNCSEPARATEI
	gpBlendFuncSeparateiARB                       C.GPBLENDFUNCSEPARATEIARB
	gpBlendFunci                                  C.GPBLENDFUNCI
	gpBlendFunciARB                               C.GPBLENDFUNCIARB
	gpBlitFramebuffer                             C.GPBLITFRAMEBUFFER
	gpBlitNamedFramebuffer                        C.GPBLITNAMEDFRAMEBUFFER
	gpBufferData                                  C.GPBUFFERDATA
	gpBufferPageCommitmentARB                     C.GPBUFFERPAGECOMMITMENTARB
	gpBufferStorage                               C.GPBUFFERSTORAGE
	gpBufferSubData                               C.GPBUFFERSUBDATA
	gpCallList                                    C.GPCALLLIST
	gpCallLists                                   C.GPCALLLISTS
	gpCheckFramebufferStatus                      C.GPCHECKFRAMEBUFFERSTATUS
	gpCheckNamedFramebufferStatus                 C.GPCHECKNAMEDFRAMEBUFFERSTATUS
	gpClampColor                                  C.GPCLAMPCOLOR
	gpClear                                       C.GPCLEAR
	gpClearAccum                                  C.GPCLEARACCUM
	gpClearBufferData                             C.GPCLEARBUFFERDATA
	gpClearBufferSubData                          C.GPCLEARBUFFERSUBDATA
	gpClearBufferfi                               C.GPCLEARBUFFERFI
	gpClearBufferfv                               C.GPCLEARBUFFERFV
	gpClearBufferiv                               C.GPCLEARBUFFERIV
	gpClearBufferuiv                              C.GPCLEARBUFFERUIV
	gpClearColor                                  C.GPCLEARCOLOR
	gpClearDepth                                  C.GPCLEARDEPTH
	gpClearDepthf                                 C.GPCLEARDEPTHF
	gpClearIndex                                  C.GPCLEARINDEX
	gpClearNamedBufferData                        C.GPCLEARNAMEDBUFFERDATA
	gpClearNamedBufferSubData                     C.GPCLEARNAMEDBUFFERSUBDATA
	gpClearNamedFramebufferfi                     C.GPCLEARNAMEDFRAMEBUFFERFI
	gpClearNamedFramebufferfv                     C.GPCLEARNAMEDFRAMEBUFFERFV
	gpClearNamedFramebufferiv                     C.GPCLEARNAMEDFRAMEBUFFERIV
	gpClearNamedFramebufferuiv                    C.GPCLEARNAMEDFRAMEBUFFERUIV
	gpClearStencil                                C.GPCLEARSTENCIL
	gpClearTexImage                               C.GPCLEARTEXIMAGE
	gpClearTexSubImage                            C.GPCLEARTEXSUBIMAGE
	gpClientActiveTexture                         C.GPCLIENTACTIVETEXTURE
	gpClientWaitSync                              C.GPCLIENTWAITSYNC
	gpClipControl                                 C.GPCLIPCONTROL
	gpClipPlane                                   C.GPCLIPPLANE
	gpColor3b                                     C.GPCOLOR3B
	gpColor3bv                                    C.GPCOLOR3BV
	gpColor3d                                     C.GPCOLOR3D
	gpColor3dv                                    C.GPCOLOR3DV
	gpColor3f                                     C.GPCOLOR3F
	gpColor3fv                                    C.GPCOLOR3FV
	gpColor3i                                     C.GPCOLOR3I
	gpColor3iv                                    C.GPCOLOR3IV
	gpColor3s                                     C.GPCOLOR3S
	gpColor3sv                                    C.GPCOLOR3SV
	gpColor3ub                                    C.GPCOLOR3UB
	gpColor3ubv                                   C.GPCOLOR3UBV
	gpColor3ui                                    C.GPCOLOR3UI
	gpColor3uiv                                   C.GPCOLOR3UIV
	gpColor3us                                    C.GPCOLOR3US
	gpColor3usv                                   C.GPCOLOR3USV
	gpColor4b                                     C.GPCOLOR4B
	gpColor4bv                                    C.GPCOLOR4BV
	gpColor4d                                     C.GPCOLOR4D
	gpColor4dv                                    C.GPCOLOR4DV
	gpColor4f                                     C.GPCOLOR4F
	gpColor4fv                                    C.GPCOLOR4FV
	gpColor4i                                     C.GPCOLOR4I
	gpColor4iv                                    C.GPCOLOR4IV
	gpColor4s                                     C.GPCOLOR4S
	gpColor4sv                                    C.GPCOLOR4SV
	gpColor4ub                                    C.GPCOLOR4UB
	gpColor4ubv                                   C.GPCOLOR4UBV
	gpColor4ui                                    C.GPCOLOR4UI
	gpColor4uiv                                   C.GPCOLOR4UIV
	gpColor4us                                    C.GPCOLOR4US
	gpColor4usv                                   C.GPCOLOR4USV
	gpColorMask                                   C.GPCOLORMASK
	gpColorMaski                                  C.GPCOLORMASKI
	gpColorMaterial                               C.GPCOLORMATERIAL
	gpColorPointer                                C.GPCOLORPOINTER
	gpCompileShader                               C.GPCOMPILESHADER
	gpCompileShaderIncludeARB                     C.GPCOMPILESHADERINCLUDEARB
	gpCompressedTexImage1D                        C.GPCOMPRESSEDTEXIMAGE1D
	gpCompressedTexImage2D                        C.GPCOMPRESSEDTEXIMAGE2D
	gpCompressedTexImage3D                        C.GPCOMPRESSEDTEXIMAGE3D
	gpCompressedTexSubImage1D                     C.GPCOMPRESSEDTEXSUBIMAGE1D
	gpCompressedTexSubImage2D                     C.GPCOMPRESSEDTEXSUBIMAGE2D
	gpCompressedTexSubImage3D                     C.GPCOMPRESSEDTEXSUBIMAGE3D
	gpCompressedTextureSubImage1D                 C.GPCOMPRESSEDTEXTURESUBIMAGE1D
	gpCompressedTextureSubImage2D                 C.GPCOMPRESSEDTEXTURESUBIMAGE2D
	gpCompressedTextureSubImage3D                 C.GPCOMPRESSEDTEXTURESUBIMAGE3D
	gpCopyBufferSubData                           C.GPCOPYBUFFERSUBDATA
	gpCopyImageSubData                            C.GPCOPYIMAGESUBDATA
	gpCopyNamedBufferSubData                      C.GPCOPYNAMEDBUFFERSUBDATA
	gpCopyPixels                                  C.GPCOPYPIXELS
	gpCopyTexImage1D                              C.GPCOPYTEXIMAGE1D
	gpCopyTexImage2D                              C.GPCOPYTEXIMAGE2D
	gpCopyTexSubImage1D                           C.GPCOPYTEXSUBIMAGE1D
	gpCopyTexSubImage2D                           C.GPCOPYTEXSUBIMAGE2D
	gpCopyTexSubImage3D                           C.GPCOPYTEXSUBIMAGE3D
	gpCopyTextureSubImage1D                       C.GPCOPYTEXTURESUBIMAGE1D
	gpCopyTextureSubImage2D                       C.GPCOPYTEXTURESUBIMAGE2D
	gpCopyTextureSubImage3D                       C.GPCOPYTEXTURESUBIMAGE3D
	gpCreateBuffers                               C.GPCREATEBUFFERS
	gpCreateFramebuffers                          C.GPCREATEFRAMEBUFFERS
	gpCreateProgram                               C.GPCREATEPROGRAM
	gpCreateProgramPipelines                      C.GPCREATEPROGRAMPIPELINES
	gpCreateQueries                               C.GPCREATEQUERIES
	gpCreateRenderbuffers                         C.GPCREATERENDERBUFFERS
	gpCreateSamplers                              C.GPCREATESAMPLERS
	gpCreateShader                                C.GPCREATESHADER
	gpCreateShaderProgramv                        C.GPCREATESHADERPROGRAMV
	gpCreateSyncFromCLeventARB                    C.GPCREATESYNCFROMCLEVENTARB
	gpCreateTextures                              C.GPCREATETEXTURES
	gpCreateTransformFeedbacks                    C.GPCREATETRANSFORMFEEDBACKS
	gpCreateVertexArrays                          C.GPCREATEVERTEXARRAYS
	gpCullFace                                    C.GPCULLFACE
	gpDebugMessageCallback                        C.GPDEBUGMESSAGECALLBACK
	gpDebugMessageCallbackARB                     C.GPDEBUGMESSAGECALLBACKARB
	gpDebugMessageCallbackKHR                     C.GPDEBUGMESSAGECALLBACKKHR
	gpDebugMessageControl                         C.GPDEBUGMESSAGECONTROL
	gpDebugMessageControlARB                      C.GPDEBUGMESSAGECONTROLARB
	gpDebugMessageControlKHR                      C.GPDEBUGMESSAGECONTROLKHR
	gpDebugMessageInsert                          C.GPDEBUGMESSAGEINSERT
	gpDebugMessageInsertARB                       C.GPDEBUGMESSAGEINSERTARB
	gpDebugMessageInsertKHR                       C.GPDEBUGMESSAGEINSERTKHR
	gpDeleteBuffers                               C.GPDELETEBUFFERS
	gpDeleteFramebuffers                          C.GPDELETEFRAMEBUFFERS
	gpDeleteLists                                 C.GPDELETELISTS
	gpDeleteNamedStringARB                        C.GPDELETENAMEDSTRINGARB
	gpDeleteProgram                               C.GPDELETEPROGRAM
	gpDeleteProgramPipelines                      C.GPDELETEPROGRAMPIPELINES
	gpDeleteQueries                               C.GPDELETEQUERIES
	gpDeleteRenderbuffers                         C.GPDELETERENDERBUFFERS
	gpDeleteSamplers                              C.GPDELETESAMPLERS
	gpDeleteShader                                C.GPDELETESHADER
	gpDeleteSync                                  C.GPDELETESYNC
	gpDeleteTextures                              C.GPDELETETEXTURES
	gpDeleteTransformFeedbacks                    C.GPDELETETRANSFORMFEEDBACKS
	gpDeleteVertexArrays                          C.GPDELETEVERTEXARRAYS
	gpDepthFunc                                   C.GPDEPTHFUNC
	gpDepthMask                                   C.GPDEPTHMASK
	gpDepthRange                                  C.GPDEPTHRANGE
	gpDepthRangeArrayv                            C.GPDEPTHRANGEARRAYV
	gpDepthRangeIndexed                           C.GPDEPTHRANGEINDEXED
	gpDepthRangef                                 C.GPDEPTHRANGEF
	gpDetachShader                                C.GPDETACHSHADER
	gpDisable                                     C.GPDISABLE
	gpDisableClientState                          C.GPDISABLECLIENTSTATE
	gpDisableVertexArrayAttrib                    C.GPDISABLEVERTEXARRAYATTRIB
	gpDisableVertexAttribArray                    C.GPDISABLEVERTEXATTRIBARRAY
	gpDisablei                                    C.GPDISABLEI
	gpDispatchCompute                             C.GPDISPATCHCOMPUTE
	gpDispatchComputeGroupSizeARB                 C.GPDISPATCHCOMPUTEGROUPSIZEARB
	gpDispatchComputeIndirect                     C.GPDISPATCHCOMPUTEINDIRECT
	gpDrawArrays                                  C.GPDRAWARRAYS
	gpDrawArraysIndirect                          C.GPDRAWARRAYSINDIRECT
	gpDrawArraysInstanced                         C.GPDRAWARRAYSINSTANCED
	gpDrawArraysInstancedBaseInstance             C.GPDRAWARRAYSINSTANCEDBASEINSTANCE
	gpDrawBuffer                                  C.GPDRAWBUFFER
	gpDrawBuffers                                 C.GPDRAWBUFFERS
	gpDrawElements                                C.GPDRAWELEMENTS
	gpDrawElementsBaseVertex                      C.GPDRAWELEMENTSBASEVERTEX
	gpDrawElementsIndirect                        C.GPDRAWELEMENTSINDIRECT
	gpDrawElementsInstanced                       C.GPDRAWELEMENTSINSTANCED
	gpDrawElementsInstancedBaseInstance           C.GPDRAWELEMENTSINSTANCEDBASEINSTANCE
	gpDrawElementsInstancedBaseVertex             C.GPDRAWELEMENTSINSTANCEDBASEVERTEX
	gpDrawElementsInstancedBaseVertexBaseInstance C.GPDRAWELEMENTSINSTANCEDBASEVERTEXBASEINSTANCE
	gpDrawPixels                                  C.GPDRAWPIXELS
	gpDrawRangeElements                           C.GPDRAWRANGEELEMENTS
	gpDrawRangeElementsBaseVertex                 C.GPDRAWRANGEELEMENTSBASEVERTEX
	gpDrawTransformFeedback                       C.GPDRAWTRANSFORMFEEDBACK
	gpDrawTransformFeedbackInstanced              C.GPDRAWTRANSFORMFEEDBACKINSTANCED
	gpDrawTransformFeedbackStream                 C.GPDRAWTRANSFORMFEEDBACKSTREAM
	gpDrawTransformFeedbackStreamInstanced        C.GPDRAWTRANSFORMFEEDBACKSTREAMINSTANCED
	gpEdgeFlag                                    C.GPEDGEFLAG
	gpEdgeFlagPointer                             C.GPEDGEFLAGPOINTER
	gpEdgeFlagv                                   C.GPEDGEFLAGV
	gpEnable                                      C.GPENABLE
	gpEnableClientState                           C.GPENABLECLIENTSTATE
	gpEnableVertexArrayAttrib                     C.GPENABLEVERTEXARRAYATTRIB
	gpEnableVertexAttribArray                     C.GPENABLEVERTEXATTRIBARRAY
	gpEnablei                                     C.GPENABLEI
	gpEnd                                         C.GPEND
	gpEndConditionalRender                        C.GPENDCONDITIONALRENDER
	gpEndList                                     C.GPENDLIST
	gpEndQuery                                    C.GPENDQUERY
	gpEndQueryIndexed                             C.GPENDQUERYINDEXED
	gpEndTransformFeedback                        C.GPENDTRANSFORMFEEDBACK
	gpEvalCoord1d                                 C.GPEVALCOORD1D
	gpEvalCoord1dv                                C.GPEVALCOORD1DV
	gpEvalCoord1f                                 C.GPEVALCOORD1F
	gpEvalCoord1fv                                C.GPEVALCOORD1FV
	gpEvalCoord2d                                 C.GPEVALCOORD2D
	gpEvalCoord2dv                                C.GPEVALCOORD2DV
	gpEvalCoord2f                                 C.GPEVALCOORD2F
	gpEvalCoord2fv                                C.GPEVALCOORD2FV
	gpEvalMesh1                                   C.GPEVALMESH1
	gpEvalMesh2                                   C.GPEVALMESH2
	gpEvalPoint1                                  C.GPEVALPOINT1
	gpEvalPoint2                                  C.GPEVALPOINT2
	gpFeedbackBuffer                              C.GPFEEDBACKBUFFER
	gpFenceSync                                   C.GPFENCESYNC
	gpFinish                                      C.GPFINISH
	gpFlush                                       C.GPFLUSH
	gpFlushMappedBufferRange                      C.GPFLUSHMAPPEDBUFFERRANGE
	gpFlushMappedNamedBufferRange                 C.GPFLUSHMAPPEDNAMEDBUFFERRANGE
	gpFogCoordPointer                             C.GPFOGCOORDPOINTER
	gpFogCoordd                                   C.GPFOGCOORDD
	gpFogCoorddv                                  C.GPFOGCOORDDV
	gpFogCoordf                                   C.GPFOGCOORDF
	gpFogCoordfv                                  C.GPFOGCOORDFV
	gpFogf                                        C.GPFOGF
	gpFogfv                                       C.GPFOGFV
	gpFogi                                        C.GPFOGI
	gpFogiv                                       C.GPFOGIV
	gpFramebufferParameteri                       C.GPFRAMEBUFFERPARAMETERI
	gpFramebufferRenderbuffer                     C.GPFRAMEBUFFERRENDERBUFFER
	gpFramebufferTexture                          C.GPFRAMEBUFFERTEXTURE
	gpFramebufferTexture1D                        C.GPFRAMEBUFFERTEXTURE1D
	gpFramebufferTexture2D                        C.GPFRAMEBUFFERTEXTURE2D
	gpFramebufferTexture3D                        C.GPFRAMEBUFFERTEXTURE3D
	gpFramebufferTextureLayer                     C.GPFRAMEBUFFERTEXTURELAYER
	gpFrontFace                                   C.GPFRONTFACE
	gpFrustum                                     C.GPFRUSTUM
	gpGenBuffers                                  C.GPGENBUFFERS
	gpGenFramebuffers                             C.GPGENFRAMEBUFFERS
	gpGenLists                                    C.GPGENLISTS
	gpGenProgramPipelines                         C.GPGENPROGRAMPIPELINES
	gpGenQueries                                  C.GPGENQUERIES
	gpGenRenderbuffers                            C.GPGENRENDERBUFFERS
	gpGenSamplers                                 C.GPGENSAMPLERS
	gpGenTextures                                 C.GPGENTEXTURES
	gpGenTransformFeedbacks                       C.GPGENTRANSFORMFEEDBACKS
	gpGenVertexArrays                             C.GPGENVERTEXARRAYS
	gpGenerateMipmap                              C.GPGENERATEMIPMAP
	gpGenerateTextureMipmap                       C.GPGENERATETEXTUREMIPMAP
	gpGetActiveAtomicCounterBufferiv              C.GPGETACTIVEATOMICCOUNTERBUFFERIV
	gpGetActiveAttrib                             C.GPGETACTIVEATTRIB
	gpGetActiveSubroutineName                     C.GPGETACTIVESUBROUTINENAME
	gpGetActiveSubroutineUniformName              C.GPGETACTIVESUBROUTINEUNIFORMNAME
	gpGetActiveSubroutineUniformiv                C.GPGETACTIVESUBROUTINEUNIFORMIV
	gpGetActiveUniform                            C.GPGETACTIVEUNIFORM
	gpGetActiveUniformBlockName                   C.GPGETACTIVEUNIFORMBLOCKNAME
	gpGetActiveUniformBlockiv                     C.GPGETACTIVEUNIFORMBLOCKIV
	gpGetActiveUniformName                        C.GPGETACTIVEUNIFORMNAME
	gpGetActiveUniformsiv                         C.GPGETACTIVEUNIFORMSIV
	gpGetAttachedShaders                          C.GPGETATTACHEDSHADERS
	gpGetAttribLocation                           C.GPGETATTRIBLOCATION
	gpGetBooleani_v                               C.GPGETBOOLEANI_V
	gpGetBooleanv                                 C.GPGETBOOLEANV
	gpGetBufferParameteri64v                      C.GPGETBUFFERPARAMETERI64V
	gpGetBufferParameteriv                        C.GPGETBUFFERPARAMETERIV
	gpGetBufferPointerv                           C.GPGETBUFFERPOINTERV
	gpGetBufferSubData                            C.GPGETBUFFERSUBDATA
	gpGetClipPlane                                C.GPGETCLIPPLANE
	gpGetCompressedTexImage                       C.GPGETCOMPRESSEDTEXIMAGE
	gpGetCompressedTextureImage                   C.GPGETCOMPRESSEDTEXTUREIMAGE
	gpGetCompressedTextureSubImage                C.GPGETCOMPRESSEDTEXTURESUBIMAGE
	gpGetDebugMessageLog                          C.GPGETDEBUGMESSAGELOG
	gpGetDebugMessageLogARB                       C.GPGETDEBUGMESSAGELOGARB
	gpGetDebugMessageLogKHR                       C.GPGETDEBUGMESSAGELOGKHR
	gpGetDoublei_v                                C.GPGETDOUBLEI_V
	gpGetDoublev                                  C.GPGETDOUBLEV
	gpGetError                                    C.GPGETERROR
	gpGetFloati_v                                 C.GPGETFLOATI_V
	gpGetFloatv                                   C.GPGETFLOATV
	gpGetFragDataIndex                            C.GPGETFRAGDATAINDEX
	gpGetFragDataLocation                         C.GPGETFRAGDATALOCATION
	gpGetFramebufferAttachmentParameteriv         C.GPGETFRAMEBUFFERATTACHMENTPARAMETERIV
	gpGetFramebufferParameteriv                   C.GPGETFRAMEBUFFERPARAMETERIV
	gpGetGraphicsResetStatus                      C.GPGETGRAPHICSRESETSTATUS
	gpGetGraphicsResetStatusARB                   C.GPGETGRAPHICSRESETSTATUSARB
	gpGetGraphicsResetStatusKHR                   C.GPGETGRAPHICSRESETSTATUSKHR
	gpGetImageHandleARB                           C.GPGETIMAGEHANDLEARB
	gpGetInteger64i_v                             C.GPGETINTEGER64I_V
	gpGetInteger64v                               C.GPGETINTEGER64V
	gpGetIntegeri_v                               C.GPGETINTEGERI_V
	gpGetIntegerv                                 C.GPGETINTEGERV
	gpGetInternalformati64v                       C.GPGETINTERNALFORMATI64V
	gpGetInternalformativ                         C.GPGETINTERNALFORMATIV
	gpGetLightfv                                  C.GPGETLIGHTFV
	gpGetLightiv                                  C.GPGETLIGHTIV
	gpGetMapdv                                    C.GPGETMAPDV
	gpGetMapfv                                    C.GPGETMAPFV
	gpGetMapiv                                    C.GPGETMAPIV
	gpGetMaterialfv                               C.GPGETMATERIALFV
	gpGetMaterialiv                               C.GPGETMATERIALIV
	gpGetMultisamplefv                            C.GPGETMULTISAMPLEFV
	gpGetNamedBufferParameteri64v                 C.GPGETNAMEDBUFFERPARAMETERI64V
	gpGetNamedBufferParameteriv                   C.GPGETNAMEDBUFFERPARAMETERIV
	gpGetNamedBufferPointerv                      C.GPGETNAMEDBUFFERPOINTERV
	gpGetNamedBufferSubData                       C.GPGETNAMEDBUFFERSUBDATA
	gpGetNamedFramebufferAttachmentParameteriv    C.GPGETNAMEDFRAMEBUFFERATTACHMENTPARAMETERIV
	gpGetNamedFramebufferParameteriv              C.GPGETNAMEDFRAMEBUFFERPARAMETERIV
	gpGetNamedRenderbufferParameteriv             C.GPGETNAMEDRENDERBUFFERPARAMETERIV
	gpGetNamedStringARB                           C.GPGETNAMEDSTRINGARB
	gpGetNamedStringivARB                         C.GPGETNAMEDSTRINGIVARB
	gpGetObjectLabel                              C.GPGETOBJECTLABEL
	gpGetObjectLabelKHR                           C.GPGETOBJECTLABELKHR
	gpGetObjectPtrLabel                           C.GPGETOBJECTPTRLABEL
	gpGetObjectPtrLabelKHR                        C.GPGETOBJECTPTRLABELKHR
	gpGetPixelMapfv                               C.GPGETPIXELMAPFV
	gpGetPixelMapuiv                              C.GPGETPIXELMAPUIV
	gpGetPixelMapusv                              C.GPGETPIXELMAPUSV
	gpGetPointerv                                 C.GPGETPOINTERV
	gpGetPointervKHR                              C.GPGETPOINTERVKHR
	gpGetPolygonStipple                           C.GPGETPOLYGONSTIPPLE
	gpGetProgramBinary                            C.GPGETPROGRAMBINARY
	gpGetProgramInfoLog                           C.GPGETPROGRAMINFOLOG
	gpGetProgramInterfaceiv                       C.GPGETPROGRAMINTERFACEIV
	gpGetProgramPipelineInfoLog                   C.GPGETPROGRAMPIPELINEINFOLOG
	gpGetProgramPipelineiv                        C.GPGETPROGRAMPIPELINEIV
	gpGetProgramResourceIndex                     C.GPGETPROGRAMRESOURCEINDEX
	gpGetProgramResourceLocation                  C.GPGETPROGRAMRESOURCELOCATION
	gpGetProgramResourceLocationIndex             C.GPGETPROGRAMRESOURCELOCATIONINDEX
	gpGetProgramResourceName                      C.GPGETPROGRAMRESOURCENAME
	gpGetProgramResourceiv                        C.GPGETPROGRAMRESOURCEIV
	gpGetProgramStageiv                           C.GPGETPROGRAMSTAGEIV
	gpGetProgramiv                                C.GPGETPROGRAMIV
	gpGetQueryIndexediv                           C.GPGETQUERYINDEXEDIV
	gpGetQueryObjecti64v                          C.GPGETQUERYOBJECTI64V
	gpGetQueryObjectiv                            C.GPGETQUERYOBJECTIV
	gpGetQueryObjectui64v                         C.GPGETQUERYOBJECTUI64V
	gpGetQueryObjectuiv                           C.GPGETQUERYOBJECTUIV
	gpGetQueryiv                                  C.GPGETQUERYIV
	gpGetRenderbufferParameteriv                  C.GPGETRENDERBUFFERPARAMETERIV
	gpGetSamplerParameterIiv                      C.GPGETSAMPLERPARAMETERIIV
	gpGetSamplerParameterIuiv                     C.GPGETSAMPLERPARAMETERIUIV
	gpGetSamplerParameterfv                       C.GPGETSAMPLERPARAMETERFV
	gpGetSamplerParameteriv                       C.GPGETSAMPLERPARAMETERIV
	gpGetShaderInfoLog                            C.GPGETSHADERINFOLOG
	gpGetShaderPrecisionFormat                    C.GPGETSHADERPRECISIONFORMAT
	gpGetShaderSource                             C.GPGETSHADERSOURCE
	gpGetShaderiv                                 C.GPGETSHADERIV
	gpGetString                                   C.GPGETSTRING
	gpGetStringi                                  C.GPGETSTRINGI
	gpGetSubroutineIndex                          C.GPGETSUBROUTINEINDEX
	gpGetSubroutineUniformLocation                C.GPGETSUBROUTINEUNIFORMLOCATION
	gpGetSynciv                                   C.GPGETSYNCIV
	gpGetTexEnvfv                                 C.GPGETTEXENVFV
	gpGetTexEnviv                                 C.GPGETTEXENVIV
	gpGetTexGendv                                 C.GPGETTEXGENDV
	gpGetTexGenfv                                 C.GPGETTEXGENFV
	gpGetTexGeniv                                 C.GPGETTEXGENIV
	gpGetTexImage                                 C.GPGETTEXIMAGE
	gpGetTexLevelParameterfv                      C.GPGETTEXLEVELPARAMETERFV
	gpGetTexLevelParameteriv                      C.GPGETTEXLEVELPARAMETERIV
	gpGetTexParameterIiv                          C.GPGETTEXPARAMETERIIV
	gpGetTexParameterIuiv                         C.GPGETTEXPARAMETERIUIV
	gpGetTexParameterfv                           C.GPGETTEXPARAMETERFV
	gpGetTexParameteriv                           C.GPGETTEXPARAMETERIV
	gpGetTextureHandleARB                         C.GPGETTEXTUREHANDLEARB
	gpGetTextureImage                             C.GPGETTEXTUREIMAGE
	gpGetTextureLevelParameterfv                  C.GPGETTEXTURELEVELPARAMETERFV
	gpGetTextureLevelParameteriv                  C.GPGETTEXTURELEVELPARAMETERIV
	gpGetTextureParameterIiv                      C.GPGETTEXTUREPARAMETERIIV
	gpGetTextureParameterIuiv                     C.GPGETTEXTUREPARAMETERIUIV
	gpGetTextureParameterfv                       C.GPGETTEXTUREPARAMETERFV
	gpGetTextureParameteriv                       C.GPGETTEXTUREPARAMETERIV
	gpGetTextureSamplerHandleARB                  C.GPGETTEXTURESAMPLERHANDLEARB
	gpGetTextureSubImage                          C.GPGETTEXTURESUBIMAGE
	gpGetTransformFeedbackVarying                 C.GPGETTRANSFORMFEEDBACKVARYING
	gpGetTransformFeedbacki64_v                   C.GPGETTRANSFORMFEEDBACKI64_V
	gpGetTransformFeedbacki_v                     C.GPGETTRANSFORMFEEDBACKI_V
	gpGetTransformFeedbackiv                      C.GPGETTRANSFORMFEEDBACKIV
	gpGetUniformBlockIndex                        C.GPGETUNIFORMBLOCKINDEX
	gpGetUniformIndices                           C.GPGETUNIFORMINDICES
	gpGetUniformLocation                          C.GPGETUNIFORMLOCATION
	gpGetUniformSubroutineuiv                     C.GPGETUNIFORMSUBROUTINEUIV
	gpGetUniformdv                                C.GPGETUNIFORMDV
	gpGetUniformfv                                C.GPGETUNIFORMFV
	gpGetUniformiv                                C.GPGETUNIFORMIV
	gpGetUniformuiv                               C.GPGETUNIFORMUIV
	gpGetVertexArrayIndexed64iv                   C.GPGETVERTEXARRAYINDEXED64IV
	gpGetVertexArrayIndexediv                     C.GPGETVERTEXARRAYINDEXEDIV
	gpGetVertexArrayiv                            C.GPGETVERTEXARRAYIV
	gpGetVertexAttribIiv                          C.GPGETVERTEXATTRIBIIV
	gpGetVertexAttribIuiv                         C.GPGETVERTEXATTRIBIUIV
	gpGetVertexAttribLdv                          C.GPGETVERTEXATTRIBLDV
	gpGetVertexAttribLui64vARB                    C.GPGETVERTEXATTRIBLUI64VARB
	gpGetVertexAttribPointerv                     C.GPGETVERTEXATTRIBPOINTERV
	gpGetVertexAttribdv                           C.GPGETVERTEXATTRIBDV
	gpGetVertexAttribfv                           C.GPGETVERTEXATTRIBFV
	gpGetVertexAttribiv                           C.GPGETVERTEXATTRIBIV
	gpGetnCompressedTexImage                      C.GPGETNCOMPRESSEDTEXIMAGE
	gpGetnCompressedTexImageARB                   C.GPGETNCOMPRESSEDTEXIMAGEARB
	gpGetnTexImage                                C.GPGETNTEXIMAGE
	gpGetnTexImageARB                             C.GPGETNTEXIMAGEARB
	gpGetnUniformdv                               C.GPGETNUNIFORMDV
	gpGetnUniformdvARB                            C.GPGETNUNIFORMDVARB
	gpGetnUniformfv                               C.GPGETNUNIFORMFV
	gpGetnUniformfvARB                            C.GPGETNUNIFORMFVARB
	gpGetnUniformfvKHR                            C.GPGETNUNIFORMFVKHR
	gpGetnUniformiv                               C.GPGETNUNIFORMIV
	gpGetnUniformivARB                            C.GPGETNUNIFORMIVARB
	gpGetnUniformivKHR                            C.GPGETNUNIFORMIVKHR
	gpGetnUniformuiv                              C.GPGETNUNIFORMUIV
	gpGetnUniformuivARB                           C.GPGETNUNIFORMUIVARB
	gpGetnUniformuivKHR                           C.GPGETNUNIFORMUIVKHR
	gpHint                                        C.GPHINT
	gpIndexMask                                   C.GPINDEXMASK
	gpIndexPointer                                C.GPINDEXPOINTER
	gpIndexd                                      C.GPINDEXD
	gpIndexdv                                     C.GPINDEXDV
	gpIndexf                                      C.GPINDEXF
	gpIndexfv                                     C.GPINDEXFV
	gpIndexi                                      C.GPINDEXI
	gpIndexiv                                     C.GPINDEXIV
	gpIndexs                                      C.GPINDEXS
	gpIndexsv                                     C.GPINDEXSV
	gpIndexub                                     C.GPINDEXUB
	gpIndexubv                                    C.GPINDEXUBV
	gpInitNames                                   C.GPINITNAMES
	gpInterleavedArrays                           C.GPINTERLEAVEDARRAYS
	gpInvalidateBufferData                        C.GPINVALIDATEBUFFERDATA
	gpInvalidateBufferSubData                     C.GPINVALIDATEBUFFERSUBDATA
	gpInvalidateFramebuffer                       C.GPINVALIDATEFRAMEBUFFER
	gpInvalidateNamedFramebufferData              C.GPINVALIDATENAMEDFRAMEBUFFERDATA
	gpInvalidateNamedFramebufferSubData           C.GPINVALIDATENAMEDFRAMEBUFFERSUBDATA
	gpInvalidateSubFramebuffer                    C.GPINVALIDATESUBFRAMEBUFFER
	gpInvalidateTexImage                          C.GPINVALIDATETEXIMAGE
	gpInvalidateTexSubImage                       C.GPINVALIDATETEXSUBIMAGE
	gpIsBuffer                                    C.GPISBUFFER
	gpIsEnabled                                   C.GPISENABLED
	gpIsEnabledi                                  C.GPISENABLEDI
	gpIsFramebuffer                               C.GPISFRAMEBUFFER
	gpIsImageHandleResidentARB                    C.GPISIMAGEHANDLERESIDENTARB
	gpIsList                                      C.GPISLIST
	gpIsNamedStringARB                            C.GPISNAMEDSTRINGARB
	gpIsProgram                                   C.GPISPROGRAM
	gpIsProgramPipeline                           C.GPISPROGRAMPIPELINE
	gpIsQuery                                     C.GPISQUERY
	gpIsRenderbuffer                              C.GPISRENDERBUFFER
	gpIsSampler                                   C.GPISSAMPLER
	gpIsShader                                    C.GPISSHADER
	gpIsSync                                      C.GPISSYNC
	gpIsTexture                                   C.GPISTEXTURE
	gpIsTextureHandleResidentARB                  C.GPISTEXTUREHANDLERESIDENTARB
	gpIsTransformFeedback                         C.GPISTRANSFORMFEEDBACK
	gpIsVertexArray                               C.GPISVERTEXARRAY
	gpLightModelf                                 C.GPLIGHTMODELF
	gpLightModelfv                                C.GPLIGHTMODELFV
	gpLightModeli                                 C.GPLIGHTMODELI
	gpLightModeliv                                C.GPLIGHTMODELIV
	gpLightf                                      C.GPLIGHTF
	gpLightfv                                     C.GPLIGHTFV
	gpLighti                                      C.GPLIGHTI
	gpLightiv                                     C.GPLIGHTIV
	gpLineStipple                                 C.GPLINESTIPPLE
	gpLineWidth                                   C.GPLINEWIDTH
	gpLinkProgram                                 C.GPLINKPROGRAM
	gpListBase                                    C.GPLISTBASE
	gpLoadIdentity                                C.GPLOADIDENTITY
	gpLoadMatrixd                                 C.GPLOADMATRIXD
	gpLoadMatrixf                                 C.GPLOADMATRIXF
	gpLoadName                                    C.GPLOADNAME
	gpLoadTransposeMatrixd                        C.GPLOADTRANSPOSEMATRIXD
	gpLoadTransposeMatrixf                        C.GPLOADTRANSPOSEMATRIXF
	gpLogicOp                                     C.GPLOGICOP
	gpMakeImageHandleNonResidentARB               C.GPMAKEIMAGEHANDLENONRESIDENTARB
	gpMakeImageHandleResidentARB                  C.GPMAKEIMAGEHANDLERESIDENTARB
	gpMakeTextureHandleNonResidentARB             C.GPMAKETEXTUREHANDLENONRESIDENTARB
	gpMakeTextureHandleResidentARB                C.GPMAKETEXTUREHANDLERESIDENTARB
	gpMap1d                                       C.GPMAP1D
	gpMap1f                                       C.GPMAP1F
	gpMap2d                                       C.GPMAP2D
	gpMap2f                                       C.GPMAP2F
	gpMapBuffer                                   C.GPMAPBUFFER
	gpMapBufferRange                              C.GPMAPBUFFERRANGE
	gpMapGrid1d                                   C.GPMAPGRID1D
	gpMapGrid1f                                   C.GPMAPGRID1F
	gpMapGrid2d                                   C.GPMAPGRID2D
	gpMapGrid2f                                   C.GPMAPGRID2F
	gpMapNamedBuffer                              C.GPMAPNAMEDBUFFER
	gpMapNamedBufferRange                         C.GPMAPNAMEDBUFFERRANGE
	gpMaterialf                                   C.GPMATERIALF
	gpMaterialfv                                  C.GPMATERIALFV
	gpMateriali                                   C.GPMATERIALI
	gpMaterialiv                                  C.GPMATERIALIV
	gpMatrixMode                                  C.GPMATRIXMODE
	gpMemoryBarrier                               C.GPMEMORYBARRIER
	gpMemoryBarrierByRegion                       C.GPMEMORYBARRIERBYREGION
	gpMinSampleShading                            C.GPMINSAMPLESHADING
	gpMinSampleShadingARB                         C.GPMINSAMPLESHADINGARB
	gpMultMatrixd                                 C.GPMULTMATRIXD
	gpMultMatrixf                                 C.GPMULTMATRIXF
	gpMultTransposeMatrixd                        C.GPMULTTRANSPOSEMATRIXD
	gpMultTransposeMatrixf                        C.GPMULTTRANSPOSEMATRIXF
	gpMultiDrawArrays                             C.GPMULTIDRAWARRAYS
	gpMultiDrawArraysIndirect                     C.GPMULTIDRAWARRAYSINDIRECT
	gpMultiDrawArraysIndirectCountARB             C.GPMULTIDRAWARRAYSINDIRECTCOUNTARB
	gpMultiDrawElements                           C.GPMULTIDRAWELEMENTS
	gpMultiDrawElementsBaseVertex                 C.GPMULTIDRAWELEMENTSBASEVERTEX
	gpMultiDrawElementsIndirect                   C.GPMULTIDRAWELEMENTSINDIRECT
	gpMultiDrawElementsIndirectCountARB           C.GPMULTIDRAWELEMENTSINDIRECTCOUNTARB
	gpMultiTexCoord1d                             C.GPMULTITEXCOORD1D
	gpMultiTexCoord1dv                            C.GPMULTITEXCOORD1DV
	gpMultiTexCoord1f                             C.GPMULTITEXCOORD1F
	gpMultiTexCoord1fv                            C.GPMULTITEXCOORD1FV
	gpMultiTexCoord1i                             C.GPMULTITEXCOORD1I
	gpMultiTexCoord1iv                            C.GPMULTITEXCOORD1IV
	gpMultiTexCoord1s                             C.GPMULTITEXCOORD1S
	gpMultiTexCoord1sv                            C.GPMULTITEXCOORD1SV
	gpMultiTexCoord2d                             C.GPMULTITEXCOORD2D
	gpMultiTexCoord2dv                            C.GPMULTITEXCOORD2DV
	gpMultiTexCoord2f                             C.GPMULTITEXCOORD2F
	gpMultiTexCoord2fv                            C.GPMULTITEXCOORD2FV
	gpMultiTexCoord2i                             C.GPMULTITEXCOORD2I
	gpMultiTexCoord2iv                            C.GPMULTITEXCOORD2IV
	gpMultiTexCoord2s                             C.GPMULTITEXCOORD2S
	gpMultiTexCoord2sv                            C.GPMULTITEXCOORD2SV
	gpMultiTexCoord3d                             C.GPMULTITEXCOORD3D
	gpMultiTexCoord3dv                            C.GPMULTITEXCOORD3DV
	gpMultiTexCoord3f                             C.GPMULTITEXCOORD3F
	gpMultiTexCoord3fv                            C.GPMULTITEXCOORD3FV
	gpMultiTexCoord3i                             C.GPMULTITEXCOORD3I
	gpMultiTexCoord3iv                            C.GPMULTITEXCOORD3IV
	gpMultiTexCoord3s                             C.GPMULTITEXCOORD3S
	gpMultiTexCoord3sv                            C.GPMULTITEXCOORD3SV
	gpMultiTexCoord4d                             C.GPMULTITEXCOORD4D
	gpMultiTexCoord4dv                            C.GPMULTITEXCOORD4DV
	gpMultiTexCoord4f                             C.GPMULTITEXCOORD4F
	gpMultiTexCoord4fv                            C.GPMULTITEXCOORD4FV
	gpMultiTexCoord4i                             C.GPMULTITEXCOORD4I
	gpMultiTexCoord4iv                            C.GPMULTITEXCOORD4IV
	gpMultiTexCoord4s                             C.GPMULTITEXCOORD4S
	gpMultiTexCoord4sv                            C.GPMULTITEXCOORD4SV
	gpNamedBufferData                             C.GPNAMEDBUFFERDATA
	gpNamedBufferPageCommitmentARB                C.GPNAMEDBUFFERPAGECOMMITMENTARB
	gpNamedBufferPageCommitmentEXT                C.GPNAMEDBUFFERPAGECOMMITMENTEXT
	gpNamedBufferStorage                          C.GPNAMEDBUFFERSTORAGE
	gpNamedBufferSubData                          C.GPNAMEDBUFFERSUBDATA
	gpNamedFramebufferDrawBuffer                  C.GPNAMEDFRAMEBUFFERDRAWBUFFER
	gpNamedFramebufferDrawBuffers                 C.GPNAMEDFRAMEBUFFERDRAWBUFFERS
	gpNamedFramebufferParameteri                  C.GPNAMEDFRAMEBUFFERPARAMETERI
	gpNamedFramebufferReadBuffer                  C.GPNAMEDFRAMEBUFFERREADBUFFER
	gpNamedFramebufferRenderbuffer                C.GPNAMEDFRAMEBUFFERRENDERBUFFER
	gpNamedFramebufferTexture                     C.GPNAMEDFRAMEBUFFERTEXTURE
	gpNamedFramebufferTextureLayer                C.GPNAMEDFRAMEBUFFERTEXTURELAYER
	gpNamedRenderbufferStorage                    C.GPNAMEDRENDERBUFFERSTORAGE
	gpNamedRenderbufferStorageMultisample         C.GPNAMEDRENDERBUFFERSTORAGEMULTISAMPLE
	gpNamedStringARB                              C.GPNAMEDSTRINGARB
	gpNewList                                     C.GPNEWLIST
	gpNormal3b                                    C.GPNORMAL3B
	gpNormal3bv                                   C.GPNORMAL3BV
	gpNormal3d                                    C.GPNORMAL3D
	gpNormal3dv                                   C.GPNORMAL3DV
	gpNormal3f                                    C.GPNORMAL3F
	gpNormal3fv                                   C.GPNORMAL3FV
	gpNormal3i                                    C.GPNORMAL3I
	gpNormal3iv                                   C.GPNORMAL3IV
	gpNormal3s                                    C.GPNORMAL3S
	gpNormal3sv                                   C.GPNORMAL3SV
	gpNormalPointer                               C.GPNORMALPOINTER
	gpObjectLabel                                 C.GPOBJECTLABEL
	gpObjectLabelKHR                              C.GPOBJECTLABELKHR
	gpObjectPtrLabel                              C.GPOBJECTPTRLABEL
	gpObjectPtrLabelKHR                           C.GPOBJECTPTRLABELKHR
	gpOrtho                                       C.GPORTHO
	gpPassThrough                                 C.GPPASSTHROUGH
	gpPatchParameterfv                            C.GPPATCHPARAMETERFV
	gpPatchParameteri                             C.GPPATCHPARAMETERI
	gpPauseTransformFeedback                      C.GPPAUSETRANSFORMFEEDBACK
	gpPixelMapfv                                  C.GPPIXELMAPFV
	gpPixelMapuiv                                 C.GPPIXELMAPUIV
	gpPixelMapusv                                 C.GPPIXELMAPUSV
	gpPixelStoref                                 C.GPPIXELSTOREF
	gpPixelStorei                                 C.GPPIXELSTOREI
	gpPixelTransferf                              C.GPPIXELTRANSFERF
	gpPixelTransferi                              C.GPPIXELTRANSFERI
	gpPixelZoom                                   C.GPPIXELZOOM
	gpPointParameterf                             C.GPPOINTPARAMETERF
	gpPointParameterfv                            C.GPPOINTPARAMETERFV
	gpPointParameteri                             C.GPPOINTPARAMETERI
	gpPointParameteriv                            C.GPPOINTPARAMETERIV
	gpPointSize                                   C.GPPOINTSIZE
	gpPolygonMode                                 C.GPPOLYGONMODE
	gpPolygonOffset                               C.GPPOLYGONOFFSET
	gpPolygonStipple                              C.GPPOLYGONSTIPPLE
	gpPopAttrib                                   C.GPPOPATTRIB
	gpPopClientAttrib                             C.GPPOPCLIENTATTRIB
	gpPopDebugGroup                               C.GPPOPDEBUGGROUP
	gpPopDebugGroupKHR                            C.GPPOPDEBUGGROUPKHR
	gpPopMatrix                                   C.GPPOPMATRIX
	gpPopName                                     C.GPPOPNAME
	gpPrimitiveRestartIndex                       C.GPPRIMITIVERESTARTINDEX
	gpPrioritizeTextures                          C.GPPRIORITIZETEXTURES
	gpProgramBinary                               C.GPPROGRAMBINARY
	gpProgramParameteri                           C.GPPROGRAMPARAMETERI
	gpProgramUniform1d                            C.GPPROGRAMUNIFORM1D
	gpProgramUniform1dv                           C.GPPROGRAMUNIFORM1DV
	gpProgramUniform1f                            C.GPPROGRAMUNIFORM1F
	gpProgramUniform1fv                           C.GPPROGRAMUNIFORM1FV
	gpProgramUniform1i                            C.GPPROGRAMUNIFORM1I
	gpProgramUniform1iv                           C.GPPROGRAMUNIFORM1IV
	gpProgramUniform1ui                           C.GPPROGRAMUNIFORM1UI
	gpProgramUniform1uiv                          C.GPPROGRAMUNIFORM1UIV
	gpProgramUniform2d                            C.GPPROGRAMUNIFORM2D
	gpProgramUniform2dv                           C.GPPROGRAMUNIFORM2DV
	gpProgramUniform2f                            C.GPPROGRAMUNIFORM2F
	gpProgramUniform2fv                           C.GPPROGRAMUNIFORM2FV
	gpProgramUniform2i                            C.GPPROGRAMUNIFORM2I
	gpProgramUniform2iv                           C.GPPROGRAMUNIFORM2IV
	gpProgramUniform2ui                           C.GPPROGRAMUNIFORM2UI
	gpProgramUniform2uiv                          C.GPPROGRAMUNIFORM2UIV
	gpProgramUniform3d                            C.GPPROGRAMUNIFORM3D
	gpProgramUniform3dv                           C.GPPROGRAMUNIFORM3DV
	gpProgramUniform3f                            C.GPPROGRAMUNIFORM3F
	gpProgramUniform3fv                           C.GPPROGRAMUNIFORM3FV
	gpProgramUniform3i                            C.GPPROGRAMUNIFORM3I
	gpProgramUniform3iv                           C.GPPROGRAMUNIFORM3IV
	gpProgramUniform3ui                           C.GPPROGRAMUNIFORM3UI
	gpProgramUniform3uiv                          C.GPPROGRAMUNIFORM3UIV
	gpProgramUniform4d                            C.GPPROGRAMUNIFORM4D
	gpProgramUniform4dv                           C.GPPROGRAMUNIFORM4DV
	gpProgramUniform4f                            C.GPPROGRAMUNIFORM4F
	gpProgramUniform4fv                           C.GPPROGRAMUNIFORM4FV
	gpProgramUniform4i                            C.GPPROGRAMUNIFORM4I
	gpProgramUniform4iv                           C.GPPROGRAMUNIFORM4IV
	gpProgramUniform4ui                           C.GPPROGRAMUNIFORM4UI
	gpProgramUniform4uiv                          C.GPPROGRAMUNIFORM4UIV
	gpProgramUniformHandleui64ARB                 C.GPPROGRAMUNIFORMHANDLEUI64ARB
	gpProgramUniformHandleui64vARB                C.GPPROGRAMUNIFORMHANDLEUI64VARB
	gpProgramUniformMatrix2dv                     C.GPPROGRAMUNIFORMMATRIX2DV
	gpProgramUniformMatrix2fv                     C.GPPROGRAMUNIFORMMATRIX2FV
	gpProgramUniformMatrix2x3dv                   C.GPPROGRAMUNIFORMMATRIX2X3DV
	gpProgramUniformMatrix2x3fv                   C.GPPROGRAMUNIFORMMATRIX2X3FV
	gpProgramUniformMatrix2x4dv                   C.GPPROGRAMUNIFORMMATRIX2X4DV
	gpProgramUniformMatrix2x4fv                   C.GPPROGRAMUNIFORMMATRIX2X4FV
	gpProgramUniformMatrix3dv                     C.GPPROGRAMUNIFORMMATRIX3DV
	gpProgramUniformMatrix3fv                     C.GPPROGRAMUNIFORMMATRIX3FV
	gpProgramUniformMatrix3x2dv                   C.GPPROGRAMUNIFORMMATRIX3X2DV
	gpProgramUniformMatrix3x2fv                   C.GPPROGRAMUNIFORMMATRIX3X2FV
	gpProgramUniformMatrix3x4dv                   C.GPPROGRAMUNIFORMMATRIX3X4DV
	gpProgramUniformMatrix3x4fv                   C.GPPROGRAMUNIFORMMATRIX3X4FV
	gpProgramUniformMatrix4dv                     C.GPPROGRAMUNIFORMMATRIX4DV
	gpProgramUniformMatrix4fv                     C.GPPROGRAMUNIFORMMATRIX4FV
	gpProgramUniformMatrix4x2dv                   C.GPPROGRAMUNIFORMMATRIX4X2DV
	gpProgramUniformMatrix4x2fv                   C.GPPROGRAMUNIFORMMATRIX4X2FV
	gpProgramUniformMatrix4x3dv                   C.GPPROGRAMUNIFORMMATRIX4X3DV
	gpProgramUniformMatrix4x3fv                   C.GPPROGRAMUNIFORMMATRIX4X3FV
	gpProvokingVertex                             C.GPPROVOKINGVERTEX
	gpPushAttrib                                  C.GPPUSHATTRIB
	gpPushClientAttrib                            C.GPPUSHCLIENTATTRIB
	gpPushDebugGroup                              C.GPPUSHDEBUGGROUP
	gpPushDebugGroupKHR                           C.GPPUSHDEBUGGROUPKHR
	gpPushMatrix                                  C.GPPUSHMATRIX
	gpPushName                                    C.GPPUSHNAME
	gpQueryCounter                                C.GPQUERYCOUNTER
	gpRasterPos2d                                 C.GPRASTERPOS2D
	gpRasterPos2dv                                C.GPRASTERPOS2DV
	gpRasterPos2f                                 C.GPRASTERPOS2F
	gpRasterPos2fv                                C.GPRASTERPOS2FV
	gpRasterPos2i                                 C.GPRASTERPOS2I
	gpRasterPos2iv                                C.GPRASTERPOS2IV
	gpRasterPos2s                                 C.GPRASTERPOS2S
	gpRasterPos2sv                                C.GPRASTERPOS2SV
	gpRasterPos3d                                 C.GPRASTERPOS3D
	gpRasterPos3dv                                C.GPRASTERPOS3DV
	gpRasterPos3f                                 C.GPRASTERPOS3F
	gpRasterPos3fv                                C.GPRASTERPOS3FV
	gpRasterPos3i                                 C.GPRASTERPOS3I
	gpRasterPos3iv                                C.GPRASTERPOS3IV
	gpRasterPos3s                                 C.GPRASTERPOS3S
	gpRasterPos3sv                                C.GPRASTERPOS3SV
	gpRasterPos4d                                 C.GPRASTERPOS4D
	gpRasterPos4dv                                C.GPRASTERPOS4DV
	gpRasterPos4f                                 C.GPRASTERPOS4F
	gpRasterPos4fv                                C.GPRASTERPOS4FV
	gpRasterPos4i                                 C.GPRASTERPOS4I
	gpRasterPos4iv                                C.GPRASTERPOS4IV
	gpRasterPos4s                                 C.GPRASTERPOS4S
	gpRasterPos4sv                                C.GPRASTERPOS4SV
	gpReadBuffer                                  C.GPREADBUFFER
	gpReadPixels                                  C.GPREADPIXELS
	gpReadnPixels                                 C.GPREADNPIXELS
	gpReadnPixelsARB                              C.GPREADNPIXELSARB
	gpReadnPixelsKHR                              C.GPREADNPIXELSKHR
	gpRectd                                       C.GPRECTD
	gpRectdv                                      C.GPRECTDV
	gpRectf                                       C.GPRECTF
	gpRectfv                                      C.GPRECTFV
	gpRecti                                       C.GPRECTI
	gpRectiv                                      C.GPRECTIV
	gpRects                                       C.GPRECTS
	gpRectsv                                      C.GPRECTSV
	gpReleaseShaderCompiler                       C.GPRELEASESHADERCOMPILER
	gpRenderMode                                  C.GPRENDERMODE
	gpRenderbufferStorage                         C.GPRENDERBUFFERSTORAGE
	gpRenderbufferStorageMultisample              C.GPRENDERBUFFERSTORAGEMULTISAMPLE
	gpResumeTransformFeedback                     C.GPRESUMETRANSFORMFEEDBACK
	gpRotated                                     C.GPROTATED
	gpRotatef                                     C.GPROTATEF
	gpSampleCoverage                              C.GPSAMPLECOVERAGE
	gpSampleMaski                                 C.GPSAMPLEMASKI
	gpSamplerParameterIiv                         C.GPSAMPLERPARAMETERIIV
	gpSamplerParameterIuiv                        C.GPSAMPLERPARAMETERIUIV
	gpSamplerParameterf                           C.GPSAMPLERPARAMETERF
	gpSamplerParameterfv                          C.GPSAMPLERPARAMETERFV
	gpSamplerParameteri                           C.GPSAMPLERPARAMETERI
	gpSamplerParameteriv                          C.GPSAMPLERPARAMETERIV
	gpScaled                                      C.GPSCALED
	gpScalef                                      C.GPSCALEF
	gpScissor                                     C.GPSCISSOR
	gpScissorArrayv                               C.GPSCISSORARRAYV
	gpScissorIndexed                              C.GPSCISSORINDEXED
	gpScissorIndexedv                             C.GPSCISSORINDEXEDV
	gpSecondaryColor3b                            C.GPSECONDARYCOLOR3B
	gpSecondaryColor3bv                           C.GPSECONDARYCOLOR3BV
	gpSecondaryColor3d                            C.GPSECONDARYCOLOR3D
	gpSecondaryColor3dv                           C.GPSECONDARYCOLOR3DV
	gpSecondaryColor3f                            C.GPSECONDARYCOLOR3F
	gpSecondaryColor3fv                           C.GPSECONDARYCOLOR3FV
	gpSecondaryColor3i                            C.GPSECONDARYCOLOR3I
	gpSecondaryColor3iv                           C.GPSECONDARYCOLOR3IV
	gpSecondaryColor3s                            C.GPSECONDARYCOLOR3S
	gpSecondaryColor3sv                           C.GPSECONDARYCOLOR3SV
	gpSecondaryColor3ub                           C.GPSECONDARYCOLOR3UB
	gpSecondaryColor3ubv                          C.GPSECONDARYCOLOR3UBV
	gpSecondaryColor3ui                           C.GPSECONDARYCOLOR3UI
	gpSecondaryColor3uiv                          C.GPSECONDARYCOLOR3UIV
	gpSecondaryColor3us                           C.GPSECONDARYCOLOR3US
	gpSecondaryColor3usv                          C.GPSECONDARYCOLOR3USV
	gpSecondaryColorPointer                       C.GPSECONDARYCOLORPOINTER
	gpSelectBuffer                                C.GPSELECTBUFFER
	gpShadeModel                                  C.GPSHADEMODEL
	gpShaderBinary                                C.GPSHADERBINARY
	gpShaderSource                                C.GPSHADERSOURCE
	gpShaderStorageBlockBinding                   C.GPSHADERSTORAGEBLOCKBINDING
	gpStencilFunc                                 C.GPSTENCILFUNC
	gpStencilFuncSeparate                         C.GPSTENCILFUNCSEPARATE
	gpStencilMask                                 C.GPSTENCILMASK
	gpStencilMaskSeparate                         C.GPSTENCILMASKSEPARATE
	gpStencilOp                                   C.GPSTENCILOP
	gpStencilOpSeparate                           C.GPSTENCILOPSEPARATE
	gpTexBuffer                                   C.GPTEXBUFFER
	gpTexBufferRange                              C.GPTEXBUFFERRANGE
	gpTexCoord1d                                  C.GPTEXCOORD1D
	gpTexCoord1dv                                 C.GPTEXCOORD1DV
	gpTexCoord1f                                  C.GPTEXCOORD1F
	gpTexCoord1fv                                 C.GPTEXCOORD1FV
	gpTexCoord1i                                  C.GPTEXCOORD1I
	gpTexCoord1iv                                 C.GPTEXCOORD1IV
	gpTexCoord1s                                  C.GPTEXCOORD1S
	gpTexCoord1sv                                 C.GPTEXCOORD1SV
	gpTexCoord2d                                  C.GPTEXCOORD2D
	gpTexCoord2dv                                 C.GPTEXCOORD2DV
	gpTexCoord2f                                  C.GPTEXCOORD2F
	gpTexCoord2fv                                 C.GPTEXCOORD2FV
	gpTexCoord2i                                  C.GPTEXCOORD2I
	gpTexCoord2iv                                 C.GPTEXCOORD2IV
	gpTexCoord2s                                  C.GPTEXCOORD2S
	gpTexCoord2sv                                 C.GPTEXCOORD2SV
	gpTexCoord3d                                  C.GPTEXCOORD3D
	gpTexCoord3dv                                 C.GPTEXCOORD3DV
	gpTexCoord3f                                  C.GPTEXCOORD3F
	gpTexCoord3fv                                 C.GPTEXCOORD3FV
	gpTexCoord3i                                  C.GPTEXCOORD3I
	gpTexCoord3iv                                 C.GPTEXCOORD3IV
	gpTexCoord3s                                  C.GPTEXCOORD3S
	gpTexCoord3sv                                 C.GPTEXCOORD3SV
	gpTexCoord4d                                  C.GPTEXCOORD4D
	gpTexCoord4dv                                 C.GPTEXCOORD4DV
	gpTexCoord4f                                  C.GPTEXCOORD4F
	gpTexCoord4fv                                 C.GPTEXCOORD4FV
	gpTexCoord4i                                  C.GPTEXCOORD4I
	gpTexCoord4iv                                 C.GPTEXCOORD4IV
	gpTexCoord4s                                  C.GPTEXCOORD4S
	gpTexCoord4sv                                 C.GPTEXCOORD4SV
	gpTexCoordPointer                             C.GPTEXCOORDPOINTER
	gpTexEnvf                                     C.GPTEXENVF
	gpTexEnvfv                                    C.GPTEXENVFV
	gpTexEnvi                                     C.GPTEXENVI
	gpTexEnviv                                    C.GPTEXENVIV
	gpTexGend                                     C.GPTEXGEND
	gpTexGendv                                    C.GPTEXGENDV
	gpTexGenf                                     C.GPTEXGENF
	gpTexGenfv                                    C.GPTEXGENFV
	gpTexGeni                                     C.GPTEXGENI
	gpTexGeniv                                    C.GPTEXGENIV
	gpTexImage1D                                  C.GPTEXIMAGE1D
	gpTexImage2D                                  C.GPTEXIMAGE2D
	gpTexImage2DMultisample                       C.GPTEXIMAGE2DMULTISAMPLE
	gpTexImage3D                                  C.GPTEXIMAGE3D
	gpTexImage3DMultisample                       C.GPTEXIMAGE3DMULTISAMPLE
	gpTexPageCommitmentARB                        C.GPTEXPAGECOMMITMENTARB
	gpTexParameterIiv                             C.GPTEXPARAMETERIIV
	gpTexParameterIuiv                            C.GPTEXPARAMETERIUIV
	gpTexParameterf                               C.GPTEXPARAMETERF
	gpTexParameterfv                              C.GPTEXPARAMETERFV
	gpTexParameteri                               C.GPTEXPARAMETERI
	gpTexParameteriv                              C.GPTEXPARAMETERIV
	gpTexStorage1D                                C.GPTEXSTORAGE1D
	gpTexStorage2D                                C.GPTEXSTORAGE2D
	gpTexStorage2DMultisample                     C.GPTEXSTORAGE2DMULTISAMPLE
	gpTexStorage3D                                C.GPTEXSTORAGE3D
	gpTexStorage3DMultisample                     C.GPTEXSTORAGE3DMULTISAMPLE
	gpTexSubImage1D                               C.GPTEXSUBIMAGE1D
	gpTexSubImage2D                               C.GPTEXSUBIMAGE2D
	gpTexSubImage3D                               C.GPTEXSUBIMAGE3D
	gpTextureBarrier                              C.GPTEXTUREBARRIER
	gpTextureBuffer                               C.GPTEXTUREBUFFER
	gpTextureBufferRange                          C.GPTEXTUREBUFFERRANGE
	gpTextureParameterIiv                         C.GPTEXTUREPARAMETERIIV
	gpTextureParameterIuiv                        C.GPTEXTUREPARAMETERIUIV
	gpTextureParameterf                           C.GPTEXTUREPARAMETERF
	gpTextureParameterfv                          C.GPTEXTUREPARAMETERFV
	gpTextureParameteri                           C.GPTEXTUREPARAMETERI
	gpTextureParameteriv                          C.GPTEXTUREPARAMETERIV
	gpTextureStorage1D                            C.GPTEXTURESTORAGE1D
	gpTextureStorage2D                            C.GPTEXTURESTORAGE2D
	gpTextureStorage2DMultisample                 C.GPTEXTURESTORAGE2DMULTISAMPLE
	gpTextureStorage3D                            C.GPTEXTURESTORAGE3D
	gpTextureStorage3DMultisample                 C.GPTEXTURESTORAGE3DMULTISAMPLE
	gpTextureSubImage1D                           C.GPTEXTURESUBIMAGE1D
	gpTextureSubImage2D                           C.GPTEXTURESUBIMAGE2D
	gpTextureSubImage3D                           C.GPTEXTURESUBIMAGE3D
	gpTextureView                                 C.GPTEXTUREVIEW
	gpTransformFeedbackBufferBase                 C.GPTRANSFORMFEEDBACKBUFFERBASE
	gpTransformFeedbackBufferRange                C.GPTRANSFORMFEEDBACKBUFFERRANGE
	gpTransformFeedbackVaryings                   C.GPTRANSFORMFEEDBACKVARYINGS
	gpTranslated                                  C.GPTRANSLATED
	gpTranslatef                                  C.GPTRANSLATEF
	gpUniform1d                                   C.GPUNIFORM1D
	gpUniform1dv                                  C.GPUNIFORM1DV
	gpUniform1f                                   C.GPUNIFORM1F
	gpUniform1fv                                  C.GPUNIFORM1FV
	gpUniform1i                                   C.GPUNIFORM1I
	gpUniform1iv                                  C.GPUNIFORM1IV
	gpUniform1ui                                  C.GPUNIFORM1UI
	gpUniform1uiv                                 C.GPUNIFORM1UIV
	gpUniform2d                                   C.GPUNIFORM2D
	gpUniform2dv                                  C.GPUNIFORM2DV
	gpUniform2f                                   C.GPUNIFORM2F
	gpUniform2fv                                  C.GPUNIFORM2FV
	gpUniform2i                                   C.GPUNIFORM2I
	gpUniform2iv                                  C.GPUNIFORM2IV
	gpUniform2ui                                  C.GPUNIFORM2UI
	gpUniform2uiv                                 C.GPUNIFORM2UIV
	gpUniform3d                                   C.GPUNIFORM3D
	gpUniform3dv                                  C.GPUNIFORM3DV
	gpUniform3f                                   C.GPUNIFORM3F
	gpUniform3fv                                  C.GPUNIFORM3FV
	gpUniform3i                                   C.GPUNIFORM3I
	gpUniform3iv                                  C.GPUNIFORM3IV
	gpUniform3ui                                  C.GPUNIFORM3UI
	gpUniform3uiv                                 C.GPUNIFORM3UIV
	gpUniform4d                                   C.GPUNIFORM4D
	gpUniform4dv                                  C.GPUNIFORM4DV
	gpUniform4f                                   C.GPUNIFORM4F
	gpUniform4fv                                  C.GPUNIFORM4FV
	gpUniform4i                                   C.GPUNIFORM4I
	gpUniform4iv                                  C.GPUNIFORM4IV
	gpUniform4ui                                  C.GPUNIFORM4UI
	gpUniform4uiv                                 C.GPUNIFORM4UIV
	gpUniformBlockBinding                         C.GPUNIFORMBLOCKBINDING
	gpUniformHandleui64ARB                        C.GPUNIFORMHANDLEUI64ARB
	gpUniformHandleui64vARB                       C.GPUNIFORMHANDLEUI64VARB
	gpUniformMatrix2dv                            C.GPUNIFORMMATRIX2DV
	gpUniformMatrix2fv                            C.GPUNIFORMMATRIX2FV
	gpUniformMatrix2x3dv                          C.GPUNIFORMMATRIX2X3DV
	gpUniformMatrix2x3fv                          C.GPUNIFORMMATRIX2X3FV
	gpUniformMatrix2x4dv                          C.GPUNIFORMMATRIX2X4DV
	gpUniformMatrix2x4fv                          C.GPUNIFORMMATRIX2X4FV
	gpUniformMatrix3dv                            C.GPUNIFORMMATRIX3DV
	gpUniformMatrix3fv                            C.GPUNIFORMMATRIX3FV
	gpUniformMatrix3x2dv                          C.GPUNIFORMMATRIX3X2DV
	gpUniformMatrix3x2fv                          C.GPUNIFORMMATRIX3X2FV
	gpUniformMatrix3x4dv                          C.GPUNIFORMMATRIX3X4DV
	gpUniformMatrix3x4fv                          C.GPUNIFORMMATRIX3X4FV
	gpUniformMatrix4dv                            C.GPUNIFORMMATRIX4DV
	gpUniformMatrix4fv                            C.GPUNIFORMMATRIX4FV
	gpUniformMatrix4x2dv                          C.GPUNIFORMMATRIX4X2DV
	gpUniformMatrix4x2fv                          C.GPUNIFORMMATRIX4X2FV
	gpUniformMatrix4x3dv                          C.GPUNIFORMMATRIX4X3DV
	gpUniformMatrix4x3fv                          C.GPUNIFORMMATRIX4X3FV
	gpUniformSubroutinesuiv                       C.GPUNIFORMSUBROUTINESUIV
	gpUnmapBuffer                                 C.GPUNMAPBUFFER
	gpUnmapNamedBuffer                            C.GPUNMAPNAMEDBUFFER
	gpUseProgram                                  C.GPUSEPROGRAM
	gpUseProgramStages                            C.GPUSEPROGRAMSTAGES
	gpValidateProgram                             C.GPVALIDATEPROGRAM
	gpValidateProgramPipeline                     C.GPVALIDATEPROGRAMPIPELINE
	gpVertex2d                                    C.GPVERTEX2D
	gpVertex2dv                                   C.GPVERTEX2DV
	gpVertex2f                                    C.GPVERTEX2F
	gpVertex2fv                                   C.GPVERTEX2FV
	gpVertex2i                                    C.GPVERTEX2I
	gpVertex2iv                                   C.GPVERTEX2IV
	gpVertex2s                                    C.GPVERTEX2S
	gpVertex2sv                                   C.GPVERTEX2SV
	gpVertex3d                                    C.GPVERTEX3D
	gpVertex3dv                                   C.GPVERTEX3DV
	gpVertex3f                                    C.GPVERTEX3F
	gpVertex3fv                                   C.GPVERTEX3FV
	gpVertex3i                                    C.GPVERTEX3I
	gpVertex3iv                                   C.GPVERTEX3IV
	gpVertex3s                                    C.GPVERTEX3S
	gpVertex3sv                                   C.GPVERTEX3SV
	gpVertex4d                                    C.GPVERTEX4D
	gpVertex4dv                                   C.GPVERTEX4DV
	gpVertex4f                                    C.GPVERTEX4F
	gpVertex4fv                                   C.GPVERTEX4FV
	gpVertex4i                                    C.GPVERTEX4I
	gpVertex4iv                                   C.GPVERTEX4IV
	gpVertex4s                                    C.GPVERTEX4S
	gpVertex4sv                                   C.GPVERTEX4SV
	gpVertexArrayAttribBinding                    C.GPVERTEXARRAYATTRIBBINDING
	gpVertexArrayAttribFormat                     C.GPVERTEXARRAYATTRIBFORMAT
	gpVertexArrayAttribIFormat                    C.GPVERTEXARRAYATTRIBIFORMAT
	gpVertexArrayAttribLFormat                    C.GPVERTEXARRAYATTRIBLFORMAT
	gpVertexArrayBindingDivisor                   C.GPVERTEXARRAYBINDINGDIVISOR
	gpVertexArrayElementBuffer                    C.GPVERTEXARRAYELEMENTBUFFER
	gpVertexArrayVertexBuffer                     C.GPVERTEXARRAYVERTEXBUFFER
	gpVertexArrayVertexBuffers                    C.GPVERTEXARRAYVERTEXBUFFERS
	gpVertexAttrib1d                              C.GPVERTEXATTRIB1D
	gpVertexAttrib1dv                             C.GPVERTEXATTRIB1DV
	gpVertexAttrib1f                              C.GPVERTEXATTRIB1F
	gpVertexAttrib1fv                             C.GPVERTEXATTRIB1FV
	gpVertexAttrib1s                              C.GPVERTEXATTRIB1S
	gpVertexAttrib1sv                             C.GPVERTEXATTRIB1SV
	gpVertexAttrib2d                              C.GPVERTEXATTRIB2D
	gpVertexAttrib2dv                             C.GPVERTEXATTRIB2DV
	gpVertexAttrib2f                              C.GPVERTEXATTRIB2F
	gpVertexAttrib2fv                             C.GPVERTEXATTRIB2FV
	gpVertexAttrib2s                              C.GPVERTEXATTRIB2S
	gpVertexAttrib2sv                             C.GPVERTEXATTRIB2SV
	gpVertexAttrib3d                              C.GPVERTEXATTRIB3D
	gpVertexAttrib3dv                             C.GPVERTEXATTRIB3DV
	gpVertexAttrib3f                              C.GPVERTEXATTRIB3F
	gpVertexAttrib3fv                             C.GPVERTEXATTRIB3FV
	gpVertexAttrib3s                              C.GPVERTEXATTRIB3S
	gpVertexAttrib3sv                             C.GPVERTEXATTRIB3SV
	gpVertexAttrib4Nbv                            C.GPVERTEXATTRIB4NBV
	gpVertexAttrib4Niv                            C.GPVERTEXATTRIB4NIV
	gpVertexAttrib4Nsv                            C.GPVERTEXATTRIB4NSV
	gpVertexAttrib4Nub                            C.GPVERTEXATTRIB4NUB
	gpVertexAttrib4Nubv                           C.GPVERTEXATTRIB4NUBV
	gpVertexAttrib4Nuiv                           C.GPVERTEXATTRIB4NUIV
	gpVertexAttrib4Nusv                           C.GPVERTEXATTRIB4NUSV
	gpVertexAttrib4bv                             C.GPVERTEXATTRIB4BV
	gpVertexAttrib4d                              C.GPVERTEXATTRIB4D
	gpVertexAttrib4dv                             C.GPVERTEXATTRIB4DV
	gpVertexAttrib4f                              C.GPVERTEXATTRIB4F
	gpVertexAttrib4fv                             C.GPVERTEXATTRIB4FV
	gpVertexAttrib4iv                             C.GPVERTEXATTRIB4IV
	gpVertexAttrib4s                              C.GPVERTEXATTRIB4S
	gpVertexAttrib4sv                             C.GPVERTEXATTRIB4SV
	gpVertexAttrib4ubv                            C.GPVERTEXATTRIB4UBV
	gpVertexAttrib4uiv                            C.GPVERTEXATTRIB4UIV
	gpVertexAttrib4usv                            C.GPVERTEXATTRIB4USV
	gpVertexAttribBinding                         C.GPVERTEXATTRIBBINDING
	gpVertexAttribDivisor                         C.GPVERTEXATTRIBDIVISOR
	gpVertexAttribFormat                          C.GPVERTEXATTRIBFORMAT
	gpVertexAttribI1i                             C.GPVERTEXATTRIBI1I
	gpVertexAttribI1iv                            C.GPVERTEXATTRIBI1IV
	gpVertexAttribI1ui                            C.GPVERTEXATTRIBI1UI
	gpVertexAttribI1uiv                           C.GPVERTEXATTRIBI1UIV
	gpVertexAttribI2i                             C.GPVERTEXATTRIBI2I
	gpVertexAttribI2iv                            C.GPVERTEXATTRIBI2IV
	gpVertexAttribI2ui                            C.GPVERTEXATTRIBI2UI
	gpVertexAttribI2uiv                           C.GPVERTEXATTRIBI2UIV
	gpVertexAttribI3i                             C.GPVERTEXATTRIBI3I
	gpVertexAttribI3iv                            C.GPVERTEXATTRIBI3IV
	gpVertexAttribI3ui                            C.GPVERTEXATTRIBI3UI
	gpVertexAttribI3uiv                           C.GPVERTEXATTRIBI3UIV
	gpVertexAttribI4bv                            C.GPVERTEXATTRIBI4BV
	gpVertexAttribI4i                             C.GPVERTEXATTRIBI4I
	gpVertexAttribI4iv                            C.GPVERTEXATTRIBI4IV
	gpVertexAttribI4sv                            C.GPVERTEXATTRIBI4SV
	gpVertexAttribI4ubv                           C.GPVERTEXATTRIBI4UBV
	gpVertexAttribI4ui                            C.GPVERTEXATTRIBI4UI
	gpVertexAttribI4uiv                           C.GPVERTEXATTRIBI4UIV
	gpVertexAttribI4usv                           C.GPVERTEXATTRIBI4USV
	gpVertexAttribIFormat                         C.GPVERTEXATTRIBIFORMAT
	gpVertexAttribIPointer                        C.GPVERTEXATTRIBIPOINTER
	gpVertexAttribL1d                             C.GPVERTEXATTRIBL1D
	gpVertexAttribL1dv                            C.GPVERTEXATTRIBL1DV
	gpVertexAttribL1ui64ARB                       C.GPVERTEXATTRIBL1UI64ARB
	gpVertexAttribL1ui64vARB                      C.GPVERTEXATTRIBL1UI64VARB
	gpVertexAttribL2d                             C.GPVERTEXATTRIBL2D
	gpVertexAttribL2dv                            C.GPVERTEXATTRIBL2DV
	gpVertexAttribL3d                             C.GPVERTEXATTRIBL3D
	gpVertexAttribL3dv                            C.GPVERTEXATTRIBL3DV
	gpVertexAttribL4d                             C.GPVERTEXATTRIBL4D
	gpVertexAttribL4dv                            C.GPVERTEXATTRIBL4DV
	gpVertexAttribLFormat                         C.GPVERTEXATTRIBLFORMAT
	gpVertexAttribLPointer                        C.GPVERTEXATTRIBLPOINTER
	gpVertexAttribP1ui                            C.GPVERTEXATTRIBP1UI
	gpVertexAttribP1uiv                           C.GPVERTEXATTRIBP1UIV
	gpVertexAttribP2ui                            C.GPVERTEXATTRIBP2UI
	gpVertexAttribP2uiv                           C.GPVERTEXATTRIBP2UIV
	gpVertexAttribP3ui                            C.GPVERTEXATTRIBP3UI
	gpVertexAttribP3uiv                           C.GPVERTEXATTRIBP3UIV
	gpVertexAttribP4ui                            C.GPVERTEXATTRIBP4UI
	gpVertexAttribP4uiv                           C.GPVERTEXATTRIBP4UIV
	gpVertexAttribPointer                         C.GPVERTEXATTRIBPOINTER
	gpVertexBindingDivisor                        C.GPVERTEXBINDINGDIVISOR
	gpVertexPointer                               C.GPVERTEXPOINTER
	gpViewport                                    C.GPVIEWPORT
	gpViewportArrayv                              C.GPVIEWPORTARRAYV
	gpViewportIndexedf                            C.GPVIEWPORTINDEXEDF
	gpViewportIndexedfv                           C.GPVIEWPORTINDEXEDFV
	gpWaitSync                                    C.GPWAITSYNC
	gpWindowPos2d                                 C.GPWINDOWPOS2D
	gpWindowPos2dv                                C.GPWINDOWPOS2DV
	gpWindowPos2f                                 C.GPWINDOWPOS2F
	gpWindowPos2fv                                C.GPWINDOWPOS2FV
	gpWindowPos2i                                 C.GPWINDOWPOS2I
	gpWindowPos2iv                                C.GPWINDOWPOS2IV
	gpWindowPos2s                                 C.GPWINDOWPOS2S
	gpWindowPos2sv                                C.GPWINDOWPOS2SV
	gpWindowPos3d                                 C.GPWINDOWPOS3D
	gpWindowPos3dv                                C.GPWINDOWPOS3DV
	gpWindowPos3f                                 C.GPWINDOWPOS3F
	gpWindowPos3fv                                C.GPWINDOWPOS3FV
	gpWindowPos3i                                 C.GPWINDOWPOS3I
	gpWindowPos3iv                                C.GPWINDOWPOS3IV
	gpWindowPos3s                                 C.GPWINDOWPOS3S
	gpWindowPos3sv                                C.GPWINDOWPOS3SV
)

// Helper functions
func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

// operate on the accumulation buffer
func Accum(op uint32, value float32) {
	C.glowAccum(gpAccum, (C.GLenum)(op), (C.GLfloat)(value))
}

// set the active program object for a program pipeline object
func ActiveShaderProgram(pipeline uint32, program uint32) {
	C.glowActiveShaderProgram(gpActiveShaderProgram, (C.GLuint)(pipeline), (C.GLuint)(program))
}

// select active texture unit
func ActiveTexture(texture uint32) {
	C.glowActiveTexture(gpActiveTexture, (C.GLenum)(texture))
}

// specify the alpha test function
func AlphaFunc(xfunc uint32, ref float32) {
	C.glowAlphaFunc(gpAlphaFunc, (C.GLenum)(xfunc), (C.GLfloat)(ref))
}

// determine if textures are loaded in texture memory
func AreTexturesResident(n int32, textures *uint32, residences *bool) bool {
	ret := C.glowAreTexturesResident(gpAreTexturesResident, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(textures)), (*C.GLboolean)(unsafe.Pointer(residences)))
	return ret == TRUE
}

// render a vertex using the specified vertex array element
func ArrayElement(i int32) {
	C.glowArrayElement(gpArrayElement, (C.GLint)(i))
}

// Attaches a shader object to a program object
func AttachShader(program uint32, shader uint32) {
	C.glowAttachShader(gpAttachShader, (C.GLuint)(program), (C.GLuint)(shader))
}

// delimit the vertices of a primitive or a group of like primitives
func Begin(mode uint32) {
	C.glowBegin(gpBegin, (C.GLenum)(mode))
}

// start conditional rendering
func BeginConditionalRender(id uint32, mode uint32) {
	C.glowBeginConditionalRender(gpBeginConditionalRender, (C.GLuint)(id), (C.GLenum)(mode))
}

// delimit the boundaries of a query object
func BeginQuery(target uint32, id uint32) {
	C.glowBeginQuery(gpBeginQuery, (C.GLenum)(target), (C.GLuint)(id))
}
func BeginQueryIndexed(target uint32, index uint32, id uint32) {
	C.glowBeginQueryIndexed(gpBeginQueryIndexed, (C.GLenum)(target), (C.GLuint)(index), (C.GLuint)(id))
}

// start transform feedback operation
func BeginTransformFeedback(primitiveMode uint32) {
	C.glowBeginTransformFeedback(gpBeginTransformFeedback, (C.GLenum)(primitiveMode))
}

// Associates a generic vertex attribute index with a named attribute variable
func BindAttribLocation(program uint32, index uint32, name *uint8) {
	C.glowBindAttribLocation(gpBindAttribLocation, (C.GLuint)(program), (C.GLuint)(index), (*C.GLchar)(unsafe.Pointer(name)))
}

// bind a named buffer object
func BindBuffer(target uint32, buffer uint32) {
	C.glowBindBuffer(gpBindBuffer, (C.GLenum)(target), (C.GLuint)(buffer))
}

// bind a buffer object to an indexed buffer target
func BindBufferBase(target uint32, index uint32, buffer uint32) {
	C.glowBindBufferBase(gpBindBufferBase, (C.GLenum)(target), (C.GLuint)(index), (C.GLuint)(buffer))
}

// bind a range within a buffer object to an indexed buffer target
func BindBufferRange(target uint32, index uint32, buffer uint32, offset int, size int) {
	C.glowBindBufferRange(gpBindBufferRange, (C.GLenum)(target), (C.GLuint)(index), (C.GLuint)(buffer), (C.GLintptr)(offset), (C.GLsizeiptr)(size))
}

// bind one or more buffer objects to a sequence of indexed buffer targets
func BindBuffersBase(target uint32, first uint32, count int32, buffers *uint32) {
	C.glowBindBuffersBase(gpBindBuffersBase, (C.GLenum)(target), (C.GLuint)(first), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(buffers)))
}

// bind ranges of one or more buffer objects to a sequence of indexed buffer targets
func BindBuffersRange(target uint32, first uint32, count int32, buffers *uint32, offsets *int, sizes *int) {
	C.glowBindBuffersRange(gpBindBuffersRange, (C.GLenum)(target), (C.GLuint)(first), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(buffers)), (*C.GLintptr)(unsafe.Pointer(offsets)), (*C.GLsizeiptr)(unsafe.Pointer(sizes)))
}

// bind a user-defined varying out variable to a fragment shader color number
func BindFragDataLocation(program uint32, color uint32, name *uint8) {
	C.glowBindFragDataLocation(gpBindFragDataLocation, (C.GLuint)(program), (C.GLuint)(color), (*C.GLchar)(unsafe.Pointer(name)))
}

// bind a user-defined varying out variable to a fragment shader color number and index
func BindFragDataLocationIndexed(program uint32, colorNumber uint32, index uint32, name *uint8) {
	C.glowBindFragDataLocationIndexed(gpBindFragDataLocationIndexed, (C.GLuint)(program), (C.GLuint)(colorNumber), (C.GLuint)(index), (*C.GLchar)(unsafe.Pointer(name)))
}

// bind a framebuffer to a framebuffer target
func BindFramebuffer(target uint32, framebuffer uint32) {
	C.glowBindFramebuffer(gpBindFramebuffer, (C.GLenum)(target), (C.GLuint)(framebuffer))
}

// bind a level of a texture to an image unit
func BindImageTexture(unit uint32, texture uint32, level int32, layered bool, layer int32, access uint32, format uint32) {
	C.glowBindImageTexture(gpBindImageTexture, (C.GLuint)(unit), (C.GLuint)(texture), (C.GLint)(level), (C.GLboolean)(boolToInt(layered)), (C.GLint)(layer), (C.GLenum)(access), (C.GLenum)(format))
}

// bind one or more named texture images to a sequence of consecutive image units
func BindImageTextures(first uint32, count int32, textures *uint32) {
	C.glowBindImageTextures(gpBindImageTextures, (C.GLuint)(first), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(textures)))
}

// bind a program pipeline to the current context
func BindProgramPipeline(pipeline uint32) {
	C.glowBindProgramPipeline(gpBindProgramPipeline, (C.GLuint)(pipeline))
}

// bind a renderbuffer to a renderbuffer target
func BindRenderbuffer(target uint32, renderbuffer uint32) {
	C.glowBindRenderbuffer(gpBindRenderbuffer, (C.GLenum)(target), (C.GLuint)(renderbuffer))
}

// bind a named sampler to a texturing target
func BindSampler(unit uint32, sampler uint32) {
	C.glowBindSampler(gpBindSampler, (C.GLuint)(unit), (C.GLuint)(sampler))
}

// bind one or more named sampler objects to a sequence of consecutive sampler units
func BindSamplers(first uint32, count int32, samplers *uint32) {
	C.glowBindSamplers(gpBindSamplers, (C.GLuint)(first), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(samplers)))
}

// bind a named texture to a texturing target
func BindTexture(target uint32, texture uint32) {
	C.glowBindTexture(gpBindTexture, (C.GLenum)(target), (C.GLuint)(texture))
}

// bind an existing texture object to the specified texture unit
func BindTextureUnit(unit uint32, texture uint32) {
	C.glowBindTextureUnit(gpBindTextureUnit, (C.GLuint)(unit), (C.GLuint)(texture))
}

// bind one or more named textures to a sequence of consecutive texture units
func BindTextures(first uint32, count int32, textures *uint32) {
	C.glowBindTextures(gpBindTextures, (C.GLuint)(first), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(textures)))
}

// bind a transform feedback object
func BindTransformFeedback(target uint32, id uint32) {
	C.glowBindTransformFeedback(gpBindTransformFeedback, (C.GLenum)(target), (C.GLuint)(id))
}

// bind a vertex array object
func BindVertexArray(array uint32) {
	C.glowBindVertexArray(gpBindVertexArray, (C.GLuint)(array))
}

// bind a buffer to a vertex buffer bind point
func BindVertexBuffer(bindingindex uint32, buffer uint32, offset int, stride int32) {
	C.glowBindVertexBuffer(gpBindVertexBuffer, (C.GLuint)(bindingindex), (C.GLuint)(buffer), (C.GLintptr)(offset), (C.GLsizei)(stride))
}

// attach multiple buffer objects to a vertex array object
func BindVertexBuffers(first uint32, count int32, buffers *uint32, offsets *int, strides *int32) {
	C.glowBindVertexBuffers(gpBindVertexBuffers, (C.GLuint)(first), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(buffers)), (*C.GLintptr)(unsafe.Pointer(offsets)), (*C.GLsizei)(unsafe.Pointer(strides)))
}

// draw a bitmap
func Bitmap(width int32, height int32, xorig float32, yorig float32, xmove float32, ymove float32, bitmap *uint8) {
	C.glowBitmap(gpBitmap, (C.GLsizei)(width), (C.GLsizei)(height), (C.GLfloat)(xorig), (C.GLfloat)(yorig), (C.GLfloat)(xmove), (C.GLfloat)(ymove), (*C.GLubyte)(unsafe.Pointer(bitmap)))
}

// set the blend color
func BlendColor(red float32, green float32, blue float32, alpha float32) {
	C.glowBlendColor(gpBlendColor, (C.GLfloat)(red), (C.GLfloat)(green), (C.GLfloat)(blue), (C.GLfloat)(alpha))
}

// specify the equation used for both the RGB blend equation and the Alpha blend equation
func BlendEquation(mode uint32) {
	C.glowBlendEquation(gpBlendEquation, (C.GLenum)(mode))
}

// set the RGB blend equation and the alpha blend equation separately
func BlendEquationSeparate(modeRGB uint32, modeAlpha uint32) {
	C.glowBlendEquationSeparate(gpBlendEquationSeparate, (C.GLenum)(modeRGB), (C.GLenum)(modeAlpha))
}
func BlendEquationSeparatei(buf uint32, modeRGB uint32, modeAlpha uint32) {
	C.glowBlendEquationSeparatei(gpBlendEquationSeparatei, (C.GLuint)(buf), (C.GLenum)(modeRGB), (C.GLenum)(modeAlpha))
}
func BlendEquationSeparateiARB(buf uint32, modeRGB uint32, modeAlpha uint32) {
	C.glowBlendEquationSeparateiARB(gpBlendEquationSeparateiARB, (C.GLuint)(buf), (C.GLenum)(modeRGB), (C.GLenum)(modeAlpha))
}
func BlendEquationi(buf uint32, mode uint32) {
	C.glowBlendEquationi(gpBlendEquationi, (C.GLuint)(buf), (C.GLenum)(mode))
}
func BlendEquationiARB(buf uint32, mode uint32) {
	C.glowBlendEquationiARB(gpBlendEquationiARB, (C.GLuint)(buf), (C.GLenum)(mode))
}

// specify pixel arithmetic
func BlendFunc(sfactor uint32, dfactor uint32) {
	C.glowBlendFunc(gpBlendFunc, (C.GLenum)(sfactor), (C.GLenum)(dfactor))
}

// specify pixel arithmetic for RGB and alpha components separately
func BlendFuncSeparate(sfactorRGB uint32, dfactorRGB uint32, sfactorAlpha uint32, dfactorAlpha uint32) {
	C.glowBlendFuncSeparate(gpBlendFuncSeparate, (C.GLenum)(sfactorRGB), (C.GLenum)(dfactorRGB), (C.GLenum)(sfactorAlpha), (C.GLenum)(dfactorAlpha))
}
func BlendFuncSeparatei(buf uint32, srcRGB uint32, dstRGB uint32, srcAlpha uint32, dstAlpha uint32) {
	C.glowBlendFuncSeparatei(gpBlendFuncSeparatei, (C.GLuint)(buf), (C.GLenum)(srcRGB), (C.GLenum)(dstRGB), (C.GLenum)(srcAlpha), (C.GLenum)(dstAlpha))
}
func BlendFuncSeparateiARB(buf uint32, srcRGB uint32, dstRGB uint32, srcAlpha uint32, dstAlpha uint32) {
	C.glowBlendFuncSeparateiARB(gpBlendFuncSeparateiARB, (C.GLuint)(buf), (C.GLenum)(srcRGB), (C.GLenum)(dstRGB), (C.GLenum)(srcAlpha), (C.GLenum)(dstAlpha))
}
func BlendFunci(buf uint32, src uint32, dst uint32) {
	C.glowBlendFunci(gpBlendFunci, (C.GLuint)(buf), (C.GLenum)(src), (C.GLenum)(dst))
}
func BlendFunciARB(buf uint32, src uint32, dst uint32) {
	C.glowBlendFunciARB(gpBlendFunciARB, (C.GLuint)(buf), (C.GLenum)(src), (C.GLenum)(dst))
}

// copy a block of pixels from one framebuffer object to another
func BlitFramebuffer(srcX0 int32, srcY0 int32, srcX1 int32, srcY1 int32, dstX0 int32, dstY0 int32, dstX1 int32, dstY1 int32, mask uint32, filter uint32) {
	C.glowBlitFramebuffer(gpBlitFramebuffer, (C.GLint)(srcX0), (C.GLint)(srcY0), (C.GLint)(srcX1), (C.GLint)(srcY1), (C.GLint)(dstX0), (C.GLint)(dstY0), (C.GLint)(dstX1), (C.GLint)(dstY1), (C.GLbitfield)(mask), (C.GLenum)(filter))
}

// copy a block of pixels from one framebuffer object to another
func BlitNamedFramebuffer(readFramebuffer uint32, drawFramebuffer uint32, srcX0 int32, srcY0 int32, srcX1 int32, srcY1 int32, dstX0 int32, dstY0 int32, dstX1 int32, dstY1 int32, mask uint32, filter uint32) {
	C.glowBlitNamedFramebuffer(gpBlitNamedFramebuffer, (C.GLuint)(readFramebuffer), (C.GLuint)(drawFramebuffer), (C.GLint)(srcX0), (C.GLint)(srcY0), (C.GLint)(srcX1), (C.GLint)(srcY1), (C.GLint)(dstX0), (C.GLint)(dstY0), (C.GLint)(dstX1), (C.GLint)(dstY1), (C.GLbitfield)(mask), (C.GLenum)(filter))
}

// creates and initializes a buffer object's data     store
func BufferData(target uint32, size int, data unsafe.Pointer, usage uint32) {
	C.glowBufferData(gpBufferData, (C.GLenum)(target), (C.GLsizeiptr)(size), data, (C.GLenum)(usage))
}
func BufferPageCommitmentARB(target uint32, offset int, size int32, commit bool) {
	C.glowBufferPageCommitmentARB(gpBufferPageCommitmentARB, (C.GLenum)(target), (C.GLintptr)(offset), (C.GLsizei)(size), (C.GLboolean)(boolToInt(commit)))
}

// creates and initializes a buffer object's immutable data     store
func BufferStorage(target uint32, size int, data unsafe.Pointer, flags uint32) {
	C.glowBufferStorage(gpBufferStorage, (C.GLenum)(target), (C.GLsizeiptr)(size), data, (C.GLbitfield)(flags))
}

// updates a subset of a buffer object's data store
func BufferSubData(target uint32, offset int, size int, data unsafe.Pointer) {
	C.glowBufferSubData(gpBufferSubData, (C.GLenum)(target), (C.GLintptr)(offset), (C.GLsizeiptr)(size), data)
}

// execute a display list
func CallList(list uint32) {
	C.glowCallList(gpCallList, (C.GLuint)(list))
}

// execute a list of display lists
func CallLists(n int32, xtype uint32, lists unsafe.Pointer) {
	C.glowCallLists(gpCallLists, (C.GLsizei)(n), (C.GLenum)(xtype), lists)
}

// check the completeness status of a framebuffer
func CheckFramebufferStatus(target uint32) uint32 {
	ret := C.glowCheckFramebufferStatus(gpCheckFramebufferStatus, (C.GLenum)(target))
	return (uint32)(ret)
}

// check the completeness status of a framebuffer
func CheckNamedFramebufferStatus(framebuffer uint32, target uint32) uint32 {
	ret := C.glowCheckNamedFramebufferStatus(gpCheckNamedFramebufferStatus, (C.GLuint)(framebuffer), (C.GLenum)(target))
	return (uint32)(ret)
}

// specify whether data read via  should be clamped
func ClampColor(target uint32, clamp uint32) {
	C.glowClampColor(gpClampColor, (C.GLenum)(target), (C.GLenum)(clamp))
}

// clear buffers to preset values
func Clear(mask uint32) {
	C.glowClear(gpClear, (C.GLbitfield)(mask))
}

// specify clear values for the accumulation buffer
func ClearAccum(red float32, green float32, blue float32, alpha float32) {
	C.glowClearAccum(gpClearAccum, (C.GLfloat)(red), (C.GLfloat)(green), (C.GLfloat)(blue), (C.GLfloat)(alpha))
}

// fill a buffer object's data store with a fixed value
func ClearBufferData(target uint32, internalformat uint32, format uint32, xtype uint32, data unsafe.Pointer) {
	C.glowClearBufferData(gpClearBufferData, (C.GLenum)(target), (C.GLenum)(internalformat), (C.GLenum)(format), (C.GLenum)(xtype), data)
}

// fill all or part of buffer object's data store with a fixed value
func ClearBufferSubData(target uint32, internalformat uint32, offset int, size int, format uint32, xtype uint32, data unsafe.Pointer) {
	C.glowClearBufferSubData(gpClearBufferSubData, (C.GLenum)(target), (C.GLenum)(internalformat), (C.GLintptr)(offset), (C.GLsizeiptr)(size), (C.GLenum)(format), (C.GLenum)(xtype), data)
}
func ClearBufferfi(buffer uint32, drawbuffer int32, depth float32, stencil int32) {
	C.glowClearBufferfi(gpClearBufferfi, (C.GLenum)(buffer), (C.GLint)(drawbuffer), (C.GLfloat)(depth), (C.GLint)(stencil))
}
func ClearBufferfv(buffer uint32, drawbuffer int32, value *float32) {
	C.glowClearBufferfv(gpClearBufferfv, (C.GLenum)(buffer), (C.GLint)(drawbuffer), (*C.GLfloat)(unsafe.Pointer(value)))
}
func ClearBufferiv(buffer uint32, drawbuffer int32, value *int32) {
	C.glowClearBufferiv(gpClearBufferiv, (C.GLenum)(buffer), (C.GLint)(drawbuffer), (*C.GLint)(unsafe.Pointer(value)))
}
func ClearBufferuiv(buffer uint32, drawbuffer int32, value *uint32) {
	C.glowClearBufferuiv(gpClearBufferuiv, (C.GLenum)(buffer), (C.GLint)(drawbuffer), (*C.GLuint)(unsafe.Pointer(value)))
}

// specify clear values for the color buffers
func ClearColor(red float32, green float32, blue float32, alpha float32) {
	C.glowClearColor(gpClearColor, (C.GLfloat)(red), (C.GLfloat)(green), (C.GLfloat)(blue), (C.GLfloat)(alpha))
}

// specify the clear value for the depth buffer
func ClearDepth(depth float64) {
	C.glowClearDepth(gpClearDepth, (C.GLdouble)(depth))
}
func ClearDepthf(d float32) {
	C.glowClearDepthf(gpClearDepthf, (C.GLfloat)(d))
}

// specify the clear value for the color index buffers
func ClearIndex(c float32) {
	C.glowClearIndex(gpClearIndex, (C.GLfloat)(c))
}

// fill a buffer object's data store with a fixed value
func ClearNamedBufferData(buffer uint32, internalformat uint32, format uint32, xtype uint32, data unsafe.Pointer) {
	C.glowClearNamedBufferData(gpClearNamedBufferData, (C.GLuint)(buffer), (C.GLenum)(internalformat), (C.GLenum)(format), (C.GLenum)(xtype), data)
}

// fill all or part of buffer object's data store with a fixed value
func ClearNamedBufferSubData(buffer uint32, internalformat uint32, offset int, size int32, format uint32, xtype uint32, data unsafe.Pointer) {
	C.glowClearNamedBufferSubData(gpClearNamedBufferSubData, (C.GLuint)(buffer), (C.GLenum)(internalformat), (C.GLintptr)(offset), (C.GLsizei)(size), (C.GLenum)(format), (C.GLenum)(xtype), data)
}
func ClearNamedFramebufferfi(framebuffer uint32, buffer uint32, depth float32, stencil int32) {
	C.glowClearNamedFramebufferfi(gpClearNamedFramebufferfi, (C.GLuint)(framebuffer), (C.GLenum)(buffer), (C.GLfloat)(depth), (C.GLint)(stencil))
}
func ClearNamedFramebufferfv(framebuffer uint32, buffer uint32, drawbuffer int32, value *float32) {
	C.glowClearNamedFramebufferfv(gpClearNamedFramebufferfv, (C.GLuint)(framebuffer), (C.GLenum)(buffer), (C.GLint)(drawbuffer), (*C.GLfloat)(unsafe.Pointer(value)))
}
func ClearNamedFramebufferiv(framebuffer uint32, buffer uint32, drawbuffer int32, value *int32) {
	C.glowClearNamedFramebufferiv(gpClearNamedFramebufferiv, (C.GLuint)(framebuffer), (C.GLenum)(buffer), (C.GLint)(drawbuffer), (*C.GLint)(unsafe.Pointer(value)))
}
func ClearNamedFramebufferuiv(framebuffer uint32, buffer uint32, drawbuffer int32, value *uint32) {
	C.glowClearNamedFramebufferuiv(gpClearNamedFramebufferuiv, (C.GLuint)(framebuffer), (C.GLenum)(buffer), (C.GLint)(drawbuffer), (*C.GLuint)(unsafe.Pointer(value)))
}

// specify the clear value for the stencil buffer
func ClearStencil(s int32) {
	C.glowClearStencil(gpClearStencil, (C.GLint)(s))
}

// fills all a texture image with a constant value
func ClearTexImage(texture uint32, level int32, format uint32, xtype uint32, data unsafe.Pointer) {
	C.glowClearTexImage(gpClearTexImage, (C.GLuint)(texture), (C.GLint)(level), (C.GLenum)(format), (C.GLenum)(xtype), data)
}

// fills all or part of a texture image with a constant value
func ClearTexSubImage(texture uint32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, format uint32, xtype uint32, data unsafe.Pointer) {
	C.glowClearTexSubImage(gpClearTexSubImage, (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLenum)(format), (C.GLenum)(xtype), data)
}

// select active texture unit
func ClientActiveTexture(texture uint32) {
	C.glowClientActiveTexture(gpClientActiveTexture, (C.GLenum)(texture))
}

// block and wait for a sync object to become signaled
func ClientWaitSync(sync uintptr, flags uint32, timeout uint64) uint32 {
	ret := C.glowClientWaitSync(gpClientWaitSync, (C.GLsync)(sync), (C.GLbitfield)(flags), (C.GLuint64)(timeout))
	return (uint32)(ret)
}

// control clip coordinate to window coordinate behavior
func ClipControl(origin uint32, depth uint32) {
	C.glowClipControl(gpClipControl, (C.GLenum)(origin), (C.GLenum)(depth))
}

// specify a plane against which all geometry is clipped
func ClipPlane(plane uint32, equation *float64) {
	C.glowClipPlane(gpClipPlane, (C.GLenum)(plane), (*C.GLdouble)(unsafe.Pointer(equation)))
}
func Color3b(red int8, green int8, blue int8) {
	C.glowColor3b(gpColor3b, (C.GLbyte)(red), (C.GLbyte)(green), (C.GLbyte)(blue))
}
func Color3bv(v *int8) {
	C.glowColor3bv(gpColor3bv, (*C.GLbyte)(unsafe.Pointer(v)))
}
func Color3d(red float64, green float64, blue float64) {
	C.glowColor3d(gpColor3d, (C.GLdouble)(red), (C.GLdouble)(green), (C.GLdouble)(blue))
}
func Color3dv(v *float64) {
	C.glowColor3dv(gpColor3dv, (*C.GLdouble)(unsafe.Pointer(v)))
}
func Color3f(red float32, green float32, blue float32) {
	C.glowColor3f(gpColor3f, (C.GLfloat)(red), (C.GLfloat)(green), (C.GLfloat)(blue))
}
func Color3fv(v *float32) {
	C.glowColor3fv(gpColor3fv, (*C.GLfloat)(unsafe.Pointer(v)))
}
func Color3i(red int32, green int32, blue int32) {
	C.glowColor3i(gpColor3i, (C.GLint)(red), (C.GLint)(green), (C.GLint)(blue))
}
func Color3iv(v *int32) {
	C.glowColor3iv(gpColor3iv, (*C.GLint)(unsafe.Pointer(v)))
}
func Color3s(red int16, green int16, blue int16) {
	C.glowColor3s(gpColor3s, (C.GLshort)(red), (C.GLshort)(green), (C.GLshort)(blue))
}
func Color3sv(v *int16) {
	C.glowColor3sv(gpColor3sv, (*C.GLshort)(unsafe.Pointer(v)))
}
func Color3ub(red uint8, green uint8, blue uint8) {
	C.glowColor3ub(gpColor3ub, (C.GLubyte)(red), (C.GLubyte)(green), (C.GLubyte)(blue))
}
func Color3ubv(v *uint8) {
	C.glowColor3ubv(gpColor3ubv, (*C.GLubyte)(unsafe.Pointer(v)))
}
func Color3ui(red uint32, green uint32, blue uint32) {
	C.glowColor3ui(gpColor3ui, (C.GLuint)(red), (C.GLuint)(green), (C.GLuint)(blue))
}
func Color3uiv(v *uint32) {
	C.glowColor3uiv(gpColor3uiv, (*C.GLuint)(unsafe.Pointer(v)))
}
func Color3us(red uint16, green uint16, blue uint16) {
	C.glowColor3us(gpColor3us, (C.GLushort)(red), (C.GLushort)(green), (C.GLushort)(blue))
}
func Color3usv(v *uint16) {
	C.glowColor3usv(gpColor3usv, (*C.GLushort)(unsafe.Pointer(v)))
}
func Color4b(red int8, green int8, blue int8, alpha int8) {
	C.glowColor4b(gpColor4b, (C.GLbyte)(red), (C.GLbyte)(green), (C.GLbyte)(blue), (C.GLbyte)(alpha))
}
func Color4bv(v *int8) {
	C.glowColor4bv(gpColor4bv, (*C.GLbyte)(unsafe.Pointer(v)))
}
func Color4d(red float64, green float64, blue float64, alpha float64) {
	C.glowColor4d(gpColor4d, (C.GLdouble)(red), (C.GLdouble)(green), (C.GLdouble)(blue), (C.GLdouble)(alpha))
}
func Color4dv(v *float64) {
	C.glowColor4dv(gpColor4dv, (*C.GLdouble)(unsafe.Pointer(v)))
}
func Color4f(red float32, green float32, blue float32, alpha float32) {
	C.glowColor4f(gpColor4f, (C.GLfloat)(red), (C.GLfloat)(green), (C.GLfloat)(blue), (C.GLfloat)(alpha))
}
func Color4fv(v *float32) {
	C.glowColor4fv(gpColor4fv, (*C.GLfloat)(unsafe.Pointer(v)))
}
func Color4i(red int32, green int32, blue int32, alpha int32) {
	C.glowColor4i(gpColor4i, (C.GLint)(red), (C.GLint)(green), (C.GLint)(blue), (C.GLint)(alpha))
}
func Color4iv(v *int32) {
	C.glowColor4iv(gpColor4iv, (*C.GLint)(unsafe.Pointer(v)))
}
func Color4s(red int16, green int16, blue int16, alpha int16) {
	C.glowColor4s(gpColor4s, (C.GLshort)(red), (C.GLshort)(green), (C.GLshort)(blue), (C.GLshort)(alpha))
}
func Color4sv(v *int16) {
	C.glowColor4sv(gpColor4sv, (*C.GLshort)(unsafe.Pointer(v)))
}
func Color4ub(red uint8, green uint8, blue uint8, alpha uint8) {
	C.glowColor4ub(gpColor4ub, (C.GLubyte)(red), (C.GLubyte)(green), (C.GLubyte)(blue), (C.GLubyte)(alpha))
}
func Color4ubv(v *uint8) {
	C.glowColor4ubv(gpColor4ubv, (*C.GLubyte)(unsafe.Pointer(v)))
}
func Color4ui(red uint32, green uint32, blue uint32, alpha uint32) {
	C.glowColor4ui(gpColor4ui, (C.GLuint)(red), (C.GLuint)(green), (C.GLuint)(blue), (C.GLuint)(alpha))
}
func Color4uiv(v *uint32) {
	C.glowColor4uiv(gpColor4uiv, (*C.GLuint)(unsafe.Pointer(v)))
}
func Color4us(red uint16, green uint16, blue uint16, alpha uint16) {
	C.glowColor4us(gpColor4us, (C.GLushort)(red), (C.GLushort)(green), (C.GLushort)(blue), (C.GLushort)(alpha))
}
func Color4usv(v *uint16) {
	C.glowColor4usv(gpColor4usv, (*C.GLushort)(unsafe.Pointer(v)))
}
func ColorMask(red bool, green bool, blue bool, alpha bool) {
	C.glowColorMask(gpColorMask, (C.GLboolean)(boolToInt(red)), (C.GLboolean)(boolToInt(green)), (C.GLboolean)(boolToInt(blue)), (C.GLboolean)(boolToInt(alpha)))
}
func ColorMaski(index uint32, r bool, g bool, b bool, a bool) {
	C.glowColorMaski(gpColorMaski, (C.GLuint)(index), (C.GLboolean)(boolToInt(r)), (C.GLboolean)(boolToInt(g)), (C.GLboolean)(boolToInt(b)), (C.GLboolean)(boolToInt(a)))
}

// cause a material color to track the current color
func ColorMaterial(face uint32, mode uint32) {
	C.glowColorMaterial(gpColorMaterial, (C.GLenum)(face), (C.GLenum)(mode))
}

// define an array of colors
func ColorPointer(size int32, xtype uint32, stride int32, pointer unsafe.Pointer) {
	C.glowColorPointer(gpColorPointer, (C.GLint)(size), (C.GLenum)(xtype), (C.GLsizei)(stride), pointer)
}

// Compiles a shader object
func CompileShader(shader uint32) {
	C.glowCompileShader(gpCompileShader, (C.GLuint)(shader))
}
func CompileShaderIncludeARB(shader uint32, count int32, path **uint8, length *int32) {
	C.glowCompileShaderIncludeARB(gpCompileShaderIncludeARB, (C.GLuint)(shader), (C.GLsizei)(count), (**C.GLchar)(unsafe.Pointer(path)), (*C.GLint)(unsafe.Pointer(length)))
}

// specify a one-dimensional texture image in a compressed format
func CompressedTexImage1D(target uint32, level int32, internalformat uint32, width int32, border int32, imageSize int32, data unsafe.Pointer) {
	C.glowCompressedTexImage1D(gpCompressedTexImage1D, (C.GLenum)(target), (C.GLint)(level), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLint)(border), (C.GLsizei)(imageSize), data)
}

// specify a two-dimensional texture image in a compressed format
func CompressedTexImage2D(target uint32, level int32, internalformat uint32, width int32, height int32, border int32, imageSize int32, data unsafe.Pointer) {
	C.glowCompressedTexImage2D(gpCompressedTexImage2D, (C.GLenum)(target), (C.GLint)(level), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLint)(border), (C.GLsizei)(imageSize), data)
}

// specify a three-dimensional texture image in a compressed format
func CompressedTexImage3D(target uint32, level int32, internalformat uint32, width int32, height int32, depth int32, border int32, imageSize int32, data unsafe.Pointer) {
	C.glowCompressedTexImage3D(gpCompressedTexImage3D, (C.GLenum)(target), (C.GLint)(level), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLint)(border), (C.GLsizei)(imageSize), data)
}

// specify a one-dimensional texture subimage in a compressed     format
func CompressedTexSubImage1D(target uint32, level int32, xoffset int32, width int32, format uint32, imageSize int32, data unsafe.Pointer) {
	C.glowCompressedTexSubImage1D(gpCompressedTexSubImage1D, (C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLsizei)(width), (C.GLenum)(format), (C.GLsizei)(imageSize), data)
}

// specify a two-dimensional texture subimage in a compressed format
func CompressedTexSubImage2D(target uint32, level int32, xoffset int32, yoffset int32, width int32, height int32, format uint32, imageSize int32, data unsafe.Pointer) {
	C.glowCompressedTexSubImage2D(gpCompressedTexSubImage2D, (C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLsizei)(imageSize), data)
}

// specify a three-dimensional texture subimage in a compressed format
func CompressedTexSubImage3D(target uint32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, format uint32, imageSize int32, data unsafe.Pointer) {
	C.glowCompressedTexSubImage3D(gpCompressedTexSubImage3D, (C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLenum)(format), (C.GLsizei)(imageSize), data)
}

// specify a one-dimensional texture subimage in a compressed     format
func CompressedTextureSubImage1D(texture uint32, level int32, xoffset int32, width int32, format uint32, imageSize int32, data unsafe.Pointer) {
	C.glowCompressedTextureSubImage1D(gpCompressedTextureSubImage1D, (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(xoffset), (C.GLsizei)(width), (C.GLenum)(format), (C.GLsizei)(imageSize), data)
}

// specify a two-dimensional texture subimage in a compressed format
func CompressedTextureSubImage2D(texture uint32, level int32, xoffset int32, yoffset int32, width int32, height int32, format uint32, imageSize int32, data unsafe.Pointer) {
	C.glowCompressedTextureSubImage2D(gpCompressedTextureSubImage2D, (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLsizei)(imageSize), data)
}

// specify a three-dimensional texture subimage in a compressed format
func CompressedTextureSubImage3D(texture uint32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, format uint32, imageSize int32, data unsafe.Pointer) {
	C.glowCompressedTextureSubImage3D(gpCompressedTextureSubImage3D, (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLenum)(format), (C.GLsizei)(imageSize), data)
}

// copy all or part of the data store of a buffer object to the data store of another buffer object
func CopyBufferSubData(readTarget uint32, writeTarget uint32, readOffset int, writeOffset int, size int) {
	C.glowCopyBufferSubData(gpCopyBufferSubData, (C.GLenum)(readTarget), (C.GLenum)(writeTarget), (C.GLintptr)(readOffset), (C.GLintptr)(writeOffset), (C.GLsizeiptr)(size))
}

// perform a raw data copy between two images
func CopyImageSubData(srcName uint32, srcTarget uint32, srcLevel int32, srcX int32, srcY int32, srcZ int32, dstName uint32, dstTarget uint32, dstLevel int32, dstX int32, dstY int32, dstZ int32, srcWidth int32, srcHeight int32, srcDepth int32) {
	C.glowCopyImageSubData(gpCopyImageSubData, (C.GLuint)(srcName), (C.GLenum)(srcTarget), (C.GLint)(srcLevel), (C.GLint)(srcX), (C.GLint)(srcY), (C.GLint)(srcZ), (C.GLuint)(dstName), (C.GLenum)(dstTarget), (C.GLint)(dstLevel), (C.GLint)(dstX), (C.GLint)(dstY), (C.GLint)(dstZ), (C.GLsizei)(srcWidth), (C.GLsizei)(srcHeight), (C.GLsizei)(srcDepth))
}

// copy all or part of the data store of a buffer object to the data store of another buffer object
func CopyNamedBufferSubData(readBuffer uint32, writeBuffer uint32, readOffset int, writeOffset int, size int32) {
	C.glowCopyNamedBufferSubData(gpCopyNamedBufferSubData, (C.GLuint)(readBuffer), (C.GLuint)(writeBuffer), (C.GLintptr)(readOffset), (C.GLintptr)(writeOffset), (C.GLsizei)(size))
}

// copy pixels in the frame buffer
func CopyPixels(x int32, y int32, width int32, height int32, xtype uint32) {
	C.glowCopyPixels(gpCopyPixels, (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(xtype))
}

// copy pixels into a 1D texture image
func CopyTexImage1D(target uint32, level int32, internalformat uint32, x int32, y int32, width int32, border int32) {
	C.glowCopyTexImage1D(gpCopyTexImage1D, (C.GLenum)(target), (C.GLint)(level), (C.GLenum)(internalformat), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLint)(border))
}

// copy pixels into a 2D texture image
func CopyTexImage2D(target uint32, level int32, internalformat uint32, x int32, y int32, width int32, height int32, border int32) {
	C.glowCopyTexImage2D(gpCopyTexImage2D, (C.GLenum)(target), (C.GLint)(level), (C.GLenum)(internalformat), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLint)(border))
}

// copy a one-dimensional texture subimage
func CopyTexSubImage1D(target uint32, level int32, xoffset int32, x int32, y int32, width int32) {
	C.glowCopyTexSubImage1D(gpCopyTexSubImage1D, (C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width))
}

// copy a two-dimensional texture subimage
func CopyTexSubImage2D(target uint32, level int32, xoffset int32, yoffset int32, x int32, y int32, width int32, height int32) {
	C.glowCopyTexSubImage2D(gpCopyTexSubImage2D, (C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}

// copy a three-dimensional texture subimage
func CopyTexSubImage3D(target uint32, level int32, xoffset int32, yoffset int32, zoffset int32, x int32, y int32, width int32, height int32) {
	C.glowCopyTexSubImage3D(gpCopyTexSubImage3D, (C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}

// copy a one-dimensional texture subimage
func CopyTextureSubImage1D(texture uint32, level int32, xoffset int32, x int32, y int32, width int32) {
	C.glowCopyTextureSubImage1D(gpCopyTextureSubImage1D, (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width))
}

// copy a two-dimensional texture subimage
func CopyTextureSubImage2D(texture uint32, level int32, xoffset int32, yoffset int32, x int32, y int32, width int32, height int32) {
	C.glowCopyTextureSubImage2D(gpCopyTextureSubImage2D, (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}

// copy a three-dimensional texture subimage
func CopyTextureSubImage3D(texture uint32, level int32, xoffset int32, yoffset int32, zoffset int32, x int32, y int32, width int32, height int32) {
	C.glowCopyTextureSubImage3D(gpCopyTextureSubImage3D, (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}

// create buffer objects
func CreateBuffers(n int32, buffers *uint32) {
	C.glowCreateBuffers(gpCreateBuffers, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(buffers)))
}

// create framebuffer objects
func CreateFramebuffers(n int32, framebuffers *uint32) {
	C.glowCreateFramebuffers(gpCreateFramebuffers, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(framebuffers)))
}

// Creates a program object
func CreateProgram() uint32 {
	ret := C.glowCreateProgram(gpCreateProgram)
	return (uint32)(ret)
}

// create program pipeline objects
func CreateProgramPipelines(n int32, pipelines *uint32) {
	C.glowCreateProgramPipelines(gpCreateProgramPipelines, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(pipelines)))
}

// create query objects
func CreateQueries(target uint32, n int32, ids *uint32) {
	C.glowCreateQueries(gpCreateQueries, (C.GLenum)(target), (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(ids)))
}

// create renderbuffer objects
func CreateRenderbuffers(n int32, renderbuffers *uint32) {
	C.glowCreateRenderbuffers(gpCreateRenderbuffers, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(renderbuffers)))
}

// create sampler objects
func CreateSamplers(n int32, samplers *uint32) {
	C.glowCreateSamplers(gpCreateSamplers, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(samplers)))
}

// Creates a shader object
func CreateShader(xtype uint32) uint32 {
	ret := C.glowCreateShader(gpCreateShader, (C.GLenum)(xtype))
	return (uint32)(ret)
}

// create a stand-alone program from an array of null-terminated source code strings
func CreateShaderProgramv(xtype uint32, count int32, strings **uint8) uint32 {
	ret := C.glowCreateShaderProgramv(gpCreateShaderProgramv, (C.GLenum)(xtype), (C.GLsizei)(count), (**C.GLchar)(unsafe.Pointer(strings)))
	return (uint32)(ret)
}
func CreateSyncFromCLeventARB(context *C.struct__cl_context, event *C.struct__cl_event, flags uint32) uintptr {
	ret := C.glowCreateSyncFromCLeventARB(gpCreateSyncFromCLeventARB, (*C.struct__cl_context)(unsafe.Pointer(context)), (*C.struct__cl_event)(unsafe.Pointer(event)), (C.GLbitfield)(flags))
	return (uintptr)(ret)
}

// create texture objects
func CreateTextures(target uint32, n int32, textures *uint32) {
	C.glowCreateTextures(gpCreateTextures, (C.GLenum)(target), (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(textures)))
}

// create transform feedback objects
func CreateTransformFeedbacks(n int32, ids *uint32) {
	C.glowCreateTransformFeedbacks(gpCreateTransformFeedbacks, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(ids)))
}

// create vertex array objects
func CreateVertexArrays(n int32, arrays *uint32) {
	C.glowCreateVertexArrays(gpCreateVertexArrays, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(arrays)))
}

// specify whether front- or back-facing facets can be culled
func CullFace(mode uint32) {
	C.glowCullFace(gpCullFace, (C.GLenum)(mode))
}

// specify a callback to receive debugging messages from the GL
func DebugMessageCallback(callback DebugProc, userParam unsafe.Pointer) {
	userDebugCallback = callback
	C.glowDebugMessageCallback(gpDebugMessageCallback, (C.GLDEBUGPROC)(unsafe.Pointer(&callback)), userParam)
}
func DebugMessageCallbackARB(callback DebugProc, userParam unsafe.Pointer) {
	userDebugCallback = callback
	C.glowDebugMessageCallbackARB(gpDebugMessageCallbackARB, (C.GLDEBUGPROCARB)(unsafe.Pointer(&callback)), userParam)
}
func DebugMessageCallbackKHR(callback DebugProc, userParam unsafe.Pointer) {
	userDebugCallback = callback
	C.glowDebugMessageCallbackKHR(gpDebugMessageCallbackKHR, (C.GLDEBUGPROCKHR)(unsafe.Pointer(&callback)), userParam)
}

// control the reporting of debug messages in a debug context
func DebugMessageControl(source uint32, xtype uint32, severity uint32, count int32, ids *uint32, enabled bool) {
	C.glowDebugMessageControl(gpDebugMessageControl, (C.GLenum)(source), (C.GLenum)(xtype), (C.GLenum)(severity), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(ids)), (C.GLboolean)(boolToInt(enabled)))
}
func DebugMessageControlARB(source uint32, xtype uint32, severity uint32, count int32, ids *uint32, enabled bool) {
	C.glowDebugMessageControlARB(gpDebugMessageControlARB, (C.GLenum)(source), (C.GLenum)(xtype), (C.GLenum)(severity), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(ids)), (C.GLboolean)(boolToInt(enabled)))
}
func DebugMessageControlKHR(source uint32, xtype uint32, severity uint32, count int32, ids *uint32, enabled bool) {
	C.glowDebugMessageControlKHR(gpDebugMessageControlKHR, (C.GLenum)(source), (C.GLenum)(xtype), (C.GLenum)(severity), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(ids)), (C.GLboolean)(boolToInt(enabled)))
}

// inject an application-supplied message into the debug message queue
func DebugMessageInsert(source uint32, xtype uint32, id uint32, severity uint32, length int32, buf *uint8) {
	C.glowDebugMessageInsert(gpDebugMessageInsert, (C.GLenum)(source), (C.GLenum)(xtype), (C.GLuint)(id), (C.GLenum)(severity), (C.GLsizei)(length), (*C.GLchar)(unsafe.Pointer(buf)))
}
func DebugMessageInsertARB(source uint32, xtype uint32, id uint32, severity uint32, length int32, buf *uint8) {
	C.glowDebugMessageInsertARB(gpDebugMessageInsertARB, (C.GLenum)(source), (C.GLenum)(xtype), (C.GLuint)(id), (C.GLenum)(severity), (C.GLsizei)(length), (*C.GLchar)(unsafe.Pointer(buf)))
}
func DebugMessageInsertKHR(source uint32, xtype uint32, id uint32, severity uint32, length int32, buf *uint8) {
	C.glowDebugMessageInsertKHR(gpDebugMessageInsertKHR, (C.GLenum)(source), (C.GLenum)(xtype), (C.GLuint)(id), (C.GLenum)(severity), (C.GLsizei)(length), (*C.GLchar)(unsafe.Pointer(buf)))
}

// delete named buffer objects
func DeleteBuffers(n int32, buffers *uint32) {
	C.glowDeleteBuffers(gpDeleteBuffers, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(buffers)))
}

// delete framebuffer objects
func DeleteFramebuffers(n int32, framebuffers *uint32) {
	C.glowDeleteFramebuffers(gpDeleteFramebuffers, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(framebuffers)))
}

// delete a contiguous group of display lists
func DeleteLists(list uint32, xrange int32) {
	C.glowDeleteLists(gpDeleteLists, (C.GLuint)(list), (C.GLsizei)(xrange))
}
func DeleteNamedStringARB(namelen int32, name *uint8) {
	C.glowDeleteNamedStringARB(gpDeleteNamedStringARB, (C.GLint)(namelen), (*C.GLchar)(unsafe.Pointer(name)))
}

// Deletes a program object
func DeleteProgram(program uint32) {
	C.glowDeleteProgram(gpDeleteProgram, (C.GLuint)(program))
}

// delete program pipeline objects
func DeleteProgramPipelines(n int32, pipelines *uint32) {
	C.glowDeleteProgramPipelines(gpDeleteProgramPipelines, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(pipelines)))
}

// delete named query objects
func DeleteQueries(n int32, ids *uint32) {
	C.glowDeleteQueries(gpDeleteQueries, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(ids)))
}

// delete renderbuffer objects
func DeleteRenderbuffers(n int32, renderbuffers *uint32) {
	C.glowDeleteRenderbuffers(gpDeleteRenderbuffers, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(renderbuffers)))
}

// delete named sampler objects
func DeleteSamplers(count int32, samplers *uint32) {
	C.glowDeleteSamplers(gpDeleteSamplers, (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(samplers)))
}

// Deletes a shader object
func DeleteShader(shader uint32) {
	C.glowDeleteShader(gpDeleteShader, (C.GLuint)(shader))
}

// delete a sync object
func DeleteSync(sync uintptr) {
	C.glowDeleteSync(gpDeleteSync, (C.GLsync)(sync))
}

// delete named textures
func DeleteTextures(n int32, textures *uint32) {
	C.glowDeleteTextures(gpDeleteTextures, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(textures)))
}

// delete transform feedback objects
func DeleteTransformFeedbacks(n int32, ids *uint32) {
	C.glowDeleteTransformFeedbacks(gpDeleteTransformFeedbacks, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(ids)))
}

// delete vertex array objects
func DeleteVertexArrays(n int32, arrays *uint32) {
	C.glowDeleteVertexArrays(gpDeleteVertexArrays, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(arrays)))
}

// specify the value used for depth buffer comparisons
func DepthFunc(xfunc uint32) {
	C.glowDepthFunc(gpDepthFunc, (C.GLenum)(xfunc))
}

// enable or disable writing into the depth buffer
func DepthMask(flag bool) {
	C.glowDepthMask(gpDepthMask, (C.GLboolean)(boolToInt(flag)))
}

// specify mapping of depth values from normalized device coordinates to window coordinates
func DepthRange(near float64, far float64) {
	C.glowDepthRange(gpDepthRange, (C.GLdouble)(near), (C.GLdouble)(far))
}
func DepthRangeArrayv(first uint32, count int32, v *float64) {
	C.glowDepthRangeArrayv(gpDepthRangeArrayv, (C.GLuint)(first), (C.GLsizei)(count), (*C.GLdouble)(unsafe.Pointer(v)))
}

// specify mapping of depth values from normalized device coordinates to window coordinates for a specified viewport
func DepthRangeIndexed(index uint32, n float64, f float64) {
	C.glowDepthRangeIndexed(gpDepthRangeIndexed, (C.GLuint)(index), (C.GLdouble)(n), (C.GLdouble)(f))
}
func DepthRangef(n float32, f float32) {
	C.glowDepthRangef(gpDepthRangef, (C.GLfloat)(n), (C.GLfloat)(f))
}

// Detaches a shader object from a program object to which it is attached
func DetachShader(program uint32, shader uint32) {
	C.glowDetachShader(gpDetachShader, (C.GLuint)(program), (C.GLuint)(shader))
}
func Disable(cap uint32) {
	C.glowDisable(gpDisable, (C.GLenum)(cap))
}
func DisableClientState(array uint32) {
	C.glowDisableClientState(gpDisableClientState, (C.GLenum)(array))
}

// Enable or disable a generic vertex attribute     array
func DisableVertexArrayAttrib(vaobj uint32, index uint32) {
	C.glowDisableVertexArrayAttrib(gpDisableVertexArrayAttrib, (C.GLuint)(vaobj), (C.GLuint)(index))
}

// Enable or disable a generic vertex attribute     array
func DisableVertexAttribArray(index uint32) {
	C.glowDisableVertexAttribArray(gpDisableVertexAttribArray, (C.GLuint)(index))
}
func Disablei(target uint32, index uint32) {
	C.glowDisablei(gpDisablei, (C.GLenum)(target), (C.GLuint)(index))
}

// launch one or more compute work groups
func DispatchCompute(num_groups_x uint32, num_groups_y uint32, num_groups_z uint32) {
	C.glowDispatchCompute(gpDispatchCompute, (C.GLuint)(num_groups_x), (C.GLuint)(num_groups_y), (C.GLuint)(num_groups_z))
}
func DispatchComputeGroupSizeARB(num_groups_x uint32, num_groups_y uint32, num_groups_z uint32, group_size_x uint32, group_size_y uint32, group_size_z uint32) {
	C.glowDispatchComputeGroupSizeARB(gpDispatchComputeGroupSizeARB, (C.GLuint)(num_groups_x), (C.GLuint)(num_groups_y), (C.GLuint)(num_groups_z), (C.GLuint)(group_size_x), (C.GLuint)(group_size_y), (C.GLuint)(group_size_z))
}

// launch one or more compute work groups using parameters stored in a buffer
func DispatchComputeIndirect(indirect int) {
	C.glowDispatchComputeIndirect(gpDispatchComputeIndirect, (C.GLintptr)(indirect))
}

// render primitives from array data
func DrawArrays(mode uint32, first int32, count int32) {
	C.glowDrawArrays(gpDrawArrays, (C.GLenum)(mode), (C.GLint)(first), (C.GLsizei)(count))
}

// render primitives from array data, taking parameters from memory
func DrawArraysIndirect(mode uint32, indirect unsafe.Pointer) {
	C.glowDrawArraysIndirect(gpDrawArraysIndirect, (C.GLenum)(mode), indirect)
}

// draw multiple instances of a range of elements
func DrawArraysInstanced(mode uint32, first int32, count int32, instancecount int32) {
	C.glowDrawArraysInstanced(gpDrawArraysInstanced, (C.GLenum)(mode), (C.GLint)(first), (C.GLsizei)(count), (C.GLsizei)(instancecount))
}

// draw multiple instances of a range of elements with offset applied to instanced attributes
func DrawArraysInstancedBaseInstance(mode uint32, first int32, count int32, instancecount int32, baseinstance uint32) {
	C.glowDrawArraysInstancedBaseInstance(gpDrawArraysInstancedBaseInstance, (C.GLenum)(mode), (C.GLint)(first), (C.GLsizei)(count), (C.GLsizei)(instancecount), (C.GLuint)(baseinstance))
}

// specify which color buffers are to be drawn into
func DrawBuffer(buf uint32) {
	C.glowDrawBuffer(gpDrawBuffer, (C.GLenum)(buf))
}

// Specifies a list of color buffers to be drawn     into
func DrawBuffers(n int32, bufs *uint32) {
	C.glowDrawBuffers(gpDrawBuffers, (C.GLsizei)(n), (*C.GLenum)(unsafe.Pointer(bufs)))
}

// render primitives from array data
func DrawElements(mode uint32, count int32, xtype uint32, indices unsafe.Pointer) {
	C.glowDrawElements(gpDrawElements, (C.GLenum)(mode), (C.GLsizei)(count), (C.GLenum)(xtype), indices)
}

// render primitives from array data with a per-element offset
func DrawElementsBaseVertex(mode uint32, count int32, xtype uint32, indices unsafe.Pointer, basevertex int32) {
	C.glowDrawElementsBaseVertex(gpDrawElementsBaseVertex, (C.GLenum)(mode), (C.GLsizei)(count), (C.GLenum)(xtype), indices, (C.GLint)(basevertex))
}

// render indexed primitives from array data, taking parameters from memory
func DrawElementsIndirect(mode uint32, xtype uint32, indirect unsafe.Pointer) {
	C.glowDrawElementsIndirect(gpDrawElementsIndirect, (C.GLenum)(mode), (C.GLenum)(xtype), indirect)
}

// draw multiple instances of a set of elements
func DrawElementsInstanced(mode uint32, count int32, xtype uint32, indices unsafe.Pointer, instancecount int32) {
	C.glowDrawElementsInstanced(gpDrawElementsInstanced, (C.GLenum)(mode), (C.GLsizei)(count), (C.GLenum)(xtype), indices, (C.GLsizei)(instancecount))
}

// draw multiple instances of a set of elements with offset applied to instanced attributes
func DrawElementsInstancedBaseInstance(mode uint32, count int32, xtype uint32, indices unsafe.Pointer, instancecount int32, baseinstance uint32) {
	C.glowDrawElementsInstancedBaseInstance(gpDrawElementsInstancedBaseInstance, (C.GLenum)(mode), (C.GLsizei)(count), (C.GLenum)(xtype), indices, (C.GLsizei)(instancecount), (C.GLuint)(baseinstance))
}

// render multiple instances of a set of primitives from array data with a per-element offset
func DrawElementsInstancedBaseVertex(mode uint32, count int32, xtype uint32, indices unsafe.Pointer, instancecount int32, basevertex int32) {
	C.glowDrawElementsInstancedBaseVertex(gpDrawElementsInstancedBaseVertex, (C.GLenum)(mode), (C.GLsizei)(count), (C.GLenum)(xtype), indices, (C.GLsizei)(instancecount), (C.GLint)(basevertex))
}

// render multiple instances of a set of primitives from array data with a per-element offset
func DrawElementsInstancedBaseVertexBaseInstance(mode uint32, count int32, xtype uint32, indices unsafe.Pointer, instancecount int32, basevertex int32, baseinstance uint32) {
	C.glowDrawElementsInstancedBaseVertexBaseInstance(gpDrawElementsInstancedBaseVertexBaseInstance, (C.GLenum)(mode), (C.GLsizei)(count), (C.GLenum)(xtype), indices, (C.GLsizei)(instancecount), (C.GLint)(basevertex), (C.GLuint)(baseinstance))
}

// write a block of pixels to the frame buffer
func DrawPixels(width int32, height int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	C.glowDrawPixels(gpDrawPixels, (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLenum)(xtype), pixels)
}

// render primitives from array data
func DrawRangeElements(mode uint32, start uint32, end uint32, count int32, xtype uint32, indices unsafe.Pointer) {
	C.glowDrawRangeElements(gpDrawRangeElements, (C.GLenum)(mode), (C.GLuint)(start), (C.GLuint)(end), (C.GLsizei)(count), (C.GLenum)(xtype), indices)
}

// render primitives from array data with a per-element offset
func DrawRangeElementsBaseVertex(mode uint32, start uint32, end uint32, count int32, xtype uint32, indices unsafe.Pointer, basevertex int32) {
	C.glowDrawRangeElementsBaseVertex(gpDrawRangeElementsBaseVertex, (C.GLenum)(mode), (C.GLuint)(start), (C.GLuint)(end), (C.GLsizei)(count), (C.GLenum)(xtype), indices, (C.GLint)(basevertex))
}

// render primitives using a count derived from a transform feedback object
func DrawTransformFeedback(mode uint32, id uint32) {
	C.glowDrawTransformFeedback(gpDrawTransformFeedback, (C.GLenum)(mode), (C.GLuint)(id))
}

// render multiple instances of primitives using a count derived from a transform feedback object
func DrawTransformFeedbackInstanced(mode uint32, id uint32, instancecount int32) {
	C.glowDrawTransformFeedbackInstanced(gpDrawTransformFeedbackInstanced, (C.GLenum)(mode), (C.GLuint)(id), (C.GLsizei)(instancecount))
}

// render primitives using a count derived from a specifed stream of a transform feedback object
func DrawTransformFeedbackStream(mode uint32, id uint32, stream uint32) {
	C.glowDrawTransformFeedbackStream(gpDrawTransformFeedbackStream, (C.GLenum)(mode), (C.GLuint)(id), (C.GLuint)(stream))
}

// render multiple instances of primitives using a count derived from a specifed stream of a transform feedback object
func DrawTransformFeedbackStreamInstanced(mode uint32, id uint32, stream uint32, instancecount int32) {
	C.glowDrawTransformFeedbackStreamInstanced(gpDrawTransformFeedbackStreamInstanced, (C.GLenum)(mode), (C.GLuint)(id), (C.GLuint)(stream), (C.GLsizei)(instancecount))
}

// flag edges as either boundary or nonboundary
func EdgeFlag(flag bool) {
	C.glowEdgeFlag(gpEdgeFlag, (C.GLboolean)(boolToInt(flag)))
}

// define an array of edge flags
func EdgeFlagPointer(stride int32, pointer unsafe.Pointer) {
	C.glowEdgeFlagPointer(gpEdgeFlagPointer, (C.GLsizei)(stride), pointer)
}
func EdgeFlagv(flag *bool) {
	C.glowEdgeFlagv(gpEdgeFlagv, (*C.GLboolean)(unsafe.Pointer(flag)))
}

// enable or disable server-side GL capabilities
func Enable(cap uint32) {
	C.glowEnable(gpEnable, (C.GLenum)(cap))
}

// enable or disable client-side capability
func EnableClientState(array uint32) {
	C.glowEnableClientState(gpEnableClientState, (C.GLenum)(array))
}

// Enable or disable a generic vertex attribute     array
func EnableVertexArrayAttrib(vaobj uint32, index uint32) {
	C.glowEnableVertexArrayAttrib(gpEnableVertexArrayAttrib, (C.GLuint)(vaobj), (C.GLuint)(index))
}

// Enable or disable a generic vertex attribute     array
func EnableVertexAttribArray(index uint32) {
	C.glowEnableVertexAttribArray(gpEnableVertexAttribArray, (C.GLuint)(index))
}
func Enablei(target uint32, index uint32) {
	C.glowEnablei(gpEnablei, (C.GLenum)(target), (C.GLuint)(index))
}
func End() {
	C.glowEnd(gpEnd)
}
func EndConditionalRender() {
	C.glowEndConditionalRender(gpEndConditionalRender)
}
func EndList() {
	C.glowEndList(gpEndList)
}
func EndQuery(target uint32) {
	C.glowEndQuery(gpEndQuery, (C.GLenum)(target))
}
func EndQueryIndexed(target uint32, index uint32) {
	C.glowEndQueryIndexed(gpEndQueryIndexed, (C.GLenum)(target), (C.GLuint)(index))
}
func EndTransformFeedback() {
	C.glowEndTransformFeedback(gpEndTransformFeedback)
}
func EvalCoord1d(u float64) {
	C.glowEvalCoord1d(gpEvalCoord1d, (C.GLdouble)(u))
}
func EvalCoord1dv(u *float64) {
	C.glowEvalCoord1dv(gpEvalCoord1dv, (*C.GLdouble)(unsafe.Pointer(u)))
}
func EvalCoord1f(u float32) {
	C.glowEvalCoord1f(gpEvalCoord1f, (C.GLfloat)(u))
}
func EvalCoord1fv(u *float32) {
	C.glowEvalCoord1fv(gpEvalCoord1fv, (*C.GLfloat)(unsafe.Pointer(u)))
}
func EvalCoord2d(u float64, v float64) {
	C.glowEvalCoord2d(gpEvalCoord2d, (C.GLdouble)(u), (C.GLdouble)(v))
}
func EvalCoord2dv(u *float64) {
	C.glowEvalCoord2dv(gpEvalCoord2dv, (*C.GLdouble)(unsafe.Pointer(u)))
}
func EvalCoord2f(u float32, v float32) {
	C.glowEvalCoord2f(gpEvalCoord2f, (C.GLfloat)(u), (C.GLfloat)(v))
}
func EvalCoord2fv(u *float32) {
	C.glowEvalCoord2fv(gpEvalCoord2fv, (*C.GLfloat)(unsafe.Pointer(u)))
}
func EvalMesh1(mode uint32, i1 int32, i2 int32) {
	C.glowEvalMesh1(gpEvalMesh1, (C.GLenum)(mode), (C.GLint)(i1), (C.GLint)(i2))
}
func EvalMesh2(mode uint32, i1 int32, i2 int32, j1 int32, j2 int32) {
	C.glowEvalMesh2(gpEvalMesh2, (C.GLenum)(mode), (C.GLint)(i1), (C.GLint)(i2), (C.GLint)(j1), (C.GLint)(j2))
}
func EvalPoint1(i int32) {
	C.glowEvalPoint1(gpEvalPoint1, (C.GLint)(i))
}
func EvalPoint2(i int32, j int32) {
	C.glowEvalPoint2(gpEvalPoint2, (C.GLint)(i), (C.GLint)(j))
}

// controls feedback mode
func FeedbackBuffer(size int32, xtype uint32, buffer *float32) {
	C.glowFeedbackBuffer(gpFeedbackBuffer, (C.GLsizei)(size), (C.GLenum)(xtype), (*C.GLfloat)(unsafe.Pointer(buffer)))
}

// create a new sync object and insert it into the GL command stream
func FenceSync(condition uint32, flags uint32) uintptr {
	ret := C.glowFenceSync(gpFenceSync, (C.GLenum)(condition), (C.GLbitfield)(flags))
	return (uintptr)(ret)
}

// block until all GL execution is complete
func Finish() {
	C.glowFinish(gpFinish)
}

// force execution of GL commands in finite time
func Flush() {
	C.glowFlush(gpFlush)
}

// indicate modifications to a range of a mapped buffer
func FlushMappedBufferRange(target uint32, offset int, length int) {
	C.glowFlushMappedBufferRange(gpFlushMappedBufferRange, (C.GLenum)(target), (C.GLintptr)(offset), (C.GLsizeiptr)(length))
}

// indicate modifications to a range of a mapped buffer
func FlushMappedNamedBufferRange(buffer uint32, offset int, length int32) {
	C.glowFlushMappedNamedBufferRange(gpFlushMappedNamedBufferRange, (C.GLuint)(buffer), (C.GLintptr)(offset), (C.GLsizei)(length))
}

// define an array of fog coordinates
func FogCoordPointer(xtype uint32, stride int32, pointer unsafe.Pointer) {
	C.glowFogCoordPointer(gpFogCoordPointer, (C.GLenum)(xtype), (C.GLsizei)(stride), pointer)
}
func FogCoordd(coord float64) {
	C.glowFogCoordd(gpFogCoordd, (C.GLdouble)(coord))
}
func FogCoorddv(coord *float64) {
	C.glowFogCoorddv(gpFogCoorddv, (*C.GLdouble)(unsafe.Pointer(coord)))
}
func FogCoordf(coord float32) {
	C.glowFogCoordf(gpFogCoordf, (C.GLfloat)(coord))
}
func FogCoordfv(coord *float32) {
	C.glowFogCoordfv(gpFogCoordfv, (*C.GLfloat)(unsafe.Pointer(coord)))
}
func Fogf(pname uint32, param float32) {
	C.glowFogf(gpFogf, (C.GLenum)(pname), (C.GLfloat)(param))
}
func Fogfv(pname uint32, params *float32) {
	C.glowFogfv(gpFogfv, (C.GLenum)(pname), (*C.GLfloat)(unsafe.Pointer(params)))
}
func Fogi(pname uint32, param int32) {
	C.glowFogi(gpFogi, (C.GLenum)(pname), (C.GLint)(param))
}
func Fogiv(pname uint32, params *int32) {
	C.glowFogiv(gpFogiv, (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}

// set a named parameter of a framebuffer object
func FramebufferParameteri(target uint32, pname uint32, param int32) {
	C.glowFramebufferParameteri(gpFramebufferParameteri, (C.GLenum)(target), (C.GLenum)(pname), (C.GLint)(param))
}

// attach a renderbuffer as a logical buffer of a framebuffer object
func FramebufferRenderbuffer(target uint32, attachment uint32, renderbuffertarget uint32, renderbuffer uint32) {
	C.glowFramebufferRenderbuffer(gpFramebufferRenderbuffer, (C.GLenum)(target), (C.GLenum)(attachment), (C.GLenum)(renderbuffertarget), (C.GLuint)(renderbuffer))
}

// attach a level of a texture object as a logical buffer of a framebuffer object
func FramebufferTexture(target uint32, attachment uint32, texture uint32, level int32) {
	C.glowFramebufferTexture(gpFramebufferTexture, (C.GLenum)(target), (C.GLenum)(attachment), (C.GLuint)(texture), (C.GLint)(level))
}
func FramebufferTexture1D(target uint32, attachment uint32, textarget uint32, texture uint32, level int32) {
	C.glowFramebufferTexture1D(gpFramebufferTexture1D, (C.GLenum)(target), (C.GLenum)(attachment), (C.GLenum)(textarget), (C.GLuint)(texture), (C.GLint)(level))
}
func FramebufferTexture2D(target uint32, attachment uint32, textarget uint32, texture uint32, level int32) {
	C.glowFramebufferTexture2D(gpFramebufferTexture2D, (C.GLenum)(target), (C.GLenum)(attachment), (C.GLenum)(textarget), (C.GLuint)(texture), (C.GLint)(level))
}
func FramebufferTexture3D(target uint32, attachment uint32, textarget uint32, texture uint32, level int32, zoffset int32) {
	C.glowFramebufferTexture3D(gpFramebufferTexture3D, (C.GLenum)(target), (C.GLenum)(attachment), (C.GLenum)(textarget), (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(zoffset))
}

// attach a single layer of a texture object as a logical buffer of a framebuffer object
func FramebufferTextureLayer(target uint32, attachment uint32, texture uint32, level int32, layer int32) {
	C.glowFramebufferTextureLayer(gpFramebufferTextureLayer, (C.GLenum)(target), (C.GLenum)(attachment), (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(layer))
}

// define front- and back-facing polygons
func FrontFace(mode uint32) {
	C.glowFrontFace(gpFrontFace, (C.GLenum)(mode))
}

// multiply the current matrix by a perspective matrix
func Frustum(left float64, right float64, bottom float64, top float64, zNear float64, zFar float64) {
	C.glowFrustum(gpFrustum, (C.GLdouble)(left), (C.GLdouble)(right), (C.GLdouble)(bottom), (C.GLdouble)(top), (C.GLdouble)(zNear), (C.GLdouble)(zFar))
}

// generate buffer object names
func GenBuffers(n int32, buffers *uint32) {
	C.glowGenBuffers(gpGenBuffers, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(buffers)))
}

// generate framebuffer object names
func GenFramebuffers(n int32, framebuffers *uint32) {
	C.glowGenFramebuffers(gpGenFramebuffers, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(framebuffers)))
}

// generate a contiguous set of empty display lists
func GenLists(xrange int32) uint32 {
	ret := C.glowGenLists(gpGenLists, (C.GLsizei)(xrange))
	return (uint32)(ret)
}

// reserve program pipeline object names
func GenProgramPipelines(n int32, pipelines *uint32) {
	C.glowGenProgramPipelines(gpGenProgramPipelines, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(pipelines)))
}

// generate query object names
func GenQueries(n int32, ids *uint32) {
	C.glowGenQueries(gpGenQueries, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(ids)))
}

// generate renderbuffer object names
func GenRenderbuffers(n int32, renderbuffers *uint32) {
	C.glowGenRenderbuffers(gpGenRenderbuffers, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(renderbuffers)))
}

// generate sampler object names
func GenSamplers(count int32, samplers *uint32) {
	C.glowGenSamplers(gpGenSamplers, (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(samplers)))
}

// generate texture names
func GenTextures(n int32, textures *uint32) {
	C.glowGenTextures(gpGenTextures, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(textures)))
}

// reserve transform feedback object names
func GenTransformFeedbacks(n int32, ids *uint32) {
	C.glowGenTransformFeedbacks(gpGenTransformFeedbacks, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(ids)))
}

// generate vertex array object names
func GenVertexArrays(n int32, arrays *uint32) {
	C.glowGenVertexArrays(gpGenVertexArrays, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(arrays)))
}

// generate mipmaps for a specified texture object
func GenerateMipmap(target uint32) {
	C.glowGenerateMipmap(gpGenerateMipmap, (C.GLenum)(target))
}

// generate mipmaps for a specified texture object
func GenerateTextureMipmap(texture uint32) {
	C.glowGenerateTextureMipmap(gpGenerateTextureMipmap, (C.GLuint)(texture))
}

// retrieve information about the set of active atomic counter buffers for a program
func GetActiveAtomicCounterBufferiv(program uint32, bufferIndex uint32, pname uint32, params *int32) {
	C.glowGetActiveAtomicCounterBufferiv(gpGetActiveAtomicCounterBufferiv, (C.GLuint)(program), (C.GLuint)(bufferIndex), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}

// Returns information about an active attribute variable for the specified program object
func GetActiveAttrib(program uint32, index uint32, bufSize int32, length *int32, size *int32, xtype *uint32, name *uint8) {
	C.glowGetActiveAttrib(gpGetActiveAttrib, (C.GLuint)(program), (C.GLuint)(index), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLint)(unsafe.Pointer(size)), (*C.GLenum)(unsafe.Pointer(xtype)), (*C.GLchar)(unsafe.Pointer(name)))
}

// query the name of an active shader subroutine
func GetActiveSubroutineName(program uint32, shadertype uint32, index uint32, bufsize int32, length *int32, name *uint8) {
	C.glowGetActiveSubroutineName(gpGetActiveSubroutineName, (C.GLuint)(program), (C.GLenum)(shadertype), (C.GLuint)(index), (C.GLsizei)(bufsize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(name)))
}

// query the name of an active shader subroutine uniform
func GetActiveSubroutineUniformName(program uint32, shadertype uint32, index uint32, bufsize int32, length *int32, name *uint8) {
	C.glowGetActiveSubroutineUniformName(gpGetActiveSubroutineUniformName, (C.GLuint)(program), (C.GLenum)(shadertype), (C.GLuint)(index), (C.GLsizei)(bufsize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(name)))
}
func GetActiveSubroutineUniformiv(program uint32, shadertype uint32, index uint32, pname uint32, values *int32) {
	C.glowGetActiveSubroutineUniformiv(gpGetActiveSubroutineUniformiv, (C.GLuint)(program), (C.GLenum)(shadertype), (C.GLuint)(index), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(values)))
}

// Returns information about an active uniform variable for the specified program object
func GetActiveUniform(program uint32, index uint32, bufSize int32, length *int32, size *int32, xtype *uint32, name *uint8) {
	C.glowGetActiveUniform(gpGetActiveUniform, (C.GLuint)(program), (C.GLuint)(index), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLint)(unsafe.Pointer(size)), (*C.GLenum)(unsafe.Pointer(xtype)), (*C.GLchar)(unsafe.Pointer(name)))
}

// retrieve the name of an active uniform block
func GetActiveUniformBlockName(program uint32, uniformBlockIndex uint32, bufSize int32, length *int32, uniformBlockName *uint8) {
	C.glowGetActiveUniformBlockName(gpGetActiveUniformBlockName, (C.GLuint)(program), (C.GLuint)(uniformBlockIndex), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(uniformBlockName)))
}
func GetActiveUniformBlockiv(program uint32, uniformBlockIndex uint32, pname uint32, params *int32) {
	C.glowGetActiveUniformBlockiv(gpGetActiveUniformBlockiv, (C.GLuint)(program), (C.GLuint)(uniformBlockIndex), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}

// query the name of an active uniform
func GetActiveUniformName(program uint32, uniformIndex uint32, bufSize int32, length *int32, uniformName *uint8) {
	C.glowGetActiveUniformName(gpGetActiveUniformName, (C.GLuint)(program), (C.GLuint)(uniformIndex), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(uniformName)))
}

// Returns information about several active uniform variables for the specified program object
func GetActiveUniformsiv(program uint32, uniformCount int32, uniformIndices *uint32, pname uint32, params *int32) {
	C.glowGetActiveUniformsiv(gpGetActiveUniformsiv, (C.GLuint)(program), (C.GLsizei)(uniformCount), (*C.GLuint)(unsafe.Pointer(uniformIndices)), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}

// Returns the handles of the shader objects attached to a program object
func GetAttachedShaders(program uint32, maxCount int32, count *int32, shaders *uint32) {
	C.glowGetAttachedShaders(gpGetAttachedShaders, (C.GLuint)(program), (C.GLsizei)(maxCount), (*C.GLsizei)(unsafe.Pointer(count)), (*C.GLuint)(unsafe.Pointer(shaders)))
}

// Returns the location of an attribute variable
func GetAttribLocation(program uint32, name *uint8) int32 {
	ret := C.glowGetAttribLocation(gpGetAttribLocation, (C.GLuint)(program), (*C.GLchar)(unsafe.Pointer(name)))
	return (int32)(ret)
}
func GetBooleani_v(target uint32, index uint32, data *bool) {
	C.glowGetBooleani_v(gpGetBooleani_v, (C.GLenum)(target), (C.GLuint)(index), (*C.GLboolean)(unsafe.Pointer(data)))
}
func GetBooleanv(pname uint32, data *bool) {
	C.glowGetBooleanv(gpGetBooleanv, (C.GLenum)(pname), (*C.GLboolean)(unsafe.Pointer(data)))
}

// return parameters of a buffer object
func GetBufferParameteri64v(target uint32, pname uint32, params *int64) {
	C.glowGetBufferParameteri64v(gpGetBufferParameteri64v, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLint64)(unsafe.Pointer(params)))
}

// return parameters of a buffer object
func GetBufferParameteriv(target uint32, pname uint32, params *int32) {
	C.glowGetBufferParameteriv(gpGetBufferParameteriv, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}

// return the pointer to a mapped buffer object's data store
func GetBufferPointerv(target uint32, pname uint32, params *unsafe.Pointer) {
	C.glowGetBufferPointerv(gpGetBufferPointerv, (C.GLenum)(target), (C.GLenum)(pname), params)
}

// returns a subset of a buffer object's data store
func GetBufferSubData(target uint32, offset int, size int, data unsafe.Pointer) {
	C.glowGetBufferSubData(gpGetBufferSubData, (C.GLenum)(target), (C.GLintptr)(offset), (C.GLsizeiptr)(size), data)
}

// return the coefficients of the specified clipping plane
func GetClipPlane(plane uint32, equation *float64) {
	C.glowGetClipPlane(gpGetClipPlane, (C.GLenum)(plane), (*C.GLdouble)(unsafe.Pointer(equation)))
}

// return a compressed texture image
func GetCompressedTexImage(target uint32, level int32, img unsafe.Pointer) {
	C.glowGetCompressedTexImage(gpGetCompressedTexImage, (C.GLenum)(target), (C.GLint)(level), img)
}

// return a compressed texture image
func GetCompressedTextureImage(texture uint32, level int32, bufSize int32, pixels unsafe.Pointer) {
	C.glowGetCompressedTextureImage(gpGetCompressedTextureImage, (C.GLuint)(texture), (C.GLint)(level), (C.GLsizei)(bufSize), pixels)
}

// retrieve a sub-region of a compressed texture image from a     compressed texture object
func GetCompressedTextureSubImage(texture uint32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, bufSize int32, pixels unsafe.Pointer) {
	C.glowGetCompressedTextureSubImage(gpGetCompressedTextureSubImage, (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLsizei)(bufSize), pixels)
}

// retrieve messages from the debug message log
func GetDebugMessageLog(count uint32, bufSize int32, sources *uint32, types *uint32, ids *uint32, severities *uint32, lengths *int32, messageLog *uint8) uint32 {
	ret := C.glowGetDebugMessageLog(gpGetDebugMessageLog, (C.GLuint)(count), (C.GLsizei)(bufSize), (*C.GLenum)(unsafe.Pointer(sources)), (*C.GLenum)(unsafe.Pointer(types)), (*C.GLuint)(unsafe.Pointer(ids)), (*C.GLenum)(unsafe.Pointer(severities)), (*C.GLsizei)(unsafe.Pointer(lengths)), (*C.GLchar)(unsafe.Pointer(messageLog)))
	return (uint32)(ret)
}
func GetDebugMessageLogARB(count uint32, bufSize int32, sources *uint32, types *uint32, ids *uint32, severities *uint32, lengths *int32, messageLog *uint8) uint32 {
	ret := C.glowGetDebugMessageLogARB(gpGetDebugMessageLogARB, (C.GLuint)(count), (C.GLsizei)(bufSize), (*C.GLenum)(unsafe.Pointer(sources)), (*C.GLenum)(unsafe.Pointer(types)), (*C.GLuint)(unsafe.Pointer(ids)), (*C.GLenum)(unsafe.Pointer(severities)), (*C.GLsizei)(unsafe.Pointer(lengths)), (*C.GLchar)(unsafe.Pointer(messageLog)))
	return (uint32)(ret)
}
func GetDebugMessageLogKHR(count uint32, bufSize int32, sources *uint32, types *uint32, ids *uint32, severities *uint32, lengths *int32, messageLog *uint8) uint32 {
	ret := C.glowGetDebugMessageLogKHR(gpGetDebugMessageLogKHR, (C.GLuint)(count), (C.GLsizei)(bufSize), (*C.GLenum)(unsafe.Pointer(sources)), (*C.GLenum)(unsafe.Pointer(types)), (*C.GLuint)(unsafe.Pointer(ids)), (*C.GLenum)(unsafe.Pointer(severities)), (*C.GLsizei)(unsafe.Pointer(lengths)), (*C.GLchar)(unsafe.Pointer(messageLog)))
	return (uint32)(ret)
}
func GetDoublei_v(target uint32, index uint32, data *float64) {
	C.glowGetDoublei_v(gpGetDoublei_v, (C.GLenum)(target), (C.GLuint)(index), (*C.GLdouble)(unsafe.Pointer(data)))
}
func GetDoublev(pname uint32, data *float64) {
	C.glowGetDoublev(gpGetDoublev, (C.GLenum)(pname), (*C.GLdouble)(unsafe.Pointer(data)))
}

// return error information
func GetError() uint32 {
	ret := C.glowGetError(gpGetError)
	return (uint32)(ret)
}
func GetFloati_v(target uint32, index uint32, data *float32) {
	C.glowGetFloati_v(gpGetFloati_v, (C.GLenum)(target), (C.GLuint)(index), (*C.GLfloat)(unsafe.Pointer(data)))
}
func GetFloatv(pname uint32, data *float32) {
	C.glowGetFloatv(gpGetFloatv, (C.GLenum)(pname), (*C.GLfloat)(unsafe.Pointer(data)))
}

// query the bindings of color indices to user-defined varying out variables
func GetFragDataIndex(program uint32, name *uint8) int32 {
	ret := C.glowGetFragDataIndex(gpGetFragDataIndex, (C.GLuint)(program), (*C.GLchar)(unsafe.Pointer(name)))
	return (int32)(ret)
}

// query the bindings of color numbers to user-defined varying out variables
func GetFragDataLocation(program uint32, name *uint8) int32 {
	ret := C.glowGetFragDataLocation(gpGetFragDataLocation, (C.GLuint)(program), (*C.GLchar)(unsafe.Pointer(name)))
	return (int32)(ret)
}

// retrieve information about attachments of a framebuffer object
func GetFramebufferAttachmentParameteriv(target uint32, attachment uint32, pname uint32, params *int32) {
	C.glowGetFramebufferAttachmentParameteriv(gpGetFramebufferAttachmentParameteriv, (C.GLenum)(target), (C.GLenum)(attachment), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}

// query a named parameter of a framebuffer object
func GetFramebufferParameteriv(target uint32, pname uint32, params *int32) {
	C.glowGetFramebufferParameteriv(gpGetFramebufferParameteriv, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}

// check if the rendering context has not been lost due to software or hardware issues
func GetGraphicsResetStatus() uint32 {
	ret := C.glowGetGraphicsResetStatus(gpGetGraphicsResetStatus)
	return (uint32)(ret)
}
func GetGraphicsResetStatusARB() uint32 {
	ret := C.glowGetGraphicsResetStatusARB(gpGetGraphicsResetStatusARB)
	return (uint32)(ret)
}
func GetGraphicsResetStatusKHR() uint32 {
	ret := C.glowGetGraphicsResetStatusKHR(gpGetGraphicsResetStatusKHR)
	return (uint32)(ret)
}
func GetImageHandleARB(texture uint32, level int32, layered bool, layer int32, format uint32) uint64 {
	ret := C.glowGetImageHandleARB(gpGetImageHandleARB, (C.GLuint)(texture), (C.GLint)(level), (C.GLboolean)(boolToInt(layered)), (C.GLint)(layer), (C.GLenum)(format))
	return (uint64)(ret)
}
func GetInteger64i_v(target uint32, index uint32, data *int64) {
	C.glowGetInteger64i_v(gpGetInteger64i_v, (C.GLenum)(target), (C.GLuint)(index), (*C.GLint64)(unsafe.Pointer(data)))
}
func GetInteger64v(pname uint32, data *int64) {
	C.glowGetInteger64v(gpGetInteger64v, (C.GLenum)(pname), (*C.GLint64)(unsafe.Pointer(data)))
}
func GetIntegeri_v(target uint32, index uint32, data *int32) {
	C.glowGetIntegeri_v(gpGetIntegeri_v, (C.GLenum)(target), (C.GLuint)(index), (*C.GLint)(unsafe.Pointer(data)))
}
func GetIntegerv(pname uint32, data *int32) {
	C.glowGetIntegerv(gpGetIntegerv, (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(data)))
}
func GetInternalformati64v(target uint32, internalformat uint32, pname uint32, bufSize int32, params *int64) {
	C.glowGetInternalformati64v(gpGetInternalformati64v, (C.GLenum)(target), (C.GLenum)(internalformat), (C.GLenum)(pname), (C.GLsizei)(bufSize), (*C.GLint64)(unsafe.Pointer(params)))
}
func GetInternalformativ(target uint32, internalformat uint32, pname uint32, bufSize int32, params *int32) {
	C.glowGetInternalformativ(gpGetInternalformativ, (C.GLenum)(target), (C.GLenum)(internalformat), (C.GLenum)(pname), (C.GLsizei)(bufSize), (*C.GLint)(unsafe.Pointer(params)))
}
func GetLightfv(light uint32, pname uint32, params *float32) {
	C.glowGetLightfv(gpGetLightfv, (C.GLenum)(light), (C.GLenum)(pname), (*C.GLfloat)(unsafe.Pointer(params)))
}
func GetLightiv(light uint32, pname uint32, params *int32) {
	C.glowGetLightiv(gpGetLightiv, (C.GLenum)(light), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetMapdv(target uint32, query uint32, v *float64) {
	C.glowGetMapdv(gpGetMapdv, (C.GLenum)(target), (C.GLenum)(query), (*C.GLdouble)(unsafe.Pointer(v)))
}
func GetMapfv(target uint32, query uint32, v *float32) {
	C.glowGetMapfv(gpGetMapfv, (C.GLenum)(target), (C.GLenum)(query), (*C.GLfloat)(unsafe.Pointer(v)))
}
func GetMapiv(target uint32, query uint32, v *int32) {
	C.glowGetMapiv(gpGetMapiv, (C.GLenum)(target), (C.GLenum)(query), (*C.GLint)(unsafe.Pointer(v)))
}
func GetMaterialfv(face uint32, pname uint32, params *float32) {
	C.glowGetMaterialfv(gpGetMaterialfv, (C.GLenum)(face), (C.GLenum)(pname), (*C.GLfloat)(unsafe.Pointer(params)))
}
func GetMaterialiv(face uint32, pname uint32, params *int32) {
	C.glowGetMaterialiv(gpGetMaterialiv, (C.GLenum)(face), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}

// retrieve the location of a sample
func GetMultisamplefv(pname uint32, index uint32, val *float32) {
	C.glowGetMultisamplefv(gpGetMultisamplefv, (C.GLenum)(pname), (C.GLuint)(index), (*C.GLfloat)(unsafe.Pointer(val)))
}

// return parameters of a buffer object
func GetNamedBufferParameteri64v(buffer uint32, pname uint32, params *int64) {
	C.glowGetNamedBufferParameteri64v(gpGetNamedBufferParameteri64v, (C.GLuint)(buffer), (C.GLenum)(pname), (*C.GLint64)(unsafe.Pointer(params)))
}

// return parameters of a buffer object
func GetNamedBufferParameteriv(buffer uint32, pname uint32, params *int32) {
	C.glowGetNamedBufferParameteriv(gpGetNamedBufferParameteriv, (C.GLuint)(buffer), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}

// return the pointer to a mapped buffer object's data store
func GetNamedBufferPointerv(buffer uint32, pname uint32, params *unsafe.Pointer) {
	C.glowGetNamedBufferPointerv(gpGetNamedBufferPointerv, (C.GLuint)(buffer), (C.GLenum)(pname), params)
}

// returns a subset of a buffer object's data store
func GetNamedBufferSubData(buffer uint32, offset int, size int32, data unsafe.Pointer) {
	C.glowGetNamedBufferSubData(gpGetNamedBufferSubData, (C.GLuint)(buffer), (C.GLintptr)(offset), (C.GLsizei)(size), data)
}

// retrieve information about attachments of a framebuffer object
func GetNamedFramebufferAttachmentParameteriv(framebuffer uint32, attachment uint32, pname uint32, params *int32) {
	C.glowGetNamedFramebufferAttachmentParameteriv(gpGetNamedFramebufferAttachmentParameteriv, (C.GLuint)(framebuffer), (C.GLenum)(attachment), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}

// query a named parameter of a framebuffer object
func GetNamedFramebufferParameteriv(framebuffer uint32, pname uint32, param *int32) {
	C.glowGetNamedFramebufferParameteriv(gpGetNamedFramebufferParameteriv, (C.GLuint)(framebuffer), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(param)))
}

// query a named parameter of a renderbuffer object
func GetNamedRenderbufferParameteriv(renderbuffer uint32, pname uint32, params *int32) {
	C.glowGetNamedRenderbufferParameteriv(gpGetNamedRenderbufferParameteriv, (C.GLuint)(renderbuffer), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetNamedStringARB(namelen int32, name *uint8, bufSize int32, stringlen *int32, xstring *uint8) {
	C.glowGetNamedStringARB(gpGetNamedStringARB, (C.GLint)(namelen), (*C.GLchar)(unsafe.Pointer(name)), (C.GLsizei)(bufSize), (*C.GLint)(unsafe.Pointer(stringlen)), (*C.GLchar)(unsafe.Pointer(xstring)))
}
func GetNamedStringivARB(namelen int32, name *uint8, pname uint32, params *int32) {
	C.glowGetNamedStringivARB(gpGetNamedStringivARB, (C.GLint)(namelen), (*C.GLchar)(unsafe.Pointer(name)), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}

// retrieve the label of a named object identified within a namespace
func GetObjectLabel(identifier uint32, name uint32, bufSize int32, length *int32, label *uint8) {
	C.glowGetObjectLabel(gpGetObjectLabel, (C.GLenum)(identifier), (C.GLuint)(name), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(label)))
}
func GetObjectLabelKHR(identifier uint32, name uint32, bufSize int32, length *int32, label *uint8) {
	C.glowGetObjectLabelKHR(gpGetObjectLabelKHR, (C.GLenum)(identifier), (C.GLuint)(name), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(label)))
}

// retrieve the label of a sync object identified by a pointer
func GetObjectPtrLabel(ptr unsafe.Pointer, bufSize int32, length *int32, label *uint8) {
	C.glowGetObjectPtrLabel(gpGetObjectPtrLabel, ptr, (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(label)))
}
func GetObjectPtrLabelKHR(ptr unsafe.Pointer, bufSize int32, length *int32, label *uint8) {
	C.glowGetObjectPtrLabelKHR(gpGetObjectPtrLabelKHR, ptr, (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(label)))
}
func GetPixelMapfv(xmap uint32, values *float32) {
	C.glowGetPixelMapfv(gpGetPixelMapfv, (C.GLenum)(xmap), (*C.GLfloat)(unsafe.Pointer(values)))
}
func GetPixelMapuiv(xmap uint32, values *uint32) {
	C.glowGetPixelMapuiv(gpGetPixelMapuiv, (C.GLenum)(xmap), (*C.GLuint)(unsafe.Pointer(values)))
}
func GetPixelMapusv(xmap uint32, values *uint16) {
	C.glowGetPixelMapusv(gpGetPixelMapusv, (C.GLenum)(xmap), (*C.GLushort)(unsafe.Pointer(values)))
}

// return the address of the specified pointer
func GetPointerv(pname uint32, params *unsafe.Pointer) {
	C.glowGetPointerv(gpGetPointerv, (C.GLenum)(pname), params)
}
func GetPointervKHR(pname uint32, params *unsafe.Pointer) {
	C.glowGetPointervKHR(gpGetPointervKHR, (C.GLenum)(pname), params)
}

// return the polygon stipple pattern
func GetPolygonStipple(mask *uint8) {
	C.glowGetPolygonStipple(gpGetPolygonStipple, (*C.GLubyte)(unsafe.Pointer(mask)))
}

// return a binary representation of a program object's compiled and linked executable source
func GetProgramBinary(program uint32, bufSize int32, length *int32, binaryFormat *uint32, binary unsafe.Pointer) {
	C.glowGetProgramBinary(gpGetProgramBinary, (C.GLuint)(program), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLenum)(unsafe.Pointer(binaryFormat)), binary)
}

// Returns the information log for a program object
func GetProgramInfoLog(program uint32, bufSize int32, length *int32, infoLog *uint8) {
	C.glowGetProgramInfoLog(gpGetProgramInfoLog, (C.GLuint)(program), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(infoLog)))
}
func GetProgramInterfaceiv(program uint32, programInterface uint32, pname uint32, params *int32) {
	C.glowGetProgramInterfaceiv(gpGetProgramInterfaceiv, (C.GLuint)(program), (C.GLenum)(programInterface), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}

// retrieve the info log string from a program pipeline object
func GetProgramPipelineInfoLog(pipeline uint32, bufSize int32, length *int32, infoLog *uint8) {
	C.glowGetProgramPipelineInfoLog(gpGetProgramPipelineInfoLog, (C.GLuint)(pipeline), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(infoLog)))
}
func GetProgramPipelineiv(pipeline uint32, pname uint32, params *int32) {
	C.glowGetProgramPipelineiv(gpGetProgramPipelineiv, (C.GLuint)(pipeline), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}

// query the index of a named resource within a program
func GetProgramResourceIndex(program uint32, programInterface uint32, name *uint8) uint32 {
	ret := C.glowGetProgramResourceIndex(gpGetProgramResourceIndex, (C.GLuint)(program), (C.GLenum)(programInterface), (*C.GLchar)(unsafe.Pointer(name)))
	return (uint32)(ret)
}

// query the location of a named resource within a program
func GetProgramResourceLocation(program uint32, programInterface uint32, name *uint8) int32 {
	ret := C.glowGetProgramResourceLocation(gpGetProgramResourceLocation, (C.GLuint)(program), (C.GLenum)(programInterface), (*C.GLchar)(unsafe.Pointer(name)))
	return (int32)(ret)
}

// query the fragment color index of a named variable within a program
func GetProgramResourceLocationIndex(program uint32, programInterface uint32, name *uint8) int32 {
	ret := C.glowGetProgramResourceLocationIndex(gpGetProgramResourceLocationIndex, (C.GLuint)(program), (C.GLenum)(programInterface), (*C.GLchar)(unsafe.Pointer(name)))
	return (int32)(ret)
}

// query the name of an indexed resource within a program
func GetProgramResourceName(program uint32, programInterface uint32, index uint32, bufSize int32, length *int32, name *uint8) {
	C.glowGetProgramResourceName(gpGetProgramResourceName, (C.GLuint)(program), (C.GLenum)(programInterface), (C.GLuint)(index), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(name)))
}
func GetProgramResourceiv(program uint32, programInterface uint32, index uint32, propCount int32, props *uint32, bufSize int32, length *int32, params *int32) {
	C.glowGetProgramResourceiv(gpGetProgramResourceiv, (C.GLuint)(program), (C.GLenum)(programInterface), (C.GLuint)(index), (C.GLsizei)(propCount), (*C.GLenum)(unsafe.Pointer(props)), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLint)(unsafe.Pointer(params)))
}
func GetProgramStageiv(program uint32, shadertype uint32, pname uint32, values *int32) {
	C.glowGetProgramStageiv(gpGetProgramStageiv, (C.GLuint)(program), (C.GLenum)(shadertype), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(values)))
}

// Returns a parameter from a program object
func GetProgramiv(program uint32, pname uint32, params *int32) {
	C.glowGetProgramiv(gpGetProgramiv, (C.GLuint)(program), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}

// return parameters of an indexed query object target
func GetQueryIndexediv(target uint32, index uint32, pname uint32, params *int32) {
	C.glowGetQueryIndexediv(gpGetQueryIndexediv, (C.GLenum)(target), (C.GLuint)(index), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetQueryObjecti64v(id uint32, pname uint32, params *int64) {
	C.glowGetQueryObjecti64v(gpGetQueryObjecti64v, (C.GLuint)(id), (C.GLenum)(pname), (*C.GLint64)(unsafe.Pointer(params)))
}
func GetQueryObjectiv(id uint32, pname uint32, params *int32) {
	C.glowGetQueryObjectiv(gpGetQueryObjectiv, (C.GLuint)(id), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetQueryObjectui64v(id uint32, pname uint32, params *uint64) {
	C.glowGetQueryObjectui64v(gpGetQueryObjectui64v, (C.GLuint)(id), (C.GLenum)(pname), (*C.GLuint64)(unsafe.Pointer(params)))
}
func GetQueryObjectuiv(id uint32, pname uint32, params *uint32) {
	C.glowGetQueryObjectuiv(gpGetQueryObjectuiv, (C.GLuint)(id), (C.GLenum)(pname), (*C.GLuint)(unsafe.Pointer(params)))
}

// return parameters of a query object target
func GetQueryiv(target uint32, pname uint32, params *int32) {
	C.glowGetQueryiv(gpGetQueryiv, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}

// query a named parameter of a renderbuffer object
func GetRenderbufferParameteriv(target uint32, pname uint32, params *int32) {
	C.glowGetRenderbufferParameteriv(gpGetRenderbufferParameteriv, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetSamplerParameterIiv(sampler uint32, pname uint32, params *int32) {
	C.glowGetSamplerParameterIiv(gpGetSamplerParameterIiv, (C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetSamplerParameterIuiv(sampler uint32, pname uint32, params *uint32) {
	C.glowGetSamplerParameterIuiv(gpGetSamplerParameterIuiv, (C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLuint)(unsafe.Pointer(params)))
}
func GetSamplerParameterfv(sampler uint32, pname uint32, params *float32) {
	C.glowGetSamplerParameterfv(gpGetSamplerParameterfv, (C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLfloat)(unsafe.Pointer(params)))
}
func GetSamplerParameteriv(sampler uint32, pname uint32, params *int32) {
	C.glowGetSamplerParameteriv(gpGetSamplerParameteriv, (C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}

// Returns the information log for a shader object
func GetShaderInfoLog(shader uint32, bufSize int32, length *int32, infoLog *uint8) {
	C.glowGetShaderInfoLog(gpGetShaderInfoLog, (C.GLuint)(shader), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(infoLog)))
}

// retrieve the range and precision for numeric formats supported by the shader compiler
func GetShaderPrecisionFormat(shadertype uint32, precisiontype uint32, xrange *int32, precision *int32) {
	C.glowGetShaderPrecisionFormat(gpGetShaderPrecisionFormat, (C.GLenum)(shadertype), (C.GLenum)(precisiontype), (*C.GLint)(unsafe.Pointer(xrange)), (*C.GLint)(unsafe.Pointer(precision)))
}

// Returns the source code string from a shader object
func GetShaderSource(shader uint32, bufSize int32, length *int32, source *uint8) {
	C.glowGetShaderSource(gpGetShaderSource, (C.GLuint)(shader), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(source)))
}

// Returns a parameter from a shader object
func GetShaderiv(shader uint32, pname uint32, params *int32) {
	C.glowGetShaderiv(gpGetShaderiv, (C.GLuint)(shader), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}

// return a string describing the current GL connection
func GetString(name uint32) *uint8 {
	ret := C.glowGetString(gpGetString, (C.GLenum)(name))
	return (*uint8)(ret)
}
func GetStringi(name uint32, index uint32) *uint8 {
	ret := C.glowGetStringi(gpGetStringi, (C.GLenum)(name), (C.GLuint)(index))
	return (*uint8)(ret)
}

// retrieve the index of a subroutine uniform of a given shader stage within a program
func GetSubroutineIndex(program uint32, shadertype uint32, name *uint8) uint32 {
	ret := C.glowGetSubroutineIndex(gpGetSubroutineIndex, (C.GLuint)(program), (C.GLenum)(shadertype), (*C.GLchar)(unsafe.Pointer(name)))
	return (uint32)(ret)
}

// retrieve the location of a subroutine uniform of a given shader stage within a program
func GetSubroutineUniformLocation(program uint32, shadertype uint32, name *uint8) int32 {
	ret := C.glowGetSubroutineUniformLocation(gpGetSubroutineUniformLocation, (C.GLuint)(program), (C.GLenum)(shadertype), (*C.GLchar)(unsafe.Pointer(name)))
	return (int32)(ret)
}

// query the properties of a sync object
func GetSynciv(sync uintptr, pname uint32, bufSize int32, length *int32, values *int32) {
	C.glowGetSynciv(gpGetSynciv, (C.GLsync)(sync), (C.GLenum)(pname), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLint)(unsafe.Pointer(values)))
}
func GetTexEnvfv(target uint32, pname uint32, params *float32) {
	C.glowGetTexEnvfv(gpGetTexEnvfv, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLfloat)(unsafe.Pointer(params)))
}
func GetTexEnviv(target uint32, pname uint32, params *int32) {
	C.glowGetTexEnviv(gpGetTexEnviv, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetTexGendv(coord uint32, pname uint32, params *float64) {
	C.glowGetTexGendv(gpGetTexGendv, (C.GLenum)(coord), (C.GLenum)(pname), (*C.GLdouble)(unsafe.Pointer(params)))
}
func GetTexGenfv(coord uint32, pname uint32, params *float32) {
	C.glowGetTexGenfv(gpGetTexGenfv, (C.GLenum)(coord), (C.GLenum)(pname), (*C.GLfloat)(unsafe.Pointer(params)))
}
func GetTexGeniv(coord uint32, pname uint32, params *int32) {
	C.glowGetTexGeniv(gpGetTexGeniv, (C.GLenum)(coord), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}

// return a texture image
func GetTexImage(target uint32, level int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	C.glowGetTexImage(gpGetTexImage, (C.GLenum)(target), (C.GLint)(level), (C.GLenum)(format), (C.GLenum)(xtype), pixels)
}
func GetTexLevelParameterfv(target uint32, level int32, pname uint32, params *float32) {
	C.glowGetTexLevelParameterfv(gpGetTexLevelParameterfv, (C.GLenum)(target), (C.GLint)(level), (C.GLenum)(pname), (*C.GLfloat)(unsafe.Pointer(params)))
}
func GetTexLevelParameteriv(target uint32, level int32, pname uint32, params *int32) {
	C.glowGetTexLevelParameteriv(gpGetTexLevelParameteriv, (C.GLenum)(target), (C.GLint)(level), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetTexParameterIiv(target uint32, pname uint32, params *int32) {
	C.glowGetTexParameterIiv(gpGetTexParameterIiv, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetTexParameterIuiv(target uint32, pname uint32, params *uint32) {
	C.glowGetTexParameterIuiv(gpGetTexParameterIuiv, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLuint)(unsafe.Pointer(params)))
}
func GetTexParameterfv(target uint32, pname uint32, params *float32) {
	C.glowGetTexParameterfv(gpGetTexParameterfv, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLfloat)(unsafe.Pointer(params)))
}
func GetTexParameteriv(target uint32, pname uint32, params *int32) {
	C.glowGetTexParameteriv(gpGetTexParameteriv, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetTextureHandleARB(texture uint32) uint64 {
	ret := C.glowGetTextureHandleARB(gpGetTextureHandleARB, (C.GLuint)(texture))
	return (uint64)(ret)
}

// return a texture image
func GetTextureImage(texture uint32, level int32, format uint32, xtype uint32, bufSize int32, pixels unsafe.Pointer) {
	C.glowGetTextureImage(gpGetTextureImage, (C.GLuint)(texture), (C.GLint)(level), (C.GLenum)(format), (C.GLenum)(xtype), (C.GLsizei)(bufSize), pixels)
}
func GetTextureLevelParameterfv(texture uint32, level int32, pname uint32, params *float32) {
	C.glowGetTextureLevelParameterfv(gpGetTextureLevelParameterfv, (C.GLuint)(texture), (C.GLint)(level), (C.GLenum)(pname), (*C.GLfloat)(unsafe.Pointer(params)))
}
func GetTextureLevelParameteriv(texture uint32, level int32, pname uint32, params *int32) {
	C.glowGetTextureLevelParameteriv(gpGetTextureLevelParameteriv, (C.GLuint)(texture), (C.GLint)(level), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetTextureParameterIiv(texture uint32, pname uint32, params *int32) {
	C.glowGetTextureParameterIiv(gpGetTextureParameterIiv, (C.GLuint)(texture), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetTextureParameterIuiv(texture uint32, pname uint32, params *uint32) {
	C.glowGetTextureParameterIuiv(gpGetTextureParameterIuiv, (C.GLuint)(texture), (C.GLenum)(pname), (*C.GLuint)(unsafe.Pointer(params)))
}
func GetTextureParameterfv(texture uint32, pname uint32, params *float32) {
	C.glowGetTextureParameterfv(gpGetTextureParameterfv, (C.GLuint)(texture), (C.GLenum)(pname), (*C.GLfloat)(unsafe.Pointer(params)))
}
func GetTextureParameteriv(texture uint32, pname uint32, params *int32) {
	C.glowGetTextureParameteriv(gpGetTextureParameteriv, (C.GLuint)(texture), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetTextureSamplerHandleARB(texture uint32, sampler uint32) uint64 {
	ret := C.glowGetTextureSamplerHandleARB(gpGetTextureSamplerHandleARB, (C.GLuint)(texture), (C.GLuint)(sampler))
	return (uint64)(ret)
}

// retrieve a sub-region of a texture image from a texture     object
func GetTextureSubImage(texture uint32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, format uint32, xtype uint32, bufSize int32, pixels unsafe.Pointer) {
	C.glowGetTextureSubImage(gpGetTextureSubImage, (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLenum)(format), (C.GLenum)(xtype), (C.GLsizei)(bufSize), pixels)
}

// retrieve information about varying variables selected for transform feedback
func GetTransformFeedbackVarying(program uint32, index uint32, bufSize int32, length *int32, size *int32, xtype *uint32, name *uint8) {
	C.glowGetTransformFeedbackVarying(gpGetTransformFeedbackVarying, (C.GLuint)(program), (C.GLuint)(index), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLsizei)(unsafe.Pointer(size)), (*C.GLenum)(unsafe.Pointer(xtype)), (*C.GLchar)(unsafe.Pointer(name)))
}
func GetTransformFeedbacki64_v(xfb uint32, pname uint32, index uint32, param *int64) {
	C.glowGetTransformFeedbacki64_v(gpGetTransformFeedbacki64_v, (C.GLuint)(xfb), (C.GLenum)(pname), (C.GLuint)(index), (*C.GLint64)(unsafe.Pointer(param)))
}
func GetTransformFeedbacki_v(xfb uint32, pname uint32, index uint32, param *int32) {
	C.glowGetTransformFeedbacki_v(gpGetTransformFeedbacki_v, (C.GLuint)(xfb), (C.GLenum)(pname), (C.GLuint)(index), (*C.GLint)(unsafe.Pointer(param)))
}

// query the state of a transform feedback object.
func GetTransformFeedbackiv(xfb uint32, pname uint32, param *int32) {
	C.glowGetTransformFeedbackiv(gpGetTransformFeedbackiv, (C.GLuint)(xfb), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(param)))
}

// retrieve the index of a named uniform block
func GetUniformBlockIndex(program uint32, uniformBlockName *uint8) uint32 {
	ret := C.glowGetUniformBlockIndex(gpGetUniformBlockIndex, (C.GLuint)(program), (*C.GLchar)(unsafe.Pointer(uniformBlockName)))
	return (uint32)(ret)
}

// retrieve the index of a named uniform block
func GetUniformIndices(program uint32, uniformCount int32, uniformNames **uint8, uniformIndices *uint32) {
	C.glowGetUniformIndices(gpGetUniformIndices, (C.GLuint)(program), (C.GLsizei)(uniformCount), (**C.GLchar)(unsafe.Pointer(uniformNames)), (*C.GLuint)(unsafe.Pointer(uniformIndices)))
}

// Returns the location of a uniform variable
func GetUniformLocation(program uint32, name *uint8) int32 {
	ret := C.glowGetUniformLocation(gpGetUniformLocation, (C.GLuint)(program), (*C.GLchar)(unsafe.Pointer(name)))
	return (int32)(ret)
}
func GetUniformSubroutineuiv(shadertype uint32, location int32, params *uint32) {
	C.glowGetUniformSubroutineuiv(gpGetUniformSubroutineuiv, (C.GLenum)(shadertype), (C.GLint)(location), (*C.GLuint)(unsafe.Pointer(params)))
}
func GetUniformdv(program uint32, location int32, params *float64) {
	C.glowGetUniformdv(gpGetUniformdv, (C.GLuint)(program), (C.GLint)(location), (*C.GLdouble)(unsafe.Pointer(params)))
}

// Returns the value of a uniform variable
func GetUniformfv(program uint32, location int32, params *float32) {
	C.glowGetUniformfv(gpGetUniformfv, (C.GLuint)(program), (C.GLint)(location), (*C.GLfloat)(unsafe.Pointer(params)))
}

// Returns the value of a uniform variable
func GetUniformiv(program uint32, location int32, params *int32) {
	C.glowGetUniformiv(gpGetUniformiv, (C.GLuint)(program), (C.GLint)(location), (*C.GLint)(unsafe.Pointer(params)))
}
func GetUniformuiv(program uint32, location int32, params *uint32) {
	C.glowGetUniformuiv(gpGetUniformuiv, (C.GLuint)(program), (C.GLint)(location), (*C.GLuint)(unsafe.Pointer(params)))
}
func GetVertexArrayIndexed64iv(vaobj uint32, index uint32, pname uint32, param *int64) {
	C.glowGetVertexArrayIndexed64iv(gpGetVertexArrayIndexed64iv, (C.GLuint)(vaobj), (C.GLuint)(index), (C.GLenum)(pname), (*C.GLint64)(unsafe.Pointer(param)))
}
func GetVertexArrayIndexediv(vaobj uint32, index uint32, pname uint32, param *int32) {
	C.glowGetVertexArrayIndexediv(gpGetVertexArrayIndexediv, (C.GLuint)(vaobj), (C.GLuint)(index), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(param)))
}

// retrieve parameters of a vertex array object
func GetVertexArrayiv(vaobj uint32, pname uint32, param *int32) {
	C.glowGetVertexArrayiv(gpGetVertexArrayiv, (C.GLuint)(vaobj), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(param)))
}

// Return a generic vertex attribute parameter
func GetVertexAttribIiv(index uint32, pname uint32, params *int32) {
	C.glowGetVertexAttribIiv(gpGetVertexAttribIiv, (C.GLuint)(index), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}

// Return a generic vertex attribute parameter
func GetVertexAttribIuiv(index uint32, pname uint32, params *uint32) {
	C.glowGetVertexAttribIuiv(gpGetVertexAttribIuiv, (C.GLuint)(index), (C.GLenum)(pname), (*C.GLuint)(unsafe.Pointer(params)))
}

// Return a generic vertex attribute parameter
func GetVertexAttribLdv(index uint32, pname uint32, params *float64) {
	C.glowGetVertexAttribLdv(gpGetVertexAttribLdv, (C.GLuint)(index), (C.GLenum)(pname), (*C.GLdouble)(unsafe.Pointer(params)))
}
func GetVertexAttribLui64vARB(index uint32, pname uint32, params *uint64) {
	C.glowGetVertexAttribLui64vARB(gpGetVertexAttribLui64vARB, (C.GLuint)(index), (C.GLenum)(pname), (*C.GLuint64EXT)(unsafe.Pointer(params)))
}

// return the address of the specified generic vertex attribute pointer
func GetVertexAttribPointerv(index uint32, pname uint32, pointer *unsafe.Pointer) {
	C.glowGetVertexAttribPointerv(gpGetVertexAttribPointerv, (C.GLuint)(index), (C.GLenum)(pname), pointer)
}

// Return a generic vertex attribute parameter
func GetVertexAttribdv(index uint32, pname uint32, params *float64) {
	C.glowGetVertexAttribdv(gpGetVertexAttribdv, (C.GLuint)(index), (C.GLenum)(pname), (*C.GLdouble)(unsafe.Pointer(params)))
}

// Return a generic vertex attribute parameter
func GetVertexAttribfv(index uint32, pname uint32, params *float32) {
	C.glowGetVertexAttribfv(gpGetVertexAttribfv, (C.GLuint)(index), (C.GLenum)(pname), (*C.GLfloat)(unsafe.Pointer(params)))
}

// Return a generic vertex attribute parameter
func GetVertexAttribiv(index uint32, pname uint32, params *int32) {
	C.glowGetVertexAttribiv(gpGetVertexAttribiv, (C.GLuint)(index), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}

// return a compressed texture image
func GetnCompressedTexImage(target uint32, lod int32, bufSize int32, pixels unsafe.Pointer) {
	C.glowGetnCompressedTexImage(gpGetnCompressedTexImage, (C.GLenum)(target), (C.GLint)(lod), (C.GLsizei)(bufSize), pixels)
}
func GetnCompressedTexImageARB(target uint32, lod int32, bufSize int32, img unsafe.Pointer) {
	C.glowGetnCompressedTexImageARB(gpGetnCompressedTexImageARB, (C.GLenum)(target), (C.GLint)(lod), (C.GLsizei)(bufSize), img)
}

// return a texture image
func GetnTexImage(target uint32, level int32, format uint32, xtype uint32, bufSize int32, pixels unsafe.Pointer) {
	C.glowGetnTexImage(gpGetnTexImage, (C.GLenum)(target), (C.GLint)(level), (C.GLenum)(format), (C.GLenum)(xtype), (C.GLsizei)(bufSize), pixels)
}
func GetnTexImageARB(target uint32, level int32, format uint32, xtype uint32, bufSize int32, img unsafe.Pointer) {
	C.glowGetnTexImageARB(gpGetnTexImageARB, (C.GLenum)(target), (C.GLint)(level), (C.GLenum)(format), (C.GLenum)(xtype), (C.GLsizei)(bufSize), img)
}
func GetnUniformdv(program uint32, location int32, bufSize int32, params *float64) {
	C.glowGetnUniformdv(gpGetnUniformdv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(bufSize), (*C.GLdouble)(unsafe.Pointer(params)))
}
func GetnUniformdvARB(program uint32, location int32, bufSize int32, params *float64) {
	C.glowGetnUniformdvARB(gpGetnUniformdvARB, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(bufSize), (*C.GLdouble)(unsafe.Pointer(params)))
}
func GetnUniformfv(program uint32, location int32, bufSize int32, params *float32) {
	C.glowGetnUniformfv(gpGetnUniformfv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(bufSize), (*C.GLfloat)(unsafe.Pointer(params)))
}
func GetnUniformfvARB(program uint32, location int32, bufSize int32, params *float32) {
	C.glowGetnUniformfvARB(gpGetnUniformfvARB, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(bufSize), (*C.GLfloat)(unsafe.Pointer(params)))
}
func GetnUniformfvKHR(program uint32, location int32, bufSize int32, params *float32) {
	C.glowGetnUniformfvKHR(gpGetnUniformfvKHR, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(bufSize), (*C.GLfloat)(unsafe.Pointer(params)))
}
func GetnUniformiv(program uint32, location int32, bufSize int32, params *int32) {
	C.glowGetnUniformiv(gpGetnUniformiv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(bufSize), (*C.GLint)(unsafe.Pointer(params)))
}
func GetnUniformivARB(program uint32, location int32, bufSize int32, params *int32) {
	C.glowGetnUniformivARB(gpGetnUniformivARB, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(bufSize), (*C.GLint)(unsafe.Pointer(params)))
}
func GetnUniformivKHR(program uint32, location int32, bufSize int32, params *int32) {
	C.glowGetnUniformivKHR(gpGetnUniformivKHR, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(bufSize), (*C.GLint)(unsafe.Pointer(params)))
}
func GetnUniformuiv(program uint32, location int32, bufSize int32, params *uint32) {
	C.glowGetnUniformuiv(gpGetnUniformuiv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(bufSize), (*C.GLuint)(unsafe.Pointer(params)))
}
func GetnUniformuivARB(program uint32, location int32, bufSize int32, params *uint32) {
	C.glowGetnUniformuivARB(gpGetnUniformuivARB, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(bufSize), (*C.GLuint)(unsafe.Pointer(params)))
}
func GetnUniformuivKHR(program uint32, location int32, bufSize int32, params *uint32) {
	C.glowGetnUniformuivKHR(gpGetnUniformuivKHR, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(bufSize), (*C.GLuint)(unsafe.Pointer(params)))
}

// specify implementation-specific hints
func Hint(target uint32, mode uint32) {
	C.glowHint(gpHint, (C.GLenum)(target), (C.GLenum)(mode))
}

// control the writing of individual bits in the color index buffers
func IndexMask(mask uint32) {
	C.glowIndexMask(gpIndexMask, (C.GLuint)(mask))
}

// define an array of color indexes
func IndexPointer(xtype uint32, stride int32, pointer unsafe.Pointer) {
	C.glowIndexPointer(gpIndexPointer, (C.GLenum)(xtype), (C.GLsizei)(stride), pointer)
}
func Indexd(c float64) {
	C.glowIndexd(gpIndexd, (C.GLdouble)(c))
}
func Indexdv(c *float64) {
	C.glowIndexdv(gpIndexdv, (*C.GLdouble)(unsafe.Pointer(c)))
}
func Indexf(c float32) {
	C.glowIndexf(gpIndexf, (C.GLfloat)(c))
}
func Indexfv(c *float32) {
	C.glowIndexfv(gpIndexfv, (*C.GLfloat)(unsafe.Pointer(c)))
}
func Indexi(c int32) {
	C.glowIndexi(gpIndexi, (C.GLint)(c))
}
func Indexiv(c *int32) {
	C.glowIndexiv(gpIndexiv, (*C.GLint)(unsafe.Pointer(c)))
}
func Indexs(c int16) {
	C.glowIndexs(gpIndexs, (C.GLshort)(c))
}
func Indexsv(c *int16) {
	C.glowIndexsv(gpIndexsv, (*C.GLshort)(unsafe.Pointer(c)))
}
func Indexub(c uint8) {
	C.glowIndexub(gpIndexub, (C.GLubyte)(c))
}
func Indexubv(c *uint8) {
	C.glowIndexubv(gpIndexubv, (*C.GLubyte)(unsafe.Pointer(c)))
}

// initialize the name stack
func InitNames() {
	C.glowInitNames(gpInitNames)
}

// simultaneously specify and enable several interleaved arrays
func InterleavedArrays(format uint32, stride int32, pointer unsafe.Pointer) {
	C.glowInterleavedArrays(gpInterleavedArrays, (C.GLenum)(format), (C.GLsizei)(stride), pointer)
}

// invalidate the content of a buffer object's data store
func InvalidateBufferData(buffer uint32) {
	C.glowInvalidateBufferData(gpInvalidateBufferData, (C.GLuint)(buffer))
}

// invalidate a region of a buffer object's data store
func InvalidateBufferSubData(buffer uint32, offset int, length int) {
	C.glowInvalidateBufferSubData(gpInvalidateBufferSubData, (C.GLuint)(buffer), (C.GLintptr)(offset), (C.GLsizeiptr)(length))
}

// invalidate the content of some or all of a framebuffer's attachments
func InvalidateFramebuffer(target uint32, numAttachments int32, attachments *uint32) {
	C.glowInvalidateFramebuffer(gpInvalidateFramebuffer, (C.GLenum)(target), (C.GLsizei)(numAttachments), (*C.GLenum)(unsafe.Pointer(attachments)))
}

// invalidate the content of some or all of a framebuffer's attachments
func InvalidateNamedFramebufferData(framebuffer uint32, numAttachments int32, attachments *uint32) {
	C.glowInvalidateNamedFramebufferData(gpInvalidateNamedFramebufferData, (C.GLuint)(framebuffer), (C.GLsizei)(numAttachments), (*C.GLenum)(unsafe.Pointer(attachments)))
}

// invalidate the content of a region of some or all of a framebuffer's attachments
func InvalidateNamedFramebufferSubData(framebuffer uint32, numAttachments int32, attachments *uint32, x int32, y int32, width int32, height int32) {
	C.glowInvalidateNamedFramebufferSubData(gpInvalidateNamedFramebufferSubData, (C.GLuint)(framebuffer), (C.GLsizei)(numAttachments), (*C.GLenum)(unsafe.Pointer(attachments)), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}

// invalidate the content of a region of some or all of a framebuffer's attachments
func InvalidateSubFramebuffer(target uint32, numAttachments int32, attachments *uint32, x int32, y int32, width int32, height int32) {
	C.glowInvalidateSubFramebuffer(gpInvalidateSubFramebuffer, (C.GLenum)(target), (C.GLsizei)(numAttachments), (*C.GLenum)(unsafe.Pointer(attachments)), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}

// invalidate the entirety a texture image
func InvalidateTexImage(texture uint32, level int32) {
	C.glowInvalidateTexImage(gpInvalidateTexImage, (C.GLuint)(texture), (C.GLint)(level))
}

// invalidate a region of a texture image
func InvalidateTexSubImage(texture uint32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32) {
	C.glowInvalidateTexSubImage(gpInvalidateTexSubImage, (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth))
}

// determine if a name corresponds to a buffer object
func IsBuffer(buffer uint32) bool {
	ret := C.glowIsBuffer(gpIsBuffer, (C.GLuint)(buffer))
	return ret == TRUE
}
func IsEnabled(cap uint32) bool {
	ret := C.glowIsEnabled(gpIsEnabled, (C.GLenum)(cap))
	return ret == TRUE
}
func IsEnabledi(target uint32, index uint32) bool {
	ret := C.glowIsEnabledi(gpIsEnabledi, (C.GLenum)(target), (C.GLuint)(index))
	return ret == TRUE
}

// determine if a name corresponds to a framebuffer object
func IsFramebuffer(framebuffer uint32) bool {
	ret := C.glowIsFramebuffer(gpIsFramebuffer, (C.GLuint)(framebuffer))
	return ret == TRUE
}
func IsImageHandleResidentARB(handle uint64) bool {
	ret := C.glowIsImageHandleResidentARB(gpIsImageHandleResidentARB, (C.GLuint64)(handle))
	return ret == TRUE
}

// determine if a name corresponds to a display list
func IsList(list uint32) bool {
	ret := C.glowIsList(gpIsList, (C.GLuint)(list))
	return ret == TRUE
}
func IsNamedStringARB(namelen int32, name *uint8) bool {
	ret := C.glowIsNamedStringARB(gpIsNamedStringARB, (C.GLint)(namelen), (*C.GLchar)(unsafe.Pointer(name)))
	return ret == TRUE
}

// Determines if a name corresponds to a program object
func IsProgram(program uint32) bool {
	ret := C.glowIsProgram(gpIsProgram, (C.GLuint)(program))
	return ret == TRUE
}

// determine if a name corresponds to a program pipeline object
func IsProgramPipeline(pipeline uint32) bool {
	ret := C.glowIsProgramPipeline(gpIsProgramPipeline, (C.GLuint)(pipeline))
	return ret == TRUE
}

// determine if a name corresponds to a query object
func IsQuery(id uint32) bool {
	ret := C.glowIsQuery(gpIsQuery, (C.GLuint)(id))
	return ret == TRUE
}

// determine if a name corresponds to a renderbuffer object
func IsRenderbuffer(renderbuffer uint32) bool {
	ret := C.glowIsRenderbuffer(gpIsRenderbuffer, (C.GLuint)(renderbuffer))
	return ret == TRUE
}

// determine if a name corresponds to a sampler object
func IsSampler(sampler uint32) bool {
	ret := C.glowIsSampler(gpIsSampler, (C.GLuint)(sampler))
	return ret == TRUE
}

// Determines if a name corresponds to a shader object
func IsShader(shader uint32) bool {
	ret := C.glowIsShader(gpIsShader, (C.GLuint)(shader))
	return ret == TRUE
}

// determine if a name corresponds to a sync object
func IsSync(sync uintptr) bool {
	ret := C.glowIsSync(gpIsSync, (C.GLsync)(sync))
	return ret == TRUE
}

// determine if a name corresponds to a texture
func IsTexture(texture uint32) bool {
	ret := C.glowIsTexture(gpIsTexture, (C.GLuint)(texture))
	return ret == TRUE
}
func IsTextureHandleResidentARB(handle uint64) bool {
	ret := C.glowIsTextureHandleResidentARB(gpIsTextureHandleResidentARB, (C.GLuint64)(handle))
	return ret == TRUE
}

// determine if a name corresponds to a transform feedback object
func IsTransformFeedback(id uint32) bool {
	ret := C.glowIsTransformFeedback(gpIsTransformFeedback, (C.GLuint)(id))
	return ret == TRUE
}

// determine if a name corresponds to a vertex array object
func IsVertexArray(array uint32) bool {
	ret := C.glowIsVertexArray(gpIsVertexArray, (C.GLuint)(array))
	return ret == TRUE
}
func LightModelf(pname uint32, param float32) {
	C.glowLightModelf(gpLightModelf, (C.GLenum)(pname), (C.GLfloat)(param))
}
func LightModelfv(pname uint32, params *float32) {
	C.glowLightModelfv(gpLightModelfv, (C.GLenum)(pname), (*C.GLfloat)(unsafe.Pointer(params)))
}
func LightModeli(pname uint32, param int32) {
	C.glowLightModeli(gpLightModeli, (C.GLenum)(pname), (C.GLint)(param))
}
func LightModeliv(pname uint32, params *int32) {
	C.glowLightModeliv(gpLightModeliv, (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func Lightf(light uint32, pname uint32, param float32) {
	C.glowLightf(gpLightf, (C.GLenum)(light), (C.GLenum)(pname), (C.GLfloat)(param))
}
func Lightfv(light uint32, pname uint32, params *float32) {
	C.glowLightfv(gpLightfv, (C.GLenum)(light), (C.GLenum)(pname), (*C.GLfloat)(unsafe.Pointer(params)))
}
func Lighti(light uint32, pname uint32, param int32) {
	C.glowLighti(gpLighti, (C.GLenum)(light), (C.GLenum)(pname), (C.GLint)(param))
}
func Lightiv(light uint32, pname uint32, params *int32) {
	C.glowLightiv(gpLightiv, (C.GLenum)(light), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}

// specify the line stipple pattern
func LineStipple(factor int32, pattern uint16) {
	C.glowLineStipple(gpLineStipple, (C.GLint)(factor), (C.GLushort)(pattern))
}

// specify the width of rasterized lines
func LineWidth(width float32) {
	C.glowLineWidth(gpLineWidth, (C.GLfloat)(width))
}

// Links a program object
func LinkProgram(program uint32) {
	C.glowLinkProgram(gpLinkProgram, (C.GLuint)(program))
}

// set the display-list base for
func ListBase(base uint32) {
	C.glowListBase(gpListBase, (C.GLuint)(base))
}

// replace the current matrix with the identity matrix
func LoadIdentity() {
	C.glowLoadIdentity(gpLoadIdentity)
}
func LoadMatrixd(m *float64) {
	C.glowLoadMatrixd(gpLoadMatrixd, (*C.GLdouble)(unsafe.Pointer(m)))
}
func LoadMatrixf(m *float32) {
	C.glowLoadMatrixf(gpLoadMatrixf, (*C.GLfloat)(unsafe.Pointer(m)))
}

// load a name onto the name stack
func LoadName(name uint32) {
	C.glowLoadName(gpLoadName, (C.GLuint)(name))
}
func LoadTransposeMatrixd(m *float64) {
	C.glowLoadTransposeMatrixd(gpLoadTransposeMatrixd, (*C.GLdouble)(unsafe.Pointer(m)))
}
func LoadTransposeMatrixf(m *float32) {
	C.glowLoadTransposeMatrixf(gpLoadTransposeMatrixf, (*C.GLfloat)(unsafe.Pointer(m)))
}

// specify a logical pixel operation for rendering
func LogicOp(opcode uint32) {
	C.glowLogicOp(gpLogicOp, (C.GLenum)(opcode))
}
func MakeImageHandleNonResidentARB(handle uint64) {
	C.glowMakeImageHandleNonResidentARB(gpMakeImageHandleNonResidentARB, (C.GLuint64)(handle))
}
func MakeImageHandleResidentARB(handle uint64, access uint32) {
	C.glowMakeImageHandleResidentARB(gpMakeImageHandleResidentARB, (C.GLuint64)(handle), (C.GLenum)(access))
}
func MakeTextureHandleNonResidentARB(handle uint64) {
	C.glowMakeTextureHandleNonResidentARB(gpMakeTextureHandleNonResidentARB, (C.GLuint64)(handle))
}
func MakeTextureHandleResidentARB(handle uint64) {
	C.glowMakeTextureHandleResidentARB(gpMakeTextureHandleResidentARB, (C.GLuint64)(handle))
}
func Map1d(target uint32, u1 float64, u2 float64, stride int32, order int32, points *float64) {
	C.glowMap1d(gpMap1d, (C.GLenum)(target), (C.GLdouble)(u1), (C.GLdouble)(u2), (C.GLint)(stride), (C.GLint)(order), (*C.GLdouble)(unsafe.Pointer(points)))
}
func Map1f(target uint32, u1 float32, u2 float32, stride int32, order int32, points *float32) {
	C.glowMap1f(gpMap1f, (C.GLenum)(target), (C.GLfloat)(u1), (C.GLfloat)(u2), (C.GLint)(stride), (C.GLint)(order), (*C.GLfloat)(unsafe.Pointer(points)))
}
func Map2d(target uint32, u1 float64, u2 float64, ustride int32, uorder int32, v1 float64, v2 float64, vstride int32, vorder int32, points *float64) {
	C.glowMap2d(gpMap2d, (C.GLenum)(target), (C.GLdouble)(u1), (C.GLdouble)(u2), (C.GLint)(ustride), (C.GLint)(uorder), (C.GLdouble)(v1), (C.GLdouble)(v2), (C.GLint)(vstride), (C.GLint)(vorder), (*C.GLdouble)(unsafe.Pointer(points)))
}
func Map2f(target uint32, u1 float32, u2 float32, ustride int32, uorder int32, v1 float32, v2 float32, vstride int32, vorder int32, points *float32) {
	C.glowMap2f(gpMap2f, (C.GLenum)(target), (C.GLfloat)(u1), (C.GLfloat)(u2), (C.GLint)(ustride), (C.GLint)(uorder), (C.GLfloat)(v1), (C.GLfloat)(v2), (C.GLint)(vstride), (C.GLint)(vorder), (*C.GLfloat)(unsafe.Pointer(points)))
}

// map all of a buffer object's data store into the client's address space
func MapBuffer(target uint32, access uint32) unsafe.Pointer {
	ret := C.glowMapBuffer(gpMapBuffer, (C.GLenum)(target), (C.GLenum)(access))
	return (unsafe.Pointer)(ret)
}

// map all or part of a buffer object's data store into the client's address space
func MapBufferRange(target uint32, offset int, length int, access uint32) unsafe.Pointer {
	ret := C.glowMapBufferRange(gpMapBufferRange, (C.GLenum)(target), (C.GLintptr)(offset), (C.GLsizeiptr)(length), (C.GLbitfield)(access))
	return (unsafe.Pointer)(ret)
}
func MapGrid1d(un int32, u1 float64, u2 float64) {
	C.glowMapGrid1d(gpMapGrid1d, (C.GLint)(un), (C.GLdouble)(u1), (C.GLdouble)(u2))
}
func MapGrid1f(un int32, u1 float32, u2 float32) {
	C.glowMapGrid1f(gpMapGrid1f, (C.GLint)(un), (C.GLfloat)(u1), (C.GLfloat)(u2))
}
func MapGrid2d(un int32, u1 float64, u2 float64, vn int32, v1 float64, v2 float64) {
	C.glowMapGrid2d(gpMapGrid2d, (C.GLint)(un), (C.GLdouble)(u1), (C.GLdouble)(u2), (C.GLint)(vn), (C.GLdouble)(v1), (C.GLdouble)(v2))
}
func MapGrid2f(un int32, u1 float32, u2 float32, vn int32, v1 float32, v2 float32) {
	C.glowMapGrid2f(gpMapGrid2f, (C.GLint)(un), (C.GLfloat)(u1), (C.GLfloat)(u2), (C.GLint)(vn), (C.GLfloat)(v1), (C.GLfloat)(v2))
}

// map all of a buffer object's data store into the client's address space
func MapNamedBuffer(buffer uint32, access uint32) unsafe.Pointer {
	ret := C.glowMapNamedBuffer(gpMapNamedBuffer, (C.GLuint)(buffer), (C.GLenum)(access))
	return (unsafe.Pointer)(ret)
}

// map all or part of a buffer object's data store into the client's address space
func MapNamedBufferRange(buffer uint32, offset int, length int32, access uint32) unsafe.Pointer {
	ret := C.glowMapNamedBufferRange(gpMapNamedBufferRange, (C.GLuint)(buffer), (C.GLintptr)(offset), (C.GLsizei)(length), (C.GLbitfield)(access))
	return (unsafe.Pointer)(ret)
}
func Materialf(face uint32, pname uint32, param float32) {
	C.glowMaterialf(gpMaterialf, (C.GLenum)(face), (C.GLenum)(pname), (C.GLfloat)(param))
}
func Materialfv(face uint32, pname uint32, params *float32) {
	C.glowMaterialfv(gpMaterialfv, (C.GLenum)(face), (C.GLenum)(pname), (*C.GLfloat)(unsafe.Pointer(params)))
}
func Materiali(face uint32, pname uint32, param int32) {
	C.glowMateriali(gpMateriali, (C.GLenum)(face), (C.GLenum)(pname), (C.GLint)(param))
}
func Materialiv(face uint32, pname uint32, params *int32) {
	C.glowMaterialiv(gpMaterialiv, (C.GLenum)(face), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}

// specify which matrix is the current matrix
func MatrixMode(mode uint32) {
	C.glowMatrixMode(gpMatrixMode, (C.GLenum)(mode))
}

// defines a barrier ordering memory transactions
func MemoryBarrier(barriers uint32) {
	C.glowMemoryBarrier(gpMemoryBarrier, (C.GLbitfield)(barriers))
}
func MemoryBarrierByRegion(barriers uint32) {
	C.glowMemoryBarrierByRegion(gpMemoryBarrierByRegion, (C.GLbitfield)(barriers))
}

// specifies minimum rate at which sample shaing takes place
func MinSampleShading(value float32) {
	C.glowMinSampleShading(gpMinSampleShading, (C.GLfloat)(value))
}
func MinSampleShadingARB(value float32) {
	C.glowMinSampleShadingARB(gpMinSampleShadingARB, (C.GLfloat)(value))
}
func MultMatrixd(m *float64) {
	C.glowMultMatrixd(gpMultMatrixd, (*C.GLdouble)(unsafe.Pointer(m)))
}
func MultMatrixf(m *float32) {
	C.glowMultMatrixf(gpMultMatrixf, (*C.GLfloat)(unsafe.Pointer(m)))
}
func MultTransposeMatrixd(m *float64) {
	C.glowMultTransposeMatrixd(gpMultTransposeMatrixd, (*C.GLdouble)(unsafe.Pointer(m)))
}
func MultTransposeMatrixf(m *float32) {
	C.glowMultTransposeMatrixf(gpMultTransposeMatrixf, (*C.GLfloat)(unsafe.Pointer(m)))
}

// render multiple sets of primitives from array data
func MultiDrawArrays(mode uint32, first *int32, count *int32, drawcount int32) {
	C.glowMultiDrawArrays(gpMultiDrawArrays, (C.GLenum)(mode), (*C.GLint)(unsafe.Pointer(first)), (*C.GLsizei)(unsafe.Pointer(count)), (C.GLsizei)(drawcount))
}

// render multiple sets of primitives from array data, taking parameters from memory
func MultiDrawArraysIndirect(mode uint32, indirect unsafe.Pointer, drawcount int32, stride int32) {
	C.glowMultiDrawArraysIndirect(gpMultiDrawArraysIndirect, (C.GLenum)(mode), indirect, (C.GLsizei)(drawcount), (C.GLsizei)(stride))
}
func MultiDrawArraysIndirectCountARB(mode uint32, indirect int, drawcount int, maxdrawcount int32, stride int32) {
	C.glowMultiDrawArraysIndirectCountARB(gpMultiDrawArraysIndirectCountARB, (C.GLenum)(mode), (C.GLintptr)(indirect), (C.GLintptr)(drawcount), (C.GLsizei)(maxdrawcount), (C.GLsizei)(stride))
}

// render multiple sets of primitives by specifying indices of array data elements
func MultiDrawElements(mode uint32, count *int32, xtype uint32, indices *unsafe.Pointer, drawcount int32) {
	C.glowMultiDrawElements(gpMultiDrawElements, (C.GLenum)(mode), (*C.GLsizei)(unsafe.Pointer(count)), (C.GLenum)(xtype), indices, (C.GLsizei)(drawcount))
}

// render multiple sets of primitives by specifying indices of array data elements and an index to apply to each index
func MultiDrawElementsBaseVertex(mode uint32, count *int32, xtype uint32, indices *unsafe.Pointer, drawcount int32, basevertex *int32) {
	C.glowMultiDrawElementsBaseVertex(gpMultiDrawElementsBaseVertex, (C.GLenum)(mode), (*C.GLsizei)(unsafe.Pointer(count)), (C.GLenum)(xtype), indices, (C.GLsizei)(drawcount), (*C.GLint)(unsafe.Pointer(basevertex)))
}

// render indexed primitives from array data, taking parameters from memory
func MultiDrawElementsIndirect(mode uint32, xtype uint32, indirect unsafe.Pointer, drawcount int32, stride int32) {
	C.glowMultiDrawElementsIndirect(gpMultiDrawElementsIndirect, (C.GLenum)(mode), (C.GLenum)(xtype), indirect, (C.GLsizei)(drawcount), (C.GLsizei)(stride))
}
func MultiDrawElementsIndirectCountARB(mode uint32, xtype uint32, indirect int, drawcount int, maxdrawcount int32, stride int32) {
	C.glowMultiDrawElementsIndirectCountARB(gpMultiDrawElementsIndirectCountARB, (C.GLenum)(mode), (C.GLenum)(xtype), (C.GLintptr)(indirect), (C.GLintptr)(drawcount), (C.GLsizei)(maxdrawcount), (C.GLsizei)(stride))
}
func MultiTexCoord1d(target uint32, s float64) {
	C.glowMultiTexCoord1d(gpMultiTexCoord1d, (C.GLenum)(target), (C.GLdouble)(s))
}
func MultiTexCoord1dv(target uint32, v *float64) {
	C.glowMultiTexCoord1dv(gpMultiTexCoord1dv, (C.GLenum)(target), (*C.GLdouble)(unsafe.Pointer(v)))
}
func MultiTexCoord1f(target uint32, s float32) {
	C.glowMultiTexCoord1f(gpMultiTexCoord1f, (C.GLenum)(target), (C.GLfloat)(s))
}
func MultiTexCoord1fv(target uint32, v *float32) {
	C.glowMultiTexCoord1fv(gpMultiTexCoord1fv, (C.GLenum)(target), (*C.GLfloat)(unsafe.Pointer(v)))
}
func MultiTexCoord1i(target uint32, s int32) {
	C.glowMultiTexCoord1i(gpMultiTexCoord1i, (C.GLenum)(target), (C.GLint)(s))
}
func MultiTexCoord1iv(target uint32, v *int32) {
	C.glowMultiTexCoord1iv(gpMultiTexCoord1iv, (C.GLenum)(target), (*C.GLint)(unsafe.Pointer(v)))
}
func MultiTexCoord1s(target uint32, s int16) {
	C.glowMultiTexCoord1s(gpMultiTexCoord1s, (C.GLenum)(target), (C.GLshort)(s))
}
func MultiTexCoord1sv(target uint32, v *int16) {
	C.glowMultiTexCoord1sv(gpMultiTexCoord1sv, (C.GLenum)(target), (*C.GLshort)(unsafe.Pointer(v)))
}
func MultiTexCoord2d(target uint32, s float64, t float64) {
	C.glowMultiTexCoord2d(gpMultiTexCoord2d, (C.GLenum)(target), (C.GLdouble)(s), (C.GLdouble)(t))
}
func MultiTexCoord2dv(target uint32, v *float64) {
	C.glowMultiTexCoord2dv(gpMultiTexCoord2dv, (C.GLenum)(target), (*C.GLdouble)(unsafe.Pointer(v)))
}
func MultiTexCoord2f(target uint32, s float32, t float32) {
	C.glowMultiTexCoord2f(gpMultiTexCoord2f, (C.GLenum)(target), (C.GLfloat)(s), (C.GLfloat)(t))
}
func MultiTexCoord2fv(target uint32, v *float32) {
	C.glowMultiTexCoord2fv(gpMultiTexCoord2fv, (C.GLenum)(target), (*C.GLfloat)(unsafe.Pointer(v)))
}
func MultiTexCoord2i(target uint32, s int32, t int32) {
	C.glowMultiTexCoord2i(gpMultiTexCoord2i, (C.GLenum)(target), (C.GLint)(s), (C.GLint)(t))
}
func MultiTexCoord2iv(target uint32, v *int32) {
	C.glowMultiTexCoord2iv(gpMultiTexCoord2iv, (C.GLenum)(target), (*C.GLint)(unsafe.Pointer(v)))
}
func MultiTexCoord2s(target uint32, s int16, t int16) {
	C.glowMultiTexCoord2s(gpMultiTexCoord2s, (C.GLenum)(target), (C.GLshort)(s), (C.GLshort)(t))
}
func MultiTexCoord2sv(target uint32, v *int16) {
	C.glowMultiTexCoord2sv(gpMultiTexCoord2sv, (C.GLenum)(target), (*C.GLshort)(unsafe.Pointer(v)))
}
func MultiTexCoord3d(target uint32, s float64, t float64, r float64) {
	C.glowMultiTexCoord3d(gpMultiTexCoord3d, (C.GLenum)(target), (C.GLdouble)(s), (C.GLdouble)(t), (C.GLdouble)(r))
}
func MultiTexCoord3dv(target uint32, v *float64) {
	C.glowMultiTexCoord3dv(gpMultiTexCoord3dv, (C.GLenum)(target), (*C.GLdouble)(unsafe.Pointer(v)))
}
func MultiTexCoord3f(target uint32, s float32, t float32, r float32) {
	C.glowMultiTexCoord3f(gpMultiTexCoord3f, (C.GLenum)(target), (C.GLfloat)(s), (C.GLfloat)(t), (C.GLfloat)(r))
}
func MultiTexCoord3fv(target uint32, v *float32) {
	C.glowMultiTexCoord3fv(gpMultiTexCoord3fv, (C.GLenum)(target), (*C.GLfloat)(unsafe.Pointer(v)))
}
func MultiTexCoord3i(target uint32, s int32, t int32, r int32) {
	C.glowMultiTexCoord3i(gpMultiTexCoord3i, (C.GLenum)(target), (C.GLint)(s), (C.GLint)(t), (C.GLint)(r))
}
func MultiTexCoord3iv(target uint32, v *int32) {
	C.glowMultiTexCoord3iv(gpMultiTexCoord3iv, (C.GLenum)(target), (*C.GLint)(unsafe.Pointer(v)))
}
func MultiTexCoord3s(target uint32, s int16, t int16, r int16) {
	C.glowMultiTexCoord3s(gpMultiTexCoord3s, (C.GLenum)(target), (C.GLshort)(s), (C.GLshort)(t), (C.GLshort)(r))
}
func MultiTexCoord3sv(target uint32, v *int16) {
	C.glowMultiTexCoord3sv(gpMultiTexCoord3sv, (C.GLenum)(target), (*C.GLshort)(unsafe.Pointer(v)))
}
func MultiTexCoord4d(target uint32, s float64, t float64, r float64, q float64) {
	C.glowMultiTexCoord4d(gpMultiTexCoord4d, (C.GLenum)(target), (C.GLdouble)(s), (C.GLdouble)(t), (C.GLdouble)(r), (C.GLdouble)(q))
}
func MultiTexCoord4dv(target uint32, v *float64) {
	C.glowMultiTexCoord4dv(gpMultiTexCoord4dv, (C.GLenum)(target), (*C.GLdouble)(unsafe.Pointer(v)))
}
func MultiTexCoord4f(target uint32, s float32, t float32, r float32, q float32) {
	C.glowMultiTexCoord4f(gpMultiTexCoord4f, (C.GLenum)(target), (C.GLfloat)(s), (C.GLfloat)(t), (C.GLfloat)(r), (C.GLfloat)(q))
}
func MultiTexCoord4fv(target uint32, v *float32) {
	C.glowMultiTexCoord4fv(gpMultiTexCoord4fv, (C.GLenum)(target), (*C.GLfloat)(unsafe.Pointer(v)))
}
func MultiTexCoord4i(target uint32, s int32, t int32, r int32, q int32) {
	C.glowMultiTexCoord4i(gpMultiTexCoord4i, (C.GLenum)(target), (C.GLint)(s), (C.GLint)(t), (C.GLint)(r), (C.GLint)(q))
}
func MultiTexCoord4iv(target uint32, v *int32) {
	C.glowMultiTexCoord4iv(gpMultiTexCoord4iv, (C.GLenum)(target), (*C.GLint)(unsafe.Pointer(v)))
}
func MultiTexCoord4s(target uint32, s int16, t int16, r int16, q int16) {
	C.glowMultiTexCoord4s(gpMultiTexCoord4s, (C.GLenum)(target), (C.GLshort)(s), (C.GLshort)(t), (C.GLshort)(r), (C.GLshort)(q))
}
func MultiTexCoord4sv(target uint32, v *int16) {
	C.glowMultiTexCoord4sv(gpMultiTexCoord4sv, (C.GLenum)(target), (*C.GLshort)(unsafe.Pointer(v)))
}

// creates and initializes a buffer object's data     store
func NamedBufferData(buffer uint32, size int32, data unsafe.Pointer, usage uint32) {
	C.glowNamedBufferData(gpNamedBufferData, (C.GLuint)(buffer), (C.GLsizei)(size), data, (C.GLenum)(usage))
}
func NamedBufferPageCommitmentARB(buffer uint32, offset int, size int32, commit bool) {
	C.glowNamedBufferPageCommitmentARB(gpNamedBufferPageCommitmentARB, (C.GLuint)(buffer), (C.GLintptr)(offset), (C.GLsizei)(size), (C.GLboolean)(boolToInt(commit)))
}
func NamedBufferPageCommitmentEXT(buffer uint32, offset int, size int32, commit bool) {
	C.glowNamedBufferPageCommitmentEXT(gpNamedBufferPageCommitmentEXT, (C.GLuint)(buffer), (C.GLintptr)(offset), (C.GLsizei)(size), (C.GLboolean)(boolToInt(commit)))
}

// creates and initializes a buffer object's immutable data     store
func NamedBufferStorage(buffer uint32, size int32, data unsafe.Pointer, flags uint32) {
	C.glowNamedBufferStorage(gpNamedBufferStorage, (C.GLuint)(buffer), (C.GLsizei)(size), data, (C.GLbitfield)(flags))
}

// updates a subset of a buffer object's data store
func NamedBufferSubData(buffer uint32, offset int, size int32, data unsafe.Pointer) {
	C.glowNamedBufferSubData(gpNamedBufferSubData, (C.GLuint)(buffer), (C.GLintptr)(offset), (C.GLsizei)(size), data)
}

// specify which color buffers are to be drawn into
func NamedFramebufferDrawBuffer(framebuffer uint32, buf uint32) {
	C.glowNamedFramebufferDrawBuffer(gpNamedFramebufferDrawBuffer, (C.GLuint)(framebuffer), (C.GLenum)(buf))
}

// Specifies a list of color buffers to be drawn     into
func NamedFramebufferDrawBuffers(framebuffer uint32, n int32, bufs *uint32) {
	C.glowNamedFramebufferDrawBuffers(gpNamedFramebufferDrawBuffers, (C.GLuint)(framebuffer), (C.GLsizei)(n), (*C.GLenum)(unsafe.Pointer(bufs)))
}

// set a named parameter of a framebuffer object
func NamedFramebufferParameteri(framebuffer uint32, pname uint32, param int32) {
	C.glowNamedFramebufferParameteri(gpNamedFramebufferParameteri, (C.GLuint)(framebuffer), (C.GLenum)(pname), (C.GLint)(param))
}

// select a color buffer source for pixels
func NamedFramebufferReadBuffer(framebuffer uint32, src uint32) {
	C.glowNamedFramebufferReadBuffer(gpNamedFramebufferReadBuffer, (C.GLuint)(framebuffer), (C.GLenum)(src))
}

// attach a renderbuffer as a logical buffer of a framebuffer object
func NamedFramebufferRenderbuffer(framebuffer uint32, attachment uint32, renderbuffertarget uint32, renderbuffer uint32) {
	C.glowNamedFramebufferRenderbuffer(gpNamedFramebufferRenderbuffer, (C.GLuint)(framebuffer), (C.GLenum)(attachment), (C.GLenum)(renderbuffertarget), (C.GLuint)(renderbuffer))
}
func NamedFramebufferTexture(framebuffer uint32, attachment uint32, texture uint32, level int32) {
	C.glowNamedFramebufferTexture(gpNamedFramebufferTexture, (C.GLuint)(framebuffer), (C.GLenum)(attachment), (C.GLuint)(texture), (C.GLint)(level))
}

// attach a single layer of a texture object as a logical buffer of a framebuffer object
func NamedFramebufferTextureLayer(framebuffer uint32, attachment uint32, texture uint32, level int32, layer int32) {
	C.glowNamedFramebufferTextureLayer(gpNamedFramebufferTextureLayer, (C.GLuint)(framebuffer), (C.GLenum)(attachment), (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(layer))
}

// establish data storage, format and dimensions of a     renderbuffer object's image
func NamedRenderbufferStorage(renderbuffer uint32, internalformat uint32, width int32, height int32) {
	C.glowNamedRenderbufferStorage(gpNamedRenderbufferStorage, (C.GLuint)(renderbuffer), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height))
}

// establish data storage, format, dimensions and sample count of     a renderbuffer object's image
func NamedRenderbufferStorageMultisample(renderbuffer uint32, samples int32, internalformat uint32, width int32, height int32) {
	C.glowNamedRenderbufferStorageMultisample(gpNamedRenderbufferStorageMultisample, (C.GLuint)(renderbuffer), (C.GLsizei)(samples), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height))
}
func NamedStringARB(xtype uint32, namelen int32, name *uint8, stringlen int32, xstring *uint8) {
	C.glowNamedStringARB(gpNamedStringARB, (C.GLenum)(xtype), (C.GLint)(namelen), (*C.GLchar)(unsafe.Pointer(name)), (C.GLint)(stringlen), (*C.GLchar)(unsafe.Pointer(xstring)))
}

// create or replace a display list
func NewList(list uint32, mode uint32) {
	C.glowNewList(gpNewList, (C.GLuint)(list), (C.GLenum)(mode))
}
func Normal3b(nx int8, ny int8, nz int8) {
	C.glowNormal3b(gpNormal3b, (C.GLbyte)(nx), (C.GLbyte)(ny), (C.GLbyte)(nz))
}
func Normal3bv(v *int8) {
	C.glowNormal3bv(gpNormal3bv, (*C.GLbyte)(unsafe.Pointer(v)))
}
func Normal3d(nx float64, ny float64, nz float64) {
	C.glowNormal3d(gpNormal3d, (C.GLdouble)(nx), (C.GLdouble)(ny), (C.GLdouble)(nz))
}
func Normal3dv(v *float64) {
	C.glowNormal3dv(gpNormal3dv, (*C.GLdouble)(unsafe.Pointer(v)))
}
func Normal3f(nx float32, ny float32, nz float32) {
	C.glowNormal3f(gpNormal3f, (C.GLfloat)(nx), (C.GLfloat)(ny), (C.GLfloat)(nz))
}
func Normal3fv(v *float32) {
	C.glowNormal3fv(gpNormal3fv, (*C.GLfloat)(unsafe.Pointer(v)))
}
func Normal3i(nx int32, ny int32, nz int32) {
	C.glowNormal3i(gpNormal3i, (C.GLint)(nx), (C.GLint)(ny), (C.GLint)(nz))
}
func Normal3iv(v *int32) {
	C.glowNormal3iv(gpNormal3iv, (*C.GLint)(unsafe.Pointer(v)))
}
func Normal3s(nx int16, ny int16, nz int16) {
	C.glowNormal3s(gpNormal3s, (C.GLshort)(nx), (C.GLshort)(ny), (C.GLshort)(nz))
}
func Normal3sv(v *int16) {
	C.glowNormal3sv(gpNormal3sv, (*C.GLshort)(unsafe.Pointer(v)))
}

// define an array of normals
func NormalPointer(xtype uint32, stride int32, pointer unsafe.Pointer) {
	C.glowNormalPointer(gpNormalPointer, (C.GLenum)(xtype), (C.GLsizei)(stride), pointer)
}

// label a named object identified within a namespace
func ObjectLabel(identifier uint32, name uint32, length int32, label *uint8) {
	C.glowObjectLabel(gpObjectLabel, (C.GLenum)(identifier), (C.GLuint)(name), (C.GLsizei)(length), (*C.GLchar)(unsafe.Pointer(label)))
}
func ObjectLabelKHR(identifier uint32, name uint32, length int32, label *uint8) {
	C.glowObjectLabelKHR(gpObjectLabelKHR, (C.GLenum)(identifier), (C.GLuint)(name), (C.GLsizei)(length), (*C.GLchar)(unsafe.Pointer(label)))
}

// label a a sync object identified by a pointer
func ObjectPtrLabel(ptr unsafe.Pointer, length int32, label *uint8) {
	C.glowObjectPtrLabel(gpObjectPtrLabel, ptr, (C.GLsizei)(length), (*C.GLchar)(unsafe.Pointer(label)))
}
func ObjectPtrLabelKHR(ptr unsafe.Pointer, length int32, label *uint8) {
	C.glowObjectPtrLabelKHR(gpObjectPtrLabelKHR, ptr, (C.GLsizei)(length), (*C.GLchar)(unsafe.Pointer(label)))
}

// multiply the current matrix with an orthographic matrix
func Ortho(left float64, right float64, bottom float64, top float64, zNear float64, zFar float64) {
	C.glowOrtho(gpOrtho, (C.GLdouble)(left), (C.GLdouble)(right), (C.GLdouble)(bottom), (C.GLdouble)(top), (C.GLdouble)(zNear), (C.GLdouble)(zFar))
}

// place a marker in the feedback buffer
func PassThrough(token float32) {
	C.glowPassThrough(gpPassThrough, (C.GLfloat)(token))
}
func PatchParameterfv(pname uint32, values *float32) {
	C.glowPatchParameterfv(gpPatchParameterfv, (C.GLenum)(pname), (*C.GLfloat)(unsafe.Pointer(values)))
}
func PatchParameteri(pname uint32, value int32) {
	C.glowPatchParameteri(gpPatchParameteri, (C.GLenum)(pname), (C.GLint)(value))
}

// pause transform feedback operations
func PauseTransformFeedback() {
	C.glowPauseTransformFeedback(gpPauseTransformFeedback)
}
func PixelMapfv(xmap uint32, mapsize int32, values *float32) {
	C.glowPixelMapfv(gpPixelMapfv, (C.GLenum)(xmap), (C.GLsizei)(mapsize), (*C.GLfloat)(unsafe.Pointer(values)))
}
func PixelMapuiv(xmap uint32, mapsize int32, values *uint32) {
	C.glowPixelMapuiv(gpPixelMapuiv, (C.GLenum)(xmap), (C.GLsizei)(mapsize), (*C.GLuint)(unsafe.Pointer(values)))
}
func PixelMapusv(xmap uint32, mapsize int32, values *uint16) {
	C.glowPixelMapusv(gpPixelMapusv, (C.GLenum)(xmap), (C.GLsizei)(mapsize), (*C.GLushort)(unsafe.Pointer(values)))
}
func PixelStoref(pname uint32, param float32) {
	C.glowPixelStoref(gpPixelStoref, (C.GLenum)(pname), (C.GLfloat)(param))
}
func PixelStorei(pname uint32, param int32) {
	C.glowPixelStorei(gpPixelStorei, (C.GLenum)(pname), (C.GLint)(param))
}
func PixelTransferf(pname uint32, param float32) {
	C.glowPixelTransferf(gpPixelTransferf, (C.GLenum)(pname), (C.GLfloat)(param))
}
func PixelTransferi(pname uint32, param int32) {
	C.glowPixelTransferi(gpPixelTransferi, (C.GLenum)(pname), (C.GLint)(param))
}

// specify the pixel zoom factors
func PixelZoom(xfactor float32, yfactor float32) {
	C.glowPixelZoom(gpPixelZoom, (C.GLfloat)(xfactor), (C.GLfloat)(yfactor))
}
func PointParameterf(pname uint32, param float32) {
	C.glowPointParameterf(gpPointParameterf, (C.GLenum)(pname), (C.GLfloat)(param))
}
func PointParameterfv(pname uint32, params *float32) {
	C.glowPointParameterfv(gpPointParameterfv, (C.GLenum)(pname), (*C.GLfloat)(unsafe.Pointer(params)))
}
func PointParameteri(pname uint32, param int32) {
	C.glowPointParameteri(gpPointParameteri, (C.GLenum)(pname), (C.GLint)(param))
}
func PointParameteriv(pname uint32, params *int32) {
	C.glowPointParameteriv(gpPointParameteriv, (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}

// specify the diameter of rasterized points
func PointSize(size float32) {
	C.glowPointSize(gpPointSize, (C.GLfloat)(size))
}

// select a polygon rasterization mode
func PolygonMode(face uint32, mode uint32) {
	C.glowPolygonMode(gpPolygonMode, (C.GLenum)(face), (C.GLenum)(mode))
}

// set the scale and units used to calculate depth values
func PolygonOffset(factor float32, units float32) {
	C.glowPolygonOffset(gpPolygonOffset, (C.GLfloat)(factor), (C.GLfloat)(units))
}

// set the polygon stippling pattern
func PolygonStipple(mask *uint8) {
	C.glowPolygonStipple(gpPolygonStipple, (*C.GLubyte)(unsafe.Pointer(mask)))
}
func PopAttrib() {
	C.glowPopAttrib(gpPopAttrib)
}
func PopClientAttrib() {
	C.glowPopClientAttrib(gpPopClientAttrib)
}

// pop the active debug group
func PopDebugGroup() {
	C.glowPopDebugGroup(gpPopDebugGroup)
}
func PopDebugGroupKHR() {
	C.glowPopDebugGroupKHR(gpPopDebugGroupKHR)
}
func PopMatrix() {
	C.glowPopMatrix(gpPopMatrix)
}
func PopName() {
	C.glowPopName(gpPopName)
}

// specify the primitive restart index
func PrimitiveRestartIndex(index uint32) {
	C.glowPrimitiveRestartIndex(gpPrimitiveRestartIndex, (C.GLuint)(index))
}

// set texture residence priority
func PrioritizeTextures(n int32, textures *uint32, priorities *float32) {
	C.glowPrioritizeTextures(gpPrioritizeTextures, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(textures)), (*C.GLfloat)(unsafe.Pointer(priorities)))
}

// load a program object with a program binary
func ProgramBinary(program uint32, binaryFormat uint32, binary unsafe.Pointer, length int32) {
	C.glowProgramBinary(gpProgramBinary, (C.GLuint)(program), (C.GLenum)(binaryFormat), binary, (C.GLsizei)(length))
}
func ProgramParameteri(program uint32, pname uint32, value int32) {
	C.glowProgramParameteri(gpProgramParameteri, (C.GLuint)(program), (C.GLenum)(pname), (C.GLint)(value))
}
func ProgramUniform1d(program uint32, location int32, v0 float64) {
	C.glowProgramUniform1d(gpProgramUniform1d, (C.GLuint)(program), (C.GLint)(location), (C.GLdouble)(v0))
}
func ProgramUniform1dv(program uint32, location int32, count int32, value *float64) {
	C.glowProgramUniform1dv(gpProgramUniform1dv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLdouble)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform1f(program uint32, location int32, v0 float32) {
	C.glowProgramUniform1f(gpProgramUniform1f, (C.GLuint)(program), (C.GLint)(location), (C.GLfloat)(v0))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform1fv(program uint32, location int32, count int32, value *float32) {
	C.glowProgramUniform1fv(gpProgramUniform1fv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform1i(program uint32, location int32, v0 int32) {
	C.glowProgramUniform1i(gpProgramUniform1i, (C.GLuint)(program), (C.GLint)(location), (C.GLint)(v0))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform1iv(program uint32, location int32, count int32, value *int32) {
	C.glowProgramUniform1iv(gpProgramUniform1iv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform1ui(program uint32, location int32, v0 uint32) {
	C.glowProgramUniform1ui(gpProgramUniform1ui, (C.GLuint)(program), (C.GLint)(location), (C.GLuint)(v0))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform1uiv(program uint32, location int32, count int32, value *uint32) {
	C.glowProgramUniform1uiv(gpProgramUniform1uiv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(value)))
}
func ProgramUniform2d(program uint32, location int32, v0 float64, v1 float64) {
	C.glowProgramUniform2d(gpProgramUniform2d, (C.GLuint)(program), (C.GLint)(location), (C.GLdouble)(v0), (C.GLdouble)(v1))
}
func ProgramUniform2dv(program uint32, location int32, count int32, value *float64) {
	C.glowProgramUniform2dv(gpProgramUniform2dv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLdouble)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform2f(program uint32, location int32, v0 float32, v1 float32) {
	C.glowProgramUniform2f(gpProgramUniform2f, (C.GLuint)(program), (C.GLint)(location), (C.GLfloat)(v0), (C.GLfloat)(v1))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform2fv(program uint32, location int32, count int32, value *float32) {
	C.glowProgramUniform2fv(gpProgramUniform2fv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform2i(program uint32, location int32, v0 int32, v1 int32) {
	C.glowProgramUniform2i(gpProgramUniform2i, (C.GLuint)(program), (C.GLint)(location), (C.GLint)(v0), (C.GLint)(v1))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform2iv(program uint32, location int32, count int32, value *int32) {
	C.glowProgramUniform2iv(gpProgramUniform2iv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform2ui(program uint32, location int32, v0 uint32, v1 uint32) {
	C.glowProgramUniform2ui(gpProgramUniform2ui, (C.GLuint)(program), (C.GLint)(location), (C.GLuint)(v0), (C.GLuint)(v1))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform2uiv(program uint32, location int32, count int32, value *uint32) {
	C.glowProgramUniform2uiv(gpProgramUniform2uiv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(value)))
}
func ProgramUniform3d(program uint32, location int32, v0 float64, v1 float64, v2 float64) {
	C.glowProgramUniform3d(gpProgramUniform3d, (C.GLuint)(program), (C.GLint)(location), (C.GLdouble)(v0), (C.GLdouble)(v1), (C.GLdouble)(v2))
}
func ProgramUniform3dv(program uint32, location int32, count int32, value *float64) {
	C.glowProgramUniform3dv(gpProgramUniform3dv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLdouble)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform3f(program uint32, location int32, v0 float32, v1 float32, v2 float32) {
	C.glowProgramUniform3f(gpProgramUniform3f, (C.GLuint)(program), (C.GLint)(location), (C.GLfloat)(v0), (C.GLfloat)(v1), (C.GLfloat)(v2))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform3fv(program uint32, location int32, count int32, value *float32) {
	C.glowProgramUniform3fv(gpProgramUniform3fv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform3i(program uint32, location int32, v0 int32, v1 int32, v2 int32) {
	C.glowProgramUniform3i(gpProgramUniform3i, (C.GLuint)(program), (C.GLint)(location), (C.GLint)(v0), (C.GLint)(v1), (C.GLint)(v2))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform3iv(program uint32, location int32, count int32, value *int32) {
	C.glowProgramUniform3iv(gpProgramUniform3iv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform3ui(program uint32, location int32, v0 uint32, v1 uint32, v2 uint32) {
	C.glowProgramUniform3ui(gpProgramUniform3ui, (C.GLuint)(program), (C.GLint)(location), (C.GLuint)(v0), (C.GLuint)(v1), (C.GLuint)(v2))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform3uiv(program uint32, location int32, count int32, value *uint32) {
	C.glowProgramUniform3uiv(gpProgramUniform3uiv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(value)))
}
func ProgramUniform4d(program uint32, location int32, v0 float64, v1 float64, v2 float64, v3 float64) {
	C.glowProgramUniform4d(gpProgramUniform4d, (C.GLuint)(program), (C.GLint)(location), (C.GLdouble)(v0), (C.GLdouble)(v1), (C.GLdouble)(v2), (C.GLdouble)(v3))
}
func ProgramUniform4dv(program uint32, location int32, count int32, value *float64) {
	C.glowProgramUniform4dv(gpProgramUniform4dv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLdouble)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform4f(program uint32, location int32, v0 float32, v1 float32, v2 float32, v3 float32) {
	C.glowProgramUniform4f(gpProgramUniform4f, (C.GLuint)(program), (C.GLint)(location), (C.GLfloat)(v0), (C.GLfloat)(v1), (C.GLfloat)(v2), (C.GLfloat)(v3))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform4fv(program uint32, location int32, count int32, value *float32) {
	C.glowProgramUniform4fv(gpProgramUniform4fv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform4i(program uint32, location int32, v0 int32, v1 int32, v2 int32, v3 int32) {
	C.glowProgramUniform4i(gpProgramUniform4i, (C.GLuint)(program), (C.GLint)(location), (C.GLint)(v0), (C.GLint)(v1), (C.GLint)(v2), (C.GLint)(v3))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform4iv(program uint32, location int32, count int32, value *int32) {
	C.glowProgramUniform4iv(gpProgramUniform4iv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform4ui(program uint32, location int32, v0 uint32, v1 uint32, v2 uint32, v3 uint32) {
	C.glowProgramUniform4ui(gpProgramUniform4ui, (C.GLuint)(program), (C.GLint)(location), (C.GLuint)(v0), (C.GLuint)(v1), (C.GLuint)(v2), (C.GLuint)(v3))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform4uiv(program uint32, location int32, count int32, value *uint32) {
	C.glowProgramUniform4uiv(gpProgramUniform4uiv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(value)))
}
func ProgramUniformHandleui64ARB(program uint32, location int32, value uint64) {
	C.glowProgramUniformHandleui64ARB(gpProgramUniformHandleui64ARB, (C.GLuint)(program), (C.GLint)(location), (C.GLuint64)(value))
}
func ProgramUniformHandleui64vARB(program uint32, location int32, count int32, values *uint64) {
	C.glowProgramUniformHandleui64vARB(gpProgramUniformHandleui64vARB, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint64)(unsafe.Pointer(values)))
}
func ProgramUniformMatrix2dv(program uint32, location int32, count int32, transpose bool, value *float64) {
	C.glowProgramUniformMatrix2dv(gpProgramUniformMatrix2dv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLdouble)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniformMatrix2fv(program uint32, location int32, count int32, transpose bool, value *float32) {
	C.glowProgramUniformMatrix2fv(gpProgramUniformMatrix2fv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func ProgramUniformMatrix2x3dv(program uint32, location int32, count int32, transpose bool, value *float64) {
	C.glowProgramUniformMatrix2x3dv(gpProgramUniformMatrix2x3dv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLdouble)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniformMatrix2x3fv(program uint32, location int32, count int32, transpose bool, value *float32) {
	C.glowProgramUniformMatrix2x3fv(gpProgramUniformMatrix2x3fv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func ProgramUniformMatrix2x4dv(program uint32, location int32, count int32, transpose bool, value *float64) {
	C.glowProgramUniformMatrix2x4dv(gpProgramUniformMatrix2x4dv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLdouble)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniformMatrix2x4fv(program uint32, location int32, count int32, transpose bool, value *float32) {
	C.glowProgramUniformMatrix2x4fv(gpProgramUniformMatrix2x4fv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func ProgramUniformMatrix3dv(program uint32, location int32, count int32, transpose bool, value *float64) {
	C.glowProgramUniformMatrix3dv(gpProgramUniformMatrix3dv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLdouble)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniformMatrix3fv(program uint32, location int32, count int32, transpose bool, value *float32) {
	C.glowProgramUniformMatrix3fv(gpProgramUniformMatrix3fv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func ProgramUniformMatrix3x2dv(program uint32, location int32, count int32, transpose bool, value *float64) {
	C.glowProgramUniformMatrix3x2dv(gpProgramUniformMatrix3x2dv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLdouble)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniformMatrix3x2fv(program uint32, location int32, count int32, transpose bool, value *float32) {
	C.glowProgramUniformMatrix3x2fv(gpProgramUniformMatrix3x2fv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func ProgramUniformMatrix3x4dv(program uint32, location int32, count int32, transpose bool, value *float64) {
	C.glowProgramUniformMatrix3x4dv(gpProgramUniformMatrix3x4dv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLdouble)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniformMatrix3x4fv(program uint32, location int32, count int32, transpose bool, value *float32) {
	C.glowProgramUniformMatrix3x4fv(gpProgramUniformMatrix3x4fv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func ProgramUniformMatrix4dv(program uint32, location int32, count int32, transpose bool, value *float64) {
	C.glowProgramUniformMatrix4dv(gpProgramUniformMatrix4dv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLdouble)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniformMatrix4fv(program uint32, location int32, count int32, transpose bool, value *float32) {
	C.glowProgramUniformMatrix4fv(gpProgramUniformMatrix4fv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func ProgramUniformMatrix4x2dv(program uint32, location int32, count int32, transpose bool, value *float64) {
	C.glowProgramUniformMatrix4x2dv(gpProgramUniformMatrix4x2dv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLdouble)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniformMatrix4x2fv(program uint32, location int32, count int32, transpose bool, value *float32) {
	C.glowProgramUniformMatrix4x2fv(gpProgramUniformMatrix4x2fv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func ProgramUniformMatrix4x3dv(program uint32, location int32, count int32, transpose bool, value *float64) {
	C.glowProgramUniformMatrix4x3dv(gpProgramUniformMatrix4x3dv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLdouble)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniformMatrix4x3fv(program uint32, location int32, count int32, transpose bool, value *float32) {
	C.glowProgramUniformMatrix4x3fv(gpProgramUniformMatrix4x3fv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}

// specifiy the vertex to be used as the source of data for flat shaded varyings
func ProvokingVertex(mode uint32) {
	C.glowProvokingVertex(gpProvokingVertex, (C.GLenum)(mode))
}

// push and pop the server attribute stack
func PushAttrib(mask uint32) {
	C.glowPushAttrib(gpPushAttrib, (C.GLbitfield)(mask))
}

// push and pop the client attribute stack
func PushClientAttrib(mask uint32) {
	C.glowPushClientAttrib(gpPushClientAttrib, (C.GLbitfield)(mask))
}

// push a named debug group into the command stream
func PushDebugGroup(source uint32, id uint32, length int32, message *uint8) {
	C.glowPushDebugGroup(gpPushDebugGroup, (C.GLenum)(source), (C.GLuint)(id), (C.GLsizei)(length), (*C.GLchar)(unsafe.Pointer(message)))
}
func PushDebugGroupKHR(source uint32, id uint32, length int32, message *uint8) {
	C.glowPushDebugGroupKHR(gpPushDebugGroupKHR, (C.GLenum)(source), (C.GLuint)(id), (C.GLsizei)(length), (*C.GLchar)(unsafe.Pointer(message)))
}

// push and pop the current matrix stack
func PushMatrix() {
	C.glowPushMatrix(gpPushMatrix)
}

// push and pop the name stack
func PushName(name uint32) {
	C.glowPushName(gpPushName, (C.GLuint)(name))
}

// record the GL time into a query object after all previous commands have reached the GL server but have not yet necessarily executed.
func QueryCounter(id uint32, target uint32) {
	C.glowQueryCounter(gpQueryCounter, (C.GLuint)(id), (C.GLenum)(target))
}
func RasterPos2d(x float64, y float64) {
	C.glowRasterPos2d(gpRasterPos2d, (C.GLdouble)(x), (C.GLdouble)(y))
}
func RasterPos2dv(v *float64) {
	C.glowRasterPos2dv(gpRasterPos2dv, (*C.GLdouble)(unsafe.Pointer(v)))
}
func RasterPos2f(x float32, y float32) {
	C.glowRasterPos2f(gpRasterPos2f, (C.GLfloat)(x), (C.GLfloat)(y))
}
func RasterPos2fv(v *float32) {
	C.glowRasterPos2fv(gpRasterPos2fv, (*C.GLfloat)(unsafe.Pointer(v)))
}
func RasterPos2i(x int32, y int32) {
	C.glowRasterPos2i(gpRasterPos2i, (C.GLint)(x), (C.GLint)(y))
}
func RasterPos2iv(v *int32) {
	C.glowRasterPos2iv(gpRasterPos2iv, (*C.GLint)(unsafe.Pointer(v)))
}
func RasterPos2s(x int16, y int16) {
	C.glowRasterPos2s(gpRasterPos2s, (C.GLshort)(x), (C.GLshort)(y))
}
func RasterPos2sv(v *int16) {
	C.glowRasterPos2sv(gpRasterPos2sv, (*C.GLshort)(unsafe.Pointer(v)))
}
func RasterPos3d(x float64, y float64, z float64) {
	C.glowRasterPos3d(gpRasterPos3d, (C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z))
}
func RasterPos3dv(v *float64) {
	C.glowRasterPos3dv(gpRasterPos3dv, (*C.GLdouble)(unsafe.Pointer(v)))
}
func RasterPos3f(x float32, y float32, z float32) {
	C.glowRasterPos3f(gpRasterPos3f, (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z))
}
func RasterPos3fv(v *float32) {
	C.glowRasterPos3fv(gpRasterPos3fv, (*C.GLfloat)(unsafe.Pointer(v)))
}
func RasterPos3i(x int32, y int32, z int32) {
	C.glowRasterPos3i(gpRasterPos3i, (C.GLint)(x), (C.GLint)(y), (C.GLint)(z))
}
func RasterPos3iv(v *int32) {
	C.glowRasterPos3iv(gpRasterPos3iv, (*C.GLint)(unsafe.Pointer(v)))
}
func RasterPos3s(x int16, y int16, z int16) {
	C.glowRasterPos3s(gpRasterPos3s, (C.GLshort)(x), (C.GLshort)(y), (C.GLshort)(z))
}
func RasterPos3sv(v *int16) {
	C.glowRasterPos3sv(gpRasterPos3sv, (*C.GLshort)(unsafe.Pointer(v)))
}
func RasterPos4d(x float64, y float64, z float64, w float64) {
	C.glowRasterPos4d(gpRasterPos4d, (C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z), (C.GLdouble)(w))
}
func RasterPos4dv(v *float64) {
	C.glowRasterPos4dv(gpRasterPos4dv, (*C.GLdouble)(unsafe.Pointer(v)))
}
func RasterPos4f(x float32, y float32, z float32, w float32) {
	C.glowRasterPos4f(gpRasterPos4f, (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z), (C.GLfloat)(w))
}
func RasterPos4fv(v *float32) {
	C.glowRasterPos4fv(gpRasterPos4fv, (*C.GLfloat)(unsafe.Pointer(v)))
}
func RasterPos4i(x int32, y int32, z int32, w int32) {
	C.glowRasterPos4i(gpRasterPos4i, (C.GLint)(x), (C.GLint)(y), (C.GLint)(z), (C.GLint)(w))
}
func RasterPos4iv(v *int32) {
	C.glowRasterPos4iv(gpRasterPos4iv, (*C.GLint)(unsafe.Pointer(v)))
}
func RasterPos4s(x int16, y int16, z int16, w int16) {
	C.glowRasterPos4s(gpRasterPos4s, (C.GLshort)(x), (C.GLshort)(y), (C.GLshort)(z), (C.GLshort)(w))
}
func RasterPos4sv(v *int16) {
	C.glowRasterPos4sv(gpRasterPos4sv, (*C.GLshort)(unsafe.Pointer(v)))
}

// select a color buffer source for pixels
func ReadBuffer(src uint32) {
	C.glowReadBuffer(gpReadBuffer, (C.GLenum)(src))
}

// read a block of pixels from the frame buffer
func ReadPixels(x int32, y int32, width int32, height int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	C.glowReadPixels(gpReadPixels, (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLenum)(xtype), pixels)
}

// read a block of pixels from the frame buffer
func ReadnPixels(x int32, y int32, width int32, height int32, format uint32, xtype uint32, bufSize int32, data unsafe.Pointer) {
	C.glowReadnPixels(gpReadnPixels, (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLenum)(xtype), (C.GLsizei)(bufSize), data)
}
func ReadnPixelsARB(x int32, y int32, width int32, height int32, format uint32, xtype uint32, bufSize int32, data unsafe.Pointer) {
	C.glowReadnPixelsARB(gpReadnPixelsARB, (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLenum)(xtype), (C.GLsizei)(bufSize), data)
}
func ReadnPixelsKHR(x int32, y int32, width int32, height int32, format uint32, xtype uint32, bufSize int32, data unsafe.Pointer) {
	C.glowReadnPixelsKHR(gpReadnPixelsKHR, (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLenum)(xtype), (C.GLsizei)(bufSize), data)
}
func Rectd(x1 float64, y1 float64, x2 float64, y2 float64) {
	C.glowRectd(gpRectd, (C.GLdouble)(x1), (C.GLdouble)(y1), (C.GLdouble)(x2), (C.GLdouble)(y2))
}
func Rectdv(v1 *float64, v2 *float64) {
	C.glowRectdv(gpRectdv, (*C.GLdouble)(unsafe.Pointer(v1)), (*C.GLdouble)(unsafe.Pointer(v2)))
}
func Rectf(x1 float32, y1 float32, x2 float32, y2 float32) {
	C.glowRectf(gpRectf, (C.GLfloat)(x1), (C.GLfloat)(y1), (C.GLfloat)(x2), (C.GLfloat)(y2))
}
func Rectfv(v1 *float32, v2 *float32) {
	C.glowRectfv(gpRectfv, (*C.GLfloat)(unsafe.Pointer(v1)), (*C.GLfloat)(unsafe.Pointer(v2)))
}
func Recti(x1 int32, y1 int32, x2 int32, y2 int32) {
	C.glowRecti(gpRecti, (C.GLint)(x1), (C.GLint)(y1), (C.GLint)(x2), (C.GLint)(y2))
}
func Rectiv(v1 *int32, v2 *int32) {
	C.glowRectiv(gpRectiv, (*C.GLint)(unsafe.Pointer(v1)), (*C.GLint)(unsafe.Pointer(v2)))
}
func Rects(x1 int16, y1 int16, x2 int16, y2 int16) {
	C.glowRects(gpRects, (C.GLshort)(x1), (C.GLshort)(y1), (C.GLshort)(x2), (C.GLshort)(y2))
}
func Rectsv(v1 *int16, v2 *int16) {
	C.glowRectsv(gpRectsv, (*C.GLshort)(unsafe.Pointer(v1)), (*C.GLshort)(unsafe.Pointer(v2)))
}

// release resources consumed by the implementation's shader compiler
func ReleaseShaderCompiler() {
	C.glowReleaseShaderCompiler(gpReleaseShaderCompiler)
}

// set rasterization mode
func RenderMode(mode uint32) int32 {
	ret := C.glowRenderMode(gpRenderMode, (C.GLenum)(mode))
	return (int32)(ret)
}

// establish data storage, format and dimensions of a     renderbuffer object's image
func RenderbufferStorage(target uint32, internalformat uint32, width int32, height int32) {
	C.glowRenderbufferStorage(gpRenderbufferStorage, (C.GLenum)(target), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height))
}

// establish data storage, format, dimensions and sample count of     a renderbuffer object's image
func RenderbufferStorageMultisample(target uint32, samples int32, internalformat uint32, width int32, height int32) {
	C.glowRenderbufferStorageMultisample(gpRenderbufferStorageMultisample, (C.GLenum)(target), (C.GLsizei)(samples), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height))
}

// resume transform feedback operations
func ResumeTransformFeedback() {
	C.glowResumeTransformFeedback(gpResumeTransformFeedback)
}
func Rotated(angle float64, x float64, y float64, z float64) {
	C.glowRotated(gpRotated, (C.GLdouble)(angle), (C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z))
}
func Rotatef(angle float32, x float32, y float32, z float32) {
	C.glowRotatef(gpRotatef, (C.GLfloat)(angle), (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z))
}

// specify multisample coverage parameters
func SampleCoverage(value float32, invert bool) {
	C.glowSampleCoverage(gpSampleCoverage, (C.GLfloat)(value), (C.GLboolean)(boolToInt(invert)))
}

// set the value of a sub-word of the sample mask
func SampleMaski(maskNumber uint32, mask uint32) {
	C.glowSampleMaski(gpSampleMaski, (C.GLuint)(maskNumber), (C.GLbitfield)(mask))
}
func SamplerParameterIiv(sampler uint32, pname uint32, param *int32) {
	C.glowSamplerParameterIiv(gpSamplerParameterIiv, (C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(param)))
}
func SamplerParameterIuiv(sampler uint32, pname uint32, param *uint32) {
	C.glowSamplerParameterIuiv(gpSamplerParameterIuiv, (C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLuint)(unsafe.Pointer(param)))
}
func SamplerParameterf(sampler uint32, pname uint32, param float32) {
	C.glowSamplerParameterf(gpSamplerParameterf, (C.GLuint)(sampler), (C.GLenum)(pname), (C.GLfloat)(param))
}
func SamplerParameterfv(sampler uint32, pname uint32, param *float32) {
	C.glowSamplerParameterfv(gpSamplerParameterfv, (C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLfloat)(unsafe.Pointer(param)))
}
func SamplerParameteri(sampler uint32, pname uint32, param int32) {
	C.glowSamplerParameteri(gpSamplerParameteri, (C.GLuint)(sampler), (C.GLenum)(pname), (C.GLint)(param))
}
func SamplerParameteriv(sampler uint32, pname uint32, param *int32) {
	C.glowSamplerParameteriv(gpSamplerParameteriv, (C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(param)))
}
func Scaled(x float64, y float64, z float64) {
	C.glowScaled(gpScaled, (C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z))
}
func Scalef(x float32, y float32, z float32) {
	C.glowScalef(gpScalef, (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z))
}

// define the scissor box
func Scissor(x int32, y int32, width int32, height int32) {
	C.glowScissor(gpScissor, (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}
func ScissorArrayv(first uint32, count int32, v *int32) {
	C.glowScissorArrayv(gpScissorArrayv, (C.GLuint)(first), (C.GLsizei)(count), (*C.GLint)(unsafe.Pointer(v)))
}

// define the scissor box for a specific viewport
func ScissorIndexed(index uint32, left int32, bottom int32, width int32, height int32) {
	C.glowScissorIndexed(gpScissorIndexed, (C.GLuint)(index), (C.GLint)(left), (C.GLint)(bottom), (C.GLsizei)(width), (C.GLsizei)(height))
}
func ScissorIndexedv(index uint32, v *int32) {
	C.glowScissorIndexedv(gpScissorIndexedv, (C.GLuint)(index), (*C.GLint)(unsafe.Pointer(v)))
}
func SecondaryColor3b(red int8, green int8, blue int8) {
	C.glowSecondaryColor3b(gpSecondaryColor3b, (C.GLbyte)(red), (C.GLbyte)(green), (C.GLbyte)(blue))
}
func SecondaryColor3bv(v *int8) {
	C.glowSecondaryColor3bv(gpSecondaryColor3bv, (*C.GLbyte)(unsafe.Pointer(v)))
}
func SecondaryColor3d(red float64, green float64, blue float64) {
	C.glowSecondaryColor3d(gpSecondaryColor3d, (C.GLdouble)(red), (C.GLdouble)(green), (C.GLdouble)(blue))
}
func SecondaryColor3dv(v *float64) {
	C.glowSecondaryColor3dv(gpSecondaryColor3dv, (*C.GLdouble)(unsafe.Pointer(v)))
}
func SecondaryColor3f(red float32, green float32, blue float32) {
	C.glowSecondaryColor3f(gpSecondaryColor3f, (C.GLfloat)(red), (C.GLfloat)(green), (C.GLfloat)(blue))
}
func SecondaryColor3fv(v *float32) {
	C.glowSecondaryColor3fv(gpSecondaryColor3fv, (*C.GLfloat)(unsafe.Pointer(v)))
}
func SecondaryColor3i(red int32, green int32, blue int32) {
	C.glowSecondaryColor3i(gpSecondaryColor3i, (C.GLint)(red), (C.GLint)(green), (C.GLint)(blue))
}
func SecondaryColor3iv(v *int32) {
	C.glowSecondaryColor3iv(gpSecondaryColor3iv, (*C.GLint)(unsafe.Pointer(v)))
}
func SecondaryColor3s(red int16, green int16, blue int16) {
	C.glowSecondaryColor3s(gpSecondaryColor3s, (C.GLshort)(red), (C.GLshort)(green), (C.GLshort)(blue))
}
func SecondaryColor3sv(v *int16) {
	C.glowSecondaryColor3sv(gpSecondaryColor3sv, (*C.GLshort)(unsafe.Pointer(v)))
}
func SecondaryColor3ub(red uint8, green uint8, blue uint8) {
	C.glowSecondaryColor3ub(gpSecondaryColor3ub, (C.GLubyte)(red), (C.GLubyte)(green), (C.GLubyte)(blue))
}
func SecondaryColor3ubv(v *uint8) {
	C.glowSecondaryColor3ubv(gpSecondaryColor3ubv, (*C.GLubyte)(unsafe.Pointer(v)))
}
func SecondaryColor3ui(red uint32, green uint32, blue uint32) {
	C.glowSecondaryColor3ui(gpSecondaryColor3ui, (C.GLuint)(red), (C.GLuint)(green), (C.GLuint)(blue))
}
func SecondaryColor3uiv(v *uint32) {
	C.glowSecondaryColor3uiv(gpSecondaryColor3uiv, (*C.GLuint)(unsafe.Pointer(v)))
}
func SecondaryColor3us(red uint16, green uint16, blue uint16) {
	C.glowSecondaryColor3us(gpSecondaryColor3us, (C.GLushort)(red), (C.GLushort)(green), (C.GLushort)(blue))
}
func SecondaryColor3usv(v *uint16) {
	C.glowSecondaryColor3usv(gpSecondaryColor3usv, (*C.GLushort)(unsafe.Pointer(v)))
}

// define an array of secondary colors
func SecondaryColorPointer(size int32, xtype uint32, stride int32, pointer unsafe.Pointer) {
	C.glowSecondaryColorPointer(gpSecondaryColorPointer, (C.GLint)(size), (C.GLenum)(xtype), (C.GLsizei)(stride), pointer)
}

// establish a buffer for selection mode values
func SelectBuffer(size int32, buffer *uint32) {
	C.glowSelectBuffer(gpSelectBuffer, (C.GLsizei)(size), (*C.GLuint)(unsafe.Pointer(buffer)))
}

// select flat or smooth shading
func ShadeModel(mode uint32) {
	C.glowShadeModel(gpShadeModel, (C.GLenum)(mode))
}

// load pre-compiled shader binaries
func ShaderBinary(count int32, shaders *uint32, binaryformat uint32, binary unsafe.Pointer, length int32) {
	C.glowShaderBinary(gpShaderBinary, (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(shaders)), (C.GLenum)(binaryformat), binary, (C.GLsizei)(length))
}

// Replaces the source code in a shader object
func ShaderSource(shader uint32, count int32, xstring **uint8, length *int32) {
	C.glowShaderSource(gpShaderSource, (C.GLuint)(shader), (C.GLsizei)(count), (**C.GLchar)(unsafe.Pointer(xstring)), (*C.GLint)(unsafe.Pointer(length)))
}

// change an active shader storage block binding
func ShaderStorageBlockBinding(program uint32, storageBlockIndex uint32, storageBlockBinding uint32) {
	C.glowShaderStorageBlockBinding(gpShaderStorageBlockBinding, (C.GLuint)(program), (C.GLuint)(storageBlockIndex), (C.GLuint)(storageBlockBinding))
}

// set front and back function and reference value for stencil testing
func StencilFunc(xfunc uint32, ref int32, mask uint32) {
	C.glowStencilFunc(gpStencilFunc, (C.GLenum)(xfunc), (C.GLint)(ref), (C.GLuint)(mask))
}

// set front and/or back function and reference value for stencil testing
func StencilFuncSeparate(face uint32, xfunc uint32, ref int32, mask uint32) {
	C.glowStencilFuncSeparate(gpStencilFuncSeparate, (C.GLenum)(face), (C.GLenum)(xfunc), (C.GLint)(ref), (C.GLuint)(mask))
}

// control the front and back writing of individual bits in the stencil planes
func StencilMask(mask uint32) {
	C.glowStencilMask(gpStencilMask, (C.GLuint)(mask))
}

// control the front and/or back writing of individual bits in the stencil planes
func StencilMaskSeparate(face uint32, mask uint32) {
	C.glowStencilMaskSeparate(gpStencilMaskSeparate, (C.GLenum)(face), (C.GLuint)(mask))
}

// set front and back stencil test actions
func StencilOp(fail uint32, zfail uint32, zpass uint32) {
	C.glowStencilOp(gpStencilOp, (C.GLenum)(fail), (C.GLenum)(zfail), (C.GLenum)(zpass))
}

// set front and/or back stencil test actions
func StencilOpSeparate(face uint32, sfail uint32, dpfail uint32, dppass uint32) {
	C.glowStencilOpSeparate(gpStencilOpSeparate, (C.GLenum)(face), (C.GLenum)(sfail), (C.GLenum)(dpfail), (C.GLenum)(dppass))
}

// attach a buffer object's data store to a buffer texture object
func TexBuffer(target uint32, internalformat uint32, buffer uint32) {
	C.glowTexBuffer(gpTexBuffer, (C.GLenum)(target), (C.GLenum)(internalformat), (C.GLuint)(buffer))
}

// attach a range of a buffer object's data store to a buffer texture object
func TexBufferRange(target uint32, internalformat uint32, buffer uint32, offset int, size int) {
	C.glowTexBufferRange(gpTexBufferRange, (C.GLenum)(target), (C.GLenum)(internalformat), (C.GLuint)(buffer), (C.GLintptr)(offset), (C.GLsizeiptr)(size))
}
func TexCoord1d(s float64) {
	C.glowTexCoord1d(gpTexCoord1d, (C.GLdouble)(s))
}
func TexCoord1dv(v *float64) {
	C.glowTexCoord1dv(gpTexCoord1dv, (*C.GLdouble)(unsafe.Pointer(v)))
}
func TexCoord1f(s float32) {
	C.glowTexCoord1f(gpTexCoord1f, (C.GLfloat)(s))
}
func TexCoord1fv(v *float32) {
	C.glowTexCoord1fv(gpTexCoord1fv, (*C.GLfloat)(unsafe.Pointer(v)))
}
func TexCoord1i(s int32) {
	C.glowTexCoord1i(gpTexCoord1i, (C.GLint)(s))
}
func TexCoord1iv(v *int32) {
	C.glowTexCoord1iv(gpTexCoord1iv, (*C.GLint)(unsafe.Pointer(v)))
}
func TexCoord1s(s int16) {
	C.glowTexCoord1s(gpTexCoord1s, (C.GLshort)(s))
}
func TexCoord1sv(v *int16) {
	C.glowTexCoord1sv(gpTexCoord1sv, (*C.GLshort)(unsafe.Pointer(v)))
}
func TexCoord2d(s float64, t float64) {
	C.glowTexCoord2d(gpTexCoord2d, (C.GLdouble)(s), (C.GLdouble)(t))
}
func TexCoord2dv(v *float64) {
	C.glowTexCoord2dv(gpTexCoord2dv, (*C.GLdouble)(unsafe.Pointer(v)))
}
func TexCoord2f(s float32, t float32) {
	C.glowTexCoord2f(gpTexCoord2f, (C.GLfloat)(s), (C.GLfloat)(t))
}
func TexCoord2fv(v *float32) {
	C.glowTexCoord2fv(gpTexCoord2fv, (*C.GLfloat)(unsafe.Pointer(v)))
}
func TexCoord2i(s int32, t int32) {
	C.glowTexCoord2i(gpTexCoord2i, (C.GLint)(s), (C.GLint)(t))
}
func TexCoord2iv(v *int32) {
	C.glowTexCoord2iv(gpTexCoord2iv, (*C.GLint)(unsafe.Pointer(v)))
}
func TexCoord2s(s int16, t int16) {
	C.glowTexCoord2s(gpTexCoord2s, (C.GLshort)(s), (C.GLshort)(t))
}
func TexCoord2sv(v *int16) {
	C.glowTexCoord2sv(gpTexCoord2sv, (*C.GLshort)(unsafe.Pointer(v)))
}
func TexCoord3d(s float64, t float64, r float64) {
	C.glowTexCoord3d(gpTexCoord3d, (C.GLdouble)(s), (C.GLdouble)(t), (C.GLdouble)(r))
}
func TexCoord3dv(v *float64) {
	C.glowTexCoord3dv(gpTexCoord3dv, (*C.GLdouble)(unsafe.Pointer(v)))
}
func TexCoord3f(s float32, t float32, r float32) {
	C.glowTexCoord3f(gpTexCoord3f, (C.GLfloat)(s), (C.GLfloat)(t), (C.GLfloat)(r))
}
func TexCoord3fv(v *float32) {
	C.glowTexCoord3fv(gpTexCoord3fv, (*C.GLfloat)(unsafe.Pointer(v)))
}
func TexCoord3i(s int32, t int32, r int32) {
	C.glowTexCoord3i(gpTexCoord3i, (C.GLint)(s), (C.GLint)(t), (C.GLint)(r))
}
func TexCoord3iv(v *int32) {
	C.glowTexCoord3iv(gpTexCoord3iv, (*C.GLint)(unsafe.Pointer(v)))
}
func TexCoord3s(s int16, t int16, r int16) {
	C.glowTexCoord3s(gpTexCoord3s, (C.GLshort)(s), (C.GLshort)(t), (C.GLshort)(r))
}
func TexCoord3sv(v *int16) {
	C.glowTexCoord3sv(gpTexCoord3sv, (*C.GLshort)(unsafe.Pointer(v)))
}
func TexCoord4d(s float64, t float64, r float64, q float64) {
	C.glowTexCoord4d(gpTexCoord4d, (C.GLdouble)(s), (C.GLdouble)(t), (C.GLdouble)(r), (C.GLdouble)(q))
}
func TexCoord4dv(v *float64) {
	C.glowTexCoord4dv(gpTexCoord4dv, (*C.GLdouble)(unsafe.Pointer(v)))
}
func TexCoord4f(s float32, t float32, r float32, q float32) {
	C.glowTexCoord4f(gpTexCoord4f, (C.GLfloat)(s), (C.GLfloat)(t), (C.GLfloat)(r), (C.GLfloat)(q))
}
func TexCoord4fv(v *float32) {
	C.glowTexCoord4fv(gpTexCoord4fv, (*C.GLfloat)(unsafe.Pointer(v)))
}
func TexCoord4i(s int32, t int32, r int32, q int32) {
	C.glowTexCoord4i(gpTexCoord4i, (C.GLint)(s), (C.GLint)(t), (C.GLint)(r), (C.GLint)(q))
}
func TexCoord4iv(v *int32) {
	C.glowTexCoord4iv(gpTexCoord4iv, (*C.GLint)(unsafe.Pointer(v)))
}
func TexCoord4s(s int16, t int16, r int16, q int16) {
	C.glowTexCoord4s(gpTexCoord4s, (C.GLshort)(s), (C.GLshort)(t), (C.GLshort)(r), (C.GLshort)(q))
}
func TexCoord4sv(v *int16) {
	C.glowTexCoord4sv(gpTexCoord4sv, (*C.GLshort)(unsafe.Pointer(v)))
}

// define an array of texture coordinates
func TexCoordPointer(size int32, xtype uint32, stride int32, pointer unsafe.Pointer) {
	C.glowTexCoordPointer(gpTexCoordPointer, (C.GLint)(size), (C.GLenum)(xtype), (C.GLsizei)(stride), pointer)
}
func TexEnvf(target uint32, pname uint32, param float32) {
	C.glowTexEnvf(gpTexEnvf, (C.GLenum)(target), (C.GLenum)(pname), (C.GLfloat)(param))
}
func TexEnvfv(target uint32, pname uint32, params *float32) {
	C.glowTexEnvfv(gpTexEnvfv, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLfloat)(unsafe.Pointer(params)))
}
func TexEnvi(target uint32, pname uint32, param int32) {
	C.glowTexEnvi(gpTexEnvi, (C.GLenum)(target), (C.GLenum)(pname), (C.GLint)(param))
}
func TexEnviv(target uint32, pname uint32, params *int32) {
	C.glowTexEnviv(gpTexEnviv, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func TexGend(coord uint32, pname uint32, param float64) {
	C.glowTexGend(gpTexGend, (C.GLenum)(coord), (C.GLenum)(pname), (C.GLdouble)(param))
}
func TexGendv(coord uint32, pname uint32, params *float64) {
	C.glowTexGendv(gpTexGendv, (C.GLenum)(coord), (C.GLenum)(pname), (*C.GLdouble)(unsafe.Pointer(params)))
}
func TexGenf(coord uint32, pname uint32, param float32) {
	C.glowTexGenf(gpTexGenf, (C.GLenum)(coord), (C.GLenum)(pname), (C.GLfloat)(param))
}
func TexGenfv(coord uint32, pname uint32, params *float32) {
	C.glowTexGenfv(gpTexGenfv, (C.GLenum)(coord), (C.GLenum)(pname), (*C.GLfloat)(unsafe.Pointer(params)))
}
func TexGeni(coord uint32, pname uint32, param int32) {
	C.glowTexGeni(gpTexGeni, (C.GLenum)(coord), (C.GLenum)(pname), (C.GLint)(param))
}
func TexGeniv(coord uint32, pname uint32, params *int32) {
	C.glowTexGeniv(gpTexGeniv, (C.GLenum)(coord), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}

// specify a one-dimensional texture image
func TexImage1D(target uint32, level int32, internalformat int32, width int32, border int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	C.glowTexImage1D(gpTexImage1D, (C.GLenum)(target), (C.GLint)(level), (C.GLint)(internalformat), (C.GLsizei)(width), (C.GLint)(border), (C.GLenum)(format), (C.GLenum)(xtype), pixels)
}

// specify a two-dimensional texture image
func TexImage2D(target uint32, level int32, internalformat int32, width int32, height int32, border int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	C.glowTexImage2D(gpTexImage2D, (C.GLenum)(target), (C.GLint)(level), (C.GLint)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLint)(border), (C.GLenum)(format), (C.GLenum)(xtype), pixels)
}

// establish the data storage, format, dimensions, and number of samples of a multisample texture's image
func TexImage2DMultisample(target uint32, samples int32, internalformat uint32, width int32, height int32, fixedsamplelocations bool) {
	C.glowTexImage2DMultisample(gpTexImage2DMultisample, (C.GLenum)(target), (C.GLsizei)(samples), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLboolean)(boolToInt(fixedsamplelocations)))
}

// specify a three-dimensional texture image
func TexImage3D(target uint32, level int32, internalformat int32, width int32, height int32, depth int32, border int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	C.glowTexImage3D(gpTexImage3D, (C.GLenum)(target), (C.GLint)(level), (C.GLint)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLint)(border), (C.GLenum)(format), (C.GLenum)(xtype), pixels)
}

// establish the data storage, format, dimensions, and number of samples of a multisample texture's image
func TexImage3DMultisample(target uint32, samples int32, internalformat uint32, width int32, height int32, depth int32, fixedsamplelocations bool) {
	C.glowTexImage3DMultisample(gpTexImage3DMultisample, (C.GLenum)(target), (C.GLsizei)(samples), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLboolean)(boolToInt(fixedsamplelocations)))
}
func TexPageCommitmentARB(target uint32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, resident bool) {
	C.glowTexPageCommitmentARB(gpTexPageCommitmentARB, (C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLboolean)(boolToInt(resident)))
}
func TexParameterIiv(target uint32, pname uint32, params *int32) {
	C.glowTexParameterIiv(gpTexParameterIiv, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func TexParameterIuiv(target uint32, pname uint32, params *uint32) {
	C.glowTexParameterIuiv(gpTexParameterIuiv, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLuint)(unsafe.Pointer(params)))
}
func TexParameterf(target uint32, pname uint32, param float32) {
	C.glowTexParameterf(gpTexParameterf, (C.GLenum)(target), (C.GLenum)(pname), (C.GLfloat)(param))
}
func TexParameterfv(target uint32, pname uint32, params *float32) {
	C.glowTexParameterfv(gpTexParameterfv, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLfloat)(unsafe.Pointer(params)))
}
func TexParameteri(target uint32, pname uint32, param int32) {
	C.glowTexParameteri(gpTexParameteri, (C.GLenum)(target), (C.GLenum)(pname), (C.GLint)(param))
}
func TexParameteriv(target uint32, pname uint32, params *int32) {
	C.glowTexParameteriv(gpTexParameteriv, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}

// simultaneously specify storage for all levels of a one-dimensional texture
func TexStorage1D(target uint32, levels int32, internalformat uint32, width int32) {
	C.glowTexStorage1D(gpTexStorage1D, (C.GLenum)(target), (C.GLsizei)(levels), (C.GLenum)(internalformat), (C.GLsizei)(width))
}

// simultaneously specify storage for all levels of a two-dimensional or one-dimensional array texture
func TexStorage2D(target uint32, levels int32, internalformat uint32, width int32, height int32) {
	C.glowTexStorage2D(gpTexStorage2D, (C.GLenum)(target), (C.GLsizei)(levels), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height))
}

// specify storage for a two-dimensional multisample texture
func TexStorage2DMultisample(target uint32, samples int32, internalformat uint32, width int32, height int32, fixedsamplelocations bool) {
	C.glowTexStorage2DMultisample(gpTexStorage2DMultisample, (C.GLenum)(target), (C.GLsizei)(samples), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLboolean)(boolToInt(fixedsamplelocations)))
}

// simultaneously specify storage for all levels of a three-dimensional, two-dimensional array or cube-map array texture
func TexStorage3D(target uint32, levels int32, internalformat uint32, width int32, height int32, depth int32) {
	C.glowTexStorage3D(gpTexStorage3D, (C.GLenum)(target), (C.GLsizei)(levels), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth))
}

// specify storage for a two-dimensional multisample array texture
func TexStorage3DMultisample(target uint32, samples int32, internalformat uint32, width int32, height int32, depth int32, fixedsamplelocations bool) {
	C.glowTexStorage3DMultisample(gpTexStorage3DMultisample, (C.GLenum)(target), (C.GLsizei)(samples), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLboolean)(boolToInt(fixedsamplelocations)))
}

// specify a one-dimensional texture subimage
func TexSubImage1D(target uint32, level int32, xoffset int32, width int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	C.glowTexSubImage1D(gpTexSubImage1D, (C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLsizei)(width), (C.GLenum)(format), (C.GLenum)(xtype), pixels)
}

// specify a two-dimensional texture subimage
func TexSubImage2D(target uint32, level int32, xoffset int32, yoffset int32, width int32, height int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	C.glowTexSubImage2D(gpTexSubImage2D, (C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLenum)(xtype), pixels)
}

// specify a three-dimensional texture subimage
func TexSubImage3D(target uint32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	C.glowTexSubImage3D(gpTexSubImage3D, (C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLenum)(format), (C.GLenum)(xtype), pixels)
}

// controls the ordering of reads and writes to rendered fragments across drawing commands
func TextureBarrier() {
	C.glowTextureBarrier(gpTextureBarrier)
}

// attach a buffer object's data store to a buffer texture object
func TextureBuffer(texture uint32, internalformat uint32, buffer uint32) {
	C.glowTextureBuffer(gpTextureBuffer, (C.GLuint)(texture), (C.GLenum)(internalformat), (C.GLuint)(buffer))
}

// attach a range of a buffer object's data store to a buffer texture object
func TextureBufferRange(texture uint32, internalformat uint32, buffer uint32, offset int, size int32) {
	C.glowTextureBufferRange(gpTextureBufferRange, (C.GLuint)(texture), (C.GLenum)(internalformat), (C.GLuint)(buffer), (C.GLintptr)(offset), (C.GLsizei)(size))
}
func TextureParameterIiv(texture uint32, pname uint32, params *int32) {
	C.glowTextureParameterIiv(gpTextureParameterIiv, (C.GLuint)(texture), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func TextureParameterIuiv(texture uint32, pname uint32, params *uint32) {
	C.glowTextureParameterIuiv(gpTextureParameterIuiv, (C.GLuint)(texture), (C.GLenum)(pname), (*C.GLuint)(unsafe.Pointer(params)))
}
func TextureParameterf(texture uint32, pname uint32, param float32) {
	C.glowTextureParameterf(gpTextureParameterf, (C.GLuint)(texture), (C.GLenum)(pname), (C.GLfloat)(param))
}
func TextureParameterfv(texture uint32, pname uint32, param *float32) {
	C.glowTextureParameterfv(gpTextureParameterfv, (C.GLuint)(texture), (C.GLenum)(pname), (*C.GLfloat)(unsafe.Pointer(param)))
}
func TextureParameteri(texture uint32, pname uint32, param int32) {
	C.glowTextureParameteri(gpTextureParameteri, (C.GLuint)(texture), (C.GLenum)(pname), (C.GLint)(param))
}
func TextureParameteriv(texture uint32, pname uint32, param *int32) {
	C.glowTextureParameteriv(gpTextureParameteriv, (C.GLuint)(texture), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(param)))
}

// simultaneously specify storage for all levels of a one-dimensional texture
func TextureStorage1D(texture uint32, levels int32, internalformat uint32, width int32) {
	C.glowTextureStorage1D(gpTextureStorage1D, (C.GLuint)(texture), (C.GLsizei)(levels), (C.GLenum)(internalformat), (C.GLsizei)(width))
}

// simultaneously specify storage for all levels of a two-dimensional or one-dimensional array texture
func TextureStorage2D(texture uint32, levels int32, internalformat uint32, width int32, height int32) {
	C.glowTextureStorage2D(gpTextureStorage2D, (C.GLuint)(texture), (C.GLsizei)(levels), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height))
}

// specify storage for a two-dimensional multisample texture
func TextureStorage2DMultisample(texture uint32, samples int32, internalformat uint32, width int32, height int32, fixedsamplelocations bool) {
	C.glowTextureStorage2DMultisample(gpTextureStorage2DMultisample, (C.GLuint)(texture), (C.GLsizei)(samples), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLboolean)(boolToInt(fixedsamplelocations)))
}

// simultaneously specify storage for all levels of a three-dimensional, two-dimensional array or cube-map array texture
func TextureStorage3D(texture uint32, levels int32, internalformat uint32, width int32, height int32, depth int32) {
	C.glowTextureStorage3D(gpTextureStorage3D, (C.GLuint)(texture), (C.GLsizei)(levels), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth))
}

// specify storage for a two-dimensional multisample array texture
func TextureStorage3DMultisample(texture uint32, samples int32, internalformat uint32, width int32, height int32, depth int32, fixedsamplelocations bool) {
	C.glowTextureStorage3DMultisample(gpTextureStorage3DMultisample, (C.GLuint)(texture), (C.GLsizei)(samples), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLboolean)(boolToInt(fixedsamplelocations)))
}

// specify a one-dimensional texture subimage
func TextureSubImage1D(texture uint32, level int32, xoffset int32, width int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	C.glowTextureSubImage1D(gpTextureSubImage1D, (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(xoffset), (C.GLsizei)(width), (C.GLenum)(format), (C.GLenum)(xtype), pixels)
}

// specify a two-dimensional texture subimage
func TextureSubImage2D(texture uint32, level int32, xoffset int32, yoffset int32, width int32, height int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	C.glowTextureSubImage2D(gpTextureSubImage2D, (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLenum)(xtype), pixels)
}

// specify a three-dimensional texture subimage
func TextureSubImage3D(texture uint32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	C.glowTextureSubImage3D(gpTextureSubImage3D, (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLenum)(format), (C.GLenum)(xtype), pixels)
}

// initialize a texture as a data alias of another texture's data store
func TextureView(texture uint32, target uint32, origtexture uint32, internalformat uint32, minlevel uint32, numlevels uint32, minlayer uint32, numlayers uint32) {
	C.glowTextureView(gpTextureView, (C.GLuint)(texture), (C.GLenum)(target), (C.GLuint)(origtexture), (C.GLenum)(internalformat), (C.GLuint)(minlevel), (C.GLuint)(numlevels), (C.GLuint)(minlayer), (C.GLuint)(numlayers))
}

// bind a buffer object to a transform feedback buffer object
func TransformFeedbackBufferBase(xfb uint32, index uint32, buffer uint32) {
	C.glowTransformFeedbackBufferBase(gpTransformFeedbackBufferBase, (C.GLuint)(xfb), (C.GLuint)(index), (C.GLuint)(buffer))
}

// bind a range within a buffer object to a transform feedback buffer object
func TransformFeedbackBufferRange(xfb uint32, index uint32, buffer uint32, offset int, size int32) {
	C.glowTransformFeedbackBufferRange(gpTransformFeedbackBufferRange, (C.GLuint)(xfb), (C.GLuint)(index), (C.GLuint)(buffer), (C.GLintptr)(offset), (C.GLsizei)(size))
}

// specify values to record in transform feedback buffers
func TransformFeedbackVaryings(program uint32, count int32, varyings **uint8, bufferMode uint32) {
	C.glowTransformFeedbackVaryings(gpTransformFeedbackVaryings, (C.GLuint)(program), (C.GLsizei)(count), (**C.GLchar)(unsafe.Pointer(varyings)), (C.GLenum)(bufferMode))
}
func Translated(x float64, y float64, z float64) {
	C.glowTranslated(gpTranslated, (C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z))
}
func Translatef(x float32, y float32, z float32) {
	C.glowTranslatef(gpTranslatef, (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z))
}
func Uniform1d(location int32, x float64) {
	C.glowUniform1d(gpUniform1d, (C.GLint)(location), (C.GLdouble)(x))
}
func Uniform1dv(location int32, count int32, value *float64) {
	C.glowUniform1dv(gpUniform1dv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLdouble)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func Uniform1f(location int32, v0 float32) {
	C.glowUniform1f(gpUniform1f, (C.GLint)(location), (C.GLfloat)(v0))
}

// Specify the value of a uniform variable for the current program object
func Uniform1fv(location int32, count int32, value *float32) {
	C.glowUniform1fv(gpUniform1fv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func Uniform1i(location int32, v0 int32) {
	C.glowUniform1i(gpUniform1i, (C.GLint)(location), (C.GLint)(v0))
}

// Specify the value of a uniform variable for the current program object
func Uniform1iv(location int32, count int32, value *int32) {
	C.glowUniform1iv(gpUniform1iv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func Uniform1ui(location int32, v0 uint32) {
	C.glowUniform1ui(gpUniform1ui, (C.GLint)(location), (C.GLuint)(v0))
}

// Specify the value of a uniform variable for the current program object
func Uniform1uiv(location int32, count int32, value *uint32) {
	C.glowUniform1uiv(gpUniform1uiv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(value)))
}
func Uniform2d(location int32, x float64, y float64) {
	C.glowUniform2d(gpUniform2d, (C.GLint)(location), (C.GLdouble)(x), (C.GLdouble)(y))
}
func Uniform2dv(location int32, count int32, value *float64) {
	C.glowUniform2dv(gpUniform2dv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLdouble)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func Uniform2f(location int32, v0 float32, v1 float32) {
	C.glowUniform2f(gpUniform2f, (C.GLint)(location), (C.GLfloat)(v0), (C.GLfloat)(v1))
}

// Specify the value of a uniform variable for the current program object
func Uniform2fv(location int32, count int32, value *float32) {
	C.glowUniform2fv(gpUniform2fv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func Uniform2i(location int32, v0 int32, v1 int32) {
	C.glowUniform2i(gpUniform2i, (C.GLint)(location), (C.GLint)(v0), (C.GLint)(v1))
}

// Specify the value of a uniform variable for the current program object
func Uniform2iv(location int32, count int32, value *int32) {
	C.glowUniform2iv(gpUniform2iv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func Uniform2ui(location int32, v0 uint32, v1 uint32) {
	C.glowUniform2ui(gpUniform2ui, (C.GLint)(location), (C.GLuint)(v0), (C.GLuint)(v1))
}

// Specify the value of a uniform variable for the current program object
func Uniform2uiv(location int32, count int32, value *uint32) {
	C.glowUniform2uiv(gpUniform2uiv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(value)))
}
func Uniform3d(location int32, x float64, y float64, z float64) {
	C.glowUniform3d(gpUniform3d, (C.GLint)(location), (C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z))
}
func Uniform3dv(location int32, count int32, value *float64) {
	C.glowUniform3dv(gpUniform3dv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLdouble)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func Uniform3f(location int32, v0 float32, v1 float32, v2 float32) {
	C.glowUniform3f(gpUniform3f, (C.GLint)(location), (C.GLfloat)(v0), (C.GLfloat)(v1), (C.GLfloat)(v2))
}

// Specify the value of a uniform variable for the current program object
func Uniform3fv(location int32, count int32, value *float32) {
	C.glowUniform3fv(gpUniform3fv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func Uniform3i(location int32, v0 int32, v1 int32, v2 int32) {
	C.glowUniform3i(gpUniform3i, (C.GLint)(location), (C.GLint)(v0), (C.GLint)(v1), (C.GLint)(v2))
}

// Specify the value of a uniform variable for the current program object
func Uniform3iv(location int32, count int32, value *int32) {
	C.glowUniform3iv(gpUniform3iv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func Uniform3ui(location int32, v0 uint32, v1 uint32, v2 uint32) {
	C.glowUniform3ui(gpUniform3ui, (C.GLint)(location), (C.GLuint)(v0), (C.GLuint)(v1), (C.GLuint)(v2))
}

// Specify the value of a uniform variable for the current program object
func Uniform3uiv(location int32, count int32, value *uint32) {
	C.glowUniform3uiv(gpUniform3uiv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(value)))
}
func Uniform4d(location int32, x float64, y float64, z float64, w float64) {
	C.glowUniform4d(gpUniform4d, (C.GLint)(location), (C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z), (C.GLdouble)(w))
}
func Uniform4dv(location int32, count int32, value *float64) {
	C.glowUniform4dv(gpUniform4dv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLdouble)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func Uniform4f(location int32, v0 float32, v1 float32, v2 float32, v3 float32) {
	C.glowUniform4f(gpUniform4f, (C.GLint)(location), (C.GLfloat)(v0), (C.GLfloat)(v1), (C.GLfloat)(v2), (C.GLfloat)(v3))
}

// Specify the value of a uniform variable for the current program object
func Uniform4fv(location int32, count int32, value *float32) {
	C.glowUniform4fv(gpUniform4fv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func Uniform4i(location int32, v0 int32, v1 int32, v2 int32, v3 int32) {
	C.glowUniform4i(gpUniform4i, (C.GLint)(location), (C.GLint)(v0), (C.GLint)(v1), (C.GLint)(v2), (C.GLint)(v3))
}

// Specify the value of a uniform variable for the current program object
func Uniform4iv(location int32, count int32, value *int32) {
	C.glowUniform4iv(gpUniform4iv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func Uniform4ui(location int32, v0 uint32, v1 uint32, v2 uint32, v3 uint32) {
	C.glowUniform4ui(gpUniform4ui, (C.GLint)(location), (C.GLuint)(v0), (C.GLuint)(v1), (C.GLuint)(v2), (C.GLuint)(v3))
}

// Specify the value of a uniform variable for the current program object
func Uniform4uiv(location int32, count int32, value *uint32) {
	C.glowUniform4uiv(gpUniform4uiv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(value)))
}

// assign a binding point to an active uniform block
func UniformBlockBinding(program uint32, uniformBlockIndex uint32, uniformBlockBinding uint32) {
	C.glowUniformBlockBinding(gpUniformBlockBinding, (C.GLuint)(program), (C.GLuint)(uniformBlockIndex), (C.GLuint)(uniformBlockBinding))
}
func UniformHandleui64ARB(location int32, value uint64) {
	C.glowUniformHandleui64ARB(gpUniformHandleui64ARB, (C.GLint)(location), (C.GLuint64)(value))
}
func UniformHandleui64vARB(location int32, count int32, value *uint64) {
	C.glowUniformHandleui64vARB(gpUniformHandleui64vARB, (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint64)(unsafe.Pointer(value)))
}
func UniformMatrix2dv(location int32, count int32, transpose bool, value *float64) {
	C.glowUniformMatrix2dv(gpUniformMatrix2dv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLdouble)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func UniformMatrix2fv(location int32, count int32, transpose bool, value *float32) {
	C.glowUniformMatrix2fv(gpUniformMatrix2fv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func UniformMatrix2x3dv(location int32, count int32, transpose bool, value *float64) {
	C.glowUniformMatrix2x3dv(gpUniformMatrix2x3dv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLdouble)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func UniformMatrix2x3fv(location int32, count int32, transpose bool, value *float32) {
	C.glowUniformMatrix2x3fv(gpUniformMatrix2x3fv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func UniformMatrix2x4dv(location int32, count int32, transpose bool, value *float64) {
	C.glowUniformMatrix2x4dv(gpUniformMatrix2x4dv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLdouble)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func UniformMatrix2x4fv(location int32, count int32, transpose bool, value *float32) {
	C.glowUniformMatrix2x4fv(gpUniformMatrix2x4fv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func UniformMatrix3dv(location int32, count int32, transpose bool, value *float64) {
	C.glowUniformMatrix3dv(gpUniformMatrix3dv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLdouble)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func UniformMatrix3fv(location int32, count int32, transpose bool, value *float32) {
	C.glowUniformMatrix3fv(gpUniformMatrix3fv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func UniformMatrix3x2dv(location int32, count int32, transpose bool, value *float64) {
	C.glowUniformMatrix3x2dv(gpUniformMatrix3x2dv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLdouble)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func UniformMatrix3x2fv(location int32, count int32, transpose bool, value *float32) {
	C.glowUniformMatrix3x2fv(gpUniformMatrix3x2fv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func UniformMatrix3x4dv(location int32, count int32, transpose bool, value *float64) {
	C.glowUniformMatrix3x4dv(gpUniformMatrix3x4dv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLdouble)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func UniformMatrix3x4fv(location int32, count int32, transpose bool, value *float32) {
	C.glowUniformMatrix3x4fv(gpUniformMatrix3x4fv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func UniformMatrix4dv(location int32, count int32, transpose bool, value *float64) {
	C.glowUniformMatrix4dv(gpUniformMatrix4dv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLdouble)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func UniformMatrix4fv(location int32, count int32, transpose bool, value *float32) {
	C.glowUniformMatrix4fv(gpUniformMatrix4fv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func UniformMatrix4x2dv(location int32, count int32, transpose bool, value *float64) {
	C.glowUniformMatrix4x2dv(gpUniformMatrix4x2dv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLdouble)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func UniformMatrix4x2fv(location int32, count int32, transpose bool, value *float32) {
	C.glowUniformMatrix4x2fv(gpUniformMatrix4x2fv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func UniformMatrix4x3dv(location int32, count int32, transpose bool, value *float64) {
	C.glowUniformMatrix4x3dv(gpUniformMatrix4x3dv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLdouble)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func UniformMatrix4x3fv(location int32, count int32, transpose bool, value *float32) {
	C.glowUniformMatrix4x3fv(gpUniformMatrix4x3fv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func UniformSubroutinesuiv(shadertype uint32, count int32, indices *uint32) {
	C.glowUniformSubroutinesuiv(gpUniformSubroutinesuiv, (C.GLenum)(shadertype), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(indices)))
}

// release the mapping of a buffer object's data store into the client's address space
func UnmapBuffer(target uint32) bool {
	ret := C.glowUnmapBuffer(gpUnmapBuffer, (C.GLenum)(target))
	return ret == TRUE
}

// release the mapping of a buffer object's data store into the client's address space
func UnmapNamedBuffer(buffer uint32) bool {
	ret := C.glowUnmapNamedBuffer(gpUnmapNamedBuffer, (C.GLuint)(buffer))
	return ret == TRUE
}

// Installs a program object as part of current rendering state
func UseProgram(program uint32) {
	C.glowUseProgram(gpUseProgram, (C.GLuint)(program))
}

// bind stages of a program object to a program pipeline
func UseProgramStages(pipeline uint32, stages uint32, program uint32) {
	C.glowUseProgramStages(gpUseProgramStages, (C.GLuint)(pipeline), (C.GLbitfield)(stages), (C.GLuint)(program))
}

// Validates a program object
func ValidateProgram(program uint32) {
	C.glowValidateProgram(gpValidateProgram, (C.GLuint)(program))
}

// validate a program pipeline object against current GL state
func ValidateProgramPipeline(pipeline uint32) {
	C.glowValidateProgramPipeline(gpValidateProgramPipeline, (C.GLuint)(pipeline))
}
func Vertex2d(x float64, y float64) {
	C.glowVertex2d(gpVertex2d, (C.GLdouble)(x), (C.GLdouble)(y))
}
func Vertex2dv(v *float64) {
	C.glowVertex2dv(gpVertex2dv, (*C.GLdouble)(unsafe.Pointer(v)))
}
func Vertex2f(x float32, y float32) {
	C.glowVertex2f(gpVertex2f, (C.GLfloat)(x), (C.GLfloat)(y))
}
func Vertex2fv(v *float32) {
	C.glowVertex2fv(gpVertex2fv, (*C.GLfloat)(unsafe.Pointer(v)))
}
func Vertex2i(x int32, y int32) {
	C.glowVertex2i(gpVertex2i, (C.GLint)(x), (C.GLint)(y))
}
func Vertex2iv(v *int32) {
	C.glowVertex2iv(gpVertex2iv, (*C.GLint)(unsafe.Pointer(v)))
}
func Vertex2s(x int16, y int16) {
	C.glowVertex2s(gpVertex2s, (C.GLshort)(x), (C.GLshort)(y))
}
func Vertex2sv(v *int16) {
	C.glowVertex2sv(gpVertex2sv, (*C.GLshort)(unsafe.Pointer(v)))
}
func Vertex3d(x float64, y float64, z float64) {
	C.glowVertex3d(gpVertex3d, (C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z))
}
func Vertex3dv(v *float64) {
	C.glowVertex3dv(gpVertex3dv, (*C.GLdouble)(unsafe.Pointer(v)))
}
func Vertex3f(x float32, y float32, z float32) {
	C.glowVertex3f(gpVertex3f, (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z))
}
func Vertex3fv(v *float32) {
	C.glowVertex3fv(gpVertex3fv, (*C.GLfloat)(unsafe.Pointer(v)))
}
func Vertex3i(x int32, y int32, z int32) {
	C.glowVertex3i(gpVertex3i, (C.GLint)(x), (C.GLint)(y), (C.GLint)(z))
}
func Vertex3iv(v *int32) {
	C.glowVertex3iv(gpVertex3iv, (*C.GLint)(unsafe.Pointer(v)))
}
func Vertex3s(x int16, y int16, z int16) {
	C.glowVertex3s(gpVertex3s, (C.GLshort)(x), (C.GLshort)(y), (C.GLshort)(z))
}
func Vertex3sv(v *int16) {
	C.glowVertex3sv(gpVertex3sv, (*C.GLshort)(unsafe.Pointer(v)))
}
func Vertex4d(x float64, y float64, z float64, w float64) {
	C.glowVertex4d(gpVertex4d, (C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z), (C.GLdouble)(w))
}
func Vertex4dv(v *float64) {
	C.glowVertex4dv(gpVertex4dv, (*C.GLdouble)(unsafe.Pointer(v)))
}
func Vertex4f(x float32, y float32, z float32, w float32) {
	C.glowVertex4f(gpVertex4f, (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z), (C.GLfloat)(w))
}
func Vertex4fv(v *float32) {
	C.glowVertex4fv(gpVertex4fv, (*C.GLfloat)(unsafe.Pointer(v)))
}
func Vertex4i(x int32, y int32, z int32, w int32) {
	C.glowVertex4i(gpVertex4i, (C.GLint)(x), (C.GLint)(y), (C.GLint)(z), (C.GLint)(w))
}
func Vertex4iv(v *int32) {
	C.glowVertex4iv(gpVertex4iv, (*C.GLint)(unsafe.Pointer(v)))
}
func Vertex4s(x int16, y int16, z int16, w int16) {
	C.glowVertex4s(gpVertex4s, (C.GLshort)(x), (C.GLshort)(y), (C.GLshort)(z), (C.GLshort)(w))
}
func Vertex4sv(v *int16) {
	C.glowVertex4sv(gpVertex4sv, (*C.GLshort)(unsafe.Pointer(v)))
}
func VertexArrayAttribBinding(vaobj uint32, attribindex uint32, bindingindex uint32) {
	C.glowVertexArrayAttribBinding(gpVertexArrayAttribBinding, (C.GLuint)(vaobj), (C.GLuint)(attribindex), (C.GLuint)(bindingindex))
}

// specify the organization of vertex arrays
func VertexArrayAttribFormat(vaobj uint32, attribindex uint32, size int32, xtype uint32, normalized bool, relativeoffset uint32) {
	C.glowVertexArrayAttribFormat(gpVertexArrayAttribFormat, (C.GLuint)(vaobj), (C.GLuint)(attribindex), (C.GLint)(size), (C.GLenum)(xtype), (C.GLboolean)(boolToInt(normalized)), (C.GLuint)(relativeoffset))
}
func VertexArrayAttribIFormat(vaobj uint32, attribindex uint32, size int32, xtype uint32, relativeoffset uint32) {
	C.glowVertexArrayAttribIFormat(gpVertexArrayAttribIFormat, (C.GLuint)(vaobj), (C.GLuint)(attribindex), (C.GLint)(size), (C.GLenum)(xtype), (C.GLuint)(relativeoffset))
}
func VertexArrayAttribLFormat(vaobj uint32, attribindex uint32, size int32, xtype uint32, relativeoffset uint32) {
	C.glowVertexArrayAttribLFormat(gpVertexArrayAttribLFormat, (C.GLuint)(vaobj), (C.GLuint)(attribindex), (C.GLint)(size), (C.GLenum)(xtype), (C.GLuint)(relativeoffset))
}

// modify the rate at which generic vertex attributes     advance
func VertexArrayBindingDivisor(vaobj uint32, bindingindex uint32, divisor uint32) {
	C.glowVertexArrayBindingDivisor(gpVertexArrayBindingDivisor, (C.GLuint)(vaobj), (C.GLuint)(bindingindex), (C.GLuint)(divisor))
}

// configures element array buffer binding of a vertex array object
func VertexArrayElementBuffer(vaobj uint32, buffer uint32) {
	C.glowVertexArrayElementBuffer(gpVertexArrayElementBuffer, (C.GLuint)(vaobj), (C.GLuint)(buffer))
}

// bind a buffer to a vertex buffer bind point
func VertexArrayVertexBuffer(vaobj uint32, bindingindex uint32, buffer uint32, offset int, stride int32) {
	C.glowVertexArrayVertexBuffer(gpVertexArrayVertexBuffer, (C.GLuint)(vaobj), (C.GLuint)(bindingindex), (C.GLuint)(buffer), (C.GLintptr)(offset), (C.GLsizei)(stride))
}

// attach multiple buffer objects to a vertex array object
func VertexArrayVertexBuffers(vaobj uint32, first uint32, count int32, buffers *uint32, offsets *int, strides *int32) {
	C.glowVertexArrayVertexBuffers(gpVertexArrayVertexBuffers, (C.GLuint)(vaobj), (C.GLuint)(first), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(buffers)), (*C.GLintptr)(unsafe.Pointer(offsets)), (*C.GLsizei)(unsafe.Pointer(strides)))
}
func VertexAttrib1d(index uint32, x float64) {
	C.glowVertexAttrib1d(gpVertexAttrib1d, (C.GLuint)(index), (C.GLdouble)(x))
}
func VertexAttrib1dv(index uint32, v *float64) {
	C.glowVertexAttrib1dv(gpVertexAttrib1dv, (C.GLuint)(index), (*C.GLdouble)(unsafe.Pointer(v)))
}
func VertexAttrib1f(index uint32, x float32) {
	C.glowVertexAttrib1f(gpVertexAttrib1f, (C.GLuint)(index), (C.GLfloat)(x))
}
func VertexAttrib1fv(index uint32, v *float32) {
	C.glowVertexAttrib1fv(gpVertexAttrib1fv, (C.GLuint)(index), (*C.GLfloat)(unsafe.Pointer(v)))
}
func VertexAttrib1s(index uint32, x int16) {
	C.glowVertexAttrib1s(gpVertexAttrib1s, (C.GLuint)(index), (C.GLshort)(x))
}
func VertexAttrib1sv(index uint32, v *int16) {
	C.glowVertexAttrib1sv(gpVertexAttrib1sv, (C.GLuint)(index), (*C.GLshort)(unsafe.Pointer(v)))
}
func VertexAttrib2d(index uint32, x float64, y float64) {
	C.glowVertexAttrib2d(gpVertexAttrib2d, (C.GLuint)(index), (C.GLdouble)(x), (C.GLdouble)(y))
}
func VertexAttrib2dv(index uint32, v *float64) {
	C.glowVertexAttrib2dv(gpVertexAttrib2dv, (C.GLuint)(index), (*C.GLdouble)(unsafe.Pointer(v)))
}
func VertexAttrib2f(index uint32, x float32, y float32) {
	C.glowVertexAttrib2f(gpVertexAttrib2f, (C.GLuint)(index), (C.GLfloat)(x), (C.GLfloat)(y))
}
func VertexAttrib2fv(index uint32, v *float32) {
	C.glowVertexAttrib2fv(gpVertexAttrib2fv, (C.GLuint)(index), (*C.GLfloat)(unsafe.Pointer(v)))
}
func VertexAttrib2s(index uint32, x int16, y int16) {
	C.glowVertexAttrib2s(gpVertexAttrib2s, (C.GLuint)(index), (C.GLshort)(x), (C.GLshort)(y))
}
func VertexAttrib2sv(index uint32, v *int16) {
	C.glowVertexAttrib2sv(gpVertexAttrib2sv, (C.GLuint)(index), (*C.GLshort)(unsafe.Pointer(v)))
}
func VertexAttrib3d(index uint32, x float64, y float64, z float64) {
	C.glowVertexAttrib3d(gpVertexAttrib3d, (C.GLuint)(index), (C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z))
}
func VertexAttrib3dv(index uint32, v *float64) {
	C.glowVertexAttrib3dv(gpVertexAttrib3dv, (C.GLuint)(index), (*C.GLdouble)(unsafe.Pointer(v)))
}
func VertexAttrib3f(index uint32, x float32, y float32, z float32) {
	C.glowVertexAttrib3f(gpVertexAttrib3f, (C.GLuint)(index), (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z))
}
func VertexAttrib3fv(index uint32, v *float32) {
	C.glowVertexAttrib3fv(gpVertexAttrib3fv, (C.GLuint)(index), (*C.GLfloat)(unsafe.Pointer(v)))
}
func VertexAttrib3s(index uint32, x int16, y int16, z int16) {
	C.glowVertexAttrib3s(gpVertexAttrib3s, (C.GLuint)(index), (C.GLshort)(x), (C.GLshort)(y), (C.GLshort)(z))
}
func VertexAttrib3sv(index uint32, v *int16) {
	C.glowVertexAttrib3sv(gpVertexAttrib3sv, (C.GLuint)(index), (*C.GLshort)(unsafe.Pointer(v)))
}
func VertexAttrib4Nbv(index uint32, v *int8) {
	C.glowVertexAttrib4Nbv(gpVertexAttrib4Nbv, (C.GLuint)(index), (*C.GLbyte)(unsafe.Pointer(v)))
}
func VertexAttrib4Niv(index uint32, v *int32) {
	C.glowVertexAttrib4Niv(gpVertexAttrib4Niv, (C.GLuint)(index), (*C.GLint)(unsafe.Pointer(v)))
}
func VertexAttrib4Nsv(index uint32, v *int16) {
	C.glowVertexAttrib4Nsv(gpVertexAttrib4Nsv, (C.GLuint)(index), (*C.GLshort)(unsafe.Pointer(v)))
}
func VertexAttrib4Nub(index uint32, x uint8, y uint8, z uint8, w uint8) {
	C.glowVertexAttrib4Nub(gpVertexAttrib4Nub, (C.GLuint)(index), (C.GLubyte)(x), (C.GLubyte)(y), (C.GLubyte)(z), (C.GLubyte)(w))
}
func VertexAttrib4Nubv(index uint32, v *uint8) {
	C.glowVertexAttrib4Nubv(gpVertexAttrib4Nubv, (C.GLuint)(index), (*C.GLubyte)(unsafe.Pointer(v)))
}
func VertexAttrib4Nuiv(index uint32, v *uint32) {
	C.glowVertexAttrib4Nuiv(gpVertexAttrib4Nuiv, (C.GLuint)(index), (*C.GLuint)(unsafe.Pointer(v)))
}
func VertexAttrib4Nusv(index uint32, v *uint16) {
	C.glowVertexAttrib4Nusv(gpVertexAttrib4Nusv, (C.GLuint)(index), (*C.GLushort)(unsafe.Pointer(v)))
}
func VertexAttrib4bv(index uint32, v *int8) {
	C.glowVertexAttrib4bv(gpVertexAttrib4bv, (C.GLuint)(index), (*C.GLbyte)(unsafe.Pointer(v)))
}
func VertexAttrib4d(index uint32, x float64, y float64, z float64, w float64) {
	C.glowVertexAttrib4d(gpVertexAttrib4d, (C.GLuint)(index), (C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z), (C.GLdouble)(w))
}
func VertexAttrib4dv(index uint32, v *float64) {
	C.glowVertexAttrib4dv(gpVertexAttrib4dv, (C.GLuint)(index), (*C.GLdouble)(unsafe.Pointer(v)))
}
func VertexAttrib4f(index uint32, x float32, y float32, z float32, w float32) {
	C.glowVertexAttrib4f(gpVertexAttrib4f, (C.GLuint)(index), (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z), (C.GLfloat)(w))
}
func VertexAttrib4fv(index uint32, v *float32) {
	C.glowVertexAttrib4fv(gpVertexAttrib4fv, (C.GLuint)(index), (*C.GLfloat)(unsafe.Pointer(v)))
}
func VertexAttrib4iv(index uint32, v *int32) {
	C.glowVertexAttrib4iv(gpVertexAttrib4iv, (C.GLuint)(index), (*C.GLint)(unsafe.Pointer(v)))
}
func VertexAttrib4s(index uint32, x int16, y int16, z int16, w int16) {
	C.glowVertexAttrib4s(gpVertexAttrib4s, (C.GLuint)(index), (C.GLshort)(x), (C.GLshort)(y), (C.GLshort)(z), (C.GLshort)(w))
}
func VertexAttrib4sv(index uint32, v *int16) {
	C.glowVertexAttrib4sv(gpVertexAttrib4sv, (C.GLuint)(index), (*C.GLshort)(unsafe.Pointer(v)))
}
func VertexAttrib4ubv(index uint32, v *uint8) {
	C.glowVertexAttrib4ubv(gpVertexAttrib4ubv, (C.GLuint)(index), (*C.GLubyte)(unsafe.Pointer(v)))
}
func VertexAttrib4uiv(index uint32, v *uint32) {
	C.glowVertexAttrib4uiv(gpVertexAttrib4uiv, (C.GLuint)(index), (*C.GLuint)(unsafe.Pointer(v)))
}
func VertexAttrib4usv(index uint32, v *uint16) {
	C.glowVertexAttrib4usv(gpVertexAttrib4usv, (C.GLuint)(index), (*C.GLushort)(unsafe.Pointer(v)))
}

// associate a vertex attribute and a vertex buffer binding for a vertex array object
func VertexAttribBinding(attribindex uint32, bindingindex uint32) {
	C.glowVertexAttribBinding(gpVertexAttribBinding, (C.GLuint)(attribindex), (C.GLuint)(bindingindex))
}

// modify the rate at which generic vertex attributes advance during instanced rendering
func VertexAttribDivisor(index uint32, divisor uint32) {
	C.glowVertexAttribDivisor(gpVertexAttribDivisor, (C.GLuint)(index), (C.GLuint)(divisor))
}

// specify the organization of vertex arrays
func VertexAttribFormat(attribindex uint32, size int32, xtype uint32, normalized bool, relativeoffset uint32) {
	C.glowVertexAttribFormat(gpVertexAttribFormat, (C.GLuint)(attribindex), (C.GLint)(size), (C.GLenum)(xtype), (C.GLboolean)(boolToInt(normalized)), (C.GLuint)(relativeoffset))
}
func VertexAttribI1i(index uint32, x int32) {
	C.glowVertexAttribI1i(gpVertexAttribI1i, (C.GLuint)(index), (C.GLint)(x))
}
func VertexAttribI1iv(index uint32, v *int32) {
	C.glowVertexAttribI1iv(gpVertexAttribI1iv, (C.GLuint)(index), (*C.GLint)(unsafe.Pointer(v)))
}
func VertexAttribI1ui(index uint32, x uint32) {
	C.glowVertexAttribI1ui(gpVertexAttribI1ui, (C.GLuint)(index), (C.GLuint)(x))
}
func VertexAttribI1uiv(index uint32, v *uint32) {
	C.glowVertexAttribI1uiv(gpVertexAttribI1uiv, (C.GLuint)(index), (*C.GLuint)(unsafe.Pointer(v)))
}
func VertexAttribI2i(index uint32, x int32, y int32) {
	C.glowVertexAttribI2i(gpVertexAttribI2i, (C.GLuint)(index), (C.GLint)(x), (C.GLint)(y))
}
func VertexAttribI2iv(index uint32, v *int32) {
	C.glowVertexAttribI2iv(gpVertexAttribI2iv, (C.GLuint)(index), (*C.GLint)(unsafe.Pointer(v)))
}
func VertexAttribI2ui(index uint32, x uint32, y uint32) {
	C.glowVertexAttribI2ui(gpVertexAttribI2ui, (C.GLuint)(index), (C.GLuint)(x), (C.GLuint)(y))
}
func VertexAttribI2uiv(index uint32, v *uint32) {
	C.glowVertexAttribI2uiv(gpVertexAttribI2uiv, (C.GLuint)(index), (*C.GLuint)(unsafe.Pointer(v)))
}
func VertexAttribI3i(index uint32, x int32, y int32, z int32) {
	C.glowVertexAttribI3i(gpVertexAttribI3i, (C.GLuint)(index), (C.GLint)(x), (C.GLint)(y), (C.GLint)(z))
}
func VertexAttribI3iv(index uint32, v *int32) {
	C.glowVertexAttribI3iv(gpVertexAttribI3iv, (C.GLuint)(index), (*C.GLint)(unsafe.Pointer(v)))
}
func VertexAttribI3ui(index uint32, x uint32, y uint32, z uint32) {
	C.glowVertexAttribI3ui(gpVertexAttribI3ui, (C.GLuint)(index), (C.GLuint)(x), (C.GLuint)(y), (C.GLuint)(z))
}
func VertexAttribI3uiv(index uint32, v *uint32) {
	C.glowVertexAttribI3uiv(gpVertexAttribI3uiv, (C.GLuint)(index), (*C.GLuint)(unsafe.Pointer(v)))
}
func VertexAttribI4bv(index uint32, v *int8) {
	C.glowVertexAttribI4bv(gpVertexAttribI4bv, (C.GLuint)(index), (*C.GLbyte)(unsafe.Pointer(v)))
}
func VertexAttribI4i(index uint32, x int32, y int32, z int32, w int32) {
	C.glowVertexAttribI4i(gpVertexAttribI4i, (C.GLuint)(index), (C.GLint)(x), (C.GLint)(y), (C.GLint)(z), (C.GLint)(w))
}
func VertexAttribI4iv(index uint32, v *int32) {
	C.glowVertexAttribI4iv(gpVertexAttribI4iv, (C.GLuint)(index), (*C.GLint)(unsafe.Pointer(v)))
}
func VertexAttribI4sv(index uint32, v *int16) {
	C.glowVertexAttribI4sv(gpVertexAttribI4sv, (C.GLuint)(index), (*C.GLshort)(unsafe.Pointer(v)))
}
func VertexAttribI4ubv(index uint32, v *uint8) {
	C.glowVertexAttribI4ubv(gpVertexAttribI4ubv, (C.GLuint)(index), (*C.GLubyte)(unsafe.Pointer(v)))
}
func VertexAttribI4ui(index uint32, x uint32, y uint32, z uint32, w uint32) {
	C.glowVertexAttribI4ui(gpVertexAttribI4ui, (C.GLuint)(index), (C.GLuint)(x), (C.GLuint)(y), (C.GLuint)(z), (C.GLuint)(w))
}
func VertexAttribI4uiv(index uint32, v *uint32) {
	C.glowVertexAttribI4uiv(gpVertexAttribI4uiv, (C.GLuint)(index), (*C.GLuint)(unsafe.Pointer(v)))
}
func VertexAttribI4usv(index uint32, v *uint16) {
	C.glowVertexAttribI4usv(gpVertexAttribI4usv, (C.GLuint)(index), (*C.GLushort)(unsafe.Pointer(v)))
}
func VertexAttribIFormat(attribindex uint32, size int32, xtype uint32, relativeoffset uint32) {
	C.glowVertexAttribIFormat(gpVertexAttribIFormat, (C.GLuint)(attribindex), (C.GLint)(size), (C.GLenum)(xtype), (C.GLuint)(relativeoffset))
}
func VertexAttribIPointer(index uint32, size int32, xtype uint32, stride int32, pointer unsafe.Pointer) {
	C.glowVertexAttribIPointer(gpVertexAttribIPointer, (C.GLuint)(index), (C.GLint)(size), (C.GLenum)(xtype), (C.GLsizei)(stride), pointer)
}
func VertexAttribL1d(index uint32, x float64) {
	C.glowVertexAttribL1d(gpVertexAttribL1d, (C.GLuint)(index), (C.GLdouble)(x))
}
func VertexAttribL1dv(index uint32, v *float64) {
	C.glowVertexAttribL1dv(gpVertexAttribL1dv, (C.GLuint)(index), (*C.GLdouble)(unsafe.Pointer(v)))
}
func VertexAttribL1ui64ARB(index uint32, x uint64) {
	C.glowVertexAttribL1ui64ARB(gpVertexAttribL1ui64ARB, (C.GLuint)(index), (C.GLuint64EXT)(x))
}
func VertexAttribL1ui64vARB(index uint32, v *uint64) {
	C.glowVertexAttribL1ui64vARB(gpVertexAttribL1ui64vARB, (C.GLuint)(index), (*C.GLuint64EXT)(unsafe.Pointer(v)))
}
func VertexAttribL2d(index uint32, x float64, y float64) {
	C.glowVertexAttribL2d(gpVertexAttribL2d, (C.GLuint)(index), (C.GLdouble)(x), (C.GLdouble)(y))
}
func VertexAttribL2dv(index uint32, v *float64) {
	C.glowVertexAttribL2dv(gpVertexAttribL2dv, (C.GLuint)(index), (*C.GLdouble)(unsafe.Pointer(v)))
}
func VertexAttribL3d(index uint32, x float64, y float64, z float64) {
	C.glowVertexAttribL3d(gpVertexAttribL3d, (C.GLuint)(index), (C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z))
}
func VertexAttribL3dv(index uint32, v *float64) {
	C.glowVertexAttribL3dv(gpVertexAttribL3dv, (C.GLuint)(index), (*C.GLdouble)(unsafe.Pointer(v)))
}
func VertexAttribL4d(index uint32, x float64, y float64, z float64, w float64) {
	C.glowVertexAttribL4d(gpVertexAttribL4d, (C.GLuint)(index), (C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z), (C.GLdouble)(w))
}
func VertexAttribL4dv(index uint32, v *float64) {
	C.glowVertexAttribL4dv(gpVertexAttribL4dv, (C.GLuint)(index), (*C.GLdouble)(unsafe.Pointer(v)))
}
func VertexAttribLFormat(attribindex uint32, size int32, xtype uint32, relativeoffset uint32) {
	C.glowVertexAttribLFormat(gpVertexAttribLFormat, (C.GLuint)(attribindex), (C.GLint)(size), (C.GLenum)(xtype), (C.GLuint)(relativeoffset))
}
func VertexAttribLPointer(index uint32, size int32, xtype uint32, stride int32, pointer unsafe.Pointer) {
	C.glowVertexAttribLPointer(gpVertexAttribLPointer, (C.GLuint)(index), (C.GLint)(size), (C.GLenum)(xtype), (C.GLsizei)(stride), pointer)
}
func VertexAttribP1ui(index uint32, xtype uint32, normalized bool, value uint32) {
	C.glowVertexAttribP1ui(gpVertexAttribP1ui, (C.GLuint)(index), (C.GLenum)(xtype), (C.GLboolean)(boolToInt(normalized)), (C.GLuint)(value))
}
func VertexAttribP1uiv(index uint32, xtype uint32, normalized bool, value *uint32) {
	C.glowVertexAttribP1uiv(gpVertexAttribP1uiv, (C.GLuint)(index), (C.GLenum)(xtype), (C.GLboolean)(boolToInt(normalized)), (*C.GLuint)(unsafe.Pointer(value)))
}
func VertexAttribP2ui(index uint32, xtype uint32, normalized bool, value uint32) {
	C.glowVertexAttribP2ui(gpVertexAttribP2ui, (C.GLuint)(index), (C.GLenum)(xtype), (C.GLboolean)(boolToInt(normalized)), (C.GLuint)(value))
}
func VertexAttribP2uiv(index uint32, xtype uint32, normalized bool, value *uint32) {
	C.glowVertexAttribP2uiv(gpVertexAttribP2uiv, (C.GLuint)(index), (C.GLenum)(xtype), (C.GLboolean)(boolToInt(normalized)), (*C.GLuint)(unsafe.Pointer(value)))
}
func VertexAttribP3ui(index uint32, xtype uint32, normalized bool, value uint32) {
	C.glowVertexAttribP3ui(gpVertexAttribP3ui, (C.GLuint)(index), (C.GLenum)(xtype), (C.GLboolean)(boolToInt(normalized)), (C.GLuint)(value))
}
func VertexAttribP3uiv(index uint32, xtype uint32, normalized bool, value *uint32) {
	C.glowVertexAttribP3uiv(gpVertexAttribP3uiv, (C.GLuint)(index), (C.GLenum)(xtype), (C.GLboolean)(boolToInt(normalized)), (*C.GLuint)(unsafe.Pointer(value)))
}
func VertexAttribP4ui(index uint32, xtype uint32, normalized bool, value uint32) {
	C.glowVertexAttribP4ui(gpVertexAttribP4ui, (C.GLuint)(index), (C.GLenum)(xtype), (C.GLboolean)(boolToInt(normalized)), (C.GLuint)(value))
}
func VertexAttribP4uiv(index uint32, xtype uint32, normalized bool, value *uint32) {
	C.glowVertexAttribP4uiv(gpVertexAttribP4uiv, (C.GLuint)(index), (C.GLenum)(xtype), (C.GLboolean)(boolToInt(normalized)), (*C.GLuint)(unsafe.Pointer(value)))
}

// define an array of generic vertex attribute data
func VertexAttribPointer(index uint32, size int32, xtype uint32, normalized bool, stride int32, pointer unsafe.Pointer) {
	C.glowVertexAttribPointer(gpVertexAttribPointer, (C.GLuint)(index), (C.GLint)(size), (C.GLenum)(xtype), (C.GLboolean)(boolToInt(normalized)), (C.GLsizei)(stride), pointer)
}

// modify the rate at which generic vertex attributes     advance
func VertexBindingDivisor(bindingindex uint32, divisor uint32) {
	C.glowVertexBindingDivisor(gpVertexBindingDivisor, (C.GLuint)(bindingindex), (C.GLuint)(divisor))
}

// define an array of vertex data
func VertexPointer(size int32, xtype uint32, stride int32, pointer unsafe.Pointer) {
	C.glowVertexPointer(gpVertexPointer, (C.GLint)(size), (C.GLenum)(xtype), (C.GLsizei)(stride), pointer)
}

// set the viewport
func Viewport(x int32, y int32, width int32, height int32) {
	C.glowViewport(gpViewport, (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}
func ViewportArrayv(first uint32, count int32, v *float32) {
	C.glowViewportArrayv(gpViewportArrayv, (C.GLuint)(first), (C.GLsizei)(count), (*C.GLfloat)(unsafe.Pointer(v)))
}
func ViewportIndexedf(index uint32, x float32, y float32, w float32, h float32) {
	C.glowViewportIndexedf(gpViewportIndexedf, (C.GLuint)(index), (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(w), (C.GLfloat)(h))
}
func ViewportIndexedfv(index uint32, v *float32) {
	C.glowViewportIndexedfv(gpViewportIndexedfv, (C.GLuint)(index), (*C.GLfloat)(unsafe.Pointer(v)))
}

// instruct the GL server to block until the specified sync object becomes signaled
func WaitSync(sync uintptr, flags uint32, timeout uint64) {
	C.glowWaitSync(gpWaitSync, (C.GLsync)(sync), (C.GLbitfield)(flags), (C.GLuint64)(timeout))
}
func WindowPos2d(x float64, y float64) {
	C.glowWindowPos2d(gpWindowPos2d, (C.GLdouble)(x), (C.GLdouble)(y))
}
func WindowPos2dv(v *float64) {
	C.glowWindowPos2dv(gpWindowPos2dv, (*C.GLdouble)(unsafe.Pointer(v)))
}
func WindowPos2f(x float32, y float32) {
	C.glowWindowPos2f(gpWindowPos2f, (C.GLfloat)(x), (C.GLfloat)(y))
}
func WindowPos2fv(v *float32) {
	C.glowWindowPos2fv(gpWindowPos2fv, (*C.GLfloat)(unsafe.Pointer(v)))
}
func WindowPos2i(x int32, y int32) {
	C.glowWindowPos2i(gpWindowPos2i, (C.GLint)(x), (C.GLint)(y))
}
func WindowPos2iv(v *int32) {
	C.glowWindowPos2iv(gpWindowPos2iv, (*C.GLint)(unsafe.Pointer(v)))
}
func WindowPos2s(x int16, y int16) {
	C.glowWindowPos2s(gpWindowPos2s, (C.GLshort)(x), (C.GLshort)(y))
}
func WindowPos2sv(v *int16) {
	C.glowWindowPos2sv(gpWindowPos2sv, (*C.GLshort)(unsafe.Pointer(v)))
}
func WindowPos3d(x float64, y float64, z float64) {
	C.glowWindowPos3d(gpWindowPos3d, (C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z))
}
func WindowPos3dv(v *float64) {
	C.glowWindowPos3dv(gpWindowPos3dv, (*C.GLdouble)(unsafe.Pointer(v)))
}
func WindowPos3f(x float32, y float32, z float32) {
	C.glowWindowPos3f(gpWindowPos3f, (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z))
}
func WindowPos3fv(v *float32) {
	C.glowWindowPos3fv(gpWindowPos3fv, (*C.GLfloat)(unsafe.Pointer(v)))
}
func WindowPos3i(x int32, y int32, z int32) {
	C.glowWindowPos3i(gpWindowPos3i, (C.GLint)(x), (C.GLint)(y), (C.GLint)(z))
}
func WindowPos3iv(v *int32) {
	C.glowWindowPos3iv(gpWindowPos3iv, (*C.GLint)(unsafe.Pointer(v)))
}
func WindowPos3s(x int16, y int16, z int16) {
	C.glowWindowPos3s(gpWindowPos3s, (C.GLshort)(x), (C.GLshort)(y), (C.GLshort)(z))
}
func WindowPos3sv(v *int16) {
	C.glowWindowPos3sv(gpWindowPos3sv, (*C.GLshort)(unsafe.Pointer(v)))
}

// Init initializes the OpenGL bindings by loading the function pointers (for
// each OpenGL function) from the active OpenGL context.
//
// It must be called under the presence of an active OpenGL context, e.g.,
// always after calling window.MakeContextCurrent() and always before calling
// any OpenGL functions exported by this package.
//
// On Windows, Init loads pointers that are context-specific (and hence you
// must re-init if switching between OpenGL contexts, although not calling Init
// again after switching between OpenGL contexts may work if the contexts belong
// to the same graphics driver/device).
//
// On OS X and Linux, the behavior is different, but code written compatible
// with the Windows behavior is compatible with OS X and Linux. That is, always
// Init under an active OpenGL context, and always re-init after switching
// graphics contexts.
//
// For information about caveats of Init, you should read the "Platform Specific
// Function Retrieval" section of https://www.opengl.org/wiki/Load_OpenGL_Functions.
func Init() error {
	return InitWithProcAddrFunc(getProcAddress)
}

// InitWithProcAddrFunc intializes the package using the specified OpenGL
// function pointer loading function. For more cases Init should be used
// instead.
func InitWithProcAddrFunc(getProcAddr func(name string) unsafe.Pointer) error {
	gpAccum = (C.GPACCUM)(getProcAddr("glAccum"))
	gpActiveShaderProgram = (C.GPACTIVESHADERPROGRAM)(getProcAddr("glActiveShaderProgram"))
	gpActiveTexture = (C.GPACTIVETEXTURE)(getProcAddr("glActiveTexture"))
	gpAlphaFunc = (C.GPALPHAFUNC)(getProcAddr("glAlphaFunc"))
	gpAreTexturesResident = (C.GPARETEXTURESRESIDENT)(getProcAddr("glAreTexturesResident"))
	gpArrayElement = (C.GPARRAYELEMENT)(getProcAddr("glArrayElement"))
	gpAttachShader = (C.GPATTACHSHADER)(getProcAddr("glAttachShader"))
	gpBegin = (C.GPBEGIN)(getProcAddr("glBegin"))
	gpBeginConditionalRender = (C.GPBEGINCONDITIONALRENDER)(getProcAddr("glBeginConditionalRender"))
	gpBeginQuery = (C.GPBEGINQUERY)(getProcAddr("glBeginQuery"))
	gpBeginQueryIndexed = (C.GPBEGINQUERYINDEXED)(getProcAddr("glBeginQueryIndexed"))
	gpBeginTransformFeedback = (C.GPBEGINTRANSFORMFEEDBACK)(getProcAddr("glBeginTransformFeedback"))
	gpBindAttribLocation = (C.GPBINDATTRIBLOCATION)(getProcAddr("glBindAttribLocation"))
	gpBindBuffer = (C.GPBINDBUFFER)(getProcAddr("glBindBuffer"))
	gpBindBufferBase = (C.GPBINDBUFFERBASE)(getProcAddr("glBindBufferBase"))
	gpBindBufferRange = (C.GPBINDBUFFERRANGE)(getProcAddr("glBindBufferRange"))
	gpBindBuffersBase = (C.GPBINDBUFFERSBASE)(getProcAddr("glBindBuffersBase"))
	gpBindBuffersRange = (C.GPBINDBUFFERSRANGE)(getProcAddr("glBindBuffersRange"))
	gpBindFragDataLocation = (C.GPBINDFRAGDATALOCATION)(getProcAddr("glBindFragDataLocation"))
	gpBindFragDataLocationIndexed = (C.GPBINDFRAGDATALOCATIONINDEXED)(getProcAddr("glBindFragDataLocationIndexed"))
	gpBindFramebuffer = (C.GPBINDFRAMEBUFFER)(getProcAddr("glBindFramebuffer"))
	gpBindImageTexture = (C.GPBINDIMAGETEXTURE)(getProcAddr("glBindImageTexture"))
	gpBindImageTextures = (C.GPBINDIMAGETEXTURES)(getProcAddr("glBindImageTextures"))
	gpBindProgramPipeline = (C.GPBINDPROGRAMPIPELINE)(getProcAddr("glBindProgramPipeline"))
	gpBindRenderbuffer = (C.GPBINDRENDERBUFFER)(getProcAddr("glBindRenderbuffer"))
	gpBindSampler = (C.GPBINDSAMPLER)(getProcAddr("glBindSampler"))
	gpBindSamplers = (C.GPBINDSAMPLERS)(getProcAddr("glBindSamplers"))
	gpBindTexture = (C.GPBINDTEXTURE)(getProcAddr("glBindTexture"))
	gpBindTextureUnit = (C.GPBINDTEXTUREUNIT)(getProcAddr("glBindTextureUnit"))
	gpBindTextures = (C.GPBINDTEXTURES)(getProcAddr("glBindTextures"))
	gpBindTransformFeedback = (C.GPBINDTRANSFORMFEEDBACK)(getProcAddr("glBindTransformFeedback"))
	gpBindVertexArray = (C.GPBINDVERTEXARRAY)(getProcAddr("glBindVertexArray"))
	gpBindVertexBuffer = (C.GPBINDVERTEXBUFFER)(getProcAddr("glBindVertexBuffer"))
	gpBindVertexBuffers = (C.GPBINDVERTEXBUFFERS)(getProcAddr("glBindVertexBuffers"))
	gpBitmap = (C.GPBITMAP)(getProcAddr("glBitmap"))
	gpBlendColor = (C.GPBLENDCOLOR)(getProcAddr("glBlendColor"))
	gpBlendEquation = (C.GPBLENDEQUATION)(getProcAddr("glBlendEquation"))
	gpBlendEquationSeparate = (C.GPBLENDEQUATIONSEPARATE)(getProcAddr("glBlendEquationSeparate"))
	gpBlendEquationSeparatei = (C.GPBLENDEQUATIONSEPARATEI)(getProcAddr("glBlendEquationSeparatei"))
	gpBlendEquationSeparateiARB = (C.GPBLENDEQUATIONSEPARATEIARB)(getProcAddr("glBlendEquationSeparateiARB"))
	gpBlendEquationi = (C.GPBLENDEQUATIONI)(getProcAddr("glBlendEquationi"))
	gpBlendEquationiARB = (C.GPBLENDEQUATIONIARB)(getProcAddr("glBlendEquationiARB"))
	gpBlendFunc = (C.GPBLENDFUNC)(getProcAddr("glBlendFunc"))
	gpBlendFuncSeparate = (C.GPBLENDFUNCSEPARATE)(getProcAddr("glBlendFuncSeparate"))
	gpBlendFuncSeparatei = (C.GPBLENDFUNCSEPARATEI)(getProcAddr("glBlendFuncSeparatei"))
	gpBlendFuncSeparateiARB = (C.GPBLENDFUNCSEPARATEIARB)(getProcAddr("glBlendFuncSeparateiARB"))
	gpBlendFunci = (C.GPBLENDFUNCI)(getProcAddr("glBlendFunci"))
	gpBlendFunciARB = (C.GPBLENDFUNCIARB)(getProcAddr("glBlendFunciARB"))
	gpBlitFramebuffer = (C.GPBLITFRAMEBUFFER)(getProcAddr("glBlitFramebuffer"))
	gpBlitNamedFramebuffer = (C.GPBLITNAMEDFRAMEBUFFER)(getProcAddr("glBlitNamedFramebuffer"))
	gpBufferData = (C.GPBUFFERDATA)(getProcAddr("glBufferData"))
	gpBufferPageCommitmentARB = (C.GPBUFFERPAGECOMMITMENTARB)(getProcAddr("glBufferPageCommitmentARB"))
	gpBufferStorage = (C.GPBUFFERSTORAGE)(getProcAddr("glBufferStorage"))
	gpBufferSubData = (C.GPBUFFERSUBDATA)(getProcAddr("glBufferSubData"))
	gpCallList = (C.GPCALLLIST)(getProcAddr("glCallList"))
	gpCallLists = (C.GPCALLLISTS)(getProcAddr("glCallLists"))
	gpCheckFramebufferStatus = (C.GPCHECKFRAMEBUFFERSTATUS)(getProcAddr("glCheckFramebufferStatus"))
	gpCheckNamedFramebufferStatus = (C.GPCHECKNAMEDFRAMEBUFFERSTATUS)(getProcAddr("glCheckNamedFramebufferStatus"))
	gpClampColor = (C.GPCLAMPCOLOR)(getProcAddr("glClampColor"))
	gpClear = (C.GPCLEAR)(getProcAddr("glClear"))
	gpClearAccum = (C.GPCLEARACCUM)(getProcAddr("glClearAccum"))
	gpClearBufferData = (C.GPCLEARBUFFERDATA)(getProcAddr("glClearBufferData"))
	gpClearBufferSubData = (C.GPCLEARBUFFERSUBDATA)(getProcAddr("glClearBufferSubData"))
	gpClearBufferfi = (C.GPCLEARBUFFERFI)(getProcAddr("glClearBufferfi"))
	gpClearBufferfv = (C.GPCLEARBUFFERFV)(getProcAddr("glClearBufferfv"))
	gpClearBufferiv = (C.GPCLEARBUFFERIV)(getProcAddr("glClearBufferiv"))
	gpClearBufferuiv = (C.GPCLEARBUFFERUIV)(getProcAddr("glClearBufferuiv"))
	gpClearColor = (C.GPCLEARCOLOR)(getProcAddr("glClearColor"))
	gpClearDepth = (C.GPCLEARDEPTH)(getProcAddr("glClearDepth"))
	gpClearDepthf = (C.GPCLEARDEPTHF)(getProcAddr("glClearDepthf"))
	gpClearIndex = (C.GPCLEARINDEX)(getProcAddr("glClearIndex"))
	gpClearNamedBufferData = (C.GPCLEARNAMEDBUFFERDATA)(getProcAddr("glClearNamedBufferData"))
	gpClearNamedBufferSubData = (C.GPCLEARNAMEDBUFFERSUBDATA)(getProcAddr("glClearNamedBufferSubData"))
	gpClearNamedFramebufferfi = (C.GPCLEARNAMEDFRAMEBUFFERFI)(getProcAddr("glClearNamedFramebufferfi"))
	gpClearNamedFramebufferfv = (C.GPCLEARNAMEDFRAMEBUFFERFV)(getProcAddr("glClearNamedFramebufferfv"))
	gpClearNamedFramebufferiv = (C.GPCLEARNAMEDFRAMEBUFFERIV)(getProcAddr("glClearNamedFramebufferiv"))
	gpClearNamedFramebufferuiv = (C.GPCLEARNAMEDFRAMEBUFFERUIV)(getProcAddr("glClearNamedFramebufferuiv"))
	gpClearStencil = (C.GPCLEARSTENCIL)(getProcAddr("glClearStencil"))
	gpClearTexImage = (C.GPCLEARTEXIMAGE)(getProcAddr("glClearTexImage"))
	gpClearTexSubImage = (C.GPCLEARTEXSUBIMAGE)(getProcAddr("glClearTexSubImage"))
	gpClientActiveTexture = (C.GPCLIENTACTIVETEXTURE)(getProcAddr("glClientActiveTexture"))
	gpClientWaitSync = (C.GPCLIENTWAITSYNC)(getProcAddr("glClientWaitSync"))
	gpClipControl = (C.GPCLIPCONTROL)(getProcAddr("glClipControl"))
	gpClipPlane = (C.GPCLIPPLANE)(getProcAddr("glClipPlane"))
	gpColor3b = (C.GPCOLOR3B)(getProcAddr("glColor3b"))
	gpColor3bv = (C.GPCOLOR3BV)(getProcAddr("glColor3bv"))
	gpColor3d = (C.GPCOLOR3D)(getProcAddr("glColor3d"))
	gpColor3dv = (C.GPCOLOR3DV)(getProcAddr("glColor3dv"))
	gpColor3f = (C.GPCOLOR3F)(getProcAddr("glColor3f"))
	gpColor3fv = (C.GPCOLOR3FV)(getProcAddr("glColor3fv"))
	gpColor3i = (C.GPCOLOR3I)(getProcAddr("glColor3i"))
	gpColor3iv = (C.GPCOLOR3IV)(getProcAddr("glColor3iv"))
	gpColor3s = (C.GPCOLOR3S)(getProcAddr("glColor3s"))
	gpColor3sv = (C.GPCOLOR3SV)(getProcAddr("glColor3sv"))
	gpColor3ub = (C.GPCOLOR3UB)(getProcAddr("glColor3ub"))
	gpColor3ubv = (C.GPCOLOR3UBV)(getProcAddr("glColor3ubv"))
	gpColor3ui = (C.GPCOLOR3UI)(getProcAddr("glColor3ui"))
	gpColor3uiv = (C.GPCOLOR3UIV)(getProcAddr("glColor3uiv"))
	gpColor3us = (C.GPCOLOR3US)(getProcAddr("glColor3us"))
	gpColor3usv = (C.GPCOLOR3USV)(getProcAddr("glColor3usv"))
	gpColor4b = (C.GPCOLOR4B)(getProcAddr("glColor4b"))
	gpColor4bv = (C.GPCOLOR4BV)(getProcAddr("glColor4bv"))
	gpColor4d = (C.GPCOLOR4D)(getProcAddr("glColor4d"))
	gpColor4dv = (C.GPCOLOR4DV)(getProcAddr("glColor4dv"))
	gpColor4f = (C.GPCOLOR4F)(getProcAddr("glColor4f"))
	gpColor4fv = (C.GPCOLOR4FV)(getProcAddr("glColor4fv"))
	gpColor4i = (C.GPCOLOR4I)(getProcAddr("glColor4i"))
	gpColor4iv = (C.GPCOLOR4IV)(getProcAddr("glColor4iv"))
	gpColor4s = (C.GPCOLOR4S)(getProcAddr("glColor4s"))
	gpColor4sv = (C.GPCOLOR4SV)(getProcAddr("glColor4sv"))
	gpColor4ub = (C.GPCOLOR4UB)(getProcAddr("glColor4ub"))
	gpColor4ubv = (C.GPCOLOR4UBV)(getProcAddr("glColor4ubv"))
	gpColor4ui = (C.GPCOLOR4UI)(getProcAddr("glColor4ui"))
	gpColor4uiv = (C.GPCOLOR4UIV)(getProcAddr("glColor4uiv"))
	gpColor4us = (C.GPCOLOR4US)(getProcAddr("glColor4us"))
	gpColor4usv = (C.GPCOLOR4USV)(getProcAddr("glColor4usv"))
	gpColorMask = (C.GPCOLORMASK)(getProcAddr("glColorMask"))
	gpColorMaski = (C.GPCOLORMASKI)(getProcAddr("glColorMaski"))
	gpColorMaterial = (C.GPCOLORMATERIAL)(getProcAddr("glColorMaterial"))
	gpColorPointer = (C.GPCOLORPOINTER)(getProcAddr("glColorPointer"))
	gpCompileShader = (C.GPCOMPILESHADER)(getProcAddr("glCompileShader"))
	gpCompileShaderIncludeARB = (C.GPCOMPILESHADERINCLUDEARB)(getProcAddr("glCompileShaderIncludeARB"))
	gpCompressedTexImage1D = (C.GPCOMPRESSEDTEXIMAGE1D)(getProcAddr("glCompressedTexImage1D"))
	gpCompressedTexImage2D = (C.GPCOMPRESSEDTEXIMAGE2D)(getProcAddr("glCompressedTexImage2D"))
	gpCompressedTexImage3D = (C.GPCOMPRESSEDTEXIMAGE3D)(getProcAddr("glCompressedTexImage3D"))
	gpCompressedTexSubImage1D = (C.GPCOMPRESSEDTEXSUBIMAGE1D)(getProcAddr("glCompressedTexSubImage1D"))
	gpCompressedTexSubImage2D = (C.GPCOMPRESSEDTEXSUBIMAGE2D)(getProcAddr("glCompressedTexSubImage2D"))
	gpCompressedTexSubImage3D = (C.GPCOMPRESSEDTEXSUBIMAGE3D)(getProcAddr("glCompressedTexSubImage3D"))
	gpCompressedTextureSubImage1D = (C.GPCOMPRESSEDTEXTURESUBIMAGE1D)(getProcAddr("glCompressedTextureSubImage1D"))
	gpCompressedTextureSubImage2D = (C.GPCOMPRESSEDTEXTURESUBIMAGE2D)(getProcAddr("glCompressedTextureSubImage2D"))
	gpCompressedTextureSubImage3D = (C.GPCOMPRESSEDTEXTURESUBIMAGE3D)(getProcAddr("glCompressedTextureSubImage3D"))
	gpCopyBufferSubData = (C.GPCOPYBUFFERSUBDATA)(getProcAddr("glCopyBufferSubData"))
	gpCopyImageSubData = (C.GPCOPYIMAGESUBDATA)(getProcAddr("glCopyImageSubData"))
	gpCopyNamedBufferSubData = (C.GPCOPYNAMEDBUFFERSUBDATA)(getProcAddr("glCopyNamedBufferSubData"))
	gpCopyPixels = (C.GPCOPYPIXELS)(getProcAddr("glCopyPixels"))
	gpCopyTexImage1D = (C.GPCOPYTEXIMAGE1D)(getProcAddr("glCopyTexImage1D"))
	gpCopyTexImage2D = (C.GPCOPYTEXIMAGE2D)(getProcAddr("glCopyTexImage2D"))
	gpCopyTexSubImage1D = (C.GPCOPYTEXSUBIMAGE1D)(getProcAddr("glCopyTexSubImage1D"))
	gpCopyTexSubImage2D = (C.GPCOPYTEXSUBIMAGE2D)(getProcAddr("glCopyTexSubImage2D"))
	gpCopyTexSubImage3D = (C.GPCOPYTEXSUBIMAGE3D)(getProcAddr("glCopyTexSubImage3D"))
	gpCopyTextureSubImage1D = (C.GPCOPYTEXTURESUBIMAGE1D)(getProcAddr("glCopyTextureSubImage1D"))
	gpCopyTextureSubImage2D = (C.GPCOPYTEXTURESUBIMAGE2D)(getProcAddr("glCopyTextureSubImage2D"))
	gpCopyTextureSubImage3D = (C.GPCOPYTEXTURESUBIMAGE3D)(getProcAddr("glCopyTextureSubImage3D"))
	gpCreateBuffers = (C.GPCREATEBUFFERS)(getProcAddr("glCreateBuffers"))
	gpCreateFramebuffers = (C.GPCREATEFRAMEBUFFERS)(getProcAddr("glCreateFramebuffers"))
	gpCreateProgram = (C.GPCREATEPROGRAM)(getProcAddr("glCreateProgram"))
	gpCreateProgramPipelines = (C.GPCREATEPROGRAMPIPELINES)(getProcAddr("glCreateProgramPipelines"))
	gpCreateQueries = (C.GPCREATEQUERIES)(getProcAddr("glCreateQueries"))
	gpCreateRenderbuffers = (C.GPCREATERENDERBUFFERS)(getProcAddr("glCreateRenderbuffers"))
	gpCreateSamplers = (C.GPCREATESAMPLERS)(getProcAddr("glCreateSamplers"))
	gpCreateShader = (C.GPCREATESHADER)(getProcAddr("glCreateShader"))
	gpCreateShaderProgramv = (C.GPCREATESHADERPROGRAMV)(getProcAddr("glCreateShaderProgramv"))
	gpCreateSyncFromCLeventARB = (C.GPCREATESYNCFROMCLEVENTARB)(getProcAddr("glCreateSyncFromCLeventARB"))
	gpCreateTextures = (C.GPCREATETEXTURES)(getProcAddr("glCreateTextures"))
	gpCreateTransformFeedbacks = (C.GPCREATETRANSFORMFEEDBACKS)(getProcAddr("glCreateTransformFeedbacks"))
	gpCreateVertexArrays = (C.GPCREATEVERTEXARRAYS)(getProcAddr("glCreateVertexArrays"))
	gpCullFace = (C.GPCULLFACE)(getProcAddr("glCullFace"))
	gpDebugMessageCallback = (C.GPDEBUGMESSAGECALLBACK)(getProcAddr("glDebugMessageCallback"))
	gpDebugMessageCallbackARB = (C.GPDEBUGMESSAGECALLBACKARB)(getProcAddr("glDebugMessageCallbackARB"))
	gpDebugMessageCallbackKHR = (C.GPDEBUGMESSAGECALLBACKKHR)(getProcAddr("glDebugMessageCallbackKHR"))
	gpDebugMessageControl = (C.GPDEBUGMESSAGECONTROL)(getProcAddr("glDebugMessageControl"))
	gpDebugMessageControlARB = (C.GPDEBUGMESSAGECONTROLARB)(getProcAddr("glDebugMessageControlARB"))
	gpDebugMessageControlKHR = (C.GPDEBUGMESSAGECONTROLKHR)(getProcAddr("glDebugMessageControlKHR"))
	gpDebugMessageInsert = (C.GPDEBUGMESSAGEINSERT)(getProcAddr("glDebugMessageInsert"))
	gpDebugMessageInsertARB = (C.GPDEBUGMESSAGEINSERTARB)(getProcAddr("glDebugMessageInsertARB"))
	gpDebugMessageInsertKHR = (C.GPDEBUGMESSAGEINSERTKHR)(getProcAddr("glDebugMessageInsertKHR"))
	gpDeleteBuffers = (C.GPDELETEBUFFERS)(getProcAddr("glDeleteBuffers"))
	gpDeleteFramebuffers = (C.GPDELETEFRAMEBUFFERS)(getProcAddr("glDeleteFramebuffers"))
	gpDeleteLists = (C.GPDELETELISTS)(getProcAddr("glDeleteLists"))
	gpDeleteNamedStringARB = (C.GPDELETENAMEDSTRINGARB)(getProcAddr("glDeleteNamedStringARB"))
	gpDeleteProgram = (C.GPDELETEPROGRAM)(getProcAddr("glDeleteProgram"))
	gpDeleteProgramPipelines = (C.GPDELETEPROGRAMPIPELINES)(getProcAddr("glDeleteProgramPipelines"))
	gpDeleteQueries = (C.GPDELETEQUERIES)(getProcAddr("glDeleteQueries"))
	gpDeleteRenderbuffers = (C.GPDELETERENDERBUFFERS)(getProcAddr("glDeleteRenderbuffers"))
	gpDeleteSamplers = (C.GPDELETESAMPLERS)(getProcAddr("glDeleteSamplers"))
	gpDeleteShader = (C.GPDELETESHADER)(getProcAddr("glDeleteShader"))
	gpDeleteSync = (C.GPDELETESYNC)(getProcAddr("glDeleteSync"))
	gpDeleteTextures = (C.GPDELETETEXTURES)(getProcAddr("glDeleteTextures"))
	gpDeleteTransformFeedbacks = (C.GPDELETETRANSFORMFEEDBACKS)(getProcAddr("glDeleteTransformFeedbacks"))
	gpDeleteVertexArrays = (C.GPDELETEVERTEXARRAYS)(getProcAddr("glDeleteVertexArrays"))
	gpDepthFunc = (C.GPDEPTHFUNC)(getProcAddr("glDepthFunc"))
	gpDepthMask = (C.GPDEPTHMASK)(getProcAddr("glDepthMask"))
	gpDepthRange = (C.GPDEPTHRANGE)(getProcAddr("glDepthRange"))
	gpDepthRangeArrayv = (C.GPDEPTHRANGEARRAYV)(getProcAddr("glDepthRangeArrayv"))
	gpDepthRangeIndexed = (C.GPDEPTHRANGEINDEXED)(getProcAddr("glDepthRangeIndexed"))
	gpDepthRangef = (C.GPDEPTHRANGEF)(getProcAddr("glDepthRangef"))
	gpDetachShader = (C.GPDETACHSHADER)(getProcAddr("glDetachShader"))
	gpDisable = (C.GPDISABLE)(getProcAddr("glDisable"))
	gpDisableClientState = (C.GPDISABLECLIENTSTATE)(getProcAddr("glDisableClientState"))
	gpDisableVertexArrayAttrib = (C.GPDISABLEVERTEXARRAYATTRIB)(getProcAddr("glDisableVertexArrayAttrib"))
	gpDisableVertexAttribArray = (C.GPDISABLEVERTEXATTRIBARRAY)(getProcAddr("glDisableVertexAttribArray"))
	gpDisablei = (C.GPDISABLEI)(getProcAddr("glDisablei"))
	gpDispatchCompute = (C.GPDISPATCHCOMPUTE)(getProcAddr("glDispatchCompute"))
	gpDispatchComputeGroupSizeARB = (C.GPDISPATCHCOMPUTEGROUPSIZEARB)(getProcAddr("glDispatchComputeGroupSizeARB"))
	gpDispatchComputeIndirect = (C.GPDISPATCHCOMPUTEINDIRECT)(getProcAddr("glDispatchComputeIndirect"))
	gpDrawArrays = (C.GPDRAWARRAYS)(getProcAddr("glDrawArrays"))
	gpDrawArraysIndirect = (C.GPDRAWARRAYSINDIRECT)(getProcAddr("glDrawArraysIndirect"))
	gpDrawArraysInstanced = (C.GPDRAWARRAYSINSTANCED)(getProcAddr("glDrawArraysInstanced"))
	gpDrawArraysInstancedBaseInstance = (C.GPDRAWARRAYSINSTANCEDBASEINSTANCE)(getProcAddr("glDrawArraysInstancedBaseInstance"))
	gpDrawBuffer = (C.GPDRAWBUFFER)(getProcAddr("glDrawBuffer"))
	gpDrawBuffers = (C.GPDRAWBUFFERS)(getProcAddr("glDrawBuffers"))
	gpDrawElements = (C.GPDRAWELEMENTS)(getProcAddr("glDrawElements"))
	gpDrawElementsBaseVertex = (C.GPDRAWELEMENTSBASEVERTEX)(getProcAddr("glDrawElementsBaseVertex"))
	gpDrawElementsIndirect = (C.GPDRAWELEMENTSINDIRECT)(getProcAddr("glDrawElementsIndirect"))
	gpDrawElementsInstanced = (C.GPDRAWELEMENTSINSTANCED)(getProcAddr("glDrawElementsInstanced"))
	gpDrawElementsInstancedBaseInstance = (C.GPDRAWELEMENTSINSTANCEDBASEINSTANCE)(getProcAddr("glDrawElementsInstancedBaseInstance"))
	gpDrawElementsInstancedBaseVertex = (C.GPDRAWELEMENTSINSTANCEDBASEVERTEX)(getProcAddr("glDrawElementsInstancedBaseVertex"))
	gpDrawElementsInstancedBaseVertexBaseInstance = (C.GPDRAWELEMENTSINSTANCEDBASEVERTEXBASEINSTANCE)(getProcAddr("glDrawElementsInstancedBaseVertexBaseInstance"))
	gpDrawPixels = (C.GPDRAWPIXELS)(getProcAddr("glDrawPixels"))
	gpDrawRangeElements = (C.GPDRAWRANGEELEMENTS)(getProcAddr("glDrawRangeElements"))
	gpDrawRangeElementsBaseVertex = (C.GPDRAWRANGEELEMENTSBASEVERTEX)(getProcAddr("glDrawRangeElementsBaseVertex"))
	gpDrawTransformFeedback = (C.GPDRAWTRANSFORMFEEDBACK)(getProcAddr("glDrawTransformFeedback"))
	gpDrawTransformFeedbackInstanced = (C.GPDRAWTRANSFORMFEEDBACKINSTANCED)(getProcAddr("glDrawTransformFeedbackInstanced"))
	gpDrawTransformFeedbackStream = (C.GPDRAWTRANSFORMFEEDBACKSTREAM)(getProcAddr("glDrawTransformFeedbackStream"))
	gpDrawTransformFeedbackStreamInstanced = (C.GPDRAWTRANSFORMFEEDBACKSTREAMINSTANCED)(getProcAddr("glDrawTransformFeedbackStreamInstanced"))
	gpEdgeFlag = (C.GPEDGEFLAG)(getProcAddr("glEdgeFlag"))
	gpEdgeFlagPointer = (C.GPEDGEFLAGPOINTER)(getProcAddr("glEdgeFlagPointer"))
	gpEdgeFlagv = (C.GPEDGEFLAGV)(getProcAddr("glEdgeFlagv"))
	gpEnable = (C.GPENABLE)(getProcAddr("glEnable"))
	gpEnableClientState = (C.GPENABLECLIENTSTATE)(getProcAddr("glEnableClientState"))
	gpEnableVertexArrayAttrib = (C.GPENABLEVERTEXARRAYATTRIB)(getProcAddr("glEnableVertexArrayAttrib"))
	gpEnableVertexAttribArray = (C.GPENABLEVERTEXATTRIBARRAY)(getProcAddr("glEnableVertexAttribArray"))
	gpEnablei = (C.GPENABLEI)(getProcAddr("glEnablei"))
	gpEnd = (C.GPEND)(getProcAddr("glEnd"))
	gpEndConditionalRender = (C.GPENDCONDITIONALRENDER)(getProcAddr("glEndConditionalRender"))
	gpEndList = (C.GPENDLIST)(getProcAddr("glEndList"))
	gpEndQuery = (C.GPENDQUERY)(getProcAddr("glEndQuery"))
	gpEndQueryIndexed = (C.GPENDQUERYINDEXED)(getProcAddr("glEndQueryIndexed"))
	gpEndTransformFeedback = (C.GPENDTRANSFORMFEEDBACK)(getProcAddr("glEndTransformFeedback"))
	gpEvalCoord1d = (C.GPEVALCOORD1D)(getProcAddr("glEvalCoord1d"))
	gpEvalCoord1dv = (C.GPEVALCOORD1DV)(getProcAddr("glEvalCoord1dv"))
	gpEvalCoord1f = (C.GPEVALCOORD1F)(getProcAddr("glEvalCoord1f"))
	gpEvalCoord1fv = (C.GPEVALCOORD1FV)(getProcAddr("glEvalCoord1fv"))
	gpEvalCoord2d = (C.GPEVALCOORD2D)(getProcAddr("glEvalCoord2d"))
	gpEvalCoord2dv = (C.GPEVALCOORD2DV)(getProcAddr("glEvalCoord2dv"))
	gpEvalCoord2f = (C.GPEVALCOORD2F)(getProcAddr("glEvalCoord2f"))
	gpEvalCoord2fv = (C.GPEVALCOORD2FV)(getProcAddr("glEvalCoord2fv"))
	gpEvalMesh1 = (C.GPEVALMESH1)(getProcAddr("glEvalMesh1"))
	gpEvalMesh2 = (C.GPEVALMESH2)(getProcAddr("glEvalMesh2"))
	gpEvalPoint1 = (C.GPEVALPOINT1)(getProcAddr("glEvalPoint1"))
	gpEvalPoint2 = (C.GPEVALPOINT2)(getProcAddr("glEvalPoint2"))
	gpFeedbackBuffer = (C.GPFEEDBACKBUFFER)(getProcAddr("glFeedbackBuffer"))
	gpFenceSync = (C.GPFENCESYNC)(getProcAddr("glFenceSync"))
	gpFinish = (C.GPFINISH)(getProcAddr("glFinish"))
	gpFlush = (C.GPFLUSH)(getProcAddr("glFlush"))
	gpFlushMappedBufferRange = (C.GPFLUSHMAPPEDBUFFERRANGE)(getProcAddr("glFlushMappedBufferRange"))
	gpFlushMappedNamedBufferRange = (C.GPFLUSHMAPPEDNAMEDBUFFERRANGE)(getProcAddr("glFlushMappedNamedBufferRange"))
	gpFogCoordPointer = (C.GPFOGCOORDPOINTER)(getProcAddr("glFogCoordPointer"))
	gpFogCoordd = (C.GPFOGCOORDD)(getProcAddr("glFogCoordd"))
	gpFogCoorddv = (C.GPFOGCOORDDV)(getProcAddr("glFogCoorddv"))
	gpFogCoordf = (C.GPFOGCOORDF)(getProcAddr("glFogCoordf"))
	gpFogCoordfv = (C.GPFOGCOORDFV)(getProcAddr("glFogCoordfv"))
	gpFogf = (C.GPFOGF)(getProcAddr("glFogf"))
	gpFogfv = (C.GPFOGFV)(getProcAddr("glFogfv"))
	gpFogi = (C.GPFOGI)(getProcAddr("glFogi"))
	gpFogiv = (C.GPFOGIV)(getProcAddr("glFogiv"))
	gpFramebufferParameteri = (C.GPFRAMEBUFFERPARAMETERI)(getProcAddr("glFramebufferParameteri"))
	gpFramebufferRenderbuffer = (C.GPFRAMEBUFFERRENDERBUFFER)(getProcAddr("glFramebufferRenderbuffer"))
	gpFramebufferTexture = (C.GPFRAMEBUFFERTEXTURE)(getProcAddr("glFramebufferTexture"))
	gpFramebufferTexture1D = (C.GPFRAMEBUFFERTEXTURE1D)(getProcAddr("glFramebufferTexture1D"))
	gpFramebufferTexture2D = (C.GPFRAMEBUFFERTEXTURE2D)(getProcAddr("glFramebufferTexture2D"))
	gpFramebufferTexture3D = (C.GPFRAMEBUFFERTEXTURE3D)(getProcAddr("glFramebufferTexture3D"))
	gpFramebufferTextureLayer = (C.GPFRAMEBUFFERTEXTURELAYER)(getProcAddr("glFramebufferTextureLayer"))
	gpFrontFace = (C.GPFRONTFACE)(getProcAddr("glFrontFace"))
	gpFrustum = (C.GPFRUSTUM)(getProcAddr("glFrustum"))
	gpGenBuffers = (C.GPGENBUFFERS)(getProcAddr("glGenBuffers"))
	gpGenFramebuffers = (C.GPGENFRAMEBUFFERS)(getProcAddr("glGenFramebuffers"))
	gpGenLists = (C.GPGENLISTS)(getProcAddr("glGenLists"))
	gpGenProgramPipelines = (C.GPGENPROGRAMPIPELINES)(getProcAddr("glGenProgramPipelines"))
	gpGenQueries = (C.GPGENQUERIES)(getProcAddr("glGenQueries"))
	gpGenRenderbuffers = (C.GPGENRENDERBUFFERS)(getProcAddr("glGenRenderbuffers"))
	gpGenSamplers = (C.GPGENSAMPLERS)(getProcAddr("glGenSamplers"))
	gpGenTextures = (C.GPGENTEXTURES)(getProcAddr("glGenTextures"))
	gpGenTransformFeedbacks = (C.GPGENTRANSFORMFEEDBACKS)(getProcAddr("glGenTransformFeedbacks"))
	gpGenVertexArrays = (C.GPGENVERTEXARRAYS)(getProcAddr("glGenVertexArrays"))
	gpGenerateMipmap = (C.GPGENERATEMIPMAP)(getProcAddr("glGenerateMipmap"))
	gpGenerateTextureMipmap = (C.GPGENERATETEXTUREMIPMAP)(getProcAddr("glGenerateTextureMipmap"))
	gpGetActiveAtomicCounterBufferiv = (C.GPGETACTIVEATOMICCOUNTERBUFFERIV)(getProcAddr("glGetActiveAtomicCounterBufferiv"))
	gpGetActiveAttrib = (C.GPGETACTIVEATTRIB)(getProcAddr("glGetActiveAttrib"))
	gpGetActiveSubroutineName = (C.GPGETACTIVESUBROUTINENAME)(getProcAddr("glGetActiveSubroutineName"))
	gpGetActiveSubroutineUniformName = (C.GPGETACTIVESUBROUTINEUNIFORMNAME)(getProcAddr("glGetActiveSubroutineUniformName"))
	gpGetActiveSubroutineUniformiv = (C.GPGETACTIVESUBROUTINEUNIFORMIV)(getProcAddr("glGetActiveSubroutineUniformiv"))
	gpGetActiveUniform = (C.GPGETACTIVEUNIFORM)(getProcAddr("glGetActiveUniform"))
	gpGetActiveUniformBlockName = (C.GPGETACTIVEUNIFORMBLOCKNAME)(getProcAddr("glGetActiveUniformBlockName"))
	gpGetActiveUniformBlockiv = (C.GPGETACTIVEUNIFORMBLOCKIV)(getProcAddr("glGetActiveUniformBlockiv"))
	gpGetActiveUniformName = (C.GPGETACTIVEUNIFORMNAME)(getProcAddr("glGetActiveUniformName"))
	gpGetActiveUniformsiv = (C.GPGETACTIVEUNIFORMSIV)(getProcAddr("glGetActiveUniformsiv"))
	gpGetAttachedShaders = (C.GPGETATTACHEDSHADERS)(getProcAddr("glGetAttachedShaders"))
	gpGetAttribLocation = (C.GPGETATTRIBLOCATION)(getProcAddr("glGetAttribLocation"))
	gpGetBooleani_v = (C.GPGETBOOLEANI_V)(getProcAddr("glGetBooleani_v"))
	gpGetBooleanv = (C.GPGETBOOLEANV)(getProcAddr("glGetBooleanv"))
	gpGetBufferParameteri64v = (C.GPGETBUFFERPARAMETERI64V)(getProcAddr("glGetBufferParameteri64v"))
	gpGetBufferParameteriv = (C.GPGETBUFFERPARAMETERIV)(getProcAddr("glGetBufferParameteriv"))
	gpGetBufferPointerv = (C.GPGETBUFFERPOINTERV)(getProcAddr("glGetBufferPointerv"))
	gpGetBufferSubData = (C.GPGETBUFFERSUBDATA)(getProcAddr("glGetBufferSubData"))
	gpGetClipPlane = (C.GPGETCLIPPLANE)(getProcAddr("glGetClipPlane"))
	gpGetCompressedTexImage = (C.GPGETCOMPRESSEDTEXIMAGE)(getProcAddr("glGetCompressedTexImage"))
	gpGetCompressedTextureImage = (C.GPGETCOMPRESSEDTEXTUREIMAGE)(getProcAddr("glGetCompressedTextureImage"))
	gpGetCompressedTextureSubImage = (C.GPGETCOMPRESSEDTEXTURESUBIMAGE)(getProcAddr("glGetCompressedTextureSubImage"))
	gpGetDebugMessageLog = (C.GPGETDEBUGMESSAGELOG)(getProcAddr("glGetDebugMessageLog"))
	gpGetDebugMessageLogARB = (C.GPGETDEBUGMESSAGELOGARB)(getProcAddr("glGetDebugMessageLogARB"))
	gpGetDebugMessageLogKHR = (C.GPGETDEBUGMESSAGELOGKHR)(getProcAddr("glGetDebugMessageLogKHR"))
	gpGetDoublei_v = (C.GPGETDOUBLEI_V)(getProcAddr("glGetDoublei_v"))
	gpGetDoublev = (C.GPGETDOUBLEV)(getProcAddr("glGetDoublev"))
	gpGetError = (C.GPGETERROR)(getProcAddr("glGetError"))
	gpGetFloati_v = (C.GPGETFLOATI_V)(getProcAddr("glGetFloati_v"))
	gpGetFloatv = (C.GPGETFLOATV)(getProcAddr("glGetFloatv"))
	gpGetFragDataIndex = (C.GPGETFRAGDATAINDEX)(getProcAddr("glGetFragDataIndex"))
	gpGetFragDataLocation = (C.GPGETFRAGDATALOCATION)(getProcAddr("glGetFragDataLocation"))
	gpGetFramebufferAttachmentParameteriv = (C.GPGETFRAMEBUFFERATTACHMENTPARAMETERIV)(getProcAddr("glGetFramebufferAttachmentParameteriv"))
	gpGetFramebufferParameteriv = (C.GPGETFRAMEBUFFERPARAMETERIV)(getProcAddr("glGetFramebufferParameteriv"))
	gpGetGraphicsResetStatus = (C.GPGETGRAPHICSRESETSTATUS)(getProcAddr("glGetGraphicsResetStatus"))
	gpGetGraphicsResetStatusARB = (C.GPGETGRAPHICSRESETSTATUSARB)(getProcAddr("glGetGraphicsResetStatusARB"))
	gpGetGraphicsResetStatusKHR = (C.GPGETGRAPHICSRESETSTATUSKHR)(getProcAddr("glGetGraphicsResetStatusKHR"))
	gpGetImageHandleARB = (C.GPGETIMAGEHANDLEARB)(getProcAddr("glGetImageHandleARB"))
	gpGetInteger64i_v = (C.GPGETINTEGER64I_V)(getProcAddr("glGetInteger64i_v"))
	gpGetInteger64v = (C.GPGETINTEGER64V)(getProcAddr("glGetInteger64v"))
	gpGetIntegeri_v = (C.GPGETINTEGERI_V)(getProcAddr("glGetIntegeri_v"))
	gpGetIntegerv = (C.GPGETINTEGERV)(getProcAddr("glGetIntegerv"))
	gpGetInternalformati64v = (C.GPGETINTERNALFORMATI64V)(getProcAddr("glGetInternalformati64v"))
	gpGetInternalformativ = (C.GPGETINTERNALFORMATIV)(getProcAddr("glGetInternalformativ"))
	gpGetLightfv = (C.GPGETLIGHTFV)(getProcAddr("glGetLightfv"))
	gpGetLightiv = (C.GPGETLIGHTIV)(getProcAddr("glGetLightiv"))
	gpGetMapdv = (C.GPGETMAPDV)(getProcAddr("glGetMapdv"))
	gpGetMapfv = (C.GPGETMAPFV)(getProcAddr("glGetMapfv"))
	gpGetMapiv = (C.GPGETMAPIV)(getProcAddr("glGetMapiv"))
	gpGetMaterialfv = (C.GPGETMATERIALFV)(getProcAddr("glGetMaterialfv"))
	gpGetMaterialiv = (C.GPGETMATERIALIV)(getProcAddr("glGetMaterialiv"))
	gpGetMultisamplefv = (C.GPGETMULTISAMPLEFV)(getProcAddr("glGetMultisamplefv"))
	gpGetNamedBufferParameteri64v = (C.GPGETNAMEDBUFFERPARAMETERI64V)(getProcAddr("glGetNamedBufferParameteri64v"))
	gpGetNamedBufferParameteriv = (C.GPGETNAMEDBUFFERPARAMETERIV)(getProcAddr("glGetNamedBufferParameteriv"))
	gpGetNamedBufferPointerv = (C.GPGETNAMEDBUFFERPOINTERV)(getProcAddr("glGetNamedBufferPointerv"))
	gpGetNamedBufferSubData = (C.GPGETNAMEDBUFFERSUBDATA)(getProcAddr("glGetNamedBufferSubData"))
	gpGetNamedFramebufferAttachmentParameteriv = (C.GPGETNAMEDFRAMEBUFFERATTACHMENTPARAMETERIV)(getProcAddr("glGetNamedFramebufferAttachmentParameteriv"))
	gpGetNamedFramebufferParameteriv = (C.GPGETNAMEDFRAMEBUFFERPARAMETERIV)(getProcAddr("glGetNamedFramebufferParameteriv"))
	gpGetNamedRenderbufferParameteriv = (C.GPGETNAMEDRENDERBUFFERPARAMETERIV)(getProcAddr("glGetNamedRenderbufferParameteriv"))
	gpGetNamedStringARB = (C.GPGETNAMEDSTRINGARB)(getProcAddr("glGetNamedStringARB"))
	gpGetNamedStringivARB = (C.GPGETNAMEDSTRINGIVARB)(getProcAddr("glGetNamedStringivARB"))
	gpGetObjectLabel = (C.GPGETOBJECTLABEL)(getProcAddr("glGetObjectLabel"))
	gpGetObjectLabelKHR = (C.GPGETOBJECTLABELKHR)(getProcAddr("glGetObjectLabelKHR"))
	gpGetObjectPtrLabel = (C.GPGETOBJECTPTRLABEL)(getProcAddr("glGetObjectPtrLabel"))
	gpGetObjectPtrLabelKHR = (C.GPGETOBJECTPTRLABELKHR)(getProcAddr("glGetObjectPtrLabelKHR"))
	gpGetPixelMapfv = (C.GPGETPIXELMAPFV)(getProcAddr("glGetPixelMapfv"))
	gpGetPixelMapuiv = (C.GPGETPIXELMAPUIV)(getProcAddr("glGetPixelMapuiv"))
	gpGetPixelMapusv = (C.GPGETPIXELMAPUSV)(getProcAddr("glGetPixelMapusv"))
	gpGetPointerv = (C.GPGETPOINTERV)(getProcAddr("glGetPointerv"))
	gpGetPointervKHR = (C.GPGETPOINTERVKHR)(getProcAddr("glGetPointervKHR"))
	gpGetPolygonStipple = (C.GPGETPOLYGONSTIPPLE)(getProcAddr("glGetPolygonStipple"))
	gpGetProgramBinary = (C.GPGETPROGRAMBINARY)(getProcAddr("glGetProgramBinary"))
	gpGetProgramInfoLog = (C.GPGETPROGRAMINFOLOG)(getProcAddr("glGetProgramInfoLog"))
	gpGetProgramInterfaceiv = (C.GPGETPROGRAMINTERFACEIV)(getProcAddr("glGetProgramInterfaceiv"))
	gpGetProgramPipelineInfoLog = (C.GPGETPROGRAMPIPELINEINFOLOG)(getProcAddr("glGetProgramPipelineInfoLog"))
	gpGetProgramPipelineiv = (C.GPGETPROGRAMPIPELINEIV)(getProcAddr("glGetProgramPipelineiv"))
	gpGetProgramResourceIndex = (C.GPGETPROGRAMRESOURCEINDEX)(getProcAddr("glGetProgramResourceIndex"))
	gpGetProgramResourceLocation = (C.GPGETPROGRAMRESOURCELOCATION)(getProcAddr("glGetProgramResourceLocation"))
	gpGetProgramResourceLocationIndex = (C.GPGETPROGRAMRESOURCELOCATIONINDEX)(getProcAddr("glGetProgramResourceLocationIndex"))
	gpGetProgramResourceName = (C.GPGETPROGRAMRESOURCENAME)(getProcAddr("glGetProgramResourceName"))
	gpGetProgramResourceiv = (C.GPGETPROGRAMRESOURCEIV)(getProcAddr("glGetProgramResourceiv"))
	gpGetProgramStageiv = (C.GPGETPROGRAMSTAGEIV)(getProcAddr("glGetProgramStageiv"))
	gpGetProgramiv = (C.GPGETPROGRAMIV)(getProcAddr("glGetProgramiv"))
	gpGetQueryIndexediv = (C.GPGETQUERYINDEXEDIV)(getProcAddr("glGetQueryIndexediv"))
	gpGetQueryObjecti64v = (C.GPGETQUERYOBJECTI64V)(getProcAddr("glGetQueryObjecti64v"))
	gpGetQueryObjectiv = (C.GPGETQUERYOBJECTIV)(getProcAddr("glGetQueryObjectiv"))
	gpGetQueryObjectui64v = (C.GPGETQUERYOBJECTUI64V)(getProcAddr("glGetQueryObjectui64v"))
	gpGetQueryObjectuiv = (C.GPGETQUERYOBJECTUIV)(getProcAddr("glGetQueryObjectuiv"))
	gpGetQueryiv = (C.GPGETQUERYIV)(getProcAddr("glGetQueryiv"))
	gpGetRenderbufferParameteriv = (C.GPGETRENDERBUFFERPARAMETERIV)(getProcAddr("glGetRenderbufferParameteriv"))
	gpGetSamplerParameterIiv = (C.GPGETSAMPLERPARAMETERIIV)(getProcAddr("glGetSamplerParameterIiv"))
	gpGetSamplerParameterIuiv = (C.GPGETSAMPLERPARAMETERIUIV)(getProcAddr("glGetSamplerParameterIuiv"))
	gpGetSamplerParameterfv = (C.GPGETSAMPLERPARAMETERFV)(getProcAddr("glGetSamplerParameterfv"))
	gpGetSamplerParameteriv = (C.GPGETSAMPLERPARAMETERIV)(getProcAddr("glGetSamplerParameteriv"))
	gpGetShaderInfoLog = (C.GPGETSHADERINFOLOG)(getProcAddr("glGetShaderInfoLog"))
	gpGetShaderPrecisionFormat = (C.GPGETSHADERPRECISIONFORMAT)(getProcAddr("glGetShaderPrecisionFormat"))
	gpGetShaderSource = (C.GPGETSHADERSOURCE)(getProcAddr("glGetShaderSource"))
	gpGetShaderiv = (C.GPGETSHADERIV)(getProcAddr("glGetShaderiv"))
	gpGetString = (C.GPGETSTRING)(getProcAddr("glGetString"))
	gpGetStringi = (C.GPGETSTRINGI)(getProcAddr("glGetStringi"))
	gpGetSubroutineIndex = (C.GPGETSUBROUTINEINDEX)(getProcAddr("glGetSubroutineIndex"))
	gpGetSubroutineUniformLocation = (C.GPGETSUBROUTINEUNIFORMLOCATION)(getProcAddr("glGetSubroutineUniformLocation"))
	gpGetSynciv = (C.GPGETSYNCIV)(getProcAddr("glGetSynciv"))
	gpGetTexEnvfv = (C.GPGETTEXENVFV)(getProcAddr("glGetTexEnvfv"))
	gpGetTexEnviv = (C.GPGETTEXENVIV)(getProcAddr("glGetTexEnviv"))
	gpGetTexGendv = (C.GPGETTEXGENDV)(getProcAddr("glGetTexGendv"))
	gpGetTexGenfv = (C.GPGETTEXGENFV)(getProcAddr("glGetTexGenfv"))
	gpGetTexGeniv = (C.GPGETTEXGENIV)(getProcAddr("glGetTexGeniv"))
	gpGetTexImage = (C.GPGETTEXIMAGE)(getProcAddr("glGetTexImage"))
	gpGetTexLevelParameterfv = (C.GPGETTEXLEVELPARAMETERFV)(getProcAddr("glGetTexLevelParameterfv"))
	gpGetTexLevelParameteriv = (C.GPGETTEXLEVELPARAMETERIV)(getProcAddr("glGetTexLevelParameteriv"))
	gpGetTexParameterIiv = (C.GPGETTEXPARAMETERIIV)(getProcAddr("glGetTexParameterIiv"))
	gpGetTexParameterIuiv = (C.GPGETTEXPARAMETERIUIV)(getProcAddr("glGetTexParameterIuiv"))
	gpGetTexParameterfv = (C.GPGETTEXPARAMETERFV)(getProcAddr("glGetTexParameterfv"))
	gpGetTexParameteriv = (C.GPGETTEXPARAMETERIV)(getProcAddr("glGetTexParameteriv"))
	gpGetTextureHandleARB = (C.GPGETTEXTUREHANDLEARB)(getProcAddr("glGetTextureHandleARB"))
	gpGetTextureImage = (C.GPGETTEXTUREIMAGE)(getProcAddr("glGetTextureImage"))
	gpGetTextureLevelParameterfv = (C.GPGETTEXTURELEVELPARAMETERFV)(getProcAddr("glGetTextureLevelParameterfv"))
	gpGetTextureLevelParameteriv = (C.GPGETTEXTURELEVELPARAMETERIV)(getProcAddr("glGetTextureLevelParameteriv"))
	gpGetTextureParameterIiv = (C.GPGETTEXTUREPARAMETERIIV)(getProcAddr("glGetTextureParameterIiv"))
	gpGetTextureParameterIuiv = (C.GPGETTEXTUREPARAMETERIUIV)(getProcAddr("glGetTextureParameterIuiv"))
	gpGetTextureParameterfv = (C.GPGETTEXTUREPARAMETERFV)(getProcAddr("glGetTextureParameterfv"))
	gpGetTextureParameteriv = (C.GPGETTEXTUREPARAMETERIV)(getProcAddr("glGetTextureParameteriv"))
	gpGetTextureSamplerHandleARB = (C.GPGETTEXTURESAMPLERHANDLEARB)(getProcAddr("glGetTextureSamplerHandleARB"))
	gpGetTextureSubImage = (C.GPGETTEXTURESUBIMAGE)(getProcAddr("glGetTextureSubImage"))
	gpGetTransformFeedbackVarying = (C.GPGETTRANSFORMFEEDBACKVARYING)(getProcAddr("glGetTransformFeedbackVarying"))
	gpGetTransformFeedbacki64_v = (C.GPGETTRANSFORMFEEDBACKI64_V)(getProcAddr("glGetTransformFeedbacki64_v"))
	gpGetTransformFeedbacki_v = (C.GPGETTRANSFORMFEEDBACKI_V)(getProcAddr("glGetTransformFeedbacki_v"))
	gpGetTransformFeedbackiv = (C.GPGETTRANSFORMFEEDBACKIV)(getProcAddr("glGetTransformFeedbackiv"))
	gpGetUniformBlockIndex = (C.GPGETUNIFORMBLOCKINDEX)(getProcAddr("glGetUniformBlockIndex"))
	gpGetUniformIndices = (C.GPGETUNIFORMINDICES)(getProcAddr("glGetUniformIndices"))
	gpGetUniformLocation = (C.GPGETUNIFORMLOCATION)(getProcAddr("glGetUniformLocation"))
	gpGetUniformSubroutineuiv = (C.GPGETUNIFORMSUBROUTINEUIV)(getProcAddr("glGetUniformSubroutineuiv"))
	gpGetUniformdv = (C.GPGETUNIFORMDV)(getProcAddr("glGetUniformdv"))
	gpGetUniformfv = (C.GPGETUNIFORMFV)(getProcAddr("glGetUniformfv"))
	gpGetUniformiv = (C.GPGETUNIFORMIV)(getProcAddr("glGetUniformiv"))
	gpGetUniformuiv = (C.GPGETUNIFORMUIV)(getProcAddr("glGetUniformuiv"))
	gpGetVertexArrayIndexed64iv = (C.GPGETVERTEXARRAYINDEXED64IV)(getProcAddr("glGetVertexArrayIndexed64iv"))
	gpGetVertexArrayIndexediv = (C.GPGETVERTEXARRAYINDEXEDIV)(getProcAddr("glGetVertexArrayIndexediv"))
	gpGetVertexArrayiv = (C.GPGETVERTEXARRAYIV)(getProcAddr("glGetVertexArrayiv"))
	gpGetVertexAttribIiv = (C.GPGETVERTEXATTRIBIIV)(getProcAddr("glGetVertexAttribIiv"))
	gpGetVertexAttribIuiv = (C.GPGETVERTEXATTRIBIUIV)(getProcAddr("glGetVertexAttribIuiv"))
	gpGetVertexAttribLdv = (C.GPGETVERTEXATTRIBLDV)(getProcAddr("glGetVertexAttribLdv"))
	gpGetVertexAttribLui64vARB = (C.GPGETVERTEXATTRIBLUI64VARB)(getProcAddr("glGetVertexAttribLui64vARB"))
	gpGetVertexAttribPointerv = (C.GPGETVERTEXATTRIBPOINTERV)(getProcAddr("glGetVertexAttribPointerv"))
	gpGetVertexAttribdv = (C.GPGETVERTEXATTRIBDV)(getProcAddr("glGetVertexAttribdv"))
	gpGetVertexAttribfv = (C.GPGETVERTEXATTRIBFV)(getProcAddr("glGetVertexAttribfv"))
	gpGetVertexAttribiv = (C.GPGETVERTEXATTRIBIV)(getProcAddr("glGetVertexAttribiv"))
	gpGetnCompressedTexImage = (C.GPGETNCOMPRESSEDTEXIMAGE)(getProcAddr("glGetnCompressedTexImage"))
	gpGetnCompressedTexImageARB = (C.GPGETNCOMPRESSEDTEXIMAGEARB)(getProcAddr("glGetnCompressedTexImageARB"))
	gpGetnTexImage = (C.GPGETNTEXIMAGE)(getProcAddr("glGetnTexImage"))
	gpGetnTexImageARB = (C.GPGETNTEXIMAGEARB)(getProcAddr("glGetnTexImageARB"))
	gpGetnUniformdv = (C.GPGETNUNIFORMDV)(getProcAddr("glGetnUniformdv"))
	gpGetnUniformdvARB = (C.GPGETNUNIFORMDVARB)(getProcAddr("glGetnUniformdvARB"))
	gpGetnUniformfv = (C.GPGETNUNIFORMFV)(getProcAddr("glGetnUniformfv"))
	gpGetnUniformfvARB = (C.GPGETNUNIFORMFVARB)(getProcAddr("glGetnUniformfvARB"))
	gpGetnUniformfvKHR = (C.GPGETNUNIFORMFVKHR)(getProcAddr("glGetnUniformfvKHR"))
	gpGetnUniformiv = (C.GPGETNUNIFORMIV)(getProcAddr("glGetnUniformiv"))
	gpGetnUniformivARB = (C.GPGETNUNIFORMIVARB)(getProcAddr("glGetnUniformivARB"))
	gpGetnUniformivKHR = (C.GPGETNUNIFORMIVKHR)(getProcAddr("glGetnUniformivKHR"))
	gpGetnUniformuiv = (C.GPGETNUNIFORMUIV)(getProcAddr("glGetnUniformuiv"))
	gpGetnUniformuivARB = (C.GPGETNUNIFORMUIVARB)(getProcAddr("glGetnUniformuivARB"))
	gpGetnUniformuivKHR = (C.GPGETNUNIFORMUIVKHR)(getProcAddr("glGetnUniformuivKHR"))
	gpHint = (C.GPHINT)(getProcAddr("glHint"))
	gpIndexMask = (C.GPINDEXMASK)(getProcAddr("glIndexMask"))
	gpIndexPointer = (C.GPINDEXPOINTER)(getProcAddr("glIndexPointer"))
	gpIndexd = (C.GPINDEXD)(getProcAddr("glIndexd"))
	gpIndexdv = (C.GPINDEXDV)(getProcAddr("glIndexdv"))
	gpIndexf = (C.GPINDEXF)(getProcAddr("glIndexf"))
	gpIndexfv = (C.GPINDEXFV)(getProcAddr("glIndexfv"))
	gpIndexi = (C.GPINDEXI)(getProcAddr("glIndexi"))
	gpIndexiv = (C.GPINDEXIV)(getProcAddr("glIndexiv"))
	gpIndexs = (C.GPINDEXS)(getProcAddr("glIndexs"))
	gpIndexsv = (C.GPINDEXSV)(getProcAddr("glIndexsv"))
	gpIndexub = (C.GPINDEXUB)(getProcAddr("glIndexub"))
	gpIndexubv = (C.GPINDEXUBV)(getProcAddr("glIndexubv"))
	gpInitNames = (C.GPINITNAMES)(getProcAddr("glInitNames"))
	gpInterleavedArrays = (C.GPINTERLEAVEDARRAYS)(getProcAddr("glInterleavedArrays"))
	gpInvalidateBufferData = (C.GPINVALIDATEBUFFERDATA)(getProcAddr("glInvalidateBufferData"))
	gpInvalidateBufferSubData = (C.GPINVALIDATEBUFFERSUBDATA)(getProcAddr("glInvalidateBufferSubData"))
	gpInvalidateFramebuffer = (C.GPINVALIDATEFRAMEBUFFER)(getProcAddr("glInvalidateFramebuffer"))
	gpInvalidateNamedFramebufferData = (C.GPINVALIDATENAMEDFRAMEBUFFERDATA)(getProcAddr("glInvalidateNamedFramebufferData"))
	gpInvalidateNamedFramebufferSubData = (C.GPINVALIDATENAMEDFRAMEBUFFERSUBDATA)(getProcAddr("glInvalidateNamedFramebufferSubData"))
	gpInvalidateSubFramebuffer = (C.GPINVALIDATESUBFRAMEBUFFER)(getProcAddr("glInvalidateSubFramebuffer"))
	gpInvalidateTexImage = (C.GPINVALIDATETEXIMAGE)(getProcAddr("glInvalidateTexImage"))
	gpInvalidateTexSubImage = (C.GPINVALIDATETEXSUBIMAGE)(getProcAddr("glInvalidateTexSubImage"))
	gpIsBuffer = (C.GPISBUFFER)(getProcAddr("glIsBuffer"))
	gpIsEnabled = (C.GPISENABLED)(getProcAddr("glIsEnabled"))
	gpIsEnabledi = (C.GPISENABLEDI)(getProcAddr("glIsEnabledi"))
	gpIsFramebuffer = (C.GPISFRAMEBUFFER)(getProcAddr("glIsFramebuffer"))
	gpIsImageHandleResidentARB = (C.GPISIMAGEHANDLERESIDENTARB)(getProcAddr("glIsImageHandleResidentARB"))
	gpIsList = (C.GPISLIST)(getProcAddr("glIsList"))
	gpIsNamedStringARB = (C.GPISNAMEDSTRINGARB)(getProcAddr("glIsNamedStringARB"))
	gpIsProgram = (C.GPISPROGRAM)(getProcAddr("glIsProgram"))
	gpIsProgramPipeline = (C.GPISPROGRAMPIPELINE)(getProcAddr("glIsProgramPipeline"))
	gpIsQuery = (C.GPISQUERY)(getProcAddr("glIsQuery"))
	gpIsRenderbuffer = (C.GPISRENDERBUFFER)(getProcAddr("glIsRenderbuffer"))
	gpIsSampler = (C.GPISSAMPLER)(getProcAddr("glIsSampler"))
	gpIsShader = (C.GPISSHADER)(getProcAddr("glIsShader"))
	gpIsSync = (C.GPISSYNC)(getProcAddr("glIsSync"))
	gpIsTexture = (C.GPISTEXTURE)(getProcAddr("glIsTexture"))
	gpIsTextureHandleResidentARB = (C.GPISTEXTUREHANDLERESIDENTARB)(getProcAddr("glIsTextureHandleResidentARB"))
	gpIsTransformFeedback = (C.GPISTRANSFORMFEEDBACK)(getProcAddr("glIsTransformFeedback"))
	gpIsVertexArray = (C.GPISVERTEXARRAY)(getProcAddr("glIsVertexArray"))
	gpLightModelf = (C.GPLIGHTMODELF)(getProcAddr("glLightModelf"))
	gpLightModelfv = (C.GPLIGHTMODELFV)(getProcAddr("glLightModelfv"))
	gpLightModeli = (C.GPLIGHTMODELI)(getProcAddr("glLightModeli"))
	gpLightModeliv = (C.GPLIGHTMODELIV)(getProcAddr("glLightModeliv"))
	gpLightf = (C.GPLIGHTF)(getProcAddr("glLightf"))
	gpLightfv = (C.GPLIGHTFV)(getProcAddr("glLightfv"))
	gpLighti = (C.GPLIGHTI)(getProcAddr("glLighti"))
	gpLightiv = (C.GPLIGHTIV)(getProcAddr("glLightiv"))
	gpLineStipple = (C.GPLINESTIPPLE)(getProcAddr("glLineStipple"))
	gpLineWidth = (C.GPLINEWIDTH)(getProcAddr("glLineWidth"))
	gpLinkProgram = (C.GPLINKPROGRAM)(getProcAddr("glLinkProgram"))
	gpListBase = (C.GPLISTBASE)(getProcAddr("glListBase"))
	gpLoadIdentity = (C.GPLOADIDENTITY)(getProcAddr("glLoadIdentity"))
	gpLoadMatrixd = (C.GPLOADMATRIXD)(getProcAddr("glLoadMatrixd"))
	gpLoadMatrixf = (C.GPLOADMATRIXF)(getProcAddr("glLoadMatrixf"))
	gpLoadName = (C.GPLOADNAME)(getProcAddr("glLoadName"))
	gpLoadTransposeMatrixd = (C.GPLOADTRANSPOSEMATRIXD)(getProcAddr("glLoadTransposeMatrixd"))
	gpLoadTransposeMatrixf = (C.GPLOADTRANSPOSEMATRIXF)(getProcAddr("glLoadTransposeMatrixf"))
	gpLogicOp = (C.GPLOGICOP)(getProcAddr("glLogicOp"))
	gpMakeImageHandleNonResidentARB = (C.GPMAKEIMAGEHANDLENONRESIDENTARB)(getProcAddr("glMakeImageHandleNonResidentARB"))
	gpMakeImageHandleResidentARB = (C.GPMAKEIMAGEHANDLERESIDENTARB)(getProcAddr("glMakeImageHandleResidentARB"))
	gpMakeTextureHandleNonResidentARB = (C.GPMAKETEXTUREHANDLENONRESIDENTARB)(getProcAddr("glMakeTextureHandleNonResidentARB"))
	gpMakeTextureHandleResidentARB = (C.GPMAKETEXTUREHANDLERESIDENTARB)(getProcAddr("glMakeTextureHandleResidentARB"))
	gpMap1d = (C.GPMAP1D)(getProcAddr("glMap1d"))
	gpMap1f = (C.GPMAP1F)(getProcAddr("glMap1f"))
	gpMap2d = (C.GPMAP2D)(getProcAddr("glMap2d"))
	gpMap2f = (C.GPMAP2F)(getProcAddr("glMap2f"))
	gpMapBuffer = (C.GPMAPBUFFER)(getProcAddr("glMapBuffer"))
	gpMapBufferRange = (C.GPMAPBUFFERRANGE)(getProcAddr("glMapBufferRange"))
	gpMapGrid1d = (C.GPMAPGRID1D)(getProcAddr("glMapGrid1d"))
	gpMapGrid1f = (C.GPMAPGRID1F)(getProcAddr("glMapGrid1f"))
	gpMapGrid2d = (C.GPMAPGRID2D)(getProcAddr("glMapGrid2d"))
	gpMapGrid2f = (C.GPMAPGRID2F)(getProcAddr("glMapGrid2f"))
	gpMapNamedBuffer = (C.GPMAPNAMEDBUFFER)(getProcAddr("glMapNamedBuffer"))
	gpMapNamedBufferRange = (C.GPMAPNAMEDBUFFERRANGE)(getProcAddr("glMapNamedBufferRange"))
	gpMaterialf = (C.GPMATERIALF)(getProcAddr("glMaterialf"))
	gpMaterialfv = (C.GPMATERIALFV)(getProcAddr("glMaterialfv"))
	gpMateriali = (C.GPMATERIALI)(getProcAddr("glMateriali"))
	gpMaterialiv = (C.GPMATERIALIV)(getProcAddr("glMaterialiv"))
	gpMatrixMode = (C.GPMATRIXMODE)(getProcAddr("glMatrixMode"))
	gpMemoryBarrier = (C.GPMEMORYBARRIER)(getProcAddr("glMemoryBarrier"))
	gpMemoryBarrierByRegion = (C.GPMEMORYBARRIERBYREGION)(getProcAddr("glMemoryBarrierByRegion"))
	gpMinSampleShading = (C.GPMINSAMPLESHADING)(getProcAddr("glMinSampleShading"))
	gpMinSampleShadingARB = (C.GPMINSAMPLESHADINGARB)(getProcAddr("glMinSampleShadingARB"))
	gpMultMatrixd = (C.GPMULTMATRIXD)(getProcAddr("glMultMatrixd"))
	gpMultMatrixf = (C.GPMULTMATRIXF)(getProcAddr("glMultMatrixf"))
	gpMultTransposeMatrixd = (C.GPMULTTRANSPOSEMATRIXD)(getProcAddr("glMultTransposeMatrixd"))
	gpMultTransposeMatrixf = (C.GPMULTTRANSPOSEMATRIXF)(getProcAddr("glMultTransposeMatrixf"))
	gpMultiDrawArrays = (C.GPMULTIDRAWARRAYS)(getProcAddr("glMultiDrawArrays"))
	gpMultiDrawArraysIndirect = (C.GPMULTIDRAWARRAYSINDIRECT)(getProcAddr("glMultiDrawArraysIndirect"))
	gpMultiDrawArraysIndirectCountARB = (C.GPMULTIDRAWARRAYSINDIRECTCOUNTARB)(getProcAddr("glMultiDrawArraysIndirectCountARB"))
	gpMultiDrawElements = (C.GPMULTIDRAWELEMENTS)(getProcAddr("glMultiDrawElements"))
	gpMultiDrawElementsBaseVertex = (C.GPMULTIDRAWELEMENTSBASEVERTEX)(getProcAddr("glMultiDrawElementsBaseVertex"))
	gpMultiDrawElementsIndirect = (C.GPMULTIDRAWELEMENTSINDIRECT)(getProcAddr("glMultiDrawElementsIndirect"))
	gpMultiDrawElementsIndirectCountARB = (C.GPMULTIDRAWELEMENTSINDIRECTCOUNTARB)(getProcAddr("glMultiDrawElementsIndirectCountARB"))
	gpMultiTexCoord1d = (C.GPMULTITEXCOORD1D)(getProcAddr("glMultiTexCoord1d"))
	gpMultiTexCoord1dv = (C.GPMULTITEXCOORD1DV)(getProcAddr("glMultiTexCoord1dv"))
	gpMultiTexCoord1f = (C.GPMULTITEXCOORD1F)(getProcAddr("glMultiTexCoord1f"))
	gpMultiTexCoord1fv = (C.GPMULTITEXCOORD1FV)(getProcAddr("glMultiTexCoord1fv"))
	gpMultiTexCoord1i = (C.GPMULTITEXCOORD1I)(getProcAddr("glMultiTexCoord1i"))
	gpMultiTexCoord1iv = (C.GPMULTITEXCOORD1IV)(getProcAddr("glMultiTexCoord1iv"))
	gpMultiTexCoord1s = (C.GPMULTITEXCOORD1S)(getProcAddr("glMultiTexCoord1s"))
	gpMultiTexCoord1sv = (C.GPMULTITEXCOORD1SV)(getProcAddr("glMultiTexCoord1sv"))
	gpMultiTexCoord2d = (C.GPMULTITEXCOORD2D)(getProcAddr("glMultiTexCoord2d"))
	gpMultiTexCoord2dv = (C.GPMULTITEXCOORD2DV)(getProcAddr("glMultiTexCoord2dv"))
	gpMultiTexCoord2f = (C.GPMULTITEXCOORD2F)(getProcAddr("glMultiTexCoord2f"))
	gpMultiTexCoord2fv = (C.GPMULTITEXCOORD2FV)(getProcAddr("glMultiTexCoord2fv"))
	gpMultiTexCoord2i = (C.GPMULTITEXCOORD2I)(getProcAddr("glMultiTexCoord2i"))
	gpMultiTexCoord2iv = (C.GPMULTITEXCOORD2IV)(getProcAddr("glMultiTexCoord2iv"))
	gpMultiTexCoord2s = (C.GPMULTITEXCOORD2S)(getProcAddr("glMultiTexCoord2s"))
	gpMultiTexCoord2sv = (C.GPMULTITEXCOORD2SV)(getProcAddr("glMultiTexCoord2sv"))
	gpMultiTexCoord3d = (C.GPMULTITEXCOORD3D)(getProcAddr("glMultiTexCoord3d"))
	gpMultiTexCoord3dv = (C.GPMULTITEXCOORD3DV)(getProcAddr("glMultiTexCoord3dv"))
	gpMultiTexCoord3f = (C.GPMULTITEXCOORD3F)(getProcAddr("glMultiTexCoord3f"))
	gpMultiTexCoord3fv = (C.GPMULTITEXCOORD3FV)(getProcAddr("glMultiTexCoord3fv"))
	gpMultiTexCoord3i = (C.GPMULTITEXCOORD3I)(getProcAddr("glMultiTexCoord3i"))
	gpMultiTexCoord3iv = (C.GPMULTITEXCOORD3IV)(getProcAddr("glMultiTexCoord3iv"))
	gpMultiTexCoord3s = (C.GPMULTITEXCOORD3S)(getProcAddr("glMultiTexCoord3s"))
	gpMultiTexCoord3sv = (C.GPMULTITEXCOORD3SV)(getProcAddr("glMultiTexCoord3sv"))
	gpMultiTexCoord4d = (C.GPMULTITEXCOORD4D)(getProcAddr("glMultiTexCoord4d"))
	gpMultiTexCoord4dv = (C.GPMULTITEXCOORD4DV)(getProcAddr("glMultiTexCoord4dv"))
	gpMultiTexCoord4f = (C.GPMULTITEXCOORD4F)(getProcAddr("glMultiTexCoord4f"))
	gpMultiTexCoord4fv = (C.GPMULTITEXCOORD4FV)(getProcAddr("glMultiTexCoord4fv"))
	gpMultiTexCoord4i = (C.GPMULTITEXCOORD4I)(getProcAddr("glMultiTexCoord4i"))
	gpMultiTexCoord4iv = (C.GPMULTITEXCOORD4IV)(getProcAddr("glMultiTexCoord4iv"))
	gpMultiTexCoord4s = (C.GPMULTITEXCOORD4S)(getProcAddr("glMultiTexCoord4s"))
	gpMultiTexCoord4sv = (C.GPMULTITEXCOORD4SV)(getProcAddr("glMultiTexCoord4sv"))
	gpNamedBufferData = (C.GPNAMEDBUFFERDATA)(getProcAddr("glNamedBufferData"))
	gpNamedBufferPageCommitmentARB = (C.GPNAMEDBUFFERPAGECOMMITMENTARB)(getProcAddr("glNamedBufferPageCommitmentARB"))
	gpNamedBufferPageCommitmentEXT = (C.GPNAMEDBUFFERPAGECOMMITMENTEXT)(getProcAddr("glNamedBufferPageCommitmentEXT"))
	gpNamedBufferStorage = (C.GPNAMEDBUFFERSTORAGE)(getProcAddr("glNamedBufferStorage"))
	gpNamedBufferSubData = (C.GPNAMEDBUFFERSUBDATA)(getProcAddr("glNamedBufferSubData"))
	gpNamedFramebufferDrawBuffer = (C.GPNAMEDFRAMEBUFFERDRAWBUFFER)(getProcAddr("glNamedFramebufferDrawBuffer"))
	gpNamedFramebufferDrawBuffers = (C.GPNAMEDFRAMEBUFFERDRAWBUFFERS)(getProcAddr("glNamedFramebufferDrawBuffers"))
	gpNamedFramebufferParameteri = (C.GPNAMEDFRAMEBUFFERPARAMETERI)(getProcAddr("glNamedFramebufferParameteri"))
	gpNamedFramebufferReadBuffer = (C.GPNAMEDFRAMEBUFFERREADBUFFER)(getProcAddr("glNamedFramebufferReadBuffer"))
	gpNamedFramebufferRenderbuffer = (C.GPNAMEDFRAMEBUFFERRENDERBUFFER)(getProcAddr("glNamedFramebufferRenderbuffer"))
	gpNamedFramebufferTexture = (C.GPNAMEDFRAMEBUFFERTEXTURE)(getProcAddr("glNamedFramebufferTexture"))
	gpNamedFramebufferTextureLayer = (C.GPNAMEDFRAMEBUFFERTEXTURELAYER)(getProcAddr("glNamedFramebufferTextureLayer"))
	gpNamedRenderbufferStorage = (C.GPNAMEDRENDERBUFFERSTORAGE)(getProcAddr("glNamedRenderbufferStorage"))
	gpNamedRenderbufferStorageMultisample = (C.GPNAMEDRENDERBUFFERSTORAGEMULTISAMPLE)(getProcAddr("glNamedRenderbufferStorageMultisample"))
	gpNamedStringARB = (C.GPNAMEDSTRINGARB)(getProcAddr("glNamedStringARB"))
	gpNewList = (C.GPNEWLIST)(getProcAddr("glNewList"))
	gpNormal3b = (C.GPNORMAL3B)(getProcAddr("glNormal3b"))
	gpNormal3bv = (C.GPNORMAL3BV)(getProcAddr("glNormal3bv"))
	gpNormal3d = (C.GPNORMAL3D)(getProcAddr("glNormal3d"))
	gpNormal3dv = (C.GPNORMAL3DV)(getProcAddr("glNormal3dv"))
	gpNormal3f = (C.GPNORMAL3F)(getProcAddr("glNormal3f"))
	gpNormal3fv = (C.GPNORMAL3FV)(getProcAddr("glNormal3fv"))
	gpNormal3i = (C.GPNORMAL3I)(getProcAddr("glNormal3i"))
	gpNormal3iv = (C.GPNORMAL3IV)(getProcAddr("glNormal3iv"))
	gpNormal3s = (C.GPNORMAL3S)(getProcAddr("glNormal3s"))
	gpNormal3sv = (C.GPNORMAL3SV)(getProcAddr("glNormal3sv"))
	gpNormalPointer = (C.GPNORMALPOINTER)(getProcAddr("glNormalPointer"))
	gpObjectLabel = (C.GPOBJECTLABEL)(getProcAddr("glObjectLabel"))
	gpObjectLabelKHR = (C.GPOBJECTLABELKHR)(getProcAddr("glObjectLabelKHR"))
	gpObjectPtrLabel = (C.GPOBJECTPTRLABEL)(getProcAddr("glObjectPtrLabel"))
	gpObjectPtrLabelKHR = (C.GPOBJECTPTRLABELKHR)(getProcAddr("glObjectPtrLabelKHR"))
	gpOrtho = (C.GPORTHO)(getProcAddr("glOrtho"))
	gpPassThrough = (C.GPPASSTHROUGH)(getProcAddr("glPassThrough"))
	gpPatchParameterfv = (C.GPPATCHPARAMETERFV)(getProcAddr("glPatchParameterfv"))
	gpPatchParameteri = (C.GPPATCHPARAMETERI)(getProcAddr("glPatchParameteri"))
	gpPauseTransformFeedback = (C.GPPAUSETRANSFORMFEEDBACK)(getProcAddr("glPauseTransformFeedback"))
	gpPixelMapfv = (C.GPPIXELMAPFV)(getProcAddr("glPixelMapfv"))
	gpPixelMapuiv = (C.GPPIXELMAPUIV)(getProcAddr("glPixelMapuiv"))
	gpPixelMapusv = (C.GPPIXELMAPUSV)(getProcAddr("glPixelMapusv"))
	gpPixelStoref = (C.GPPIXELSTOREF)(getProcAddr("glPixelStoref"))
	gpPixelStorei = (C.GPPIXELSTOREI)(getProcAddr("glPixelStorei"))
	gpPixelTransferf = (C.GPPIXELTRANSFERF)(getProcAddr("glPixelTransferf"))
	gpPixelTransferi = (C.GPPIXELTRANSFERI)(getProcAddr("glPixelTransferi"))
	gpPixelZoom = (C.GPPIXELZOOM)(getProcAddr("glPixelZoom"))
	gpPointParameterf = (C.GPPOINTPARAMETERF)(getProcAddr("glPointParameterf"))
	gpPointParameterfv = (C.GPPOINTPARAMETERFV)(getProcAddr("glPointParameterfv"))
	gpPointParameteri = (C.GPPOINTPARAMETERI)(getProcAddr("glPointParameteri"))
	gpPointParameteriv = (C.GPPOINTPARAMETERIV)(getProcAddr("glPointParameteriv"))
	gpPointSize = (C.GPPOINTSIZE)(getProcAddr("glPointSize"))
	gpPolygonMode = (C.GPPOLYGONMODE)(getProcAddr("glPolygonMode"))
	gpPolygonOffset = (C.GPPOLYGONOFFSET)(getProcAddr("glPolygonOffset"))
	gpPolygonStipple = (C.GPPOLYGONSTIPPLE)(getProcAddr("glPolygonStipple"))
	gpPopAttrib = (C.GPPOPATTRIB)(getProcAddr("glPopAttrib"))
	gpPopClientAttrib = (C.GPPOPCLIENTATTRIB)(getProcAddr("glPopClientAttrib"))
	gpPopDebugGroup = (C.GPPOPDEBUGGROUP)(getProcAddr("glPopDebugGroup"))
	gpPopDebugGroupKHR = (C.GPPOPDEBUGGROUPKHR)(getProcAddr("glPopDebugGroupKHR"))
	gpPopMatrix = (C.GPPOPMATRIX)(getProcAddr("glPopMatrix"))
	gpPopName = (C.GPPOPNAME)(getProcAddr("glPopName"))
	gpPrimitiveRestartIndex = (C.GPPRIMITIVERESTARTINDEX)(getProcAddr("glPrimitiveRestartIndex"))
	gpPrioritizeTextures = (C.GPPRIORITIZETEXTURES)(getProcAddr("glPrioritizeTextures"))
	gpProgramBinary = (C.GPPROGRAMBINARY)(getProcAddr("glProgramBinary"))
	gpProgramParameteri = (C.GPPROGRAMPARAMETERI)(getProcAddr("glProgramParameteri"))
	gpProgramUniform1d = (C.GPPROGRAMUNIFORM1D)(getProcAddr("glProgramUniform1d"))
	gpProgramUniform1dv = (C.GPPROGRAMUNIFORM1DV)(getProcAddr("glProgramUniform1dv"))
	gpProgramUniform1f = (C.GPPROGRAMUNIFORM1F)(getProcAddr("glProgramUniform1f"))
	gpProgramUniform1fv = (C.GPPROGRAMUNIFORM1FV)(getProcAddr("glProgramUniform1fv"))
	gpProgramUniform1i = (C.GPPROGRAMUNIFORM1I)(getProcAddr("glProgramUniform1i"))
	gpProgramUniform1iv = (C.GPPROGRAMUNIFORM1IV)(getProcAddr("glProgramUniform1iv"))
	gpProgramUniform1ui = (C.GPPROGRAMUNIFORM1UI)(getProcAddr("glProgramUniform1ui"))
	gpProgramUniform1uiv = (C.GPPROGRAMUNIFORM1UIV)(getProcAddr("glProgramUniform1uiv"))
	gpProgramUniform2d = (C.GPPROGRAMUNIFORM2D)(getProcAddr("glProgramUniform2d"))
	gpProgramUniform2dv = (C.GPPROGRAMUNIFORM2DV)(getProcAddr("glProgramUniform2dv"))
	gpProgramUniform2f = (C.GPPROGRAMUNIFORM2F)(getProcAddr("glProgramUniform2f"))
	gpProgramUniform2fv = (C.GPPROGRAMUNIFORM2FV)(getProcAddr("glProgramUniform2fv"))
	gpProgramUniform2i = (C.GPPROGRAMUNIFORM2I)(getProcAddr("glProgramUniform2i"))
	gpProgramUniform2iv = (C.GPPROGRAMUNIFORM2IV)(getProcAddr("glProgramUniform2iv"))
	gpProgramUniform2ui = (C.GPPROGRAMUNIFORM2UI)(getProcAddr("glProgramUniform2ui"))
	gpProgramUniform2uiv = (C.GPPROGRAMUNIFORM2UIV)(getProcAddr("glProgramUniform2uiv"))
	gpProgramUniform3d = (C.GPPROGRAMUNIFORM3D)(getProcAddr("glProgramUniform3d"))
	gpProgramUniform3dv = (C.GPPROGRAMUNIFORM3DV)(getProcAddr("glProgramUniform3dv"))
	gpProgramUniform3f = (C.GPPROGRAMUNIFORM3F)(getProcAddr("glProgramUniform3f"))
	gpProgramUniform3fv = (C.GPPROGRAMUNIFORM3FV)(getProcAddr("glProgramUniform3fv"))
	gpProgramUniform3i = (C.GPPROGRAMUNIFORM3I)(getProcAddr("glProgramUniform3i"))
	gpProgramUniform3iv = (C.GPPROGRAMUNIFORM3IV)(getProcAddr("glProgramUniform3iv"))
	gpProgramUniform3ui = (C.GPPROGRAMUNIFORM3UI)(getProcAddr("glProgramUniform3ui"))
	gpProgramUniform3uiv = (C.GPPROGRAMUNIFORM3UIV)(getProcAddr("glProgramUniform3uiv"))
	gpProgramUniform4d = (C.GPPROGRAMUNIFORM4D)(getProcAddr("glProgramUniform4d"))
	gpProgramUniform4dv = (C.GPPROGRAMUNIFORM4DV)(getProcAddr("glProgramUniform4dv"))
	gpProgramUniform4f = (C.GPPROGRAMUNIFORM4F)(getProcAddr("glProgramUniform4f"))
	gpProgramUniform4fv = (C.GPPROGRAMUNIFORM4FV)(getProcAddr("glProgramUniform4fv"))
	gpProgramUniform4i = (C.GPPROGRAMUNIFORM4I)(getProcAddr("glProgramUniform4i"))
	gpProgramUniform4iv = (C.GPPROGRAMUNIFORM4IV)(getProcAddr("glProgramUniform4iv"))
	gpProgramUniform4ui = (C.GPPROGRAMUNIFORM4UI)(getProcAddr("glProgramUniform4ui"))
	gpProgramUniform4uiv = (C.GPPROGRAMUNIFORM4UIV)(getProcAddr("glProgramUniform4uiv"))
	gpProgramUniformHandleui64ARB = (C.GPPROGRAMUNIFORMHANDLEUI64ARB)(getProcAddr("glProgramUniformHandleui64ARB"))
	gpProgramUniformHandleui64vARB = (C.GPPROGRAMUNIFORMHANDLEUI64VARB)(getProcAddr("glProgramUniformHandleui64vARB"))
	gpProgramUniformMatrix2dv = (C.GPPROGRAMUNIFORMMATRIX2DV)(getProcAddr("glProgramUniformMatrix2dv"))
	gpProgramUniformMatrix2fv = (C.GPPROGRAMUNIFORMMATRIX2FV)(getProcAddr("glProgramUniformMatrix2fv"))
	gpProgramUniformMatrix2x3dv = (C.GPPROGRAMUNIFORMMATRIX2X3DV)(getProcAddr("glProgramUniformMatrix2x3dv"))
	gpProgramUniformMatrix2x3fv = (C.GPPROGRAMUNIFORMMATRIX2X3FV)(getProcAddr("glProgramUniformMatrix2x3fv"))
	gpProgramUniformMatrix2x4dv = (C.GPPROGRAMUNIFORMMATRIX2X4DV)(getProcAddr("glProgramUniformMatrix2x4dv"))
	gpProgramUniformMatrix2x4fv = (C.GPPROGRAMUNIFORMMATRIX2X4FV)(getProcAddr("glProgramUniformMatrix2x4fv"))
	gpProgramUniformMatrix3dv = (C.GPPROGRAMUNIFORMMATRIX3DV)(getProcAddr("glProgramUniformMatrix3dv"))
	gpProgramUniformMatrix3fv = (C.GPPROGRAMUNIFORMMATRIX3FV)(getProcAddr("glProgramUniformMatrix3fv"))
	gpProgramUniformMatrix3x2dv = (C.GPPROGRAMUNIFORMMATRIX3X2DV)(getProcAddr("glProgramUniformMatrix3x2dv"))
	gpProgramUniformMatrix3x2fv = (C.GPPROGRAMUNIFORMMATRIX3X2FV)(getProcAddr("glProgramUniformMatrix3x2fv"))
	gpProgramUniformMatrix3x4dv = (C.GPPROGRAMUNIFORMMATRIX3X4DV)(getProcAddr("glProgramUniformMatrix3x4dv"))
	gpProgramUniformMatrix3x4fv = (C.GPPROGRAMUNIFORMMATRIX3X4FV)(getProcAddr("glProgramUniformMatrix3x4fv"))
	gpProgramUniformMatrix4dv = (C.GPPROGRAMUNIFORMMATRIX4DV)(getProcAddr("glProgramUniformMatrix4dv"))
	gpProgramUniformMatrix4fv = (C.GPPROGRAMUNIFORMMATRIX4FV)(getProcAddr("glProgramUniformMatrix4fv"))
	gpProgramUniformMatrix4x2dv = (C.GPPROGRAMUNIFORMMATRIX4X2DV)(getProcAddr("glProgramUniformMatrix4x2dv"))
	gpProgramUniformMatrix4x2fv = (C.GPPROGRAMUNIFORMMATRIX4X2FV)(getProcAddr("glProgramUniformMatrix4x2fv"))
	gpProgramUniformMatrix4x3dv = (C.GPPROGRAMUNIFORMMATRIX4X3DV)(getProcAddr("glProgramUniformMatrix4x3dv"))
	gpProgramUniformMatrix4x3fv = (C.GPPROGRAMUNIFORMMATRIX4X3FV)(getProcAddr("glProgramUniformMatrix4x3fv"))
	gpProvokingVertex = (C.GPPROVOKINGVERTEX)(getProcAddr("glProvokingVertex"))
	gpPushAttrib = (C.GPPUSHATTRIB)(getProcAddr("glPushAttrib"))
	gpPushClientAttrib = (C.GPPUSHCLIENTATTRIB)(getProcAddr("glPushClientAttrib"))
	gpPushDebugGroup = (C.GPPUSHDEBUGGROUP)(getProcAddr("glPushDebugGroup"))
	gpPushDebugGroupKHR = (C.GPPUSHDEBUGGROUPKHR)(getProcAddr("glPushDebugGroupKHR"))
	gpPushMatrix = (C.GPPUSHMATRIX)(getProcAddr("glPushMatrix"))
	gpPushName = (C.GPPUSHNAME)(getProcAddr("glPushName"))
	gpQueryCounter = (C.GPQUERYCOUNTER)(getProcAddr("glQueryCounter"))
	gpRasterPos2d = (C.GPRASTERPOS2D)(getProcAddr("glRasterPos2d"))
	gpRasterPos2dv = (C.GPRASTERPOS2DV)(getProcAddr("glRasterPos2dv"))
	gpRasterPos2f = (C.GPRASTERPOS2F)(getProcAddr("glRasterPos2f"))
	gpRasterPos2fv = (C.GPRASTERPOS2FV)(getProcAddr("glRasterPos2fv"))
	gpRasterPos2i = (C.GPRASTERPOS2I)(getProcAddr("glRasterPos2i"))
	gpRasterPos2iv = (C.GPRASTERPOS2IV)(getProcAddr("glRasterPos2iv"))
	gpRasterPos2s = (C.GPRASTERPOS2S)(getProcAddr("glRasterPos2s"))
	gpRasterPos2sv = (C.GPRASTERPOS2SV)(getProcAddr("glRasterPos2sv"))
	gpRasterPos3d = (C.GPRASTERPOS3D)(getProcAddr("glRasterPos3d"))
	gpRasterPos3dv = (C.GPRASTERPOS3DV)(getProcAddr("glRasterPos3dv"))
	gpRasterPos3f = (C.GPRASTERPOS3F)(getProcAddr("glRasterPos3f"))
	gpRasterPos3fv = (C.GPRASTERPOS3FV)(getProcAddr("glRasterPos3fv"))
	gpRasterPos3i = (C.GPRASTERPOS3I)(getProcAddr("glRasterPos3i"))
	gpRasterPos3iv = (C.GPRASTERPOS3IV)(getProcAddr("glRasterPos3iv"))
	gpRasterPos3s = (C.GPRASTERPOS3S)(getProcAddr("glRasterPos3s"))
	gpRasterPos3sv = (C.GPRASTERPOS3SV)(getProcAddr("glRasterPos3sv"))
	gpRasterPos4d = (C.GPRASTERPOS4D)(getProcAddr("glRasterPos4d"))
	gpRasterPos4dv = (C.GPRASTERPOS4DV)(getProcAddr("glRasterPos4dv"))
	gpRasterPos4f = (C.GPRASTERPOS4F)(getProcAddr("glRasterPos4f"))
	gpRasterPos4fv = (C.GPRASTERPOS4FV)(getProcAddr("glRasterPos4fv"))
	gpRasterPos4i = (C.GPRASTERPOS4I)(getProcAddr("glRasterPos4i"))
	gpRasterPos4iv = (C.GPRASTERPOS4IV)(getProcAddr("glRasterPos4iv"))
	gpRasterPos4s = (C.GPRASTERPOS4S)(getProcAddr("glRasterPos4s"))
	gpRasterPos4sv = (C.GPRASTERPOS4SV)(getProcAddr("glRasterPos4sv"))
	gpReadBuffer = (C.GPREADBUFFER)(getProcAddr("glReadBuffer"))
	gpReadPixels = (C.GPREADPIXELS)(getProcAddr("glReadPixels"))
	gpReadnPixels = (C.GPREADNPIXELS)(getProcAddr("glReadnPixels"))
	gpReadnPixelsARB = (C.GPREADNPIXELSARB)(getProcAddr("glReadnPixelsARB"))
	gpReadnPixelsKHR = (C.GPREADNPIXELSKHR)(getProcAddr("glReadnPixelsKHR"))
	gpRectd = (C.GPRECTD)(getProcAddr("glRectd"))
	gpRectdv = (C.GPRECTDV)(getProcAddr("glRectdv"))
	gpRectf = (C.GPRECTF)(getProcAddr("glRectf"))
	gpRectfv = (C.GPRECTFV)(getProcAddr("glRectfv"))
	gpRecti = (C.GPRECTI)(getProcAddr("glRecti"))
	gpRectiv = (C.GPRECTIV)(getProcAddr("glRectiv"))
	gpRects = (C.GPRECTS)(getProcAddr("glRects"))
	gpRectsv = (C.GPRECTSV)(getProcAddr("glRectsv"))
	gpReleaseShaderCompiler = (C.GPRELEASESHADERCOMPILER)(getProcAddr("glReleaseShaderCompiler"))
	gpRenderMode = (C.GPRENDERMODE)(getProcAddr("glRenderMode"))
	gpRenderbufferStorage = (C.GPRENDERBUFFERSTORAGE)(getProcAddr("glRenderbufferStorage"))
	gpRenderbufferStorageMultisample = (C.GPRENDERBUFFERSTORAGEMULTISAMPLE)(getProcAddr("glRenderbufferStorageMultisample"))
	gpResumeTransformFeedback = (C.GPRESUMETRANSFORMFEEDBACK)(getProcAddr("glResumeTransformFeedback"))
	gpRotated = (C.GPROTATED)(getProcAddr("glRotated"))
	gpRotatef = (C.GPROTATEF)(getProcAddr("glRotatef"))
	gpSampleCoverage = (C.GPSAMPLECOVERAGE)(getProcAddr("glSampleCoverage"))
	gpSampleMaski = (C.GPSAMPLEMASKI)(getProcAddr("glSampleMaski"))
	gpSamplerParameterIiv = (C.GPSAMPLERPARAMETERIIV)(getProcAddr("glSamplerParameterIiv"))
	gpSamplerParameterIuiv = (C.GPSAMPLERPARAMETERIUIV)(getProcAddr("glSamplerParameterIuiv"))
	gpSamplerParameterf = (C.GPSAMPLERPARAMETERF)(getProcAddr("glSamplerParameterf"))
	gpSamplerParameterfv = (C.GPSAMPLERPARAMETERFV)(getProcAddr("glSamplerParameterfv"))
	gpSamplerParameteri = (C.GPSAMPLERPARAMETERI)(getProcAddr("glSamplerParameteri"))
	gpSamplerParameteriv = (C.GPSAMPLERPARAMETERIV)(getProcAddr("glSamplerParameteriv"))
	gpScaled = (C.GPSCALED)(getProcAddr("glScaled"))
	gpScalef = (C.GPSCALEF)(getProcAddr("glScalef"))
	gpScissor = (C.GPSCISSOR)(getProcAddr("glScissor"))
	gpScissorArrayv = (C.GPSCISSORARRAYV)(getProcAddr("glScissorArrayv"))
	gpScissorIndexed = (C.GPSCISSORINDEXED)(getProcAddr("glScissorIndexed"))
	gpScissorIndexedv = (C.GPSCISSORINDEXEDV)(getProcAddr("glScissorIndexedv"))
	gpSecondaryColor3b = (C.GPSECONDARYCOLOR3B)(getProcAddr("glSecondaryColor3b"))
	gpSecondaryColor3bv = (C.GPSECONDARYCOLOR3BV)(getProcAddr("glSecondaryColor3bv"))
	gpSecondaryColor3d = (C.GPSECONDARYCOLOR3D)(getProcAddr("glSecondaryColor3d"))
	gpSecondaryColor3dv = (C.GPSECONDARYCOLOR3DV)(getProcAddr("glSecondaryColor3dv"))
	gpSecondaryColor3f = (C.GPSECONDARYCOLOR3F)(getProcAddr("glSecondaryColor3f"))
	gpSecondaryColor3fv = (C.GPSECONDARYCOLOR3FV)(getProcAddr("glSecondaryColor3fv"))
	gpSecondaryColor3i = (C.GPSECONDARYCOLOR3I)(getProcAddr("glSecondaryColor3i"))
	gpSecondaryColor3iv = (C.GPSECONDARYCOLOR3IV)(getProcAddr("glSecondaryColor3iv"))
	gpSecondaryColor3s = (C.GPSECONDARYCOLOR3S)(getProcAddr("glSecondaryColor3s"))
	gpSecondaryColor3sv = (C.GPSECONDARYCOLOR3SV)(getProcAddr("glSecondaryColor3sv"))
	gpSecondaryColor3ub = (C.GPSECONDARYCOLOR3UB)(getProcAddr("glSecondaryColor3ub"))
	gpSecondaryColor3ubv = (C.GPSECONDARYCOLOR3UBV)(getProcAddr("glSecondaryColor3ubv"))
	gpSecondaryColor3ui = (C.GPSECONDARYCOLOR3UI)(getProcAddr("glSecondaryColor3ui"))
	gpSecondaryColor3uiv = (C.GPSECONDARYCOLOR3UIV)(getProcAddr("glSecondaryColor3uiv"))
	gpSecondaryColor3us = (C.GPSECONDARYCOLOR3US)(getProcAddr("glSecondaryColor3us"))
	gpSecondaryColor3usv = (C.GPSECONDARYCOLOR3USV)(getProcAddr("glSecondaryColor3usv"))
	gpSecondaryColorPointer = (C.GPSECONDARYCOLORPOINTER)(getProcAddr("glSecondaryColorPointer"))
	gpSelectBuffer = (C.GPSELECTBUFFER)(getProcAddr("glSelectBuffer"))
	gpShadeModel = (C.GPSHADEMODEL)(getProcAddr("glShadeModel"))
	gpShaderBinary = (C.GPSHADERBINARY)(getProcAddr("glShaderBinary"))
	gpShaderSource = (C.GPSHADERSOURCE)(getProcAddr("glShaderSource"))
	gpShaderStorageBlockBinding = (C.GPSHADERSTORAGEBLOCKBINDING)(getProcAddr("glShaderStorageBlockBinding"))
	gpStencilFunc = (C.GPSTENCILFUNC)(getProcAddr("glStencilFunc"))
	gpStencilFuncSeparate = (C.GPSTENCILFUNCSEPARATE)(getProcAddr("glStencilFuncSeparate"))
	gpStencilMask = (C.GPSTENCILMASK)(getProcAddr("glStencilMask"))
	gpStencilMaskSeparate = (C.GPSTENCILMASKSEPARATE)(getProcAddr("glStencilMaskSeparate"))
	gpStencilOp = (C.GPSTENCILOP)(getProcAddr("glStencilOp"))
	gpStencilOpSeparate = (C.GPSTENCILOPSEPARATE)(getProcAddr("glStencilOpSeparate"))
	gpTexBuffer = (C.GPTEXBUFFER)(getProcAddr("glTexBuffer"))
	gpTexBufferRange = (C.GPTEXBUFFERRANGE)(getProcAddr("glTexBufferRange"))
	gpTexCoord1d = (C.GPTEXCOORD1D)(getProcAddr("glTexCoord1d"))
	gpTexCoord1dv = (C.GPTEXCOORD1DV)(getProcAddr("glTexCoord1dv"))
	gpTexCoord1f = (C.GPTEXCOORD1F)(getProcAddr("glTexCoord1f"))
	gpTexCoord1fv = (C.GPTEXCOORD1FV)(getProcAddr("glTexCoord1fv"))
	gpTexCoord1i = (C.GPTEXCOORD1I)(getProcAddr("glTexCoord1i"))
	gpTexCoord1iv = (C.GPTEXCOORD1IV)(getProcAddr("glTexCoord1iv"))
	gpTexCoord1s = (C.GPTEXCOORD1S)(getProcAddr("glTexCoord1s"))
	gpTexCoord1sv = (C.GPTEXCOORD1SV)(getProcAddr("glTexCoord1sv"))
	gpTexCoord2d = (C.GPTEXCOORD2D)(getProcAddr("glTexCoord2d"))
	gpTexCoord2dv = (C.GPTEXCOORD2DV)(getProcAddr("glTexCoord2dv"))
	gpTexCoord2f = (C.GPTEXCOORD2F)(getProcAddr("glTexCoord2f"))
	gpTexCoord2fv = (C.GPTEXCOORD2FV)(getProcAddr("glTexCoord2fv"))
	gpTexCoord2i = (C.GPTEXCOORD2I)(getProcAddr("glTexCoord2i"))
	gpTexCoord2iv = (C.GPTEXCOORD2IV)(getProcAddr("glTexCoord2iv"))
	gpTexCoord2s = (C.GPTEXCOORD2S)(getProcAddr("glTexCoord2s"))
	gpTexCoord2sv = (C.GPTEXCOORD2SV)(getProcAddr("glTexCoord2sv"))
	gpTexCoord3d = (C.GPTEXCOORD3D)(getProcAddr("glTexCoord3d"))
	gpTexCoord3dv = (C.GPTEXCOORD3DV)(getProcAddr("glTexCoord3dv"))
	gpTexCoord3f = (C.GPTEXCOORD3F)(getProcAddr("glTexCoord3f"))
	gpTexCoord3fv = (C.GPTEXCOORD3FV)(getProcAddr("glTexCoord3fv"))
	gpTexCoord3i = (C.GPTEXCOORD3I)(getProcAddr("glTexCoord3i"))
	gpTexCoord3iv = (C.GPTEXCOORD3IV)(getProcAddr("glTexCoord3iv"))
	gpTexCoord3s = (C.GPTEXCOORD3S)(getProcAddr("glTexCoord3s"))
	gpTexCoord3sv = (C.GPTEXCOORD3SV)(getProcAddr("glTexCoord3sv"))
	gpTexCoord4d = (C.GPTEXCOORD4D)(getProcAddr("glTexCoord4d"))
	gpTexCoord4dv = (C.GPTEXCOORD4DV)(getProcAddr("glTexCoord4dv"))
	gpTexCoord4f = (C.GPTEXCOORD4F)(getProcAddr("glTexCoord4f"))
	gpTexCoord4fv = (C.GPTEXCOORD4FV)(getProcAddr("glTexCoord4fv"))
	gpTexCoord4i = (C.GPTEXCOORD4I)(getProcAddr("glTexCoord4i"))
	gpTexCoord4iv = (C.GPTEXCOORD4IV)(getProcAddr("glTexCoord4iv"))
	gpTexCoord4s = (C.GPTEXCOORD4S)(getProcAddr("glTexCoord4s"))
	gpTexCoord4sv = (C.GPTEXCOORD4SV)(getProcAddr("glTexCoord4sv"))
	gpTexCoordPointer = (C.GPTEXCOORDPOINTER)(getProcAddr("glTexCoordPointer"))
	gpTexEnvf = (C.GPTEXENVF)(getProcAddr("glTexEnvf"))
	gpTexEnvfv = (C.GPTEXENVFV)(getProcAddr("glTexEnvfv"))
	gpTexEnvi = (C.GPTEXENVI)(getProcAddr("glTexEnvi"))
	gpTexEnviv = (C.GPTEXENVIV)(getProcAddr("glTexEnviv"))
	gpTexGend = (C.GPTEXGEND)(getProcAddr("glTexGend"))
	gpTexGendv = (C.GPTEXGENDV)(getProcAddr("glTexGendv"))
	gpTexGenf = (C.GPTEXGENF)(getProcAddr("glTexGenf"))
	gpTexGenfv = (C.GPTEXGENFV)(getProcAddr("glTexGenfv"))
	gpTexGeni = (C.GPTEXGENI)(getProcAddr("glTexGeni"))
	gpTexGeniv = (C.GPTEXGENIV)(getProcAddr("glTexGeniv"))
	gpTexImage1D = (C.GPTEXIMAGE1D)(getProcAddr("glTexImage1D"))
	gpTexImage2D = (C.GPTEXIMAGE2D)(getProcAddr("glTexImage2D"))
	gpTexImage2DMultisample = (C.GPTEXIMAGE2DMULTISAMPLE)(getProcAddr("glTexImage2DMultisample"))
	gpTexImage3D = (C.GPTEXIMAGE3D)(getProcAddr("glTexImage3D"))
	gpTexImage3DMultisample = (C.GPTEXIMAGE3DMULTISAMPLE)(getProcAddr("glTexImage3DMultisample"))
	gpTexPageCommitmentARB = (C.GPTEXPAGECOMMITMENTARB)(getProcAddr("glTexPageCommitmentARB"))
	gpTexParameterIiv = (C.GPTEXPARAMETERIIV)(getProcAddr("glTexParameterIiv"))
	gpTexParameterIuiv = (C.GPTEXPARAMETERIUIV)(getProcAddr("glTexParameterIuiv"))
	gpTexParameterf = (C.GPTEXPARAMETERF)(getProcAddr("glTexParameterf"))
	gpTexParameterfv = (C.GPTEXPARAMETERFV)(getProcAddr("glTexParameterfv"))
	gpTexParameteri = (C.GPTEXPARAMETERI)(getProcAddr("glTexParameteri"))
	gpTexParameteriv = (C.GPTEXPARAMETERIV)(getProcAddr("glTexParameteriv"))
	gpTexStorage1D = (C.GPTEXSTORAGE1D)(getProcAddr("glTexStorage1D"))
	gpTexStorage2D = (C.GPTEXSTORAGE2D)(getProcAddr("glTexStorage2D"))
	gpTexStorage2DMultisample = (C.GPTEXSTORAGE2DMULTISAMPLE)(getProcAddr("glTexStorage2DMultisample"))
	gpTexStorage3D = (C.GPTEXSTORAGE3D)(getProcAddr("glTexStorage3D"))
	gpTexStorage3DMultisample = (C.GPTEXSTORAGE3DMULTISAMPLE)(getProcAddr("glTexStorage3DMultisample"))
	gpTexSubImage1D = (C.GPTEXSUBIMAGE1D)(getProcAddr("glTexSubImage1D"))
	gpTexSubImage2D = (C.GPTEXSUBIMAGE2D)(getProcAddr("glTexSubImage2D"))
	gpTexSubImage3D = (C.GPTEXSUBIMAGE3D)(getProcAddr("glTexSubImage3D"))
	gpTextureBarrier = (C.GPTEXTUREBARRIER)(getProcAddr("glTextureBarrier"))
	gpTextureBuffer = (C.GPTEXTUREBUFFER)(getProcAddr("glTextureBuffer"))
	gpTextureBufferRange = (C.GPTEXTUREBUFFERRANGE)(getProcAddr("glTextureBufferRange"))
	gpTextureParameterIiv = (C.GPTEXTUREPARAMETERIIV)(getProcAddr("glTextureParameterIiv"))
	gpTextureParameterIuiv = (C.GPTEXTUREPARAMETERIUIV)(getProcAddr("glTextureParameterIuiv"))
	gpTextureParameterf = (C.GPTEXTUREPARAMETERF)(getProcAddr("glTextureParameterf"))
	gpTextureParameterfv = (C.GPTEXTUREPARAMETERFV)(getProcAddr("glTextureParameterfv"))
	gpTextureParameteri = (C.GPTEXTUREPARAMETERI)(getProcAddr("glTextureParameteri"))
	gpTextureParameteriv = (C.GPTEXTUREPARAMETERIV)(getProcAddr("glTextureParameteriv"))
	gpTextureStorage1D = (C.GPTEXTURESTORAGE1D)(getProcAddr("glTextureStorage1D"))
	gpTextureStorage2D = (C.GPTEXTURESTORAGE2D)(getProcAddr("glTextureStorage2D"))
	gpTextureStorage2DMultisample = (C.GPTEXTURESTORAGE2DMULTISAMPLE)(getProcAddr("glTextureStorage2DMultisample"))
	gpTextureStorage3D = (C.GPTEXTURESTORAGE3D)(getProcAddr("glTextureStorage3D"))
	gpTextureStorage3DMultisample = (C.GPTEXTURESTORAGE3DMULTISAMPLE)(getProcAddr("glTextureStorage3DMultisample"))
	gpTextureSubImage1D = (C.GPTEXTURESUBIMAGE1D)(getProcAddr("glTextureSubImage1D"))
	gpTextureSubImage2D = (C.GPTEXTURESUBIMAGE2D)(getProcAddr("glTextureSubImage2D"))
	gpTextureSubImage3D = (C.GPTEXTURESUBIMAGE3D)(getProcAddr("glTextureSubImage3D"))
	gpTextureView = (C.GPTEXTUREVIEW)(getProcAddr("glTextureView"))
	gpTransformFeedbackBufferBase = (C.GPTRANSFORMFEEDBACKBUFFERBASE)(getProcAddr("glTransformFeedbackBufferBase"))
	gpTransformFeedbackBufferRange = (C.GPTRANSFORMFEEDBACKBUFFERRANGE)(getProcAddr("glTransformFeedbackBufferRange"))
	gpTransformFeedbackVaryings = (C.GPTRANSFORMFEEDBACKVARYINGS)(getProcAddr("glTransformFeedbackVaryings"))
	gpTranslated = (C.GPTRANSLATED)(getProcAddr("glTranslated"))
	gpTranslatef = (C.GPTRANSLATEF)(getProcAddr("glTranslatef"))
	gpUniform1d = (C.GPUNIFORM1D)(getProcAddr("glUniform1d"))
	gpUniform1dv = (C.GPUNIFORM1DV)(getProcAddr("glUniform1dv"))
	gpUniform1f = (C.GPUNIFORM1F)(getProcAddr("glUniform1f"))
	gpUniform1fv = (C.GPUNIFORM1FV)(getProcAddr("glUniform1fv"))
	gpUniform1i = (C.GPUNIFORM1I)(getProcAddr("glUniform1i"))
	gpUniform1iv = (C.GPUNIFORM1IV)(getProcAddr("glUniform1iv"))
	gpUniform1ui = (C.GPUNIFORM1UI)(getProcAddr("glUniform1ui"))
	gpUniform1uiv = (C.GPUNIFORM1UIV)(getProcAddr("glUniform1uiv"))
	gpUniform2d = (C.GPUNIFORM2D)(getProcAddr("glUniform2d"))
	gpUniform2dv = (C.GPUNIFORM2DV)(getProcAddr("glUniform2dv"))
	gpUniform2f = (C.GPUNIFORM2F)(getProcAddr("glUniform2f"))
	gpUniform2fv = (C.GPUNIFORM2FV)(getProcAddr("glUniform2fv"))
	gpUniform2i = (C.GPUNIFORM2I)(getProcAddr("glUniform2i"))
	gpUniform2iv = (C.GPUNIFORM2IV)(getProcAddr("glUniform2iv"))
	gpUniform2ui = (C.GPUNIFORM2UI)(getProcAddr("glUniform2ui"))
	gpUniform2uiv = (C.GPUNIFORM2UIV)(getProcAddr("glUniform2uiv"))
	gpUniform3d = (C.GPUNIFORM3D)(getProcAddr("glUniform3d"))
	gpUniform3dv = (C.GPUNIFORM3DV)(getProcAddr("glUniform3dv"))
	gpUniform3f = (C.GPUNIFORM3F)(getProcAddr("glUniform3f"))
	gpUniform3fv = (C.GPUNIFORM3FV)(getProcAddr("glUniform3fv"))
	gpUniform3i = (C.GPUNIFORM3I)(getProcAddr("glUniform3i"))
	gpUniform3iv = (C.GPUNIFORM3IV)(getProcAddr("glUniform3iv"))
	gpUniform3ui = (C.GPUNIFORM3UI)(getProcAddr("glUniform3ui"))
	gpUniform3uiv = (C.GPUNIFORM3UIV)(getProcAddr("glUniform3uiv"))
	gpUniform4d = (C.GPUNIFORM4D)(getProcAddr("glUniform4d"))
	gpUniform4dv = (C.GPUNIFORM4DV)(getProcAddr("glUniform4dv"))
	gpUniform4f = (C.GPUNIFORM4F)(getProcAddr("glUniform4f"))
	gpUniform4fv = (C.GPUNIFORM4FV)(getProcAddr("glUniform4fv"))
	gpUniform4i = (C.GPUNIFORM4I)(getProcAddr("glUniform4i"))
	gpUniform4iv = (C.GPUNIFORM4IV)(getProcAddr("glUniform4iv"))
	gpUniform4ui = (C.GPUNIFORM4UI)(getProcAddr("glUniform4ui"))
	gpUniform4uiv = (C.GPUNIFORM4UIV)(getProcAddr("glUniform4uiv"))
	gpUniformBlockBinding = (C.GPUNIFORMBLOCKBINDING)(getProcAddr("glUniformBlockBinding"))
	gpUniformHandleui64ARB = (C.GPUNIFORMHANDLEUI64ARB)(getProcAddr("glUniformHandleui64ARB"))
	gpUniformHandleui64vARB = (C.GPUNIFORMHANDLEUI64VARB)(getProcAddr("glUniformHandleui64vARB"))
	gpUniformMatrix2dv = (C.GPUNIFORMMATRIX2DV)(getProcAddr("glUniformMatrix2dv"))
	gpUniformMatrix2fv = (C.GPUNIFORMMATRIX2FV)(getProcAddr("glUniformMatrix2fv"))
	gpUniformMatrix2x3dv = (C.GPUNIFORMMATRIX2X3DV)(getProcAddr("glUniformMatrix2x3dv"))
	gpUniformMatrix2x3fv = (C.GPUNIFORMMATRIX2X3FV)(getProcAddr("glUniformMatrix2x3fv"))
	gpUniformMatrix2x4dv = (C.GPUNIFORMMATRIX2X4DV)(getProcAddr("glUniformMatrix2x4dv"))
	gpUniformMatrix2x4fv = (C.GPUNIFORMMATRIX2X4FV)(getProcAddr("glUniformMatrix2x4fv"))
	gpUniformMatrix3dv = (C.GPUNIFORMMATRIX3DV)(getProcAddr("glUniformMatrix3dv"))
	gpUniformMatrix3fv = (C.GPUNIFORMMATRIX3FV)(getProcAddr("glUniformMatrix3fv"))
	gpUniformMatrix3x2dv = (C.GPUNIFORMMATRIX3X2DV)(getProcAddr("glUniformMatrix3x2dv"))
	gpUniformMatrix3x2fv = (C.GPUNIFORMMATRIX3X2FV)(getProcAddr("glUniformMatrix3x2fv"))
	gpUniformMatrix3x4dv = (C.GPUNIFORMMATRIX3X4DV)(getProcAddr("glUniformMatrix3x4dv"))
	gpUniformMatrix3x4fv = (C.GPUNIFORMMATRIX3X4FV)(getProcAddr("glUniformMatrix3x4fv"))
	gpUniformMatrix4dv = (C.GPUNIFORMMATRIX4DV)(getProcAddr("glUniformMatrix4dv"))
	gpUniformMatrix4fv = (C.GPUNIFORMMATRIX4FV)(getProcAddr("glUniformMatrix4fv"))
	gpUniformMatrix4x2dv = (C.GPUNIFORMMATRIX4X2DV)(getProcAddr("glUniformMatrix4x2dv"))
	gpUniformMatrix4x2fv = (C.GPUNIFORMMATRIX4X2FV)(getProcAddr("glUniformMatrix4x2fv"))
	gpUniformMatrix4x3dv = (C.GPUNIFORMMATRIX4X3DV)(getProcAddr("glUniformMatrix4x3dv"))
	gpUniformMatrix4x3fv = (C.GPUNIFORMMATRIX4X3FV)(getProcAddr("glUniformMatrix4x3fv"))
	gpUniformSubroutinesuiv = (C.GPUNIFORMSUBROUTINESUIV)(getProcAddr("glUniformSubroutinesuiv"))
	gpUnmapBuffer = (C.GPUNMAPBUFFER)(getProcAddr("glUnmapBuffer"))
	gpUnmapNamedBuffer = (C.GPUNMAPNAMEDBUFFER)(getProcAddr("glUnmapNamedBuffer"))
	gpUseProgram = (C.GPUSEPROGRAM)(getProcAddr("glUseProgram"))
	gpUseProgramStages = (C.GPUSEPROGRAMSTAGES)(getProcAddr("glUseProgramStages"))
	gpValidateProgram = (C.GPVALIDATEPROGRAM)(getProcAddr("glValidateProgram"))
	gpValidateProgramPipeline = (C.GPVALIDATEPROGRAMPIPELINE)(getProcAddr("glValidateProgramPipeline"))
	gpVertex2d = (C.GPVERTEX2D)(getProcAddr("glVertex2d"))
	gpVertex2dv = (C.GPVERTEX2DV)(getProcAddr("glVertex2dv"))
	gpVertex2f = (C.GPVERTEX2F)(getProcAddr("glVertex2f"))
	gpVertex2fv = (C.GPVERTEX2FV)(getProcAddr("glVertex2fv"))
	gpVertex2i = (C.GPVERTEX2I)(getProcAddr("glVertex2i"))
	gpVertex2iv = (C.GPVERTEX2IV)(getProcAddr("glVertex2iv"))
	gpVertex2s = (C.GPVERTEX2S)(getProcAddr("glVertex2s"))
	gpVertex2sv = (C.GPVERTEX2SV)(getProcAddr("glVertex2sv"))
	gpVertex3d = (C.GPVERTEX3D)(getProcAddr("glVertex3d"))
	gpVertex3dv = (C.GPVERTEX3DV)(getProcAddr("glVertex3dv"))
	gpVertex3f = (C.GPVERTEX3F)(getProcAddr("glVertex3f"))
	gpVertex3fv = (C.GPVERTEX3FV)(getProcAddr("glVertex3fv"))
	gpVertex3i = (C.GPVERTEX3I)(getProcAddr("glVertex3i"))
	gpVertex3iv = (C.GPVERTEX3IV)(getProcAddr("glVertex3iv"))
	gpVertex3s = (C.GPVERTEX3S)(getProcAddr("glVertex3s"))
	gpVertex3sv = (C.GPVERTEX3SV)(getProcAddr("glVertex3sv"))
	gpVertex4d = (C.GPVERTEX4D)(getProcAddr("glVertex4d"))
	gpVertex4dv = (C.GPVERTEX4DV)(getProcAddr("glVertex4dv"))
	gpVertex4f = (C.GPVERTEX4F)(getProcAddr("glVertex4f"))
	gpVertex4fv = (C.GPVERTEX4FV)(getProcAddr("glVertex4fv"))
	gpVertex4i = (C.GPVERTEX4I)(getProcAddr("glVertex4i"))
	gpVertex4iv = (C.GPVERTEX4IV)(getProcAddr("glVertex4iv"))
	gpVertex4s = (C.GPVERTEX4S)(getProcAddr("glVertex4s"))
	gpVertex4sv = (C.GPVERTEX4SV)(getProcAddr("glVertex4sv"))
	gpVertexArrayAttribBinding = (C.GPVERTEXARRAYATTRIBBINDING)(getProcAddr("glVertexArrayAttribBinding"))
	gpVertexArrayAttribFormat = (C.GPVERTEXARRAYATTRIBFORMAT)(getProcAddr("glVertexArrayAttribFormat"))
	gpVertexArrayAttribIFormat = (C.GPVERTEXARRAYATTRIBIFORMAT)(getProcAddr("glVertexArrayAttribIFormat"))
	gpVertexArrayAttribLFormat = (C.GPVERTEXARRAYATTRIBLFORMAT)(getProcAddr("glVertexArrayAttribLFormat"))
	gpVertexArrayBindingDivisor = (C.GPVERTEXARRAYBINDINGDIVISOR)(getProcAddr("glVertexArrayBindingDivisor"))
	gpVertexArrayElementBuffer = (C.GPVERTEXARRAYELEMENTBUFFER)(getProcAddr("glVertexArrayElementBuffer"))
	gpVertexArrayVertexBuffer = (C.GPVERTEXARRAYVERTEXBUFFER)(getProcAddr("glVertexArrayVertexBuffer"))
	gpVertexArrayVertexBuffers = (C.GPVERTEXARRAYVERTEXBUFFERS)(getProcAddr("glVertexArrayVertexBuffers"))
	gpVertexAttrib1d = (C.GPVERTEXATTRIB1D)(getProcAddr("glVertexAttrib1d"))
	gpVertexAttrib1dv = (C.GPVERTEXATTRIB1DV)(getProcAddr("glVertexAttrib1dv"))
	gpVertexAttrib1f = (C.GPVERTEXATTRIB1F)(getProcAddr("glVertexAttrib1f"))
	gpVertexAttrib1fv = (C.GPVERTEXATTRIB1FV)(getProcAddr("glVertexAttrib1fv"))
	gpVertexAttrib1s = (C.GPVERTEXATTRIB1S)(getProcAddr("glVertexAttrib1s"))
	gpVertexAttrib1sv = (C.GPVERTEXATTRIB1SV)(getProcAddr("glVertexAttrib1sv"))
	gpVertexAttrib2d = (C.GPVERTEXATTRIB2D)(getProcAddr("glVertexAttrib2d"))
	gpVertexAttrib2dv = (C.GPVERTEXATTRIB2DV)(getProcAddr("glVertexAttrib2dv"))
	gpVertexAttrib2f = (C.GPVERTEXATTRIB2F)(getProcAddr("glVertexAttrib2f"))
	gpVertexAttrib2fv = (C.GPVERTEXATTRIB2FV)(getProcAddr("glVertexAttrib2fv"))
	gpVertexAttrib2s = (C.GPVERTEXATTRIB2S)(getProcAddr("glVertexAttrib2s"))
	gpVertexAttrib2sv = (C.GPVERTEXATTRIB2SV)(getProcAddr("glVertexAttrib2sv"))
	gpVertexAttrib3d = (C.GPVERTEXATTRIB3D)(getProcAddr("glVertexAttrib3d"))
	gpVertexAttrib3dv = (C.GPVERTEXATTRIB3DV)(getProcAddr("glVertexAttrib3dv"))
	gpVertexAttrib3f = (C.GPVERTEXATTRIB3F)(getProcAddr("glVertexAttrib3f"))
	gpVertexAttrib3fv = (C.GPVERTEXATTRIB3FV)(getProcAddr("glVertexAttrib3fv"))
	gpVertexAttrib3s = (C.GPVERTEXATTRIB3S)(getProcAddr("glVertexAttrib3s"))
	gpVertexAttrib3sv = (C.GPVERTEXATTRIB3SV)(getProcAddr("glVertexAttrib3sv"))
	gpVertexAttrib4Nbv = (C.GPVERTEXATTRIB4NBV)(getProcAddr("glVertexAttrib4Nbv"))
	gpVertexAttrib4Niv = (C.GPVERTEXATTRIB4NIV)(getProcAddr("glVertexAttrib4Niv"))
	gpVertexAttrib4Nsv = (C.GPVERTEXATTRIB4NSV)(getProcAddr("glVertexAttrib4Nsv"))
	gpVertexAttrib4Nub = (C.GPVERTEXATTRIB4NUB)(getProcAddr("glVertexAttrib4Nub"))
	gpVertexAttrib4Nubv = (C.GPVERTEXATTRIB4NUBV)(getProcAddr("glVertexAttrib4Nubv"))
	gpVertexAttrib4Nuiv = (C.GPVERTEXATTRIB4NUIV)(getProcAddr("glVertexAttrib4Nuiv"))
	gpVertexAttrib4Nusv = (C.GPVERTEXATTRIB4NUSV)(getProcAddr("glVertexAttrib4Nusv"))
	gpVertexAttrib4bv = (C.GPVERTEXATTRIB4BV)(getProcAddr("glVertexAttrib4bv"))
	gpVertexAttrib4d = (C.GPVERTEXATTRIB4D)(getProcAddr("glVertexAttrib4d"))
	gpVertexAttrib4dv = (C.GPVERTEXATTRIB4DV)(getProcAddr("glVertexAttrib4dv"))
	gpVertexAttrib4f = (C.GPVERTEXATTRIB4F)(getProcAddr("glVertexAttrib4f"))
	gpVertexAttrib4fv = (C.GPVERTEXATTRIB4FV)(getProcAddr("glVertexAttrib4fv"))
	gpVertexAttrib4iv = (C.GPVERTEXATTRIB4IV)(getProcAddr("glVertexAttrib4iv"))
	gpVertexAttrib4s = (C.GPVERTEXATTRIB4S)(getProcAddr("glVertexAttrib4s"))
	gpVertexAttrib4sv = (C.GPVERTEXATTRIB4SV)(getProcAddr("glVertexAttrib4sv"))
	gpVertexAttrib4ubv = (C.GPVERTEXATTRIB4UBV)(getProcAddr("glVertexAttrib4ubv"))
	gpVertexAttrib4uiv = (C.GPVERTEXATTRIB4UIV)(getProcAddr("glVertexAttrib4uiv"))
	gpVertexAttrib4usv = (C.GPVERTEXATTRIB4USV)(getProcAddr("glVertexAttrib4usv"))
	gpVertexAttribBinding = (C.GPVERTEXATTRIBBINDING)(getProcAddr("glVertexAttribBinding"))
	gpVertexAttribDivisor = (C.GPVERTEXATTRIBDIVISOR)(getProcAddr("glVertexAttribDivisor"))
	gpVertexAttribFormat = (C.GPVERTEXATTRIBFORMAT)(getProcAddr("glVertexAttribFormat"))
	gpVertexAttribI1i = (C.GPVERTEXATTRIBI1I)(getProcAddr("glVertexAttribI1i"))
	gpVertexAttribI1iv = (C.GPVERTEXATTRIBI1IV)(getProcAddr("glVertexAttribI1iv"))
	gpVertexAttribI1ui = (C.GPVERTEXATTRIBI1UI)(getProcAddr("glVertexAttribI1ui"))
	gpVertexAttribI1uiv = (C.GPVERTEXATTRIBI1UIV)(getProcAddr("glVertexAttribI1uiv"))
	gpVertexAttribI2i = (C.GPVERTEXATTRIBI2I)(getProcAddr("glVertexAttribI2i"))
	gpVertexAttribI2iv = (C.GPVERTEXATTRIBI2IV)(getProcAddr("glVertexAttribI2iv"))
	gpVertexAttribI2ui = (C.GPVERTEXATTRIBI2UI)(getProcAddr("glVertexAttribI2ui"))
	gpVertexAttribI2uiv = (C.GPVERTEXATTRIBI2UIV)(getProcAddr("glVertexAttribI2uiv"))
	gpVertexAttribI3i = (C.GPVERTEXATTRIBI3I)(getProcAddr("glVertexAttribI3i"))
	gpVertexAttribI3iv = (C.GPVERTEXATTRIBI3IV)(getProcAddr("glVertexAttribI3iv"))
	gpVertexAttribI3ui = (C.GPVERTEXATTRIBI3UI)(getProcAddr("glVertexAttribI3ui"))
	gpVertexAttribI3uiv = (C.GPVERTEXATTRIBI3UIV)(getProcAddr("glVertexAttribI3uiv"))
	gpVertexAttribI4bv = (C.GPVERTEXATTRIBI4BV)(getProcAddr("glVertexAttribI4bv"))
	gpVertexAttribI4i = (C.GPVERTEXATTRIBI4I)(getProcAddr("glVertexAttribI4i"))
	gpVertexAttribI4iv = (C.GPVERTEXATTRIBI4IV)(getProcAddr("glVertexAttribI4iv"))
	gpVertexAttribI4sv = (C.GPVERTEXATTRIBI4SV)(getProcAddr("glVertexAttribI4sv"))
	gpVertexAttribI4ubv = (C.GPVERTEXATTRIBI4UBV)(getProcAddr("glVertexAttribI4ubv"))
	gpVertexAttribI4ui = (C.GPVERTEXATTRIBI4UI)(getProcAddr("glVertexAttribI4ui"))
	gpVertexAttribI4uiv = (C.GPVERTEXATTRIBI4UIV)(getProcAddr("glVertexAttribI4uiv"))
	gpVertexAttribI4usv = (C.GPVERTEXATTRIBI4USV)(getProcAddr("glVertexAttribI4usv"))
	gpVertexAttribIFormat = (C.GPVERTEXATTRIBIFORMAT)(getProcAddr("glVertexAttribIFormat"))
	gpVertexAttribIPointer = (C.GPVERTEXATTRIBIPOINTER)(getProcAddr("glVertexAttribIPointer"))
	gpVertexAttribL1d = (C.GPVERTEXATTRIBL1D)(getProcAddr("glVertexAttribL1d"))
	gpVertexAttribL1dv = (C.GPVERTEXATTRIBL1DV)(getProcAddr("glVertexAttribL1dv"))
	gpVertexAttribL1ui64ARB = (C.GPVERTEXATTRIBL1UI64ARB)(getProcAddr("glVertexAttribL1ui64ARB"))
	gpVertexAttribL1ui64vARB = (C.GPVERTEXATTRIBL1UI64VARB)(getProcAddr("glVertexAttribL1ui64vARB"))
	gpVertexAttribL2d = (C.GPVERTEXATTRIBL2D)(getProcAddr("glVertexAttribL2d"))
	gpVertexAttribL2dv = (C.GPVERTEXATTRIBL2DV)(getProcAddr("glVertexAttribL2dv"))
	gpVertexAttribL3d = (C.GPVERTEXATTRIBL3D)(getProcAddr("glVertexAttribL3d"))
	gpVertexAttribL3dv = (C.GPVERTEXATTRIBL3DV)(getProcAddr("glVertexAttribL3dv"))
	gpVertexAttribL4d = (C.GPVERTEXATTRIBL4D)(getProcAddr("glVertexAttribL4d"))
	gpVertexAttribL4dv = (C.GPVERTEXATTRIBL4DV)(getProcAddr("glVertexAttribL4dv"))
	gpVertexAttribLFormat = (C.GPVERTEXATTRIBLFORMAT)(getProcAddr("glVertexAttribLFormat"))
	gpVertexAttribLPointer = (C.GPVERTEXATTRIBLPOINTER)(getProcAddr("glVertexAttribLPointer"))
	gpVertexAttribP1ui = (C.GPVERTEXATTRIBP1UI)(getProcAddr("glVertexAttribP1ui"))
	gpVertexAttribP1uiv = (C.GPVERTEXATTRIBP1UIV)(getProcAddr("glVertexAttribP1uiv"))
	gpVertexAttribP2ui = (C.GPVERTEXATTRIBP2UI)(getProcAddr("glVertexAttribP2ui"))
	gpVertexAttribP2uiv = (C.GPVERTEXATTRIBP2UIV)(getProcAddr("glVertexAttribP2uiv"))
	gpVertexAttribP3ui = (C.GPVERTEXATTRIBP3UI)(getProcAddr("glVertexAttribP3ui"))
	gpVertexAttribP3uiv = (C.GPVERTEXATTRIBP3UIV)(getProcAddr("glVertexAttribP3uiv"))
	gpVertexAttribP4ui = (C.GPVERTEXATTRIBP4UI)(getProcAddr("glVertexAttribP4ui"))
	gpVertexAttribP4uiv = (C.GPVERTEXATTRIBP4UIV)(getProcAddr("glVertexAttribP4uiv"))
	gpVertexAttribPointer = (C.GPVERTEXATTRIBPOINTER)(getProcAddr("glVertexAttribPointer"))
	gpVertexBindingDivisor = (C.GPVERTEXBINDINGDIVISOR)(getProcAddr("glVertexBindingDivisor"))
	gpVertexPointer = (C.GPVERTEXPOINTER)(getProcAddr("glVertexPointer"))
	gpViewport = (C.GPVIEWPORT)(getProcAddr("glViewport"))
	gpViewportArrayv = (C.GPVIEWPORTARRAYV)(getProcAddr("glViewportArrayv"))
	gpViewportIndexedf = (C.GPVIEWPORTINDEXEDF)(getProcAddr("glViewportIndexedf"))
	gpViewportIndexedfv = (C.GPVIEWPORTINDEXEDFV)(getProcAddr("glViewportIndexedfv"))
	gpWaitSync = (C.GPWAITSYNC)(getProcAddr("glWaitSync"))
	gpWindowPos2d = (C.GPWINDOWPOS2D)(getProcAddr("glWindowPos2d"))
	gpWindowPos2dv = (C.GPWINDOWPOS2DV)(getProcAddr("glWindowPos2dv"))
	gpWindowPos2f = (C.GPWINDOWPOS2F)(getProcAddr("glWindowPos2f"))
	gpWindowPos2fv = (C.GPWINDOWPOS2FV)(getProcAddr("glWindowPos2fv"))
	gpWindowPos2i = (C.GPWINDOWPOS2I)(getProcAddr("glWindowPos2i"))
	gpWindowPos2iv = (C.GPWINDOWPOS2IV)(getProcAddr("glWindowPos2iv"))
	gpWindowPos2s = (C.GPWINDOWPOS2S)(getProcAddr("glWindowPos2s"))
	gpWindowPos2sv = (C.GPWINDOWPOS2SV)(getProcAddr("glWindowPos2sv"))
	gpWindowPos3d = (C.GPWINDOWPOS3D)(getProcAddr("glWindowPos3d"))
	gpWindowPos3dv = (C.GPWINDOWPOS3DV)(getProcAddr("glWindowPos3dv"))
	gpWindowPos3f = (C.GPWINDOWPOS3F)(getProcAddr("glWindowPos3f"))
	gpWindowPos3fv = (C.GPWINDOWPOS3FV)(getProcAddr("glWindowPos3fv"))
	gpWindowPos3i = (C.GPWINDOWPOS3I)(getProcAddr("glWindowPos3i"))
	gpWindowPos3iv = (C.GPWINDOWPOS3IV)(getProcAddr("glWindowPos3iv"))
	gpWindowPos3s = (C.GPWINDOWPOS3S)(getProcAddr("glWindowPos3s"))
	gpWindowPos3sv = (C.GPWINDOWPOS3SV)(getProcAddr("glWindowPos3sv"))
	return nil
}
