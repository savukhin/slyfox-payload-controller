#include <opencv2/imgcodecs/imgcodecs_c.h>
#include <opencv2/imgproc.hpp>
#include <opencv2/opencv.hpp>
#include <opencv2/core/types.hpp>
#include <gtest/gtest.h>
#include "../ImageAnalyzer/ImageAnalyzer.hpp"
#include <gmock/gmock.h>
#include <tuple>
#include <map>
#include <vector>
#include <opencv2/core.hpp>

TEST(ImageAnalyzer, img1) {
    auto img = cv::imread("suitable_img.png", cv::IMREAD_GRAYSCALE);
    EXPECT_FALSE(img.empty());

    slyfox::ImageAnalyzer analyzer({img.rows, img.cols});
    slyfox::Clusters clusters = analyzer.Analyze(img);

    ASSERT_EQ(clusters.size(), 41);

}


TEST(ImageAnalyzer, img2) {
    auto img = cv::imread("img.png", cv::IMREAD_GRAYSCALE);

    cv::resize(img, img, {500, 500});
    cv::threshold(img, img, 100, 255, cv::THRESH_BINARY);

    slyfox::ImageAnalyzer analyzer({img.rows, img.cols});
    slyfox::Clusters clusters = analyzer.Analyze(img);


    ASSERT_EQ(clusters.size(), 179);

}
