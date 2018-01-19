<?php
/**
 * Created by PhpStorm.
 * User: puzanov
 * Date: 7/7/17
 * Time: 7:55 PM
 */

while (true) {
  $accounts = array();
  if (getenv('INSTAGRAM_ACCOUNTS')) {
    $accounts = explode(',', getenv('INSTAGRAM_ACCOUNTS'));
  } else {
    print("No Instagram accounts\n");
    exit;
  }

  foreach ($accounts as $account) {
    $url = "https://www.instagram.com/$account/?__a=1";

    $json = file_get_contents($url);
    $data = json_decode($json);

    print("Instagram followed_by " . $data->user->followed_by->count . "\n");
    print("Instagram follows " . $data->user->follows->count . "\n");
    print("Instagram photos " . $data->user->media->count . "\n");
  }
  
  sleep(60);
}
