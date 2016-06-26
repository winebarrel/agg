require 'formula'

class Agg < Formula
  VERSION = '0.1.0'

  homepage 'https://github.com/winebarrel/agg'
  url "https://github.com/winebarrel/agg/releases/download/v#{VERSION}/agg-v#{VERSION}-darwin-amd64.gz"
  sha256 'bdbbd5962c257d2ccd3241a202243582311e2a625a353e829744c1ce2b44a2c2'
  version VERSION
  head 'https://github.com/winebarrel/agg.git', :branch => 'master'

  def install
    system "mv agg-v#{VERSION}-darwin-amd64 agg"
    bin.install 'agg'
  end
end
