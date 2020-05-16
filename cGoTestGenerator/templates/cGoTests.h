// +build cGoTests

/// \file
/// \brief This ANSI-C Header file provides the ANSI-C based cGoTest testing 
/// framework. 
///
/// Package description:
///   {{ .BriefDesc }}
///
/// This file is automatically (re)generated changes made to this file will 
/// be lost. 

#ifndef CGO_TESTS_H
#define CGO_TESTS_H

#include <stdio.h>
#include <string.h>
#include <memory.h>

/// \brief Log a message to the GoLang testing systems using sprintf. 
///
#define cGoTestLogf(...)          \
do {                              \
  char strBuff[1024];             \
  memset(strBuff, 0, 1024);       \
  sprintf(strBuff, __VA_ARGS__);  \
  cGoTestLog(strBuff);            \
} while (0)

/// \brief Log the current line and file to the GoLang testing system. 
///
#define cGoTestLogLineFile(aLine, aFile)    \
cGoTestLogf("  at line: %d in: %s", aLine, aFile)

/// \brief Assert the test `test` is true. If `test` is false, log the 
/// message. 
///
#define cGoTest(message, test)    \
if (!(test)) cGoTestLog(message);

/// \brief Assert the test `test` is true. If `test` is false, log the 
/// message and FAIL all further tests. 
///
#define cGoTest_MayFail(message, test) \
if (!(test)) return message

/// \brief Assert that `aPtr` is not nil. Log a message if `aPtr` is nil. 
/// Capture the calling line and file. 
///
#define cGoTest_NotNil(message, aPtr) \
cGoTest_NotNil_LineFile(message, aPtr, __LINE__, __FILE__)

/// \brief Assert that `aPtr` is not nil. Log a message if `aPtr` is nil. 
///
#define cGoTest_NotNil_LineFile(message, aPtr, aLine, aFile)  \
if ((aPtr) == 0) {                                            \
  cGoTestLog(message);                                        \
  cGoTestLogLineFile(aLine, aFile);                           \
}

/// \brief Assert that `aPtr` is not nil. Log a message if `aPtr` is nil. 
/// Capture the calling line and file. 
///
#define cGoTest_NotNil_MayFail(message, aPtr) \
cGoTest_NotNil_MayFail_LineFile(message, aPtr, __LINE__, __FILE__)

/// \brief Assert that `aPtr` is not nil. Log a message if `aPtr` is nil. 
///
#define cGoTest_NotNil_MayFail_LineFile(message, aPtr, aLine, aFile) \
if ((aPtr) == 0) {                                                   \
  cGoTestLog(message);                                               \
  cGoTestLogLineFile(aLine, aFile);                                  \
  return message;                                                    \
}

/// \brief Assert that `aPtr` is nil. Log a message if `aPtr` is not nil. 
/// Capture the calling line and file. 
///
#define cGoTest_Nil(message, aPtr) \
cGoTest_Nil_LineFile(message, aPtr, __LINE__, __FILE__)

/// \brief Assert that `aPtr` is nil. Log a message if `aPtr` is not nil. 
///
#define cGoTest_Nil_LineFile(message, aPtr, aLine, aFile)  \
if ((aPtr) != 0) {                                                  \
  cGoTestLog(message);                                                \
  cGoTestLogLineFile(aLine, aFile);                                   \
}

#define cGoUInt unsigned long

/// \brief Assert that `aUInt` == `bUInt`. Log a message if they are not 
/// equal. Capture the calling line and file. 
///
#define cGoTest_UIntEquals(message, aUInt, bUInt)                       \
cGoTest_UIntEquals_LineFile(message, aUInt, bUInt, __LINE__, __FILE__)

/// \brief Assert that `aUInt` == `bUInt`. Log a message if they are not 
/// equal. 
///
#define cGoTest_UIntEquals_LineFile(message, aUInt, bUInt, aLine, aFile)  \
if ((aUInt) != (bUInt)) {                                                 \
  cGoTestLog(message);                                                    \
  cGoTestLogf("  aUInt: %lu", ((cGoUInt)aUInt));                          \
  cGoTestLogf("  bUInt: %lu", ((cGoUInt)bUInt));                          \
  cGoTestLogLineFile(aLine, aFile);                                       \
}

/// \brief Assert that `aUInt` != `bUInt`. Log a message if they are 
/// equal. Capture the calling line and file. 
///
#define cGoTest_UIntNotEquals(message, aUInt, bUInt)                       \
cGoTest_UIntNotEquals_LineFile(message, aUInt, bUInt, __LINE__, __FILE__)

/// \brief Assert that `aUInt` == `bUInt`. Log a message if they are not 
/// equal. 
///
#define cGoTest_UIntNotEquals_LineFile(message, aUInt, bUInt, aLine, aFile)  \
if ((aUInt) == (bUInt)) {                                                    \
  cGoTestLog(message);                                                       \
  cGoTestLogf("  aUInt: %lu", ((cGoUInt)aUInt));                             \
  cGoTestLogf("  bUInt: %lu", ((cGoUInt)bUInt));                             \
  cGoTestLogLineFile(aLine, aFile);                                          \
}

#define cGoStr const char*

/// \brief Assert that `theStr` contains `aWord`. Log a message if 
/// `theStr` does not contain `aWord`. 
///
#define cGoTest_StrContains(message, theStr, aWord)                     \
cGoTest_StrContains_LineFile(message, theStr, aWord, __LINE__, __FILE__)

/// \brief Assert that `theStr` contains `aWord`. Log a message if 
/// `theStr` does not contain `aWord`. 
///
#define cGoTest_StrContains_LineFile(message, theStr, aWord, aLine, aFile) \
if (strstr((theStr), (aWord)) == NULL) {                                           \
  cGoTestLog(message);                                                     \
  cGoTestLogf("  theStr: %s", ((cGoStr)theStr));                           \
  cGoTestLogf("   aWord: %s", ((cGoStr)aWord));                            \
  cGoTestLogLineFile(aLine, aFile);                                        \
}

/// \brief Assert that `theStr` does not contain the `aWord`. Log a 
/// message if `theStr` does contain the `aWord`. 
///
#define cGoTest_StrNotContains(message, theStr, aWord)                     \
cGoTest_StrNotContains_LineFile(message, theStr, aWord, __LINE__, __FILE__)

/// \brief Assert that `theStr` does not contain the `aWord`. Log a 
/// message if `theStr` does contain the `aWord`. 
///
#define cGoTest_StrNotContains_LineFile(message, theStr, aWord, aLine, aFile) \
if (strstr((theStr), (aWord)) != NULL) {                                           \
  cGoTestLog(message);                                                     \
  cGoTestLogf("  theStr: %s", ((cGoStr)theStr));                           \
  cGoTestLogf("   aWord: %s", ((cGoStr)aWord));                            \
  cGoTestLogLineFile(aLine, aFile);                                        \
}

/// \brief Assert that `aStr` equals `bStr`. Log a message if they 
/// are not equal. 
///
#define cGoTest_StrEquals(message, aStr, bStr)                       \
cGoTest_StrEquals_LineFile(message, aStr, bStr, __LINE__, __FILE__)

/// \brief Assert that `aStr` equals `bStr`. Log a message if they 
/// are not equal. 
///
#define cGoTest_StrEquals_LineFile(message, aStr, bStr, aLine, aFile) \
if (strcmp((aStr), (bStr)) != 0) {                                    \
  cGoTestLog(message);                                                \
  cGoTestLogf("  aStr: %s", ((cGoStr)aStr));                          \
  cGoTestLogf("  bStr: %s", ((cGoStr)bStr));                          \
  cGoTestLogLineFile(aLine, aFile);                                   \
}

/// \brief Assert that `aStr` does not equal `bStr`. Log a message if they 
/// are equal. 
///
#define cGoTest_StrNotEquals(message, aStr, bStr)                       \
cGoTest_StrNotEquals_LineFile(message, aStr, bStr, __LINE__, __FILE__)

/// \brief Assert that `aStr` does not equal `bStr`. Log a message if they 
/// are equal. 
///
#define cGoTest_StrNotEquals_LineFile(message, aStr, bStr, aLine, aFile) \
if (strcmp((aStr), (bStr)) == 0) {                                       \
  cGoTestLog(message);                                                   \
  cGoTestLogf("  aStr: %s", ((cGoStr)aStr));                             \
  cGoTestLogf("  bStr: %s", ((cGoStr)bStr));                             \
  cGoTestLogLineFile(aLine, aFile);                                      \
}

#endif