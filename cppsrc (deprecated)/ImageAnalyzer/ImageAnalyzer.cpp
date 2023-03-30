#include "ImageAnalyzer.hpp"
#include <chrono>
#include <opencv2/core/hal/interface.h>

namespace slyfox {
    ImageAnalyzer::ImageAnalyzer(CameraResolution resolution): resolution(resolution) {
    }

    int ImageAnalyzer::GetValue(int value) {
        return value;
    }

    Clusters ImageAnalyzer::Analyze(cv::Mat img_thresholded, int min_area, int max_area) {
        std::vector<std::vector<cv::Point>> contours;
        cv::findContours(img_thresholded, contours, cv::RETR_TREE, cv::CHAIN_APPROX_SIMPLE);

        Clusters clusters;

        for (const auto& contour : contours)
        {
            cv::Rect rect = cv::boundingRect(contour);
            auto area = rect.area();
            if (area < min_area || area > max_area) {
                continue;
            }
            
            clusters.emplace_back(rect.x + rect.width / 2., rect.y + rect.height / 2.);
        }

        return clusters;
    }
}