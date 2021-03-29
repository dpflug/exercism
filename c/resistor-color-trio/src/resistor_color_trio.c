#include "resistor_color_trio.h"
#include <math.h>
#include <stdio.h>

const char *OHMS = "ohms";
const char *KILOOHMS = "kiloohms";

struct resistor_value_t color_code(const resistor_band_t colors[static 3]) {
  resistor_value_t result = {.value = ((colors[0] * 10) + colors[1]) *
                                      pow(10, colors[2]),
                             .unit = OHMS};
  if (result.value % 1000 == 0) {
    result.value /= 1000;
    result.unit = KILOOHMS;
  }
  return result;
}
