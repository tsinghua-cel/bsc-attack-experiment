import argparse


def count_consecutive_trues(file_path, min_height=250, max_height=2250):
    try:
        with open(file_path, 'r') as file:
            lines = file.readlines()

        valid_status = []

        # Extracts states within a specified height range and validates the height format
        for line in lines:
            parts = line.strip().split(',')
            if len(parts) < 3:
                continue  # Skip incorrectly formatted lines

            try:
                height = int(parts[0])
                if min_height <= height <= max_height:
                    valid_status.append(parts[2].lower() == 'true')
            except ValueError:
                continue  # Skip rows whose heights cannot be converted to integers

        # Counting the number of Finalized block
        consecutive_count = 0
        for i in range(1, len(valid_status)):
            if valid_status[i - 1] and valid_status[i]:
                consecutive_count += 1

        # Calculating FR
        total_pairs = max(0, len(valid_status) - 2)
        if total_pairs == 0:
            ratio = 0.0
        else:
            ratio = consecutive_count / total_pairs

        return {
            'consecutive_count': consecutive_count,
            'total_pairs': total_pairs,
            'ratio': ratio
        }

    except FileNotFoundError:
        print(f"Error: File ‘{file_path}’ does not exist")
        return None
    except Exception as e:
        print(f"Error: An exception occurred - {e}")
        return None


if __name__ == "__main__":
    parser = argparse.ArgumentParser(description='Calculating FR')
    parser.add_argument('file_path', type=str, help='input file path')
    parser.add_argument('--min', type=int, default=251, help='minimum height (default: 251)')
    parser.add_argument('--max', type=int, default=2251, help='maximum height (default: 2251)')
    args = parser.parse_args()

    result = count_consecutive_trues(args.file_path, args.min, args.max)

    if result:
        print(f"Calculated Height Range {args.min}-{args.max}:")
        print(f"  Number of Finalized blocks without attacks : {result['total_pairs']}")
        print(f"  Finalized blocks: {result['consecutive_count']}")
        print(f"  FR : {result['ratio']:.4f} ({result['ratio'] * 100:.2f}%)")