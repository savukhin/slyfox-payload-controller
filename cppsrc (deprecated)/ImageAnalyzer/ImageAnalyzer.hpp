#pragma once

#include <opencv2/opencv.hpp>

#include <utility>
#include <vector>

namespace slyfox {
    using CameraPoint = cv::Point;
    using CameraResolution = CameraPoint;
    using Clusters = std::vector<CameraPoint>;

    class ImageAnalyzer {
    private:
        CameraResolution resolution;
    public:
        ImageAnalyzer() = default;
        ImageAnalyzer(CameraResolution resolution);

        static int GetValue(int value);

        Clusters Analyze(cv::Mat img_thresholded, int min_area=20, int max_area=100);
    };
}
