// +build tests

/// \file
/// \brief This ANSI-C Header file provides the ANSI-C based cTest testing 
/// framework. 

/// \brief Log a message to the GoLang testing systems using sprintf. 
///
#define cTestLogf(...)            \
do {                              \
  char strBuff[1024];             \
  memset(strBuff, 0, 1024);       \
  sprintf(strBuff, __VA_ARGS__);  \
  cTestLog(strBuff);              \
} while (0)

/// \brief Log the current line and file to the GoLang testing system. 
///
#define cTestLogLineFile(aLine, aFile)    \
cTestLogf("  at line: %d in: %s", aLine, aFile)

/// \brief Assert the test `test` is true. If `test` is false, log the 
/// message. 
///
#define cTest(message, test)    \
if (!(test)) cTestLog(message);

/// \brief Assert the test `test` is true. If `test` is false, log the 
/// message and FAIL all further tests. 
///
#define cTest_MayFail(message, test) \
if (!(test)) return message

/// \brief Assert that `aPtr` is not nil. Log a message if `aPtr` is nil. 
/// Capture the calling line and file. 
///
#define cTest_NotNil_MayFail(message, aPtr) \
cTest_NotNil_MayFail_LineFile(message, aPtr, __LINE__, __FILE__)

/// \brief Assert that `aPtr` is not nil. Log a message if `aPtr` is nil. 
///
#define cTest_NotNil_MayFail_LineFile(message, aPtr, aLine, aFile)  \
if ((aPtr) == 0) {                                                  \
  cTestLog(message);                                                \
  cTestLogLineFile(aLine, aFile);                                   \
  return message;                                                   \
}

/// \brief Assert that `aUInt` == `bUInt`. Log a message if they are not 
/// equal. Capture the calling line and file. 
///
#define cTest_UIntEquals(message, aUInt, bUInt)                       \
cTest_UIntEquals_LineFile(message, aUInt, bUInt, __LINE__, __FILE__)

/// \brief Assert that `aUInt` == `bUInt`. Log a message if they are not 
/// equal. 
///
#define cTest_UIntEquals_LineFile(message, aUInt, bUInt, aLine, aFile)  \
if ((aUInt) != (bUInt)) {                                               \
  cTestLog(message);                                                    \
  cTestLogf("  aUInt: %lu", aUInt);                                     \
  cTestLogf("  bUInt: %lu", bUInt);                                     \
  cTestLogLineFile(aLine, aFile);                                       \
}

