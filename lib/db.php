<?php
/**
 * Created by PhpStorm.
 * User: puzanov
 * Date: 7/9/17
 * Time: 2:38 PM
 */

function getDB() {
  $mysqli = new mysqli("mysql", "root", "", "glaza");

  if (mysqli_connect_errno()) {
      printf("Mysql connection error: %s\n", mysqli_connect_error());
      exit();
  }

  return $mysqli;
}
