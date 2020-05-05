#include "darts.h"
#include <math.h>

int score(coordinate_t point) {
  float dist = sqrt(pow(point.x, 2) + pow(point.y, 2));
  if (dist > 10.0) {
    return 0;
  } else if (dist > 5.0) {
    return 1;
  } else if (dist > 1) {
    return 5;
  } else {
    return 10;
  }
}
