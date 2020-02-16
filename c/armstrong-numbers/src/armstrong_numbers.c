#include "armstrong_numbers.h"
#include <math.h>

int is_armstrong_number(int candidate)
{
  int clen = (int)log10(candidate)+1;
  int sum = 0;
  for (int i=1; i<=clen; i++) {
    int place = (int)pow(10, i-1);
    // Get digits at/below where we're looking
    int low = candidate % (place*10);
    // and divide up to place we're considering currently
    int digit = low / place;
    sum += pow(digit, clen);
  }
  return candidate == sum;
}
