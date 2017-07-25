<?php
/**
 * Created by PhpStorm.
 * User: puzanov
 * Date: 7/7/17
 * Time: 7:55 PM
 */

include_once "../lib/db.php";

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
