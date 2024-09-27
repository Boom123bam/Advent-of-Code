import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.util.*;

public class Day05 {

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
        // newArray[1] = 12;
        // newArray[2] = 2;
        run(newArray, 8);
    }

    public static int run(int[] nums, int inputNum) {
        Boolean endProgram = false;
        for (int i = 0; i < nums.length;) {
            int instruction = nums[i];
            int opcode = instruction % 100;
            int[] paramModes = { 0, 0, 0 };
            instruction /= 100;
            for (int j = 0; j < 3; j++) {
                paramModes[j] = instruction % 10;
                instruction /= 10;
            }
            instruction = nums[i];
            // System.out.println(nums[i] + "," + i);
            int param1 = i + 1 < nums.length ? nums[i + 1] : 0;
            int param2 = i + 2 < nums.length ? nums[i + 2] : 0;
            int param3 = i + 3 < nums.length ? nums[i + 3] : 0;
            int first, second;
            // System.out.println(opcode);
            switch (opcode) {
                case 1:
                case 2:
                case 5:
                case 6:
                case 7:
                case 8:
                    first = (paramModes[0] == 0 ? nums[param1] : param1);
                    second = (paramModes[1] == 0 ? nums[param2] : param2);
                    switch (opcode) {
                        case 1:
                            nums[param3] = first + second;
                            i += 4;
                            break;
                        case 2:
                            nums[param3] = first * second;
                            i += 4;
                            break;
                        case 5:
                            if (first != 0) {
                                // nums[i] = second;
                                // i = nums[i];
                                i = second;
                                break;
                            }
                            i += 3;
                            break;
                        case 6:
                            if (first == 0) {
                                // nums[i] = second;
                                // i = nums[i];
                                i = second;
                                break;
                            }
                            i += 3;
                            break;
                        case 7:
                            nums[param3] = first < second ? 1 : 0;
                            i += 3;
                            break;
                        case 8:
                            nums[param3] = first == second ? 1 : 0;
                            i += 3;
                            break;
                    }
                    break;
                case 3:
                    nums[param1] = inputNum;
                    i += 2;
                    break;
                case 4:
                    System.out.println(nums[param1]);
                    // if part 2 check if instruction modified
                    i += 2;
                    break;
                case 99:
                    endProgram = true;
                    break;
                default:
                    System.out.println("Not opcode");
                    return -1;
            }
            if (endProgram) {
                break;
            }
        }
        return nums[0];
    }
}
