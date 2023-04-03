#pragma once

#include <math.h>

namespace slyfox {
    class Calculator {
    public:
        // All the values is the projection on one axis
        // H - height of Kurama
        // h - height of payload
        // alpha - angle between Kurama direction and horizon
        // beta - angle between payload direction and horizon
        // d1 - delta in camera on this axis
        static double GetGammaAngleProjection(double H, double h, double alpha, double beta, double d1) {
            auto sina = sin(alpha);
            auto sinb = sin(beta);
            auto ctga = 1. / tan(alpha);
            auto ctgb = 1. / tan(beta);

            double BD = (d1 / sina) + h * (ctgb - ctga);
            double BC = sqrt(h*h + pow(d1/sina - h*ctga, 2));

            return BD / BC;
        }
    };
}