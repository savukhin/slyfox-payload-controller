cmake_minimum_required(VERSION 3.23)

set (CMAKE_CXX_STANDARD 20)


add_subdirectory(ImageAnalyzer)
# add_subdirectory(Devices)


add_executable(main main.cpp)
# target_link_libraries(main opencv_core opencv_highgui)

add_subdirectory(tests)

# target_include_directories(ocv_test PRIVATE
#         ${OPENCV_CONFIG_FILE_INCLUDE_DIR}
#         ${OPENCV_MODULE_opencv_core_LOCATION}/include
#         ${OPENCV_MODULE_opencv_highgui_LOCATION}/include
    # )
target_link_libraries(main ${MAIN_DEPENDENCIES})

# target_link_libraries( main ${OpenCV_LIBS} )
