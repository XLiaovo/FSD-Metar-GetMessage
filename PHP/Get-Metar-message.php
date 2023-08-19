<?php
function Get_Metars() {
    $url = "https://metar.vatsim.net/all";
    $response = file_get_contents($url);

    if ($response !== false) {
        if (file_put_contents("metar.txt", $response) !== false) {
            $output = array(
                "code" => 200,
                "msg" => "OK"
            );
        } else {
            $output = array(
                "code" => 500,
                "msg" => "Failed to save the content to file"
            );
        }
    } else {
        $output = array(
            "code" => 500,
            "msg" => "Failed to fetch data from URL"
        );
    }

    echo json_encode($output);
}

Get_Metars();
?>
