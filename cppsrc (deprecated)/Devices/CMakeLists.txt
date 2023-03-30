project(lib_image_analyzer)

set(SRC
    Devices.cpp
)

set(MAIN_DEPENDENCIES
    ${MAIN_DEPENDENCIES}
    devices
    PARENT_SCOPE
)

add_library(devices STATIC ${SRC})
