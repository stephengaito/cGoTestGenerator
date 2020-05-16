// +build cGoTests

// A simple cGoTest

#include "_cgo_export.h"
#include "testsCGoTests.h"
#include "cGoTests.h"

/// \testFixture testFixture A test fixture

/// \brief Setup the testFixture
/// \inFixture testFixture
///
void *testFixtureCGoTestSetup(void) {
  return NULL;
}

/// \brief Tear down the testFixture
/// \inFixture testFixture
///
void testFixtureCGoTestTeardown(void *data) {
  // do nothing
}

/// \brief Test the cGoTest pointer assertions
///
/// \inFixture testFixture
///
char *pointerCGoTest(void* data) {
  cGoTest_NotNil_MayFail("a pointer is not nil", pointerCGoTest);
  cGoTest_Nil("a pointer is nil", NULL);
  
  cGoTest_NotNil_MayFail("data is nil so SHOULD fail", data);

  return NULL;
}

/// \brief Test the UInt assertions
///
/// \inFixture testFixture
///
char *uintCGoTest(void* data) {
  cGoTest_UIntEquals("two uints are equal", 42, 42);
  cGoTest_UIntNotEquals("two uints are not equal", 42, 2);
  
  return NULL;
}

char *strCGoTest(void* data) {
  cGoTest_StrContains(
    "Should contain word",
    "This is a test", "test"
  );
  cGoTest_StrNotContains(
    "Should not contain word",
    "This is a test", "contain"
  );

  cGoTest_StrEquals(
    "Strings should be equal",
    "This is a test", "This is a test"
  );
  cGoTest_StrNotEquals(
    "Strings should not be equal",
    "This is a test", "this is a second test"
  );
  
  return NULL;
}