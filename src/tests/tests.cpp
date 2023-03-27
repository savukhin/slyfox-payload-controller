#include <gtest/gtest.h>
#include <gmock/gmock.h>
#include "ImageAnalyzer.hpp"
#include <iostream>

TEST(TestGroupName, Subtest_1) {
  EXPECT_EQ(1, 1); // логи покажут тут ошибку
  std::cout << "continue test" << std::endl; // при этом будет выведено на экран данное сообщение
}

int main(int argc, char **argv)
{
  ::testing::InitGoogleTest(&argc, argv);
  ::testing::InitGoogleMock(&argc, argv);
  
  return RUN_ALL_TESTS();
}