<!doctype HTML>
<html>
<body>
<form action="instagram.php" method="get">
<input type="text" name="username">
<input type="submit" value="Go">
</form>
</body>
</html>

<?php

$var = $_GET["username"]; 
$json = file_get_contents("https://www.instagram.com/$var/?__a=1");
if (empty($json)){
	$message = "no found";
	echo "<script type='text/javascript'>alert('$message');</script>";
}
$arr = json_decode($json, true);

$post =  $arr['graphql']['user']['edge_owner_to_timeline_media']['count'];
$fwer = $arr['graphql']['user']['edge_followed_by']['count'];
$fwing = $arr['graphql']['user']['edge_follow']['count'];

echo '<br/>';
echo 'Posts: ' .  $post;
echo ' <br/>';
echo 'Followers: ' .  $fwer;
echo ' <br/>';
echo 'Following: ' . $fwing;


$data = array(
	"posts" =>  $post,
	"followers" => $fwer,
	"following" => $fwing
);


$con = new MongoClient();
$list= $con-> test-> listCollections();
foreach ($list as $collection) {
    echo "<p>" . $collection . "</p>";
}
$con-> test;
$collection= $con-> test-> $var;
if (!empty($json)){
$collection->insert($data);
}
$con-> close();

?>
