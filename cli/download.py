import argparse
import os
import shutil
import zipfile
from pathlib import Path


def parse_args():
    p = argparse.ArgumentParser()
    p.add_argument("--competition", default=os.getenv("COMPETITION"))
    p.add_argument("--dest", default=os.getenv("KAGGLE_DATA_DIR", "data/raw"))
    p.add_argument("--no-unzip", action="store_true")
    return p.parse_args()


def unzip_all(dest: Path):
    for z in dest.glob("*.zip"):
        with zipfile.ZipFile(z, "r") as zr:
            zr.extractall(dest)


def main():
    args = parse_args()
    if not args.competition:
        print("ERROR: COMPETITION is not set.")
        return 2
    dest = Path(args.dest)
    dest.mkdir(parents=True, exist_ok=True)

    # 1) kagglehub（新トークン対応）
    try:
        import kagglehub

        cache_path = Path(kagglehub.competition_download(args.competition))
        if cache_path.is_file():
            shutil.copy2(cache_path, dest / cache_path.name)
        else:
            for p in cache_path.rglob("*"):
                if p.is_file():
                    shutil.copy2(p, dest / p.name)
        if not args.no_unzip:
            unzip_all(dest)
        print(f"Downloaded competition data to {dest} (via kagglehub)")
        return 0
    except Exception as e:
        print(f"INFO: kagglehub failed, falling back. Details: {e}")

    # 2) legacy kaggle（kaggle.json の username/key 前提）
    from kaggle.api.kaggle_api_extended import KaggleApi

    api = KaggleApi()
    api.authenticate()
    api.competition_download_files(args.competition, path=str(dest))
    if not args.no_unzip:
        unzip_all(dest)
    print(f"Downloaded competition data to {dest} (via kaggle legacy)")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
