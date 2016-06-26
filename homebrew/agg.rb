require 'formula'

class Agg < Formula
  VERSION = '0.1.0'

  homepage 'https://github.com/winebarrel/agg'
  url "https://github.com/winebarrel/agg/releases/download/v#{VERSION}/agg-v#{VERSION}-darwin-amd64.gz"
  sha256 'b948d3d885ae5b7bae973f061d5d742cc09931541b1e64217cd6cf3cfeb98077'
  version VERSION
  head 'https://github.com/winebarrel/agg.git', :branch => 'master'

  def install
    system "mv agg-v#{VERSION}-darwin-amd64 agg"
    bin.install 'agg'
  end
end
