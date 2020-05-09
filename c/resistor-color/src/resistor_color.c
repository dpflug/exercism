#include "resistor_color.h"

int color_code(resistor_band_t color) {
  return color;
}

resistor_band_t * colors() {
  static resistor_band_t output[11];
  for (int i=0; i<10; i++) {
    output[i] = i;
  }
  return output;
}
