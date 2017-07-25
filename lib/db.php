<?php
/**
 * Created by PhpStorm.
 * User: puzanov
 * Date: 7/9/17
 * Time: 2:38 PM
 */

function saveLog($type, $value) {
    $status = true;
    $mysqli = new mysqli("127.0.0.1", "root", "", "glaza");

    if (mysqli_connect_errno()) {
        printf("Mysql connection error: %s\n", mysqli_connect_error());
        exit();
    }

    $query = "INSERT INTO log (type, value) VALUES (?,?)";
    $stmt = $mysqli->prepare($query);
    $stmt->bind_param("ss", $typeBind, $valueBind);
    $typeBind = $type;
    $valueBind = $value;
    if (!$stmt->execute()) {
        var_dump($stmt->error);
        $status = false;
    }
    $stmt->close();
    $mysqli->close();
    return $status;
}
