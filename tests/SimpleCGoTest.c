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