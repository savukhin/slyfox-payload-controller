#pragma once

#include <opencv2/opencv.hpp>

#include <utility>
#include <vector>

namespace slyfox {
    struct CameraPoint {
        int x;
        int y;
    };

    using CameraResolution = CameraPoint;

    struct Clusters: public std::vector<CameraPoint> {
        // std::vector<CameraPoint> points;
    };

    class ImageAnalyzer {
    private:
        CameraResolution resolution;
    public:
        ImageAnalyzer() = default;
        ImageAnalyzer(CameraResolution resolution);

        static int GetValue(int value);

        Clusters Analyze(int a);
    };
}
