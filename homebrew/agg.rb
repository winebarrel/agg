require 'formula'

class Agg < Formula
  VERSION = '0.1.0'

  homepage 'https://github.com/winebarrel/agg'
  url "https://github.com/winebarrel/agg/releases/download/v#{VERSION}/agg-v#{VERSION}-darwin-amd64.gz"
  sha256 '8bbcb9a78147dcd29f7eef748d7ab03a4cfbb60ae7d798768fe416251e2d5941'
  version VERSION
  head 'https://github.com/winebarrel/agg.git', :branch => 'master'

  def install
    system "mv agg-v#{VERSION}-darwin-amd64 agg"
    bin.install 'agg'
  end
end
