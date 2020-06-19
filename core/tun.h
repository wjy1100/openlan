//
// Created by albert on 6/19/2020.
//

#ifndef CORE_TUN_H
#define CORE_TUN_H

#include <unistd.h>

int create_tap(char *name);
ssize_t read_tap(int fd, void *buf, ssize_t size);
ssize_t write_tap(int fd, const void *buf, ssize_t size);

#endif //CORE_TUN_H
