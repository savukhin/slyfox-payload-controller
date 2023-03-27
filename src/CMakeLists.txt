project(hello_world)

# include_directories(Devices)
include_directories(ImageAnalyzer)
# include_directories(SensorAnalyzer)
# include_directories(MovementController)

add_subdirectory(ImageAnalyzer)


set(MAIN_DEPENDENCIES
    image_analyzer
)

add_subdirectory(tests)

add_executable(main main.cpp)
target_link_libraries(main ${MAIN_DEPENDENCIES})
