#include <gtest/gtest.h>
#include <gmock/gmock.h>

#include "../Crsf/Crsf.hpp"

TEST(Crsf, channels) {
    slyfox::Crsf crsf;
    crsf.handleByte(1);
}
