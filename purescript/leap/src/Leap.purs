module Leap
  ( isLeapYear
  ) where

import Prelude

isLeapYear :: Int -> Boolean
isLeapYear num = 
   (num `mod` 400 == 0) || (num `mod` 4 == 0 && num `mod` 100 /= 0)