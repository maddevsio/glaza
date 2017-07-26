<?php
/**
 * Created by PhpStorm.
 * User: puzanov
 * Date: 7/7/17
 * Time: 7:55 PM
 */

set_include_path(get_include_path() . PATH_SEPARATOR . '../lib' . PATH_SEPARATOR . './lib');

include_once "db.php";

while (true) {
  $url = "https://www.instagram.com/hippohae/?__a=1";
  $json = file_get_contents($url);
  $data = json_decode($json);
  print("Instagram followed_by " . $data->user->followed_by->count . "\n");
  if (saveLog("instagram_hippohae_followed_by", strval($data->user->followed_by->count))) {
    print("Saved\n");
  }
  sleep(60);
}
