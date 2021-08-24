<?php

function createPhoneNumber($numbersArray) {
  $phone = "";
  foreach($numbersArray as $key => $value) {
    if($key == 0){
      $phone .= "(".$value;
    } elseif ($key == 2) {
      $phone .= $value.") ";
    } elseif ($key == 6) {
      $phone .= "-".$value;
    } else {
      $phone .= $value;
    }
  }

  return $phone;
}

createPhoneNumber([1, 2, 3, 4, 5, 6, 7, 8, 9, 0]);
