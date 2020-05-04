// A simple cGoTest 

/// \testFixture testFixture A test fixture

/// \testFixture sillyFixture A very silly fixture

/// \brief Test the JBlock utilities
/// \silly This is a silly comment
///
/// \inFixture sillyFixture
///
char *newJBlockCGoTest(void* data) {
  JBlock *aJBlock = newJBlock(100, 2);
  
  cTest_NotNil_MayFail("aJBlock nil", aJBlock);
  cTest_UIntEquals("wrong aJBlock.size", aJBlock->size, (size_t)(100*2));
  
  return 0;
}

/// \brief Test the RM64 utilities
///
/// \inFixture testFixture
///
char *newJRM64CGoTest(void* data) {
  JRM64 *aRegisterMachine = newJRM64();
  
  cTest_NotNil_MayFail("aRegisterMachine nil", aRegisterMachine);
  
  return 0;
}