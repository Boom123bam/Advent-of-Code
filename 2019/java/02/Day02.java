import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.util.*;

public class Day02 {

    public static void main(String[] args) {
        String line = "";
        String filePath = "input.txt";
        try (BufferedReader br = new BufferedReader(new FileReader(filePath))) {
            line = br.readLine();
        } catch (IOException e) {
            e.printStackTrace();
        }

        int[] intArray = Arrays.stream(line.split(","))
            .mapToInt(Integer::parseInt)
            .toArray();

        // part 1
        int[] newArray = intArray.clone();
        newArray[1] = 12;
        newArray[2] = 2;
        System.out.println(run(newArray));

        // part 2
        for (int noun = 0; noun < intArray.length; noun++) {
            for (int verb = 0; verb < intArray.length; verb++) {
                newArray = intArray.clone();
                newArray[1] = noun;
                newArray[2] = verb;
                if (run(newArray) == 19690720) {
                    System.out.println(100 * noun + verb);
                }
            }
        }
    }

    public static int run(int[] nums) {
        for (int i = 0; i < nums.length; i += 4) {
            if (nums[i] == 99) {
                break;
            }
            int pos1 = nums[i + 1], pos2 = nums[i + 2], pos3 = nums[i + 3];
            if (nums[i] == 1) {
                nums[pos3] = nums[pos1] + nums[pos2];
            } else if (nums[i] == 2) {
                nums[pos3] = nums[pos1] * nums[pos2];
            } else {
                System.out.println("Not opcode");
            }
        }
        return nums[0];
    }
}
