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
  $mysqli = getDB();

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

    $query = "INSERT INTO instagram (account,followed_by,follows,photos) VALUES (?,?,?,?)";
    if ($stmt = $mysqli->prepare($query)) {
      $stmt->bind_param("siii",
        $account,
        $data->user->followed_by->count,
        $data->user->follows->count,
        $data->user->media->count
      );
      if (!$stmt->execute()) {
          var_dump($stmt->error);
      } else {
        print("Saved\n");
      }
      $stmt->close();
    } else {
      var_dump($mysqli->error);
    }
  }
  $mysqli->close();
  sleep(60);
}
