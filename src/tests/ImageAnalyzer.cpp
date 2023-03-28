#include <opencv2/imgcodecs/imgcodecs_c.h>
#include <gtest/gtest.h>
#include "../ImageAnalyzer/ImageAnalyzer.hpp"
#include <gmock/gmock.h>

TEST(TestGroupName2, Subtest_2) {
  auto img = cvLoadImageM("img.png");
  std::cout << img->rows << " " << img->cols << std::endl; // при этом будет выведено на экран данное сообщение
  
}