import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.util.*;

public class Day08 {

    public static void main(String[] args) {
        int WIDTH = 25;
        int HEIGHT = 6;
        int layerSize = WIDTH * HEIGHT;
        String line = "";
        String filePath = "input.txt";
        try (BufferedReader br = new BufferedReader(new FileReader(filePath))) {
            line = br.readLine();
        } catch (IOException e) {
            e.printStackTrace();
        }

        List<Integer> imgData = new ArrayList<>();
        for (int i = 0; i < line.length(); i++) {
            int digit = line.charAt(i) - '0';
            imgData.add(digit);
        }

        List<List<Integer>> layers = new ArrayList<>();
        for (int i = 0; i < imgData.size(); i += layerSize) {
            layers.add(imgData.subList(i, i + layerSize));
        }

        int min0s = 9999999;
        int minLayer = -1;
        int layerIdx = 0;
        for (List<Integer> layer : layers) {
            if (countOccurances(layer, 0) < min0s) {
                min0s = countOccurances(layer, 0);
                minLayer = layerIdx;
            }
            layerIdx++;
        }

        System.out.println(
            countOccurances(layers.get(minLayer), 1) *
            countOccurances(layers.get(minLayer), 2)
        );

        List<Integer> resultImg = new ArrayList<>();
        for (int i = 0; i < layerSize; i++) {
            // loop over layers
            for (List<Integer> layer : layers) {
                if (layer.get(i) != 2) {
                    resultImg.add(layer.get(i));
                    break;
                }
            }
        }
        assert resultImg.size() == layerSize;

        // print image
        for (int i = 0; i < layerSize; i++) {
            System.out.print(resultImg.get(i) == 0 ? " " : "█");
            if ((i + 1) % WIDTH == 0) {
                System.out.println();
            }
        }
    }

    public static int countOccurances(List<Integer> nums, int target) {
        int count = 0;
        for (int num : nums) {
            if (num == target) {
                count++;
            }
        }
        return count;
    }
}
