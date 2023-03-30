include(FetchContent)

# project(lib_crsf)

# FetchContent_Declare(
#     wiringPi
#     GIT_REPOSITORY https://github.com/WiringPi/WiringPi.git
#     GIT_TAG        2.61-1
#     # GIT_SHALLOW TRUE
#     # GIT_PROGRESS TRUE
# )
# FetchContent_GetProperties(wiringPi)
# if (NOT googletest_POPULATED)
#     FetchContent_Populate(wiringPi)

#     # add_subdirectory(${wiringPi_SOURCE_DIR} ${wiringPi_BINARY_DIR})

# endif ()
# FetchContent_MakeAvailable(wiringPi)
# add_library(GTest::GTest INTERFACE IMPORTED)
# target_link_libraries(GTest::GTest INTERFACE gtest_main)

# find_library(wiringPi_LIB wiringPi)


# include(FetchContent)
# FetchContent_Declare(
#   Boost
#   GIT_REPOSITORY https://github.com/boostorg/boost.git
#   GIT_TAG boost-1.80.0
# )
# FetchContent_MakeAvailable(Boost)

# FIND_PACKAGE( Boost COMPONENTS asio REQUIRED )
# INCLUDE_DIRECTORIES( ${Boost_INCLUDE_DIR} )

# ADD_EXECUTABLE( anyExecutable myMain.cpp )

# TARGET_LINK_LIBRARIES( anyExecutable LINK_PUBLIC ${Boost_LIBRARIES} )

include(FetchContent)

FetchContent_Declare(
  serial
  GIT_REPOSITORY https://github.com/wjwwood/serial.git
  GIT_TAG        1.2.1
  # GIT_SHALLOW TRUE
  # GIT_PROGRESS TRUE
)
FetchContent_GetProperties(serial)
if (NOT googletest_POPULATED)
    FetchContent_Populate(serial)

    add_subdirectory(${serial_SOURCE_DIR} ${serial_BINARY_DIR})
endif ()
FetchContent_MakeAvailable(serial)
add_library(GTest::GTest INTERFACE IMPORTED)
target_link_libraries(GTest::GTest INTERFACE gtest_main)

set(SRC
    Crsf.cpp
)

set(MAIN_DEPENDENCIES
    ${MAIN_DEPENDENCIES}
    crsf
    PARENT_SCOPE
)

add_library(crsf STATIC ${SRC})
# target_link_libraries(crsf LINK_PUBLIC ${Boost_LIBRARIES})
target_link_libraries(crsf LINK_PUBLIC ${serial_LIBRARIES})

# target_link_libraries(crsf ${wiringPi_LIB})
