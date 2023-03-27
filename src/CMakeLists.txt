include_directories(Devices)
include_directories(ImageAnalyzer)
include_directories(SensorAnalyzer)
include_directories(MovementController)

add_subdirectory(tests)

add_executable(main main.cpp)
