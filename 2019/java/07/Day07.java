import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.util.*;

public class Day07 {

    public static void main(String[] args) {
        String line = "";
        String filePath = "input.txt";
        try (BufferedReader br = new BufferedReader(new FileReader(filePath))) {
            line = br.readLine();
        } catch (IOException e) {
            e.printStackTrace();
        }

        int[] P = Arrays.stream(line.split(","))
            .mapToInt(Integer::parseInt)
            .toArray();

        // part 1
        int max = 0;
        int[] nums = { 0, 1, 2, 3, 4 };
        List<List<Integer>> perms = new ArrayList<>();
        permute(nums, 0, perms);
        for (List<Integer> perm : perms) {
            int output = runAmps(makeAmpList(P, perm));
            if (output > max) {
                max = output;
            }
        }
        System.out.println(max);

        // part 2
        max = 0;
        nums = new int[] { 5, 6, 7, 8, 9 };
        perms = new ArrayList<>();
        permute(nums, 0, perms);
        for (List<Integer> perm : perms) {
            int output = runAmps2(makeAmpList(P, perm));
            if (output > max) {
                max = output;
            }
        }
        System.out.println(max);
        // int arst = runAmps2(makeAmpList(P, List.of(9, 7, 8, 5, 6)));
        // System.out.println(arst);
        // newArray = originalArray.clone();
        // System.out.println(run(newArray, new int[] { 5 }));
        // System.out.println(runAmps(originalArray, new int[] { 1, 0, 4, 3, 2 }));
    }

    public static int runAmps2(List<Amp> amps) {
        int output = 0;
        for (Amp amp : amps) {
            output = amp.run(new int[] { amp.phase, output }, true);
        }
        while (true) {
            for (Amp amp : amps) {
                int result = amp.run(new int[] { output }, true);
                if (result == -1) {
                    return output;
                }
                output = result;
            }
        }
    }

    public static int runAmps(List<Amp> amps) {
        int output = 0;
        for (Amp amp : amps) {
            output = amp.run(new int[] { amp.phase, output }, false);
        }
        return output;
    }

    public static List<Amp> makeAmpList(int[] P, List<Integer> phases) {
        List<Amp> amps = new ArrayList<>();
        for (int i = 0; i < 5; i++) {
            amps.add(new Amp(P.clone(), phases.get(i)));
        }
        return amps;
    }

    public static void permute(
        int[] nums,
        int start,
        List<List<Integer>> result
    ) {
        if (start == nums.length) {
            // Add current permutation to result list
            List<Integer> current = new ArrayList<>();
            for (int num : nums) {
                current.add(num);
            }
            result.add(current);
        } else {
            for (int i = start; i < nums.length; i++) {
                swap(nums, start, i); // Swap the elements
                permute(nums, start + 1, result); // Recurse with the next element
                swap(nums, start, i); // Backtrack
            }
        }
    }

    public static void swap(int[] nums, int i, int j) {
        int temp = nums[i];
        nums[i] = nums[j];
        nums[j] = temp;
    }
}
