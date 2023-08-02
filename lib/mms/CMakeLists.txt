cmake_minimum_required(VERSION 3.16)
project(mms C)

set(CMAKE_C_STANDARD 99)

find_package(PkgConfig REQUIRED)

pkg_check_modules(GLIB REQUIRED glib-2.0)
pkg_check_modules(LIBAVFORMAT REQUIRED libavformat)
pkg_check_modules(LIBAVCODEC REQUIRED libavcodec)
pkg_check_modules(LIBAVUTIL REQUIRED libavutil)


add_executable(mms main.c video_reader.c video_reader.h repo.c repo.h rtp.h proc.c proc.h)
add_library(mms_media SHARED video_reader.c video_reader.h rtp.c rtp.h)

target_include_directories(mms PUBLIC ${LIBAVFORMAT_INCLUDE_DIRS} ${LIBAVCODEC_INCLUDE_DIRS} ${LIBAVUTIL_INCLUDE_DIRS} ${GLIB_INCLUDE_DIRS}})
target_link_libraries(mms ${LIBAVFORMAT_LINK_LIBRARIES} ${LIBAVUTIL_LINK_LIBRARIES} ${LIBAVCODEC_LINK_LIBRARIES} ${GLIB_LINK_LIBRARIES})
target_link_libraries(mms_media ${LIBAVFORMAT_LINK_LIBRARIES} ${LIBAVUTIL_LINK_LIBRARIES} ${LIBAVCODEC_LINK_LIBRARIES} ${GLIB_LINK_LIBRARIES})
set_target_properties(mms_media PROPERTIES PUBLIC_HEADER "video_reader.h;rtp.h")

install(TARGETS mms_media DESTINATION /usr/lib PUBLIC_HEADER DESTINATION usr/include)