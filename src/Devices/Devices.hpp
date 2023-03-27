#pragma once
#include "Gyroscope.hpp"


namespace slyfox {
    class BaseDevices {
    public:
        BaseDevices() = default;

        virtual double GetAngleToHorizon() = 0;
        virtual double GetHeight() = 0;
    };

    class TestDevices {
    private:
        double angleToHorizon = 0;
        double height = 0;

    public:
        double GetAngleToHorizon() { return angleToHorizon; } 
        double GetHeight() { return height; }

        double SetAngleToHorizon(double angle) { this->angleToHorizon = angle; }
        double SetHeight(double height) { this->height = height; }
    };
}