#include <gtest/gtest.h>
#include <gmock/gmock.h>
#include "../ImageAnalyzer/ImageAnalyzer.hpp"
#include <iostream>

TEST(TestGroupName, Subtest_1) {
  auto analyzer = slyfox::ImageAnalyzer();
  EXPECT_EQ(1, analyzer.GetValue(1)); // логи покажут тут ошибку
  std::cout << "continue test" << std::endl; // при этом будет выведено на экран данное сообщение
}

int main(int argc, char **argv)
{
  ::testing::InitGoogleTest(&argc, argv);
  ::testing::InitGoogleMock(&argc, argv);
  
  return RUN_ALL_TESTS();
}